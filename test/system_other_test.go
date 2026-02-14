package test

import (
	"BanglaCode/src/object"
	"os"
	"testing"
)

// Environment variable tests
func TestPoribesh(t *testing.T) {
	// Set a test environment variable
	os.Setenv("BANGLACODE_TEST_VAR", "test_value")
	defer os.Unsetenv("BANGLACODE_TEST_VAR")

	evaluated := testEvalBuiltin(`poribesh("BANGLACODE_TEST_VAR")`)
	testBuiltinStringObject(t, evaluated, "test_value")
}

func TestPoribeshSet(t *testing.T) {
	evaluated := testEvalBuiltin(`poribesh_set("BANGLACODE_TEST_SET", "new_value")`)

	// Should return NULL
	if evaluated.Type() != "NULL" {
		t.Errorf("Expected NULL, got %s", evaluated.Type())
	}

	// Verify environment variable was set
	value := os.Getenv("BANGLACODE_TEST_SET")
	if value != "new_value" {
		t.Errorf("Expected 'new_value', got '%s'", value)
	}
	os.Unsetenv("BANGLACODE_TEST_SET")
}

func TestPoribeshShokal(t *testing.T) {
	evaluated := testEvalBuiltin(`poribesh_shokal()`)

	// Should return an array
	if evaluated.Type() != "ARRAY" {
		t.Errorf("Expected ARRAY, got %s", evaluated.Type())
		return
	}

	// Should have at least some environment variables
	arr := evaluated.(*object.Array)
	if len(arr.Elements) < 1 {
		t.Error("Expected at least 1 environment variable")
	}
}

func TestPoribeshMuke(t *testing.T) {
	// Set then unset
	os.Setenv("BANGLACODE_TEST_UNSET", "value")
	evaluated := testEvalBuiltin(`poribesh_muke("BANGLACODE_TEST_UNSET")`)

	// Should return NULL
	if evaluated.Type() != "NULL" {
		t.Errorf("Expected NULL, got %s", evaluated.Type())
	}

	// Verify it was unset
	value := os.Getenv("BANGLACODE_TEST_UNSET")
	if value != "" {
		t.Error("Environment variable was not unset")
	}
}

// Path tests
func TestPathJoro(t *testing.T) {
	evaluated := testEvalBuiltin(`path_joro("/home", "user", "file.txt")`)

	// Should return a string
	if evaluated.Type() != "STRING" {
		t.Errorf("Expected STRING, got %s", evaluated.Type())
		return
	}

	str := evaluated.(*object.String)
	// Result should contain all parts
	if !contains(str.Value, "home") || !contains(str.Value, "user") || !contains(str.Value, "file.txt") {
		t.Errorf("Path joining failed: %s", str.Value)
	}
}

func TestPathNaam(t *testing.T) {
	evaluated := testEvalBuiltin(`path_naam("/home/user/file.txt")`)
	testBuiltinStringObject(t, evaluated, "file.txt")
}

func TestDirectoryNaam(t *testing.T) {
	evaluated := testEvalBuiltin(`directory_naam("/home/user/file.txt")`)

	if evaluated.Type() != "STRING" {
		t.Errorf("Expected STRING, got %s", evaluated.Type())
		return
	}

	str := evaluated.(*object.String)
	if !contains(str.Value, "user") {
		t.Errorf("Expected directory containing 'user', got '%s'", str.Value)
	}
}

func TestFileExt(t *testing.T) {
	evaluated := testEvalBuiltin(`file_ext("/home/user/file.txt")`)
	testBuiltinStringObject(t, evaluated, ".txt")
}

func TestPathMatch(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{`path_match("*.txt", "file.txt")`, true},
		{`path_match("*.txt", "file.pdf")`, false},
		{`path_match("test_*", "test_123")`, true},
	}

	for _, tt := range tests {
		evaluated := testEvalBuiltin(tt.input)
		testBuiltinBooleanObject(t, evaluated, tt.expected)
	}
}

