package main

import "testing"

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
			_, err := FileReader{FilePath: test.filePath}.ReadLines()

			if test.expectError && err == nil {
				t.Errorf("Expected an error for file: %s, but got nil", test.filePath)
			}
			if !test.expectError && err != nil {
				t.Errorf("Did not expect an error for file: %s, but got: %v", test.filePath, err)
			}
		})
	}
}

func TestParseLevelsFromLine(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expectError bool
		expectedLen int
	}{
		{
			name:        "Valid numbers line",
			input:       "10 11 12 13 14 15",
			expectError: false,
			expectedLen: 6,
		},
		{
			name:        "Invalid characters line",
			input:       "a b c d e f",
			expectError: true,
			expectedLen: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			levels, err := ParseLevelsFromLine(test.input)

			if test.expectError && err == nil {
				t.Errorf("Expected an error for input: %s, but got nil", test.input)
			}
			if !test.expectError && err != nil {
				t.Errorf("Did not expect an error for input: %s, but got: %v", test.input, err)
			}
			if len(levels) != test.expectedLen {
				t.Errorf("Expected %d levels for input: %s, but got %d", test.expectedLen, test.input, len(levels))
			}
		})
	}
}
