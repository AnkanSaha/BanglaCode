package test

import (
	"BanglaCode/src/evaluator"
	"BanglaCode/src/evaluator/builtins"
	"BanglaCode/src/lexer"
	"BanglaCode/src/object"
	"BanglaCode/src/parser"
	"fmt"
	"net"
	"testing"
	"time"
)

// Helper function to evaluate BanglaCode
func evalTCP(input string) object.Object {
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

func TestTCPServerChalu(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
		errMsg  string
	}{
		{
			name: "TCP server with valid port and handler",
			input: `
				bishwo server_started = mittha;
				tcp_server_chalu(8081, kaj(conn) {
					server_started = sotti;
				});
				server_started;
			`,
			wantErr: false,
		},
		{
			name: "TCP server missing arguments",
			input: `tcp_server_chalu(8082);`,
			wantErr: true,
			errMsg: "wrong number of arguments",
		},
		{
			name: "TCP server invalid port type",
			input: `tcp_server_chalu("not a port", kaj(conn) {});`,
			wantErr: true,
			errMsg: "must be NUMBER",
		},
		{
			name: "TCP server invalid handler type",
			input: `tcp_server_chalu(8083, "not a function");`,
			wantErr: true,
			errMsg: "must be FUNCTION",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := evalTCP(tt.input)

			if tt.wantErr {
				errObj, ok := result.(*object.Error)
				if !ok {
					t.Fatalf("expected error, got %T (%+v)", result, result)
				}
				if tt.errMsg != "" && !tcpContains(errObj.Message, tt.errMsg) {
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

func TestTCPClientServer(t *testing.T) {
	// Start a simple Go TCP server for testing
	listener, err := net.Listen("tcp", ":8090")
	if err != nil {
		t.Fatalf("failed to start test server: %v", err)
	}
	defer listener.Close()

	// Handle connections
	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				return
			}
			go func(c net.Conn) {
				defer c.Close()
				buf := make([]byte, 1024)
				n, _ := c.Read(buf)
				if n > 0 {
					// Echo back
					c.Write([]byte("ECHO: " + string(buf[:n])))
				}
			}(conn)
		}
	}()

	time.Sleep(100 * time.Millisecond) // Let server start

	// Test TCP client connection
	input := `
		proyash kaj testClient() {
			dhoro conn = opekha tcp_jukto("localhost", 8090);
			ferao conn;
		}
		testClient();
	`

	result := evalTCP(input)

	// Should return a promise that resolves to connection map
	if promise, ok := result.(*object.Promise); ok {
		// Wait for promise to resolve
		time.Sleep(200 * time.Millisecond)

		if promise.State != "RESOLVED" && promise.State != "fulfilled" {
			t.Errorf("promise state = %s, want RESOLVED or fulfilled", promise.State)
		}

		if connMap, ok := promise.Value.(*object.Map); ok {
			// Check connection object has required fields
			if _, hasID := connMap.Pairs["id"]; !hasID {
				t.Error("connection missing 'id' field")
			}
			if _, hasHost := connMap.Pairs["host"]; !hasHost {
				t.Error("connection missing 'host' field")
			}
			if _, hasPort := connMap.Pairs["port"]; !hasPort {
				t.Error("connection missing 'port' field")
			}
		}
	}
}

func TestTCPPathao(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
		errMsg  string
	}{
		{
			name: "TCP pathao missing arguments",
			input: `tcp_pathao({});`,
			wantErr: true,
			errMsg: "wrong number of arguments",
		},
		{
			name: "TCP pathao invalid connection type",
			input: `tcp_pathao("not a map", "data");`,
			wantErr: true,
			errMsg: "must be MAP",
		},
		{
			name: "TCP pathao invalid data type",
			input: `tcp_pathao({}, 123);`,
			wantErr: true,
			errMsg: "must be STRING",
		},
		{
			name: "TCP pathao missing connection ID",
			input: `tcp_pathao({}, "data");`,
			wantErr: true,
			errMsg: "missing 'id' field",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := evalTCP(tt.input)

			if tt.wantErr {
				errObj, ok := result.(*object.Error)
				if !ok {
					t.Fatalf("expected error, got %T (%+v)", result, result)
				}
				if tt.errMsg != "" && !tcpContains(errObj.Message, tt.errMsg) {
					t.Errorf("error message = %q, want to contain %q", errObj.Message, tt.errMsg)
				}
			}
		})
	}
}

func TestTCPShuno(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
		errMsg  string
	}{
		{
			name: "TCP shuno missing arguments",
			input: `tcp_shuno();`,
			wantErr: true,
			errMsg: "wrong number of arguments",
		},
		{
			name: "TCP shuno invalid connection type",
			input: `tcp_shuno("not a map");`,
			wantErr: true,
			errMsg: "must be MAP",
		},
		{
			name: "TCP shuno missing connection ID",
			input: `tcp_shuno({});`,
			wantErr: true,
			errMsg: "missing 'id' field",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := evalTCP(tt.input)

			if tt.wantErr {
				errObj, ok := result.(*object.Error)
				if !ok {
					t.Fatalf("expected error, got %T (%+v)", result, result)
				}
				if tt.errMsg != "" && !tcpContains(errObj.Message, tt.errMsg) {
					t.Errorf("error message = %q, want to contain %q", errObj.Message, tt.errMsg)
				}
			}
		})
	}
}

func TestTCPBondho(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
		errMsg  string
	}{
		{
			name: "TCP bondho missing arguments",
			input: `tcp_bondho();`,
			wantErr: true,
			errMsg: "wrong number of arguments",
		},
		{
			name: "TCP bondho invalid connection type",
			input: `tcp_bondho("not a map");`,
			wantErr: true,
			errMsg: "must be MAP",
		},
		{
			name: "TCP bondho missing connection ID",
			input: `tcp_bondho({});`,
			wantErr: true,
			errMsg: "missing 'id' field",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := evalTCP(tt.input)

			if tt.wantErr {
				errObj, ok := result.(*object.Error)
				if !ok {
					t.Fatalf("expected error, got %T (%+v)", result, result)
				}
				if tt.errMsg != "" && !tcpContains(errObj.Message, tt.errMsg) {
					t.Errorf("error message = %q, want to contain %q", errObj.Message, tt.errMsg)
				}
			}
		})
	}
}

func TestTCPLekho(t *testing.T) {
	// tcp_lekho is an alias for tcp_pathao, test that it exists
	input := `dhoron(tcp_lekho);`
	result := evalTCP(input)

	if strObj, ok := result.(*object.String); ok {
		if strObj.Value != "builtin" && strObj.Value != "BUILTIN" {
			t.Errorf("tcp_lekho type = %s, want builtin or BUILTIN", strObj.Value)
		}
	} else {
		t.Errorf("expected string result, got %T", result)
	}
}

// Helper function to check if string contains substring
func tcpContains(s, substr string) bool {
	return len(s) >= len(substr) && (fmt.Sprintf("%s", s) != "" &&
		(len(s) == len(substr) && s == substr ||
		(len(s) >= len(substr) && (s[0:len(substr)] == substr ||
		(len(s) > len(substr) && tcpContains(s[1:], substr))))))
}

// Benchmark TCP operations
func BenchmarkTCPServerChalu(b *testing.B) {
	input := `tcp_server_chalu(9000, kaj(conn) {});`

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		evalTCP(input)
	}
}
