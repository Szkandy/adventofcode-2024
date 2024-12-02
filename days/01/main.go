package main

import (
	"fmt"
	"slices"
	"strings"
	"szkandy/adventofcode-2024/shared"
)

func pt1(arr1, arr2 []int) {
	distance := 0
	for i := 0; i < len(arr1); i++ {
		distance += shared.Abs(arr1[i] - arr2[i])
	}

	fmt.Println("Distance: ", distance)
}

func pt2(arr1, arr2 []int) {
	occurrences := make(map[int]int)
	for i := 0; i < len(arr2); i++ {
		occurrences[arr2[i]]++
	}

	score := 0

	for i := 0; i < len(arr1); i++ {
		if occurrences[arr1[i]] > 0 {
			score += arr1[i] * occurrences[arr1[i]]
		}
	}

	fmt.Println("Similarity Score: ", score)
}

func LoadFile() (arr1 []int, arr2 []int) {
	lines := shared.LoadFile("./input.txt")

	arr1 = []int{}
	arr2 = []int{}

	for _, line := range lines {
		if line == "" {
			continue
		}

		parts := strings.Split(line, "   ")

		arr1 = append(arr1, shared.ToInt(parts[0]))
		arr2 = append(arr2, shared.ToInt(parts[1]))
	}

	slices.Sort(arr1)
	slices.Sort(arr2)

	return
}

func main() {
	arr1, arr2 := LoadFile()

	pt1(arr1, arr2)
	pt2(arr1, arr2)
}
