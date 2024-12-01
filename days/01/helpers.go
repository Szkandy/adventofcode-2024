package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func LoadFile() (arr1 []int, arr2 []int) {
	file, err := os.ReadFile("./input.txt")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	lines := strings.Split(string(file), "\n")

	arr1 = []int{}
	arr2 = []int{}

	for _, line := range lines {
		if line == "" {
			continue
		}
		parts := strings.Split(line, "   ")
		a, _ := strconv.Atoi(parts[0])
		b, _ := strconv.Atoi(parts[1])

		arr1 = append(arr1, a)
		arr2 = append(arr2, b)
	}

	slices.Sort(arr1)
	slices.Sort(arr2)

	return
}
