package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")

	if err != nil {
		fmt.Println("Error opening input.txt: ", err)
	}

	group1, group2, err := parseFile(file)

	if err != nil {
		fmt.Println("Error parsing input.txt:", err)
	}

	puzzleOne(group1, group2)
	puzzleTwo(group1, group2)
}

func puzzleOne(group1 []int, group2 []int) {
	totalDistance := 0

	for index, locationIdFromGroup1 := range group1 {
		locationIdFromGroup2 := group2[index]

		if locationIdFromGroup1 > locationIdFromGroup2 {
			totalDistance += locationIdFromGroup1 - locationIdFromGroup2
		} else {
			totalDistance += locationIdFromGroup2 - locationIdFromGroup1
		}
	}

	fmt.Printf("Total distance: %d\n", totalDistance)
}

func puzzleTwo(group1 []int, group2 []int) {
	totalSimilarity := 0
	occurrencesInGroup2 := make(map[int]int)

	for _, locationIdFromGroup2 := range group2 {
		occurrencesInGroup2[locationIdFromGroup2]++
	}

	for _, locationIdFromGroup1 := range group1 {
		if similarities, exists := occurrencesInGroup2[locationIdFromGroup1]; exists {
			totalSimilarity += locationIdFromGroup1 * similarities
		}
	}

	fmt.Printf("Total similarity: %d\n", totalSimilarity)
}

func parseFile(file *os.File) ([]int, []int, error) {
	group1 := []int{}
	group2 := []int{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		locationIds := strings.Split(line, "   ")

		if len(locationIds) < 2 {
			return nil, nil, fmt.Errorf("expected 2 location IDs %q", line)
		}

		locationIdForGroup1, err := strconv.Atoi(locationIds[0])

		if err != nil {
			fmt.Println("Error converting locationId for group 1: ", err)
		}

		locationIdForGroup2, err := strconv.Atoi(locationIds[1])

		if err != nil {
			fmt.Println("Error converting locationId for group 2: ", err)
		}

		group1 = append(group1, locationIdForGroup1)
		group2 = append(group2, locationIdForGroup2)
	}

	slices.Sort(group1)
	slices.Sort(group2)

	return group1, group2, nil
}
