package test

import (
	"BanglaCode/src/evaluator"
	"BanglaCode/src/evaluator/builtins"
	"BanglaCode/src/lexer"
	"BanglaCode/src/object"
	"BanglaCode/src/parser"
	"net"
	"testing"
	"time"
)

// Helper function to evaluate BanglaCode
func evalUDP(input string) object.Object {
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

func TestUDPServerChalu(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
		errMsg  string
	}{
		{
			name: "UDP server with valid port and handler",
			input: `
				bishwo udp_started = mittha;
				udp_server_chalu(9001, kaj(packet) {
					udp_started = sotti;
				});
				udp_started;
			`,
			wantErr: false,
		},
		{
			name:    "UDP server missing arguments",
			input:   `udp_server_chalu(9002);`,
			wantErr: true,
			errMsg:  "wrong number of arguments",
		},
		{
			name:    "UDP server invalid port type",
			input:   `udp_server_chalu("not a port", kaj(packet) {});`,
			wantErr: true,
			errMsg:  "must be NUMBER",
		},
		{
			name:    "UDP server invalid handler type",
			input:   `udp_server_chalu(9003, "not a function");`,
			wantErr: true,
			errMsg:  "must be FUNCTION",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := evalUDP(tt.input)

			if tt.wantErr {
				errObj, ok := result.(*object.Error)
				if !ok {
					t.Fatalf("expected error, got %T (%+v)", result, result)
				}
				if tt.errMsg != "" && !stringContains(errObj.Message, tt.errMsg) {
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

func TestUDPPathao(t *testing.T) {
	// Start a simple Go UDP server for testing
	addr := &net.UDPAddr{Port: 9010, IP: net.ParseIP("127.0.0.1")}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		t.Fatalf("failed to start test UDP server: %v", err)
	}
	defer conn.Close()

	// Handle packets
	go func() {
		buf := make([]byte, 1024)
		for {
			n, remoteAddr, err := conn.ReadFromUDP(buf)
			if err != nil {
				return
			}
			if n > 0 {
				// Echo back
				conn.WriteToUDP([]byte("RECEIVED: "+string(buf[:n])), remoteAddr)
			}
		}
	}()

	time.Sleep(100 * time.Millisecond) // Let server start

	tests := []struct {
		name    string
		input   string
		wantErr bool
		errMsg  string
	}{
		{
			name: "UDP pathao valid send",
			input: `
				proyash kaj send() {
					opekha udp_pathao("127.0.0.1", 9010, "test packet");
				}
				send();
			`,
			wantErr: false,
		},
		{
			name:    "UDP pathao missing arguments",
			input:   `udp_pathao("localhost", 9010);`,
			wantErr: true,
			errMsg:  "wrong number of arguments",
		},
		{
			name:    "UDP pathao invalid host type",
			input:   `udp_pathao(123, 9010, "data");`,
			wantErr: true,
			errMsg:  "must be STRING",
		},
		{
			name:    "UDP pathao invalid port type",
			input:   `udp_pathao("localhost", "port", "data");`,
			wantErr: true,
			errMsg:  "must be NUMBER",
		},
		{
			name:    "UDP pathao invalid data type",
			input:   `udp_pathao("localhost", 9010, 123);`,
			wantErr: true,
			errMsg:  "must be STRING",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := evalUDP(tt.input)

			if tt.wantErr {
				errObj, ok := result.(*object.Error)
				if !ok {
					// Might be a promise that rejects
					if promise, ok := result.(*object.Promise); ok {
						time.Sleep(100 * time.Millisecond)
						if promise.State == "rejected" {
							if errVal, ok := promise.Value.(*object.Error); ok {
								if tt.errMsg != "" && !stringContains(errVal.Message, tt.errMsg) {
									t.Errorf("error message = %q, want to contain %q", errVal.Message, tt.errMsg)
								}
								return
							}
						}
					}
					t.Fatalf("expected error, got %T (%+v)", result, result)
				}
				if tt.errMsg != "" && !stringContains(errObj.Message, tt.errMsg) {
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

func TestUDPUttor(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
		errMsg  string
	}{
		{
			name:    "UDP uttor missing arguments",
			input:   `udp_uttor({});`,
			wantErr: true,
			errMsg:  "wrong number of arguments",
		},
		{
			name:    "UDP uttor invalid packet type",
			input:   `udp_uttor("not a map", "data");`,
			wantErr: true,
			errMsg:  "must be MAP",
		},
		{
			name:    "UDP uttor invalid data type",
			input:   `udp_uttor({}, 123);`,
			wantErr: true,
			errMsg:  "must be STRING",
		},
		{
			name:    "UDP uttor missing packet ID",
			input:   `udp_uttor({}, "data");`,
			wantErr: true,
			errMsg:  "missing 'id' field",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := evalUDP(tt.input)

			if tt.wantErr {
				errObj, ok := result.(*object.Error)
				if !ok {
					t.Fatalf("expected error, got %T (%+v)", result, result)
				}
				if tt.errMsg != "" && !stringContains(errObj.Message, tt.errMsg) {
					t.Errorf("error message = %q, want to contain %q", errObj.Message, tt.errMsg)
				}
			}
		})
	}
}

func TestUDPShuno(t *testing.T) {
	// udp_shuno is an alias for udp_server_chalu, test that it exists
	input := `dhoron(udp_shuno);`
	result := evalUDP(input)

	if strObj, ok := result.(*object.String); ok {
		if strObj.Value != "builtin" && strObj.Value != "BUILTIN" {
			t.Errorf("udp_shuno type = %s, want builtin or BUILTIN", strObj.Value)
		}
	} else {
		t.Errorf("expected string result, got %T", result)
	}
}

func TestUDPBondho(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
		errMsg  string
	}{
		{
			name:    "UDP bondho missing arguments",
			input:   `udp_bondho();`,
			wantErr: true,
			errMsg:  "wrong number of arguments",
		},
		{
			name:    "UDP bondho invalid connection type",
			input:   `udp_bondho("not a map");`,
			wantErr: true,
			errMsg:  "must be MAP",
		},
		{
			name:    "UDP bondho missing connection ID",
			input:   `udp_bondho({});`,
			wantErr: true,
			errMsg:  "missing 'id' field",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := evalUDP(tt.input)

			if tt.wantErr {
				errObj, ok := result.(*object.Error)
				if !ok {
					t.Fatalf("expected error, got %T (%+v)", result, result)
				}
				if tt.errMsg != "" && !stringContains(errObj.Message, tt.errMsg) {
					t.Errorf("error message = %q, want to contain %q", errObj.Message, tt.errMsg)
				}
			}
		})
	}
}

// Helper function to check if string contains substring
func stringContains(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

// Benchmark UDP operations
func BenchmarkUDPServerChalu(b *testing.B) {
	input := `udp_server_chalu(9020, kaj(packet) {});`

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		evalUDP(input)
	}
}

func BenchmarkUDPPathao(b *testing.B) {
	// Start UDP server
	addr := &net.UDPAddr{Port: 9021, IP: net.ParseIP("127.0.0.1")}
	conn, _ := net.ListenUDP("udp", addr)
	defer conn.Close()

	go func() {
		buf := make([]byte, 1024)
		for {
			conn.ReadFromUDP(buf)
		}
	}()

	time.Sleep(50 * time.Millisecond)

	input := `
		proyash kaj send() {
			opekha udp_pathao("127.0.0.1", 9021, "benchmark packet");
		}
		send();
	`

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		evalUDP(input)
	}
}
