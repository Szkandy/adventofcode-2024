package main

import (
	"fmt"
)

func pt1() {
	arr1, arr2 := LoadFile()

	distance := 0
	for i := 0; i < len(arr1); i++ {
		d := arr1[i] - arr2[i]
		if d < 0 {
			d = -d
		}

		distance += d
	}

	fmt.Println("Distance: ", distance)
}

func pt2() {
	arr1, arr2 := LoadFile()

	occurences := make(map[int]int)
	for i := 0; i < len(arr2); i++ {
		occurences[arr2[i]]++
	}

	score := 0

	for i := 0; i < len(arr1); i++ {
		if occurences[arr1[i]] > 0 {
			score += arr1[i] * occurences[arr1[i]]
		}
	}

	fmt.Println("Similarity Score: ", score)
}

func main() {
	pt1()
	pt2()
}
