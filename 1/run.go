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
	group1 := []int{}
	group2 := []int{}

	file, err := os.Open("./input.txt")

	if err != nil {
		fmt.Println("Error opening input.txt: ", err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		locationIds := strings.Split(scanner.Text(), "   ")
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
