package main

import (
	"fmt"
	"strings"
	"szkandy/adventofcode-2024/shared"
)

func getRulesAndUpdates(printQueue []string) (rules [][2]int, updates [][]int) {
	rules = [][2]int{}
	updates = [][]int{}

	for _, line := range printQueue {
		if strings.Contains(line, "|") {
			parts := strings.Split(line, "|")
			rules = append(rules, [2]int{shared.ToInt(parts[0]), shared.ToInt(parts[1])})
		} else if strings.Contains(line, ",") {
			parts := strings.Split(line, ",")
			update := []int{}
			for _, part := range parts {
				update = append(update, shared.ToInt(part))
			}

			updates = append(updates, update)
		}
	}

	return
}

func isValid(i, j int, update []int, rules [][2]int) bool {
	for _, rule := range rules {
		if update[i] == rule[0] && update[j] == rule[1] && i > j {
			return false
		}

		if update[i] == rule[1] && update[j] == rule[0] && i < j {
			return false
		}
	}

	return true
}

func getUpdates(rules [][2]int, updates [][]int, valid bool) [][]int {
	result := [][]int{}

	for _, update := range updates {
		v := true

		for i := 0; i < len(update); i++ {
			for j := i + 1; j < len(update); j++ {
				if !isValid(i, j, update, rules) {
					v = false
					break
				}
			}
		}

		if v == valid {
			result = append(result, update)
		}
	}

	return result
}

func reorder(update []int, rules [][2]int) []int {
	reorder := make([]int, len(update))
	copy(reorder, update)

	for i := 0; i < len(update); i++ {
		for j := i + 1; j < len(update); j++ {
			if !isValid(i, j, reorder, rules) {
				reorder[i], reorder[j] = reorder[j], reorder[i]
			}
		}
	}

	return reorder
}

func pt1(rules [][2]int, updates [][]int) {
	sum := 0

	validUpdates := getUpdates(rules, updates, true)
	for _, update := range validUpdates {
		sum += update[len(update)/2]
	}

	fmt.Println("Part 1 Sum: ", sum)
}

func pt2(rules [][2]int, updates [][]int) {
	sum := 0

	invalidUpdates := getUpdates(rules, updates, false)
	for _, update := range invalidUpdates {
		reordered := reorder(update, rules)
		sum += reordered[len(reordered)/2]
	}

	fmt.Println("Part 2 Sum: ", sum)
}

func main() {
	printQueue := shared.LoadFileRows("./input.txt")
	rules, updates := getRulesAndUpdates(printQueue)

	pt1(rules, updates)
	pt2(rules, updates)
}
