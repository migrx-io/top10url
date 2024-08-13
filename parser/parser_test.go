package parser

import (
	"os"
	"testing"
)

func TestParseFile(t *testing.T) {
	tests := []struct {
		name           string
		fileContent    string
		topN           int
		expectedResult []string
		expectError    bool
	}{
		{
			name: "Normal case",
			fileContent: `https://example.com/1 10
https://example.com/1 30
https://example.com/2 20
https://example.com/3 25`,
			topN: 3,
			expectedResult: []string{
				"https://example.com/1",
				"https://example.com/3",
				"https://example.com/2",
			},
			expectError: false,
		},
		{
			name:           "Empty file",
			fileContent:    ``,
			topN:           3,
			expectedResult: []string{},
			expectError:    false,
		},
		{
			name: "Invalid data",
			fileContent: `https://example.com/1 10
invalid line
https://example.com/2 abc`,
			topN:           2,
			expectedResult: []string{"https://example.com/1"},
			expectError:    false,
		},
		{
			name: "Fewer entries than topN",
			fileContent: `https://example.com/1 10
https://example.com/2 30`,
			topN: 5,
			expectedResult: []string{
				"https://example.com/2",
				"https://example.com/1",
			},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a temporary file with the specified content
			tmpfile, err := os.CreateTemp("", "testfile")
			if err != nil {
				t.Fatal(err)
			}
			defer os.Remove(tmpfile.Name())

			if _, err := tmpfile.WriteString(tt.fileContent); err != nil {
				t.Fatal(err)
			}
			if err := tmpfile.Close(); err != nil {
				t.Fatal(err)
			}

			// Call the ParseFile function
			result, err := ParseFile(tmpfile.Name(), tt.topN)

			if (err != nil) != tt.expectError {
				t.Fatalf("expected error: %v, got: %v", tt.expectError, err)
			}

			if len(result) != len(tt.expectedResult) {
				t.Fatalf("expected result length: %d, got: %d", len(tt.expectedResult), len(result))
			}

			for i, url := range result {
				if url != tt.expectedResult[i] {
					t.Errorf("expected URL at index %d: %s, got: %s", i, tt.expectedResult[i], url)
				}
			}
		})
	}
}
