package file

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFileInDir(t *testing.T) {
	// Create a temporary directory for testing
	tmpDir := t.TempDir()

	// Create some files in the temporary directory
	fileNames := []string{"testFile.txt", "anotherFile.txt", "someFile.txt"}
	for _, fileName := range fileNames {
		filePath := filepath.Join(tmpDir, fileName)
		_, err := os.Create(filePath)
		if err != nil {
			t.Fatalf("Failed to create test file: %v", err)
		}
	}

	// Test searching for a file that exists in the directory
	existingFileName := "testfile.txt"
	result := FileInDir(tmpDir, existingFileName)
	if result != "testFile.txt" {
		t.Errorf("Expected 'testFile.txt', got '%s'", result)
	}

	// Test searching for a file that does not exist in the directory
	nonExistingFileName := "nonexistent.txt"
	result = FileInDir(tmpDir, nonExistingFileName)
	if result != "" {
		t.Errorf("Expected empty string, got '%s'", result)
	}
}

func TestFileExtensionRemove(t *testing.T) {
	// Test a file name with extension
	fileName := "example.txt"
	result := FileExtensionRemove(fileName)
	expected := "example"
	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}

	// Test a file name without extension
	fileName = "example"
	result = FileExtensionRemove(fileName)
	if result != fileName {
		t.Errorf("Expected '%s', got '%s'", fileName, result)
	}
}

func TestFileSimplifyName(t *testing.T) {
	// Test a file name with dashes and underscores
	fileName := "file-with-dashes_and_underscores.txt"
	result := FileSimplifyName(fileName)
	expected := "filewithdashesandunderscores"
	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}

	// Test a file name without dashes and underscores
	fileName = "simplefilename.txt"
	result = FileSimplifyName(fileName)
	if result != fileName {
		t.Errorf("Expected '%s', got '%s'", fileName, result)
	}
}