// System info tests
func TestOsNaam(t *testing.T) {
	evaluated := testEvalBuiltin(`os_naam()`)

	// Should return a string (linux/darwin/windows)
	if evaluated.Type() != "STRING" {
		t.Errorf("Expected STRING, got %s", evaluated.Type())
		return
	}

	str := evaluated.(*object.String)
	validOS := []string{"linux", "darwin", "windows", "freebsd"}
	found := false
	for _, os := range validOS {
		if str.Value == os {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("Unexpected OS name: %s", str.Value)
	}
}

func TestBibhag(t *testing.T) {
	evaluated := testEvalBuiltin(`bibhag()`)

	// Should return a string (amd64/arm64/etc)
	if evaluated.Type() != "STRING" {
		t.Errorf("Expected STRING, got %s", evaluated.Type())
	}
}

func TestHostname(t *testing.T) {
	evaluated := testEvalBuiltin(`hostname()`)

	// Should return a non-empty string
	if evaluated.Type() != "STRING" {
		t.Errorf("Expected STRING, got %s", evaluated.Type())
		return
	}

	str := evaluated.(*object.String)
	if str.Value == "" {
		t.Error("Hostname should not be empty")
	}
}

func TestCpuSonkha(t *testing.T) {
	evaluated := testEvalBuiltin(`cpu_sonkha()`)

	// Should return a positive number
	if evaluated.Type() != "NUMBER" {
		t.Errorf("Expected NUMBER, got %s", evaluated.Type())
		return
	}

	num := evaluated.(*object.Number)
	if num.Value < 1 {
		t.Errorf("Expected at least 1 CPU, got %f", num.Value)
	}
}

func TestBebosthokNaam(t *testing.T) {
	evaluated := testEvalBuiltin(`bebosthok_naam()`)

	// Should return a non-empty string
	if evaluated.Type() != "STRING" {
		t.Errorf("Expected STRING, got %s", evaluated.Type())
	}
}

func TestBariDirectory(t *testing.T) {
	evaluated := testEvalBuiltin(`bari_directory()`)

	// Should return a non-empty string
	if evaluated.Type() != "STRING" {
		t.Errorf("Expected STRING, got %s", evaluated.Type())
		return
	}

	str := evaluated.(*object.String)
	if str.Value == "" {
		t.Error("Home directory should not be empty")
	}
}

// Temp file tests
func TestTempDirectory(t *testing.T) {
	evaluated := testEvalBuiltin(`temp_directory()`)

	// Should return a string
	if evaluated.Type() != "STRING" {
		t.Errorf("Expected STRING, got %s", evaluated.Type())
		return
	}

	str := evaluated.(*object.String)
	if str.Value == "" {
		t.Error("Temp directory should not be empty")
	}
}

func TestTempFile(t *testing.T) {
	evaluated := testEvalBuiltin(`temp_file("test-")`)

	// Should return a string (path to temp file)
	if evaluated.Type() != "STRING" {
		t.Errorf("Expected STRING, got %s", evaluated.Type())
		return
	}

	str := evaluated.(*object.String)

	// Verify file exists
	if _, err := os.Stat(str.Value); os.IsNotExist(err) {
		t.Errorf("Temp file was not created: %s", str.Value)
	}

	// Clean up
	os.Remove(str.Value)
}

func TestTempFolder(t *testing.T) {
	evaluated := testEvalBuiltin(`temp_folder("test-")`)

	// Should return a string (path to temp directory)
	if evaluated.Type() != "STRING" {
		t.Errorf("Expected STRING, got %s", evaluated.Type())
		return
	}

	str := evaluated.(*object.String)

	// Verify directory exists
	if info, err := os.Stat(str.Value); os.IsNotExist(err) || !info.IsDir() {
		t.Errorf("Temp directory was not created: %s", str.Value)
	}

	// Clean up
	os.RemoveAll(str.Value)
}

// Time tests
func TestShomoyEkhon(t *testing.T) {
	evaluated := testEvalBuiltin(`shomoy_ekhon()`)

	// Should return a number (Unix timestamp)
	if evaluated.Type() != "NUMBER" {
		t.Errorf("Expected NUMBER, got %s", evaluated.Type())
		return
	}

	num := evaluated.(*object.Number)
	// Should be a reasonable timestamp (after 2020)
	if num.Value < 1577836800 { // Jan 1, 2020
		t.Errorf("Timestamp seems incorrect: %f", num.Value)
	}
}

func TestTimezone(t *testing.T) {
	evaluated := testEvalBuiltin(`timezone()`)

	// Should return a string
	if evaluated.Type() != "STRING" {
		t.Errorf("Expected STRING, got %s", evaluated.Type())
	}
}

// Symlink tests
func TestSymlinkBanao(t *testing.T) {
	target := "/tmp/banglacode_symlink_target.txt"
	link := "/tmp/banglacode_symlink_link.txt"

	// Create target file
	os.WriteFile(target, []byte("test"), 0644)
	defer os.Remove(target)
	defer os.Remove(link)

	evaluated := testEvalBuiltin(`symlink_banao("/tmp/banglacode_symlink_target.txt", "/tmp/banglacode_symlink_link.txt")`)

	// Should return NULL
	if evaluated.Type() != "NULL" {
		t.Errorf("Expected NULL, got %s", evaluated.Type())
	}

	// Verify symlink was created
	info, err := os.Lstat(link)
	if err != nil {
		t.Errorf("Symlink was not created: %v", err)
	}
	if info.Mode()&os.ModeSymlink == 0 {
		t.Error("Created file is not a symlink")
	}
}

func TestSymlinkKi(t *testing.T) {
	target := "/tmp/banglacode_symlink_check_target.txt"
	link := "/tmp/banglacode_symlink_check_link.txt"

	os.WriteFile(target, []byte("test"), 0644)
	os.Symlink(target, link)
	defer os.Remove(target)
	defer os.Remove(link)

	tests := []struct {
		input    string
		expected bool
	}{
		{`symlink_ki("/tmp/banglacode_symlink_check_link.txt")`, true},
		{`symlink_ki("/tmp/banglacode_symlink_check_target.txt")`, false},
	}

	for _, tt := range tests {
		evaluated := testEvalBuiltin(tt.input)
		testBuiltinBooleanObject(t, evaluated, tt.expected)
	}
}

// Helper function
func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > len(substr) && findSubstring(s, substr))
}

func findSubstring(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
