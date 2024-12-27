package validator_test

import (
	"testing"

	"github.com/lukepaoloni/advent-of-code-2024/2/internal/validator"
)

func TestBasicLevelIsSafe(t *testing.T) {
	tests := []struct {
		name         string
		input        []int
		expectIsSafe bool
	}{
		{
			name:         "Safe because the levels are all decreasing by 1 or 2.",
			input:        []int{7, 6, 4, 2, 1},
			expectIsSafe: true,
		},
		{
			name:         "Unsafe because 2 7 is an increase of 5.",
			input:        []int{1, 2, 7, 8, 9},
			expectIsSafe: false,
		},
		{
			name:         "Unsafe because 6 2 is a decrease of 4.",
			input:        []int{9, 7, 6, 2, 1},
			expectIsSafe: false,
		},
		{
			name:         "Unsafe because 1 3 is increasing but 3 2 is decreasing.",
			input:        []int{1, 3, 2, 4, 5},
			expectIsSafe: false,
		},
		{
			name:         "Unsafe because 4 4 is neither an increase or a decrease.",
			input:        []int{8, 6, 4, 4, 1},
			expectIsSafe: false,
		},
		{
			name:         "Safe because the levels are all increasing by 1, 2, or 3.",
			input:        []int{1, 3, 6, 7, 9},
			expectIsSafe: true,
		},
	}

	basicLevelValidator := validator.BasicLevelValidator{}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			safe := basicLevelValidator.IsSafe(test.input)

			if test.expectIsSafe && !safe {
				t.Errorf("Expected levels to be safe for input: %v, but was deemed as unsafe", test.input)
			}

			if !test.expectIsSafe && safe {
				t.Errorf("Expected levels to be unsafe for input: %v, but was deemed as safe", test.input)
			}
		})
	}
}

func TestDampenedLevelIsSafe(t *testing.T) {
	tests := []struct {
		name         string
		input        []int
		expectIsSafe bool
	}{
		{
			name:         "Safe without removing any level",
			input:        []int{7, 6, 4, 2, 1},
			expectIsSafe: true,
		},
		{
			name:         "Unsafe regardless of which level is removed",
			input:        []int{1, 2, 7, 8, 9},
			expectIsSafe: false,
		},
		{
			name:         "Unsafe regardless of which level is removed",
			input:        []int{9, 7, 6, 2, 1},
			expectIsSafe: false,
		},
		{
			name:         "Safe by removing the second level, 3",
			input:        []int{1, 3, 2, 4, 5},
			expectIsSafe: true,
		},
		{
			name:         "Safe by removing the third level, 4",
			input:        []int{8, 6, 4, 4, 1},
			expectIsSafe: true,
		},
		{
			name:         "Safe by removing the fourth level, 4",
			input:        []int{1, 4, 7, 4, 10},
			expectIsSafe: true,
		},
		{
			name:         "Safe without removing any level",
			input:        []int{1, 3, 6, 7, 9},
			expectIsSafe: true,
		},
		{
			name:         "Safe without removing any level",
			input:        []int{7, 6, 4, 2, 1},
			expectIsSafe: true,
		},
		{
			name:         "Unsafe regardless of which level is removed",
			input:        []int{1, 2, 7, 8, 9},
			expectIsSafe: false,
		},
		{
			name:         "Unsafe regardless of which level is removed",
			input:        []int{9, 7, 6, 2, 1},
			expectIsSafe: false,
		},
		{
			name:         "Safe by removing the second level, 3",
			input:        []int{1, 3, 2, 4, 5},
			expectIsSafe: true,
		},
		{
			name:         "Safe by removing the third level, 4",
			input:        []int{8, 6, 4, 4, 1},
			expectIsSafe: true,
		},
		{
			name:         "Safe without removing any level",
			input:        []int{1, 3, 6, 7, 9},
			expectIsSafe: true,
		},
		{
			name:         "Safe by removing middle number in short sequence",
			input:        []int{1, 5, 3},
			expectIsSafe: true,
		},
		{
			name:         "Safe by removing first number",
			input:        []int{5, 1, 2, 3},
			expectIsSafe: true,
		},
		{
			name:         "Safe by removing last number",
			input:        []int{5, 4, 3, 1},
			expectIsSafe: true,
		},
		{
			name:         "Unsafe with multiple violations",
			input:        []int{1, 5, 2, 6, 3},
			expectIsSafe: false,
		},
		{
			name:         "Safe by removing middle number in alternating sequence",
			input:        []int{1, 2, 5, 4, 5},
			expectIsSafe: true,
		},
		{
			name:         "Edge case with two numbers",
			input:        []int{1, 5},
			expectIsSafe: false,
		},
	}

	dampenedLevelValidator := validator.DampenedLevelValidator{}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			safe := dampenedLevelValidator.IsSafe(test.input)

			if test.expectIsSafe && !safe {
				t.Errorf("\n%s:\nExpected levels to be safe for input: %v, but was deemed as unsafe", test.name, test.input)
			}

			if !test.expectIsSafe && safe {
				t.Errorf("\n%s:\nExpected levels to be unsafe for input: %v, but was deemed as safe", test.name, test.input)
			}
		})
	}
}
