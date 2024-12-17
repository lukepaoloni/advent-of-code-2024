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
	file, err := os.Open(reader.FilePath)

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

func ParseLevelsFromLine(line string) ([]int, error) {
	strLevels := strings.Split(line, " ")
	intLevels := make([]int, len(strLevels))
	for i, str := range strLevels {
		level, err := strconv.Atoi(str)
		if err != nil {
			return nil, fmt.Errorf("error parsing level %s as integer: %w", str, err)
		}
		intLevels[i] = level
	}
	return intLevels, nil
}

func main() {
	lines, err := FileReader{FilePath: "./input.txt"}.ReadLines()

	if err != nil {
		fmt.Println("Error processing safety reports:", err)
		return
	}
	safeReports := []string{}

	for i, line := range lines {
		levels, err := ParseLevelsFromLine(line)
		safe := false

		if err != nil {
			fmt.Println("Error parsing first level as integer.")
			continue
		}

		if len(levels) <= 1 {
			fmt.Printf("Line %d does not have enough levels reported.", i)
			return
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

		if safe {
			safeReports = append(safeReports, line)
		}
	}

	fmt.Printf("Safe Reports: %d\n", len(safeReports))
}
