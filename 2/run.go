package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type FileReader struct {
	FilePath string
}

func (reader FileReader) ReadLines() ([]string, error) {
	file, err := os.Open("./input.txt")

	if err != nil {
		return nil, fmt.Errorf("error opening %s: %w", reader.FilePath, err)
	}

	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	return lines, scanner.Err()
}

func main() {
	lines, err := FileReader{FilePath: "./input.txt"}.ReadLines()

	if err != nil {
		fmt.Println("Error processing safety reports:", err)
		return
	}
	safeReports := []string{}

	for _, line := range lines {
		levels := strings.Split(line, " ")
		safe := false

		firstLevel, err := strconv.Atoi(levels[0])

		if err != nil {
			fmt.Println("Error parsing first level as integer.")
			continue
		}

		secondLevel, err := strconv.Atoi(levels[1])

		if err != nil {
			fmt.Println("Error parsing second level as integer.")
			continue
		}

		diff := secondLevel - firstLevel
		allIncreasing := diff > 0

		if firstLevel > secondLevel {
			diff = firstLevel - secondLevel
		}
		safe = diff >= 1 && diff <= 3

		for i := 1; i < len(levels)-1 && safe; i++ {
			level, err := strconv.Atoi(levels[i])

			if err != nil {
				fmt.Println("Error parsing level as integer.")
				safe = false
				continue
			}

			nextLevel, err := strconv.Atoi(levels[i+1])

			if err != nil {
				fmt.Println("Error parsing next level as integer.")
				safe = false
				continue
			}

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

		if safe {
			safeReports = append(safeReports, line)
		}
	}

	fmt.Printf("Safe Reports: %d\n", len(safeReports))
}
