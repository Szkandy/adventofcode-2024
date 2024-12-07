package main

import (
	"fmt"
	"math"
	"strconv"
	"szkandy/adventofcode-2024/shared"
)

func createMasks(length int, base int) []string {
	totalCombinations := int(math.Pow(float64(base), float64(length)))
	combinations := make([]string, totalCombinations)

	for i := 0; i < totalCombinations; i++ {
		mask := strconv.FormatInt(int64(i), base)

		for len(mask) < length {
			mask = "0" + mask
		}

		combinations[i] = mask
	}

	return combinations
}

func isCorrect(equation []int, mask string) bool {
	expected := equation[0]
	result := equation[1]

	for i := 2; i < len(equation); i++ {
		operator := shared.GetChar(mask, i-2)

		if operator == '0' {
			result += equation[i]
		} else if operator == '1' {
			result *= equation[i]
		} else {
			result = shared.ToInt(shared.ToString(result) + shared.ToString(equation[i]))
		}
	}

	return result == expected
}

func printEquation(equation []int, mask string) {

	for i, v := range equation {
		if i == 0 {
			fmt.Print(v, ": ")
			continue
		}

		if i > 1 {
			operator := shared.GetChar(mask, i-2)

			if operator == '0' {
				fmt.Print("+")
			} else if operator == '1' {
				fmt.Print("+")
			} else {
				fmt.Print("||")
			}
		}
		fmt.Print(v)
	}
	fmt.Println()
}

func testEquation(equation []int, base int) bool {
	masks := createMasks(len(equation)-2, base)
	for _, mask := range masks {
		if isCorrect(equation, mask) {
			printEquation(equation, mask)
			return true
		}
	}

	return false
}

func pt1(calibrations [][]int) {
	var sum int = 0
	count := 0
	for _, calibration := range calibrations {
		if testEquation(calibration, 2) {
			sum += calibration[0]
			count++
		}
	}

	fmt.Println("Count: ", count)
	fmt.Println("Part 1: ", sum)
}

func pt2(calibrations [][]int) {
	var sum int = 0
	count := 0
	for _, calibration := range calibrations {
		if testEquation(calibration, 3) {
			sum += calibration[0]
			count++
		}
	}

	fmt.Println("Count: ", count)
	fmt.Println("Part 2: ", sum)
}

func main() {
	calibrations := shared.LoadFileIntMatrix("./input.txt")
	pt1(calibrations)
	pt2(calibrations)
}
