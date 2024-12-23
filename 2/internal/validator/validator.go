package validator

import (
	"github.com/lukepaoloni/advent-of-code-2024/2/internal/parser"
	"github.com/lukepaoloni/advent-of-code-2024/2/internal/reader"
)

type SafeValidator interface {
	IsSafe(levels []int) bool
}

type ReportValidator struct {
	Reader    reader.FileReader
	Validator SafeValidator
}

func (reportValidator ReportValidator) SafeReports() ([]string, error) {
	lines, err := reportValidator.Reader.ReadLines()

	if err != nil {
		return nil, err
	}
	safeReports := []string{}

	for _, line := range lines {
		levels, err := parser.ParseLevelsFromLine(line)

		if err != nil {
			continue
		}

		if safe := reportValidator.Validator.IsSafe(levels); safe {
			safeReports = append(safeReports, line)
		}
	}

	return safeReports, nil
}

type BasicLevelValidator struct{}

func (validator BasicLevelValidator) IsSafe(levels []int) bool {
	safe := false

	if len(levels) <= 1 {
		return false
	}

	firstLevel := levels[0]
	secondLevel := levels[1]

	diff := secondLevel - firstLevel
	allIncreasing := diff > 0

	if firstLevel > secondLevel {
		diff = firstLevel - secondLevel
	}
	safe = diff >= 1 && diff <= 3

	for i := 1; i < len(levels)-1 && safe; i++ {
		level := levels[i]
		nextLevel := levels[i+1]

		diff = nextLevel - level

		increasing := diff > 0

		if !increasing && diff < 0 {
			diff = level - nextLevel
		}

		eitherAllIncreasingOrDecreasing := allIncreasing && increasing || !allIncreasing && !increasing

		if !eitherAllIncreasingOrDecreasing {
			safe = false
		}

		if diff < 1 || diff > 3 {
			safe = false
		}
	}

	return safe
}

type DampenedLevelValidator struct{}

func (validator DampenedLevelValidator) IsSafe(levels []int) bool {
	// First check if the sequence is already safe without dampening
	if isSequenceSafe(levels) {
		return true
	}

	// Try removing each number one at a time
	for i := range levels {
		// Create a new slice without the current number
		dampened := make([]int, 0, len(levels)-1)
		dampened = append(dampened, levels[:i]...)
		dampened = append(dampened, levels[i+1:]...)

		// Check if this dampened sequence is safe
		if isSequenceSafe(dampened) {
			return true
		}
	}

	return false
}

func isSequenceSafe(levels []int) bool {
	if len(levels) <= 1 {
		return false
	}

	// Determine if sequence should be increasing or decreasing
	increasing := levels[1] > levels[0]

	for i := 0; i < len(levels)-1; i++ {
		diff := levels[i+1] - levels[i]

		// Check if direction matches and difference is within bounds
		if (increasing && diff <= 0) || (!increasing && diff >= 0) {
			return false
		}

		absDiff := diff
		if diff < 0 {
			absDiff = -diff
		}

		if absDiff < 1 || absDiff > 3 {
			return false
		}
	}

	return true
}
