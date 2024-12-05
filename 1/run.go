package main

import (
	"bufio"
	"fmt"
	"os"
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
	}
}
