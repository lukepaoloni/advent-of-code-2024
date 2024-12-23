package reader_test

import (
	"testing"

	"github.com/lukepaoloni/advent-of-code-2024/2/internal/reader"
)

func TestFileReader_ReadLines(t *testing.T) {
	tests := []struct {
		name        string
		filePath    string
		expectError bool
	}{
		{
			name:        "File exists",
			filePath:    "./testdata/input.txt",
			expectError: false,
		},
		{
			name:        "File does not exist",
			filePath:    "./testdata/missing.txt",
			expectError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := reader.FileReader{FilePath: test.filePath}.ReadLines()

			if test.expectError && err == nil {
				t.Errorf("Expected an error for file: %s, but got nil", test.filePath)
			}
			if !test.expectError && err != nil {
				t.Errorf("Did not expect an error for file: %s, but got: %v", test.filePath, err)
			}
		})
	}
}
