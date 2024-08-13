package cmd

import (
	"os"
	"testing"
)

// TestParseArgs tests the ParseArgs function.
func TestParseArgs(t *testing.T) {

	// Test case: file and topN provided
	os.Args = []string{"cmd", "-file=testfile.txt", "-top=5"}
	filePath, topN := ParseArgs()
	if filePath != "testfile.txt" {
		t.Errorf("Expected filePath 'testfile.txt', but got '%s'", filePath)
	}
	if topN != 5 {
		t.Errorf("Expected topN 5, but got %d", topN)
	}

}
