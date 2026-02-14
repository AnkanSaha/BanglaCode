package test

import (
	"BanglaCode/src/evaluator"
	"BanglaCode/src/evaluator/builtins"
	"BanglaCode/src/lexer"
	"BanglaCode/src/object"
	"BanglaCode/src/parser"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gorilla/websocket"
)

// Helper function to evaluate BanglaCode
func evalWS(input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	env := object.NewEnvironment()

	// Set EvalFunc for builtins
	builtins.EvalFunc = func(fn *object.Function, args []object.Object) object.Object {
		return evaluator.Eval(fn.Body, object.NewEnclosedEnvironment(fn.Env))
	}

	return evaluator.Eval(program, env)
}

func TestWebSocketServerChalu(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
		errMsg  string
	}{
		{
			name: "WebSocket server with valid port and handler",
			input: `
				bishwo ws_started = mittha;
				websocket_server_chalu(3001, kaj(conn) {
					ws_started = sotti;
				});
				ws_started;
			`,
			wantErr: false,
		},
		{
			name:    "WebSocket server missing arguments",
			input:   `websocket_server_chalu(3002);`,
			wantErr: true,
			errMsg:  "wrong number of arguments",
		},
		{
			name:    "WebSocket server invalid port type",
			input:   `websocket_server_chalu("not a port", kaj(conn) {});`,
			wantErr: true,
			errMsg:  "must be NUMBER",
		},
		{
			name:    "WebSocket server invalid handler type",
			input:   `websocket_server_chalu(3003, "not a function");`,
			wantErr: true,
			errMsg:  "must be FUNCTION",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := evalWS(tt.input)

			if tt.wantErr {
				errObj, ok := result.(*object.Error)
				if !ok {
					t.Fatalf("expected error, got %T (%+v)", result, result)
				}
				if tt.errMsg != "" && !wsContains(errObj.Message, tt.errMsg) {
					t.Errorf("error message = %q, want to contain %q", errObj.Message, tt.errMsg)
				}
			} else {
				if _, ok := result.(*object.Error); ok {
					t.Fatalf("unexpected error: %s", result.Inspect())
				}
			}
		})
	}
}

func TestWebSocketJukto(t *testing.T) {
	// Create a test WebSocket server
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			return
		}
		defer conn.Close()

		// Echo messages
		for {
			messageType, message, err := conn.ReadMessage()
			if err != nil {
				break
			}
			conn.WriteMessage(messageType, message)
		}
	}))
	defer server.Close()

	// Convert HTTP URL to WebSocket URL
	wsURL := "ws" + strings.TrimPrefix(server.URL, "http")

	tests := []struct {
		name    string
		input   string
		wantErr bool
		errMsg  string
	}{
		{
			name: "WebSocket jukto valid connection",
			input: `
				proyash kaj connect() {
					dhoro ws = opekha websocket_jukto("` + wsURL + `");
					ferao ws;
				}
				connect();
			`,
			wantErr: false,
		},
		{
			name:    "WebSocket jukto missing arguments",
			input:   `websocket_jukto();`,
			wantErr: true,
			errMsg:  "wrong number of arguments",
		},
		{
			name:    "WebSocket jukto invalid URL type",
			input:   `websocket_jukto(123);`,
			wantErr: true,
			errMsg:  "must be STRING",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := evalWS(tt.input)

			if tt.wantErr {
				errObj, ok := result.(*object.Error)
				if !ok {
					t.Fatalf("expected error, got %T (%+v)", result, result)
				}
				if tt.errMsg != "" && !wsContains(errObj.Message, tt.errMsg) {
					t.Errorf("error message = %q, want to contain %q", errObj.Message, tt.errMsg)
				}
			} else {
				if _, ok := result.(*object.Error); ok {
					t.Fatalf("unexpected error: %s", result.Inspect())
				}

				// For async operations, check if it's a promise
				if promise, ok := result.(*object.Promise); ok {
					// Wait for promise to resolve
					time.Sleep(300 * time.Millisecond)

					if promise.State != "RESOLVED" && promise.State != "fulfilled" {
						if promise.State == "rejected" {
							if errVal, ok := promise.Value.(*object.Error); ok {
								t.Errorf("promise rejected with error: %s", errVal.Message)
							} else {
								t.Errorf("promise rejected with value: %s", promise.Value.Inspect())
							}
						} else {
							t.Errorf("promise state = %s, want fulfilled", promise.State)
						}
						return
					}

					if connMap, ok := promise.Value.(*object.Map); ok {
						// Check connection object has required fields
						if _, hasID := connMap.Pairs["id"]; !hasID {
							t.Error("connection missing 'id' field")
						}
						if _, hasURL := connMap.Pairs["url"]; !hasURL {
							t.Error("connection missing 'url' field")
						}
						if _, hasConnected := connMap.Pairs["connected"]; !hasConnected {
							t.Error("connection missing 'connected' field")
						}
					} else {
						t.Errorf("expected connection map, got %T", promise.Value)
					}
				}
			}
		})
	}
}

