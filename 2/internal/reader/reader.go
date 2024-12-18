package reader

import (
	"bufio"
	"fmt"
	"os"
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
