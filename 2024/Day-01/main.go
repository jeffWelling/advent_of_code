package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

// declare a type of a list of ints
// eg.
// inputList := [
//
//	[183, 1292, 129283, 23823],
//	[1928, 3298, 924, 0823]
//
// ]
type inputList [][]int

func parseInput() (error, *inputList) {
	input := make(inputList, 2)

	// Err if file input.txt doesn't exist
	_, err := os.Stat("input.txt")
	if os.IsNotExist(err) {
		return err, nil
	}

	// open file
	file, err := os.Open("input.txt")
	if err != nil {
		return err, nil
	}
	defer file.Close()

	// For each line in file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}

		part_0, err := strconv.Atoi(parts[0])
		if err != nil {
			return err, nil
		}
		part_1, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Println(err)
			return err, nil
		}

		input[0] = append(input[0], part_0)
		input[1] = append(input[1], part_1)
	}

	return nil, &input
}

func main() {
	err, input := parseInput()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	// Sort input
	sortInput(input)

	// Count the difference
	err, total := countDifference(input)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Println("Total difference: ", total)

	count := make(map[int]int)
	for _, num := range (*input)[1] {
		count[num]++
	}

	result := 0
	for _, num := range (*input)[0] {
		// If count contains num
		if count[num] > 0 {
			result += num * count[num]
		}
	}

	fmt.Println("Result: ", result)

}

// sortInput() takes a pointer to an inputList, and sorts each of the slices
// within, making the assumption that the values of the nested slices are
// integers.
func sortInput(input *inputList) error {
	sort.Ints((*input)[0])
	sort.Ints((*input)[1])
	return nil
}

// Count the difference
func countDifference(input *inputList) (error, int) {
	total := 0
	length_limit := len((*input)[0])
	for i := 0; i < length_limit; i++ {
		total += difference((*input)[0][i], (*input)[1][i])
	}
	return nil, total
}

// difference() takes two ints, and counts the ints between them
func difference(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
