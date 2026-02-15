package database

import (
	"BanglaCode/src/object"
	"database/sql"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// Global pool registry (thread-safe)
var (
	pools         = make(map[string]*ConnectionPool)
	poolsMutex    sync.RWMutex
	poolIDCounter int64
)

// PoolConfig defines connection pool configuration
type PoolConfig struct {
	MaxConns       int           // Maximum connections (default: 10)
	MinConns       int           // Minimum idle connections (default: 2)
	MaxIdleTime    time.Duration // Max idle time before closing (default: 5min)
	ConnectTimeout time.Duration // Connection timeout (default: 10s)
}

// DefaultPoolConfig returns default pool configuration
func DefaultPoolConfig() PoolConfig {
	return PoolConfig{
		MaxConns:       10,
		MinConns:       2,
		MaxIdleTime:    5 * time.Minute,
		ConnectTimeout: 10 * time.Second,
	}
}

// ConnectionPool manages a pool of database connections
type ConnectionPool struct {
	id          string
	dbType      string
	config      PoolConfig
	connConfig  map[string]interface{}    // Connection configuration
	conns       chan *object.DBConnection // Buffered channel for available connections
	activeConns int64                     // Atomic counter for active connections
	totalConns  int64                     // Atomic counter for total connections created
	mu          sync.RWMutex
	closed      bool
	closeChan   chan struct{} // Signal to stop cleanup goroutine
}

// NewConnectionPool creates a new connection pool
func NewConnectionPool(dbType string, connConfig map[string]interface{}, maxConns int) (*ConnectionPool, error) {
	if maxConns <= 0 {
		maxConns = 10
	}

	id := atomic.AddInt64(&poolIDCounter, 1)
	poolID := fmt.Sprintf("%s-pool-%d", dbType, id)

	config := DefaultPoolConfig()
	config.MaxConns = maxConns

	pool := &ConnectionPool{
		id:         poolID,
		dbType:     dbType,
		config:     config,
		connConfig: connConfig,
		conns:      make(chan *object.DBConnection, maxConns), // Buffered channel
		closeChan:  make(chan struct{}),
	}

	// Pre-allocate minimum connections
	for i := 0; i < config.MinConns; i++ {
		conn, err := pool.createConnection()
		if err != nil {
			// Clean up any created connections
			pool.Close()
			return nil, fmt.Errorf("failed to pre-allocate connection: %v", err)
		}
		pool.conns <- conn
	}

	// Start cleanup goroutine
	go pool.cleanupIdleConnections()

	// Register pool globally
	poolsMutex.Lock()
	pools[poolID] = pool
	poolsMutex.Unlock()

	return pool, nil
}

// GetPool retrieves a pool by ID
func GetPool(poolID string) (*ConnectionPool, error) {
	poolsMutex.RLock()
	pool, ok := pools[poolID]
	poolsMutex.RUnlock()

	if !ok {
		return nil, fmt.Errorf("pool %s not found", poolID)
	}
	return pool, nil
}

// createConnection creates a new database connection
func (p *ConnectionPool) createConnection() (*object.DBConnection, error) {
	// This will be implemented by specific database connectors
	// For now, create a placeholder connection
	connID := generateConnID(p.dbType)

	conn := &object.DBConnection{
		ID:       connID,
		DBType:   p.dbType,
		PoolID:   p.id,
		Metadata: make(map[string]object.Object),
	}

	atomic.AddInt64(&p.totalConns, 1)
	return conn, nil
}

// Get retrieves a connection from the pool (blocks if all connections in use)
func (p *ConnectionPool) Get() (*object.DBConnection, error) {
	p.mu.RLock()
	if p.closed {
		p.mu.RUnlock()
		return nil, fmt.Errorf("pool %s is closed", p.id)
	}
	p.mu.RUnlock()

	select {
	case conn := <-p.conns:
		// Got connection from pool
		atomic.AddInt64(&p.activeConns, 1)

		// Validate connection (ping if SQL)
		if err := p.validateConnection(conn); err != nil {
			// Connection invalid, create new one
			atomic.AddInt64(&p.activeConns, -1)
			newConn, createErr := p.createConnection()
			if createErr != nil {
				return nil, fmt.Errorf("failed to create new connection: %v", createErr)
			}
			atomic.AddInt64(&p.activeConns, 1)
			return newConn, nil
		}

		return conn, nil

	default:
		// No idle connections, check if we can create more
		total := atomic.LoadInt64(&p.totalConns)
		if int(total) < p.config.MaxConns {
			// Create new connection
			conn, err := p.createConnection()
			if err != nil {
				return nil, fmt.Errorf("failed to create connection: %v", err)
			}
			atomic.AddInt64(&p.activeConns, 1)
			return conn, nil
		}

		// Pool exhausted, block until connection available
		conn := <-p.conns
		atomic.AddInt64(&p.activeConns, 1)

		// Validate connection
		if err := p.validateConnection(conn); err != nil {
			atomic.AddInt64(&p.activeConns, -1)
			newConn, createErr := p.createConnection()
			if createErr != nil {
				return nil, fmt.Errorf("failed to create new connection: %v", createErr)
			}
			atomic.AddInt64(&p.activeConns, 1)
			return newConn, nil
		}

		return conn, nil
	}
}

// Return returns a connection to the pool
func (p *ConnectionPool) Return(conn *object.DBConnection) error {
	p.mu.RLock()
	if p.closed {
		p.mu.RUnlock()
		// Pool closed, close the connection
		return p.closeConnection(conn)
	}
	p.mu.RUnlock()

	if conn == nil {
		return fmt.Errorf("cannot return nil connection")
	}

	if conn.PoolID != p.id {
		return fmt.Errorf("connection belongs to different pool")
	}

	atomic.AddInt64(&p.activeConns, -1)

	// Try to return to pool (non-blocking)
	select {
	case p.conns <- conn:
		// Successfully returned to pool
		return nil
	default:
		// Pool full (shouldn't happen normally), close connection
		atomic.AddInt64(&p.totalConns, -1)
		return p.closeConnection(conn)
	}
}

// validateConnection validates a connection (pings if SQL database)
func (p *ConnectionPool) validateConnection(conn *object.DBConnection) error {
	if conn.Native == nil {
		return fmt.Errorf("connection has no native handle")
	}

	// For SQL databases, ping to validate
	if p.dbType == "postgres" || p.dbType == "mysql" {
		if db, ok := conn.Native.(*sql.DB); ok {
			return db.Ping()
		}
	}

	// For other databases, assume valid (MongoDB/Redis have their own health checks)
	return nil
}

// closeConnection closes a database connection
func (p *ConnectionPool) closeConnection(conn *object.DBConnection) error {
	if conn.Native == nil {
		return nil
	}

	// Close native connection based on type
	switch p.dbType {
	case "postgres", "mysql":
		if db, ok := conn.Native.(*sql.DB); ok {
			return db.Close()
		}
		// MongoDB and Redis closers will be added when those connectors are implemented
	}

	return nil
}

// cleanupIdleConnections periodically closes idle connections
func (p *ConnectionPool) cleanupIdleConnections() {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			// Check if we have more than MinConns idle
			p.mu.RLock()
			if p.closed {
				p.mu.RUnlock()
				return
			}
			p.mu.RUnlock()

			// Try to close excess idle connections
			idleCount := len(p.conns)
			if idleCount > p.config.MinConns {
				excess := idleCount - p.config.MinConns
				for i := 0; i < excess; i++ {
					select {
					case conn := <-p.conns:
						atomic.AddInt64(&p.totalConns, -1)
						p.closeConnection(conn)
					default:
						// No more idle connections
						break
					}
				}
			}

		case <-p.closeChan:
			// Pool is closing
			return
		}
	}
}

// Close closes the connection pool and all connections
func (p *ConnectionPool) Close() error {
	p.mu.Lock()
	if p.closed {
		p.mu.Unlock()
		return fmt.Errorf("pool already closed")
	}
	p.closed = true
	p.mu.Unlock()

	// Signal cleanup goroutine to stop
	close(p.closeChan)

	// Close all idle connections
	close(p.conns)
	for conn := range p.conns {
		p.closeConnection(conn)
	}

	// Remove from global registry
	poolsMutex.Lock()
	delete(pools, p.id)
	poolsMutex.Unlock()

	return nil
}

// Stats returns pool statistics
func (p *ConnectionPool) Stats() map[string]interface{} {
	p.mu.RLock()
	defer p.mu.RUnlock()

	return map[string]interface{}{
		"id":           p.id,
		"type":         p.dbType,
		"max_conns":    p.config.MaxConns,
		"total_conns":  atomic.LoadInt64(&p.totalConns),
		"active_conns": atomic.LoadInt64(&p.activeConns),
		"idle_conns":   len(p.conns),
		"closed":       p.closed,
	}
}
