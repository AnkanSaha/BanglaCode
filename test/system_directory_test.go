package test

import (
	"BanglaCode/src/object"
	"os"
	"testing"
)

func TestDirectoryTaliika(t *testing.T) {
	// Create test directory with files
	testDir := "/tmp/banglacode_test_list"
	os.MkdirAll(testDir, 0755)
	defer os.RemoveAll(testDir)

	// Create some test files
	os.WriteFile(testDir+"/file1.txt", []byte("test1"), 0644)
	os.WriteFile(testDir+"/file2.txt", []byte("test2"), 0644)
	os.Mkdir(testDir+"/subdir", 0755)

	evaluated := testEvalBuiltin(`directory_taliika("/tmp/banglacode_test_list")`)

	// Should return an array
	if evaluated.Type() != "ARRAY" {
		t.Errorf("Expected ARRAY, got %s", evaluated.Type())
		return
	}

	// Should have 3 items (2 files + 1 directory)
	arr := evaluated.(*object.Array)
	if len(arr.Elements) != 3 {
		t.Errorf("Expected 3 items, got %d", len(arr.Elements))
	}
}

func TestDirectoryGhumao(t *testing.T) {
	// Create nested directory structure
	testDir := "/tmp/banglacode_test_walk"
	os.MkdirAll(testDir+"/sub1/sub2", 0755)
	defer os.RemoveAll(testDir)

	os.WriteFile(testDir+"/root.txt", []byte("test"), 0644)
	os.WriteFile(testDir+"/sub1/file1.txt", []byte("test"), 0644)
	os.WriteFile(testDir+"/sub1/sub2/file2.txt", []byte("test"), 0644)

	evaluated := testEvalBuiltin(`directory_ghumao("/tmp/banglacode_test_walk")`)

	// Should return an array
	if evaluated.Type() != "ARRAY" {
		t.Errorf("Expected ARRAY, got %s", evaluated.Type())
		return
	}

	// Should include all files and directories
	arr := evaluated.(*object.Array)
	if len(arr.Elements) < 6 { // root dir + 2 subdirs + 3 files
		t.Errorf("Expected at least 6 items, got %d", len(arr.Elements))
	}
}

func TestDirectoryKhaliKi(t *testing.T) {
	// Create empty directory
	emptyDir := "/tmp/banglacode_test_empty"
	os.MkdirAll(emptyDir, 0755)
	defer os.RemoveAll(emptyDir)

	// Create non-empty directory
	fullDir := "/tmp/banglacode_test_full"
	os.MkdirAll(fullDir, 0755)
	defer os.RemoveAll(fullDir)
	os.WriteFile(fullDir+"/file.txt", []byte("test"), 0644)

	tests := []struct {
		input    string
		expected bool
	}{
		{`directory_khali_ki("/tmp/banglacode_test_empty")`, true},
		{`directory_khali_ki("/tmp/banglacode_test_full")`, false},
	}

	for _, tt := range tests {
		evaluated := testEvalBuiltin(tt.input)
		testBuiltinBooleanObject(t, evaluated, tt.expected)
	}
}

func TestDirectoryAkar(t *testing.T) {
	testDir := "/tmp/banglacode_test_size"
	os.MkdirAll(testDir, 0755)
	defer os.RemoveAll(testDir)

	// Create files with known sizes
	content1 := "Hello"  // 5 bytes
	content2 := "World!" // 6 bytes
	os.WriteFile(testDir+"/file1.txt", []byte(content1), 0644)
	os.WriteFile(testDir+"/file2.txt", []byte(content2), 0644)

	evaluated := testEvalBuiltin(`directory_akar("/tmp/banglacode_test_size")`)

	// Should return total size (11 bytes)
	if evaluated.Type() != "NUMBER" {
		t.Errorf("Expected NUMBER, got %s", evaluated.Type())
		return
	}

	num := evaluated.(*object.Number)
	expectedSize := float64(len(content1) + len(content2))
	if num.Value != expectedSize {
		t.Errorf("Expected size %f, got %f", expectedSize, num.Value)
	}
}
