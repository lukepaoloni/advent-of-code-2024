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
	return isSequenceSafe(levels)
}

type DampenedLevelValidator struct{}

func (validator DampenedLevelValidator) IsSafe(levels []int) bool {
	if isSequenceSafe(levels) {
		return true
	}

	for i := range levels {
		dampened := []int{}
		dampened = append(dampened, levels[:i]...)
		dampened = append(dampened, levels[i+1:]...)

		if isSequenceSafe(dampened) {
			return true
		}
	}

	return false
}

func isSequenceSafe(levels []int) bool {
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
