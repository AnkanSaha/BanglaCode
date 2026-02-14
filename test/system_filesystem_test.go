package test

import (
	"os"
	"testing"
)

// Test file metadata functions
func TestFileAkar(t *testing.T) {
	// Create a temporary test file
	testFile := "/tmp/banglacode_test_file.txt"
	content := "Hello BanglaCode!"
	err := os.WriteFile(testFile, []byte(content), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	defer os.Remove(testFile)

	tests := []struct {
		input    string
		expected float64
	}{
		{`file_akar("/tmp/banglacode_test_file.txt")`, float64(len(content))},
	}

	for _, tt := range tests {
		evaluated := testEvalBuiltin(tt.input)
		testBuiltinNumberObject(t, evaluated, tt.expected)
	}
}

func TestFilePermission(t *testing.T) {
	// Create a temporary test file with specific permissions
	testFile := "/tmp/banglacode_test_perms.txt"
	err := os.WriteFile(testFile, []byte("test"), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	defer os.Remove(testFile)

	evaluated := testEvalBuiltin(`file_permission("/tmp/banglacode_test_perms.txt")`)
	testBuiltinStringObject(t, evaluated, "0644")
}

func TestFilePermissionSet(t *testing.T) {
	testFile := "/tmp/banglacode_test_chmod.txt"
	err := os.WriteFile(testFile, []byte("test"), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	defer os.Remove(testFile)

	// Change permissions
	evaluated := testEvalBuiltin(`file_permission_set("/tmp/banglacode_test_chmod.txt", "0755")`)
	if evaluated != nil && evaluated.Type() == "ERROR" {
		t.Errorf("Expected NULL, got error: %s", evaluated.Inspect())
	}

	// Verify permissions changed
	info, _ := os.Stat(testFile)
	mode := info.Mode().Perm()
	if mode != 0755 {
		t.Errorf("Expected permissions 0755, got %o", mode)
	}
}

func TestFileDhoron(t *testing.T) {
	// Create test file
	testFile := "/tmp/banglacode_test_type.txt"
	os.WriteFile(testFile, []byte("test"), 0644)
	defer os.Remove(testFile)

	// Create test directory
	testDir := "/tmp/banglacode_test_dir"
	os.Mkdir(testDir, 0755)
	defer os.RemoveAll(testDir)

	tests := []struct {
		input    string
		expected string
	}{
		{`file_dhoron("/tmp/banglacode_test_type.txt")`, "file"},
		{`file_dhoron("/tmp/banglacode_test_dir")`, "directory"},
	}

	for _, tt := range tests {
		evaluated := testEvalBuiltin(tt.input)
		testBuiltinStringObject(t, evaluated, tt.expected)
	}
}

func TestFileRename(t *testing.T) {
	oldPath := "/tmp/banglacode_old.txt"
	newPath := "/tmp/banglacode_new.txt"

	// Create original file
	os.WriteFile(oldPath, []byte("test"), 0644)
	defer os.Remove(newPath) // Clean up new path

	// Rename file
	evaluated := testEvalBuiltin(`file_rename("/tmp/banglacode_old.txt", "/tmp/banglacode_new.txt")`)
	if evaluated != nil && evaluated.Type() == "ERROR" {
		t.Errorf("Expected NULL, got error: %s", evaluated.Inspect())
	}

	// Verify old file doesn't exist
	if _, err := os.Stat(oldPath); !os.IsNotExist(err) {
		t.Error("Old file still exists after rename")
	}

	// Verify new file exists
	if _, err := os.Stat(newPath); os.IsNotExist(err) {
		t.Error("New file doesn't exist after rename")
	}
}

func TestFileShomoyPoribortito(t *testing.T) {
	testFile := "/tmp/banglacode_test_mtime.txt"
	os.WriteFile(testFile, []byte("test"), 0644)
	defer os.Remove(testFile)

	evaluated := testEvalBuiltin(`file_shomoy_poribortito("/tmp/banglacode_test_mtime.txt")`)

	// Should return a number (Unix timestamp)
	if evaluated.Type() != "NUMBER" {
		t.Errorf("Expected NUMBER, got %s", evaluated.Type())
	}
}

func TestAcheKi(t *testing.T) {
	testFile := "/tmp/banglacode_exists_test.txt"
	os.WriteFile(testFile, []byte("test"), 0644)
	defer os.Remove(testFile)

	tests := []struct {
		input    string
		expected bool
	}{
		{`ache_ki("/tmp/banglacode_exists_test.txt")`, true},
		{`ache_ki("/tmp/nonexistent_file_12345.txt")`, false},
	}

	for _, tt := range tests {
		evaluated := testEvalBuiltin(tt.input)
		testBuiltinBooleanObject(t, evaluated, tt.expected)
	}
}

func TestFolderBanao(t *testing.T) {
	testDir := "/tmp/banglacode_test_mkdir"
	defer os.RemoveAll(testDir)

	evaluated := testEvalBuiltin(`folder_banao("/tmp/banglacode_test_mkdir")`)
	if evaluated != nil && evaluated.Type() == "ERROR" {
		t.Errorf("Expected NULL, got error: %s", evaluated.Inspect())
	}

	// Verify directory exists
	if _, err := os.Stat(testDir); os.IsNotExist(err) {
		t.Error("Directory was not created")
	}
}

func TestFolderBanaoShokal(t *testing.T) {
	testDir := "/tmp/banglacode_test/nested/deep/dir"
	defer os.RemoveAll("/tmp/banglacode_test")

	evaluated := testEvalBuiltin(`folder_banao_shokal("/tmp/banglacode_test/nested/deep/dir")`)
	if evaluated != nil && evaluated.Type() == "ERROR" {
		t.Errorf("Expected NULL, got error: %s", evaluated.Inspect())
	}

	// Verify directory exists
	if _, err := os.Stat(testDir); os.IsNotExist(err) {
		t.Error("Nested directory was not created")
	}
}

func TestMukeFelo(t *testing.T) {
	testFile := "/tmp/banglacode_test_remove.txt"
	os.WriteFile(testFile, []byte("test"), 0644)

	evaluated := testEvalBuiltin(`muke_felo("/tmp/banglacode_test_remove.txt")`)
	if evaluated != nil && evaluated.Type() == "ERROR" {
		t.Errorf("Expected NULL, got error: %s", evaluated.Inspect())
	}

	// Verify file was removed
	if _, err := os.Stat(testFile); !os.IsNotExist(err) {
		t.Error("File still exists after removal")
	}
}
