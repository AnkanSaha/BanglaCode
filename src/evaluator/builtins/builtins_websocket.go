package builtins

import (
	"BanglaCode/src/object"
	"fmt"
	"net/http"
	"sync"
	"sync/atomic"

	"github.com/gorilla/websocket"
)

// WebSocket connection registry with thread-safe access
var (
	wsConnections = make(map[string]*websocket.Conn)
	wsMutex       sync.RWMutex
	wsCounter     int64
)

// WebSocket upgrader with permissive settings
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all origins
	},
	ReadBufferSize:  4096,
	WriteBufferSize: 4096,
}

// generateWSConnectionID creates a unique connection identifier
func generateWSConnectionID() string {
	atomic.AddInt64(&wsCounter, 1)
	return fmt.Sprintf("ws_conn_%d", wsCounter)
}

// getWSConnection retrieves a WebSocket connection by ID
func getWSConnection(id string) (*websocket.Conn, bool) {
	wsMutex.RLock()
	defer wsMutex.RUnlock()
	conn, ok := wsConnections[id]
	return conn, ok
}

// storeWSConnection stores a WebSocket connection with a unique ID
func storeWSConnection(id string, conn *websocket.Conn) {
	wsMutex.Lock()
	defer wsMutex.Unlock()
	wsConnections[id] = conn
}

// removeWSConnection removes and closes a WebSocket connection
func removeWSConnection(id string) {
	wsMutex.Lock()
	defer wsMutex.Unlock()
	if conn, ok := wsConnections[id]; ok {
		conn.Close()
		delete(wsConnections, id)
	}
}

// handleWebSocketConnection handles incoming WebSocket connections
func handleWebSocketConnection(conn *websocket.Conn, handler *object.Function) {
	// Create connection object
	connObj := &object.Map{Pairs: make(map[string]object.Object)}
	connID := generateWSConnectionID()
	storeWSConnection(connID, conn)

	connObj.Pairs["id"] = &object.String{Value: connID}
	connObj.Pairs["remote_addr"] = &object.String{Value: conn.RemoteAddr().String()}
	connObj.Pairs["local_addr"] = &object.String{Value: conn.LocalAddr().String()}
	connObj.Pairs["connected"] = &object.Boolean{Value: true}

	// Read messages loop
	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			// Connection closed or error
			connObj.Pairs["connected"] = &object.Boolean{Value: false}
			removeWSConnection(connID)
			break
		}

		// Determine message type
		msgType := "text"
		if messageType == websocket.BinaryMessage {
			msgType = "binary"
		} else if messageType == websocket.CloseMessage {
			msgType = "close"
		} else if messageType == websocket.PingMessage {
			msgType = "ping"
		} else if messageType == websocket.PongMessage {
			msgType = "pong"
		}

		// Update connection object with message data
		connObj.Pairs["message"] = &object.String{Value: string(message)}
		connObj.Pairs["type"] = &object.String{Value: msgType}

		// Call user handler
		if EvalFunc != nil {
			EvalFunc(handler, []object.Object{connObj})
		}

		// Break if close message
		if messageType == websocket.CloseMessage {
			break
		}
	}
}

