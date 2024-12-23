package main

import (
	"fmt"

	"github.com/lukepaoloni/advent-of-code-2024/2/internal/reader"
	"github.com/lukepaoloni/advent-of-code-2024/2/internal/validator"
)

func main() {
	puzzleOne := validator.ReportValidator{
		Reader:    reader.FileReader{FilePath: "./input.txt"},
		Validator: validator.BasicLevelValidator{},
	}
	puzzleOneSafeReports, err := puzzleOne.SafeReports()

	if err != nil {
		fmt.Println("Error processing safety reports:", err)
		return
	}

	fmt.Printf("Puzzle One - Safe Reports: %d\n", len(puzzleOneSafeReports))

	puzzleTwo := validator.ReportValidator{
		Reader:    reader.FileReader{FilePath: "./input.txt"},
		Validator: validator.DampenedLevelValidator{},
	}
	puzzleTwoSafeReports, err := puzzleTwo.SafeReports()

	if err != nil {
		fmt.Println("Error processing safety reports:", err)
		return
	}

	fmt.Printf("Puzzle Two - Safe Reports: %d\n", len(puzzleTwoSafeReports))
}
