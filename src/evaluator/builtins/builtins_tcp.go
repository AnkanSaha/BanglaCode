package builtins

import (
	"BanglaCode/src/object"
	"fmt"
	"net"
	"sync"
	"sync/atomic"
)

// TCP connection registry with thread-safe access
var (
	tcpConnections = make(map[string]net.Conn)
	tcpMutex       sync.RWMutex
	tcpCounter     int64
)

// generateTCPConnectionID creates a unique connection identifier
func generateTCPConnectionID() string {
	atomic.AddInt64(&tcpCounter, 1)
	return fmt.Sprintf("tcp_conn_%d", tcpCounter)
}

// getTCPConnection retrieves a TCP connection by ID
func getTCPConnection(id string) (net.Conn, bool) {
	tcpMutex.RLock()
	defer tcpMutex.RUnlock()
	conn, ok := tcpConnections[id]
	return conn, ok
}

// storeTCPConnection stores a TCP connection with a unique ID
func storeTCPConnection(id string, conn net.Conn) {
	tcpMutex.Lock()
	defer tcpMutex.Unlock()
	tcpConnections[id] = conn
}

// removeTCPConnection removes and closes a TCP connection
func removeTCPConnection(id string) {
	tcpMutex.Lock()
	defer tcpMutex.Unlock()
	if conn, ok := tcpConnections[id]; ok {
		conn.Close()
		delete(tcpConnections, id)
	}
}

// handleTCPConnection handles incoming TCP connections with callback
func handleTCPConnection(conn net.Conn, handler *object.Function) {
	// Create connection object
	connObj := &object.Map{Pairs: make(map[string]object.Object)}
	connID := generateTCPConnectionID()
	storeTCPConnection(connID, conn)

	connObj.Pairs["id"] = &object.String{Value: connID}
	connObj.Pairs["remote_addr"] = &object.String{Value: conn.RemoteAddr().String()}
	connObj.Pairs["local_addr"] = &object.String{Value: conn.LocalAddr().String()}

	// Read data loop
	buffer := make([]byte, 4096)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			// Connection closed or error
			removeTCPConnection(connID)
			break
		}

		if n > 0 {
			// Update connection object with received data
			connObj.Pairs["data"] = &object.String{Value: string(buffer[:n])}

			// Call user handler
			if EvalFunc != nil {
				EvalFunc(handler, []object.Object{connObj})
			}
		}
	}
}

