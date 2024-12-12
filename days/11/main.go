package main

import (
	"fmt"
	"szkandy/adventofcode-2024/shared"
)

type BlinkCache struct {
	results map[int]int
}

var cache = map[int]BlinkCache{}

func addToCache(stone int, nrSteps int, result int) {
	if _, ok := cache[stone]; !ok {
		cache[stone] = BlinkCache{
			results: map[int]int{},
		}
	}

	stoneCache := cache[stone]
	stoneCache.results[nrSteps] = result
}

func getFromCache(stone int, nrOfSteps int) int {
	if _, ok := cache[stone]; !ok {
		return 0
	}

	stoneCache := cache[stone]

	if len(stoneCache.results) == 0 {
		return 0
	}

	if res, ok := stoneCache.results[nrOfSteps]; ok {
		return res
	}

	return 0
}

func blink(stone int, step int, stopAt int) int {
	if step == stopAt {
		return 1
	}

	if cachedResult := getFromCache(stone, stopAt-step); cachedResult > 0 {
		return cachedResult
	}

	if stone == 0 {
		res := blink(1, step+1, stopAt)
		addToCache(stone, stopAt-step, res)
		return res
	}

	strStone := shared.ToString(stone)
	if len(strStone)%2 == 0 {
		firstHalf := shared.ToInt(strStone[:len(strStone)/2])
		secondHalf := shared.ToInt(strStone[len(strStone)/2:])

		res := blink(firstHalf, step+1, stopAt) + blink(secondHalf, step+1, stopAt)
		addToCache(stone, stopAt-step, res)
		return res
	}

	res := blink(stone*2024, step+1, stopAt)
	addToCache(stone, stopAt-step, res)
	return res
}

func pt1(stones []int) {
	score := 0

	for _, stone := range stones {
		score += blink(stone, 0, 25)
	}

	fmt.Println("Part 1: ", score)
}

func pt2(stones []int) {
	score := 0

	for _, stone := range stones {
		score += blink(stone, 0, 75)
	}

	fmt.Println("Part 2: ", score)
}

func main() {
	stones := shared.LoadFileIntRow("./input.txt", " ")

	pt1(stones)
	pt2(stones)
}
