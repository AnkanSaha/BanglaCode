package builtins

import (
	"BanglaCode/src/object"
	"fmt"
	"net"
	"sync"
	"sync/atomic"
)

// UDP connection wrapper to store both connection and remote address
type UDPConnection struct {
	Conn       *net.UDPConn
	RemoteAddr *net.UDPAddr
}

// UDP connection registry with thread-safe access
var (
	udpConnections = make(map[string]*UDPConnection)
	udpMutex       sync.RWMutex
	udpCounter     int64
)

// generateUDPConnectionID creates a unique connection identifier
func generateUDPConnectionID() string {
	atomic.AddInt64(&udpCounter, 1)
	return fmt.Sprintf("udp_conn_%d", udpCounter)
}

// getUDPConnection retrieves a UDP connection by ID
func getUDPConnection(id string) (*UDPConnection, bool) {
	udpMutex.RLock()
	defer udpMutex.RUnlock()
	conn, ok := udpConnections[id]
	return conn, ok
}

// storeUDPConnection stores a UDP connection with a unique ID
func storeUDPConnection(id string, conn *UDPConnection) {
	udpMutex.Lock()
	defer udpMutex.Unlock()
	udpConnections[id] = conn
}

// removeUDPConnection removes and closes a UDP connection
func removeUDPConnection(id string) {
	udpMutex.Lock()
	defer udpMutex.Unlock()
	if conn, ok := udpConnections[id]; ok {
		conn.Conn.Close()
		delete(udpConnections, id)
	}
}

func init() {
	// udp_server_chalu(port, handler) - Start UDP server
	// Example: udp_server_chalu(9000, kaj(packet) { dekho("Received:", packet["data"]); });
	Builtins["udp_server_chalu"] = &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			// Validate arguments
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=2", len(args))
			}

			// Validate port (number)
			if args[0].Type() != object.NUMBER_OBJ {
				return newError("argument 1 to 'udp_server_chalu' must be NUMBER, got %s", args[0].Type())
			}

			// Validate handler (function)
			if args[1].Type() != object.FUNCTION_OBJ {
				return newError("argument 2 to 'udp_server_chalu' must be FUNCTION, got %s", args[1].Type())
			}

			port := int(args[0].(*object.Number).Value)
			handler := args[1].(*object.Function)

			// Create UDP listener
			addr := &net.UDPAddr{Port: port, IP: net.ParseIP("0.0.0.0")}
			conn, err := net.ListenUDP("udp", addr)
			if err != nil {
				return newError("UDP server error: %s", err.Error())
			}

			// Listen for packets in goroutine
			go func() {
				buffer := make([]byte, 4096)
				for {
					n, remoteAddr, err := conn.ReadFromUDP(buffer)
					if err != nil {
						continue
					}

					if n > 0 {
						// Create packet object
						packet := &object.Map{Pairs: make(map[string]object.Object)}
						connID := generateUDPConnectionID()

						// Store connection for response capability
						storeUDPConnection(connID, &UDPConnection{
							Conn:       conn,
							RemoteAddr: remoteAddr,
						})

						packet.Pairs["id"] = &object.String{Value: connID}
						packet.Pairs["data"] = &object.String{Value: string(buffer[:n])}
						packet.Pairs["remote_addr"] = &object.String{Value: remoteAddr.String()}
						packet.Pairs["local_addr"] = &object.String{Value: conn.LocalAddr().String()}

						// Call user handler
						if EvalFunc != nil {
							EvalFunc(handler, []object.Object{packet})
						}
					}
				}
			}()

			return object.NULL
		},
	}

	// udp_uttor(connection, data) - Send UDP response to client
	// Example: udp_uttor(packet, "Response message");
	Builtins["udp_uttor"] = &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			// Validate arguments
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=2", len(args))
			}

			// Validate connection (map)
			if args[0].Type() != object.MAP_OBJ {
				return newError("argument 1 to 'udp_uttor' must be MAP, got %s", args[0].Type())
			}

			// Validate data (string)
			if args[1].Type() != object.STRING_OBJ {
				return newError("argument 2 to 'udp_uttor' must be STRING, got %s", args[1].Type())
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

			// Get UDP connection
			udpConn, ok := getUDPConnection(connID)
			if !ok {
				return newError("UDP connection not found or closed")
			}

			// Send response to remote address
			_, err := udpConn.Conn.WriteToUDP([]byte(data), udpConn.RemoteAddr)
			if err != nil {
				return newError("UDP send error: %s", err.Error())
			}

			return object.NULL
		},
	}

	// udp_pathao(host, port, data) - Send UDP packet (async, returns promise)
	// Example: opekha udp_pathao("localhost", 9000, "Hello UDP!");
	Builtins["udp_pathao"] = &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			// Validate arguments
			if len(args) != 3 {
				return newError("wrong number of arguments. got=%d, want=3", len(args))
			}

			// Validate host (string)
			if args[0].Type() != object.STRING_OBJ {
				return newError("argument 1 to 'udp_pathao' must be STRING, got %s", args[0].Type())
			}

			// Validate port (number)
			if args[1].Type() != object.NUMBER_OBJ {
				return newError("argument 2 to 'udp_pathao' must be NUMBER, got %s", args[1].Type())
			}

			// Validate data (string)
			if args[2].Type() != object.STRING_OBJ {
				return newError("argument 3 to 'udp_pathao' must be STRING, got %s", args[2].Type())
			}

			host := args[0].(*object.String).Value
			port := int(args[1].(*object.Number).Value)
			data := args[2].(*object.String).Value

			// Create promise
			promise := object.CreatePromise()

			// Send asynchronously
			go func() {
				addr := &net.UDPAddr{Port: port, IP: net.ParseIP(host)}
				conn, err := net.DialUDP("udp", nil, addr)
				if err != nil {
					object.RejectPromise(promise, newError("UDP dial failed: %s", err.Error()))
					return
				}
				defer conn.Close()

				_, err = conn.Write([]byte(data))
				if err != nil {
					object.RejectPromise(promise, newError("UDP send failed: %s", err.Error()))
					return
				}

				object.ResolvePromise(promise, object.NULL)
			}()

			return promise
		},
	}

	// udp_shuno(port, handler) - Listen for UDP packets on port (alias for udp_server_chalu)
	// Example: udp_shuno(9000, kaj(packet) { dekho("Got:", packet["data"]); });
	Builtins["udp_shuno"] = &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			// Just call udp_server_chalu
			return Builtins["udp_server_chalu"].Fn(args...)
		},
	}

	// udp_bondho(connection) - Close UDP connection
	// Example: udp_bondho(packet);
	Builtins["udp_bondho"] = &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			// Validate arguments
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			// Validate connection (map)
			if args[0].Type() != object.MAP_OBJ {
				return newError("argument to 'udp_bondho' must be MAP, got %s", args[0].Type())
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
			removeUDPConnection(connID)

			return object.NULL
		},
	}
}