func init() {
	// tcp_server_chalu(port, handler) - Start TCP server
	// Example: tcp_server_chalu(8080, kaj(conn) { dekho("Connected:", conn["remote_addr"]); })
	Builtins["tcp_server_chalu"] = &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			// Validate arguments
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=2", len(args))
			}

			// Validate port (number)
			if args[0].Type() != object.NUMBER_OBJ {
				return newError("argument 1 to 'tcp_server_chalu' must be NUMBER, got %s", args[0].Type())
			}

			// Validate handler (function)
			if args[1].Type() != object.FUNCTION_OBJ {
				return newError("argument 2 to 'tcp_server_chalu' must be FUNCTION, got %s", args[1].Type())
			}

			port := int(args[0].(*object.Number).Value)
			handler := args[1].(*object.Function)

			// Create TCP listener
			listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
			if err != nil {
				return newError("TCP server error: %s", err.Error())
			}

			// Accept connections in goroutine
			go func() {
				for {
					conn, err := listener.Accept()
					if err != nil {
						continue
					}

					// Handle each connection in separate goroutine
					go handleTCPConnection(conn, handler)
				}
			}()

			return object.NULL
		},
	}

	// tcp_jukto(host, port) - Connect to TCP server (async, returns promise)
	// Example: dhoro conn = opekha tcp_jukto("localhost", 8080);
	Builtins["tcp_jukto"] = &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			// Validate arguments
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=2", len(args))
			}

			// Validate host (string)
			if args[0].Type() != object.STRING_OBJ {
				return newError("argument 1 to 'tcp_jukto' must be STRING, got %s", args[0].Type())
			}

			// Validate port (number)
			if args[1].Type() != object.NUMBER_OBJ {
				return newError("argument 2 to 'tcp_jukto' must be NUMBER, got %s", args[1].Type())
			}

			host := args[0].(*object.String).Value
			port := int(args[1].(*object.Number).Value)

			// Create promise
			promise := object.CreatePromise()

			// Connect asynchronously
			go func() {
				// Use net.JoinHostPort for proper IPv6 support
				addr := net.JoinHostPort(host, fmt.Sprintf("%d", port))
				conn, err := net.Dial("tcp", addr)
				if err != nil {
					object.RejectPromise(promise, newError("TCP connection failed: %s", err.Error()))
					return
				}

				// Create connection object
				connObj := &object.Map{Pairs: make(map[string]object.Object)}
				connID := generateTCPConnectionID()
				storeTCPConnection(connID, conn)

				connObj.Pairs["id"] = &object.String{Value: connID}
				connObj.Pairs["host"] = &object.String{Value: host}
				connObj.Pairs["port"] = &object.Number{Value: float64(port)}
				connObj.Pairs["remote_addr"] = &object.String{Value: conn.RemoteAddr().String()}
				connObj.Pairs["local_addr"] = &object.String{Value: conn.LocalAddr().String()}

				object.ResolvePromise(promise, connObj)
			}()

			return promise
		},
	}

	// tcp_pathao(connection, data) - Send data on TCP connection
	// Example: tcp_pathao(conn, "Hello!");
	Builtins["tcp_pathao"] = &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			// Validate arguments
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=2", len(args))
			}

			// Validate connection (map)
			if args[0].Type() != object.MAP_OBJ {
				return newError("argument 1 to 'tcp_pathao' must be MAP, got %s", args[0].Type())
			}

			// Validate data (string)
			if args[1].Type() != object.STRING_OBJ {
				return newError("argument 2 to 'tcp_pathao' must be STRING, got %s", args[1].Type())
			}

			connMap := args[0].(*object.Map)
			data := args[1].(*object.String).Value

			// Get connection ID
			idObj, ok := connMap.Pairs["id"]
			if !ok {
				return newError("connection object missing 'id' field")
			}

			if idObj.Type() != object.STRING_OBJ {
				return newError("connection 'id' must be STRING")
			}

			connID := idObj.(*object.String).Value

			// Get TCP connection
			conn, ok := getTCPConnection(connID)
			if !ok {
				return newError("TCP connection not found or closed")
			}

			// Send data
			_, err := conn.Write([]byte(data))
			if err != nil {
				return newError("TCP send error: %s", err.Error())
			}

			return object.NULL
		},
	}

	// tcp_lekho(connection, data) - Write data to TCP connection (alias for tcp_pathao)
	// Example: tcp_lekho(conn, "Message");
	Builtins["tcp_lekho"] = &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			// Just call tcp_pathao
			return Builtins["tcp_pathao"].Fn(args...)
		},
	}

	// tcp_shuno(connection) - Read data from TCP connection (async, returns promise)
	// Example: dhoro data = opekha tcp_shuno(conn);
	Builtins["tcp_shuno"] = &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			// Validate arguments
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			// Validate connection (map)
			if args[0].Type() != object.MAP_OBJ {
				return newError("argument to 'tcp_shuno' must be MAP, got %s", args[0].Type())
			}

			connMap := args[0].(*object.Map)

			// Get connection ID
			idObj, ok := connMap.Pairs["id"]
			if !ok {
				return newError("connection object missing 'id' field")
			}

			if idObj.Type() != object.STRING_OBJ {
				return newError("connection 'id' must be STRING")
			}

			connID := idObj.(*object.String).Value

			// Get TCP connection
			conn, ok := getTCPConnection(connID)
			if !ok {
				return newError("TCP connection not found or closed")
			}

			// Create promise
			promise := object.CreatePromise()

			// Read asynchronously
			go func() {
				buffer := make([]byte, 4096)
				n, err := conn.Read(buffer)
				if err != nil {
					object.RejectPromise(promise, newError("TCP read error: %s", err.Error()))
					return
				}

				data := &object.String{Value: string(buffer[:n])}
				object.ResolvePromise(promise, data)
			}()

			return promise
		},
	}

	// tcp_bondho(connection) - Close TCP connection
	// Example: tcp_bondho(conn);
	Builtins["tcp_bondho"] = &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			// Validate arguments
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			// Validate connection (map)
			if args[0].Type() != object.MAP_OBJ {
				return newError("argument to 'tcp_bondho' must be MAP, got %s", args[0].Type())
			}

			connMap := args[0].(*object.Map)

			// Get connection ID
			idObj, ok := connMap.Pairs["id"]
			if !ok {
				return newError("connection object missing 'id' field")
			}

			if idObj.Type() != object.STRING_OBJ {
				return newError("connection 'id' must be STRING")
			}

			connID := idObj.(*object.String).Value

			// Remove and close connection
			removeTCPConnection(connID)

			return object.NULL
		},
	}
}