func TestWebSocketPathao(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
		errMsg  string
	}{
		{
			name:    "WebSocket pathao missing arguments",
			input:   `websocket_pathao({});`,
			wantErr: true,
			errMsg:  "wrong number of arguments",
		},
		{
			name:    "WebSocket pathao invalid connection type",
			input:   `websocket_pathao("not a map", "message");`,
			wantErr: true,
			errMsg:  "must be MAP",
		},
		{
			name:    "WebSocket pathao invalid message type",
			input:   `websocket_pathao({}, 123);`,
			wantErr: true,
			errMsg:  "must be STRING",
		},
		{
			name:    "WebSocket pathao missing connection ID",
			input:   `websocket_pathao({}, "message");`,
			wantErr: true,
			errMsg:  "missing 'id' field",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := evalWS(tt.input)

			if tt.wantErr {
				errObj, ok := result.(*object.Error)
				if !ok {
					t.Fatalf("expected error, got %T (%+v)", result, result)
				}
				if tt.errMsg != "" && !wsContains(errObj.Message, tt.errMsg) {
					t.Errorf("error message = %q, want to contain %q", errObj.Message, tt.errMsg)
				}
			}
		})
	}
}

func TestWebSocketBondho(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
		errMsg  string
	}{
		{
			name:    "WebSocket bondho missing arguments",
			input:   `websocket_bondho();`,
			wantErr: true,
			errMsg:  "wrong number of arguments",
		},
		{
			name:    "WebSocket bondho invalid connection type",
			input:   `websocket_bondho("not a map");`,
			wantErr: true,
			errMsg:  "must be MAP",
		},
		{
			name:    "WebSocket bondho missing connection ID",
			input:   `websocket_bondho({});`,
			wantErr: true,
			errMsg:  "missing 'id' field",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := evalWS(tt.input)

			if tt.wantErr {
				errObj, ok := result.(*object.Error)
				if !ok {
					t.Fatalf("expected error, got %T (%+v)", result, result)
				}
				if tt.errMsg != "" && !wsContains(errObj.Message, tt.errMsg) {
					t.Errorf("error message = %q, want to contain %q", errObj.Message, tt.errMsg)
				}
			}
		})
	}
}

func TestWebSocketFunctions(t *testing.T) {
	// Test that all WebSocket functions are registered
	functions := []string{
		"websocket_server_chalu",
		"websocket_jukto",
		"websocket_pathao",
		"websocket_bondho",
	}

	for _, fn := range functions {
		t.Run("Function exists: "+fn, func(t *testing.T) {
			input := `dhoron(` + fn + `);`
			result := evalWS(input)

			if strObj, ok := result.(*object.String); ok {
				if strObj.Value != "builtin" && strObj.Value != "BUILTIN" {
					t.Errorf("%s type = %s, want builtin or BUILTIN", fn, strObj.Value)
				}
			} else {
				t.Errorf("expected string result for %s, got %T", fn, result)
			}
		})
	}
}

// Helper function to check if string contains substring
func wsContains(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

// Benchmark WebSocket operations
func BenchmarkWebSocketServerChalu(b *testing.B) {
	input := `websocket_server_chalu(4000, kaj(conn) {});`

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		evalWS(input)
	}
}

func BenchmarkWebSocketJukto(b *testing.B) {
	// Create a test WebSocket server
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			return
		}
		conn.Close()
	}))
	defer server.Close()

	wsURL := "ws" + strings.TrimPrefix(server.URL, "http")

	input := `
		proyash kaj connect() {
			dhoro ws = opekha websocket_jukto("` + wsURL + `");
			ferao ws;
		}
		connect();
	`

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		evalWS(input)
	}
}
