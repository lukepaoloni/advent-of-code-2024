package parser_test

import (
	"testing"

	"github.com/lukepaoloni/advent-of-code-2024/2/internal/parser"
)

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
			levels, err := parser.ParseLevelsFromLine(test.input)

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