func init() {
	// websocket_server_chalu(port, handler) - Start WebSocket server
	// Example: websocket_server_chalu(3000, kaj(conn) { dekho("Message:", conn["message"]); });
	Builtins["websocket_server_chalu"] = &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			// Validate arguments
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=2", len(args))
			}

			// Validate port (number)
			if args[0].Type() != object.NUMBER_OBJ {
				return newError("argument 1 to 'websocket_server_chalu' must be NUMBER, got %s", args[0].Type())
			}

			// Validate handler (function)
			if args[1].Type() != object.FUNCTION_OBJ {
				return newError("argument 2 to 'websocket_server_chalu' must be FUNCTION, got %s", args[1].Type())
			}

			port := int(args[0].(*object.Number).Value)
			handler := args[1].(*object.Function)

			// Create HTTP handler for WebSocket upgrade
			http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
				// Upgrade HTTP connection to WebSocket
				conn, err := upgrader.Upgrade(w, r, nil)
				if err != nil {
					return
				}

				// Handle connection in goroutine
				go handleWebSocketConnection(conn, handler)
			})

			// Start server in goroutine
			go func() {
				addr := fmt.Sprintf(":%d", port)
				if err := http.ListenAndServe(addr, nil); err != nil {
					// Server error (ignore for now as it's in goroutine)
				}
			}()

			return object.NULL
		},
	}

	// websocket_jukto(url) - Connect to WebSocket server (async, returns promise)
	// Example: dhoro ws = opekha websocket_jukto("ws://localhost:3000");
	Builtins["websocket_jukto"] = &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			// Validate arguments
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			// Validate URL (string)
			if args[0].Type() != object.STRING_OBJ {
				return newError("argument to 'websocket_jukto' must be STRING, got %s", args[0].Type())
			}

			url := args[0].(*object.String).Value

			// Create promise
			promise := object.CreatePromise()

			// Connect asynchronously
			go func() {
				conn, _, err := websocket.DefaultDialer.Dial(url, nil)
				if err != nil {
					object.RejectPromise(promise, newError("WebSocket connection failed: %s", err.Error()))
					return
				}

				// Create connection object
				connObj := &object.Map{Pairs: make(map[string]object.Object)}
				connID := generateWSConnectionID()
				storeWSConnection(connID, conn)

				connObj.Pairs["id"] = &object.String{Value: connID}
				connObj.Pairs["url"] = &object.String{Value: url}
				connObj.Pairs["connected"] = &object.Boolean{Value: true}
				connObj.Pairs["remote_addr"] = &object.String{Value: conn.RemoteAddr().String()}
				connObj.Pairs["local_addr"] = &object.String{Value: conn.LocalAddr().String()}

				object.ResolvePromise(promise, connObj)
			}()

			return promise
		},
	}

	// websocket_pathao(connection, message) - Send WebSocket message
	// Example: websocket_pathao(ws, "Hello WebSocket!");
	Builtins["websocket_pathao"] = &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			// Validate arguments
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=2", len(args))
			}

			// Validate connection (map)
			if args[0].Type() != object.MAP_OBJ {
				return newError("argument 1 to 'websocket_pathao' must be MAP, got %s", args[0].Type())
			}

			// Validate message (string)
			if args[1].Type() != object.STRING_OBJ {
				return newError("argument 2 to 'websocket_pathao' must be STRING, got %s", args[1].Type())
			}

			connMap := args[0].(*object.Map)
			message := args[1].(*object.String).Value

			// Get connection ID
			idObj, ok := connMap.Pairs["id"]
			if !ok {
				return newError("connection object missing 'id' field")
			}

			if idObj.Type() != object.STRING_OBJ {
				return newError("connection 'id' must be STRING")
			}

			connID := idObj.(*object.String).Value

			// Get WebSocket connection
			conn, ok := getWSConnection(connID)
			if !ok {
				return newError("WebSocket connection not found or closed")
			}

			// Send text message
			err := conn.WriteMessage(websocket.TextMessage, []byte(message))
			if err != nil {
				return newError("WebSocket send error: %s", err.Error())
			}

			return object.NULL
		},
	}

	// websocket_bondho(connection) - Close WebSocket connection
	// Example: websocket_bondho(ws);
	Builtins["websocket_bondho"] = &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			// Validate arguments
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			// Validate connection (map)
			if args[0].Type() != object.MAP_OBJ {
				return newError("argument to 'websocket_bondho' must be MAP, got %s", args[0].Type())
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

			// Get WebSocket connection
			conn, ok := getWSConnection(connID)
			if !ok {
				return newError("WebSocket connection not found or already closed")
			}

			// Send close message
			err := conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err == nil {
				// Wait a bit for close handshake
				conn.Close()
			}

			// Remove from registry
			removeWSConnection(connID)

			return object.NULL
		},
	}
}
