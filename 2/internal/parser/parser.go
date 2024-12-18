package parser

import (
	"fmt"
	"strconv"
	"strings"
)

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
