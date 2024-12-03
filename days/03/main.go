package main

import (
	"fmt"
	"regexp"
	"sort"
	"szkandy/adventofcode-2024/shared"
)

type Mul struct {
	num1  int
	num2  int
	index int
}

type Cond struct {
	do    bool
	index int
}

func getMultiplications(memory string) []Mul {
	re := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	matches := re.FindAllStringIndex(memory, -1)

	muls := []Mul{}
	for _, match := range matches {
		numRe := regexp.MustCompile(`\d{1,3}`)
		nums := numRe.FindAllString(memory[match[0]:match[1]], -1)

		num1 := shared.ToInt(nums[0])
		num2 := shared.ToInt(nums[1])

		muls = append(muls, Mul{num1, num2, match[0]})
	}

	return muls
}

func getConditions(memory string) []Cond {
	re := regexp.MustCompile(`do\(\)`)
	matches := re.FindAllStringIndex(memory, -1)

	conds := []Cond{}
	for _, match := range matches {
		conds = append(conds, Cond{true, match[0]})
	}

	re = regexp.MustCompile(`don't\(\)`)
	matches = re.FindAllStringIndex(memory, -1)

	for _, match := range matches {
		conds = append(conds, Cond{false, match[0]})
	}

	sort.Slice(conds, func(i, j int) bool {
		return conds[i].index < conds[j].index
	})

	return conds
}

func getPreviousCondition(conds []Cond, index int) bool {
	for i := len(conds) - 1; i >= 0; i-- {
		if conds[i].index < index {
			return conds[i].do
		}
	}

	return true
}

func pt1(memory string) {
	sum := 0
	muls := getMultiplications(memory)

	for _, mul := range muls {
		sum += mul.num1 * mul.num2
	}

	fmt.Println("Part 1 Sum: ", sum)
}

func pt2(memory string) {
	muls := getMultiplications(memory)
	conds := getConditions(memory)
	sum := 0

	for _, mul := range muls {
		cond := getPreviousCondition(conds, mul.index)
		if cond {
			sum += mul.num1 * mul.num2
		}
	}

	fmt.Println("Part 2 Sum: ", sum)
}

func main() {
	memory := shared.LoadFile("./input.txt")

	pt1(memory)
	pt2(memory)
}
