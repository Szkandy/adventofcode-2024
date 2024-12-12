package main

import (
	"fmt"
	"szkandy/adventofcode-2024/shared"
)

func expand(diskMap []int) (expanded []int) {
	expanded = []int{}

	for i, block := range diskMap {
		for j := 0; j < block; j++ {
			val := i / 2
			if i%2 != 0 {
				val = -1
			}

			expanded = append(expanded, val)
		}
	}

	return
}

func print(diskMap []int) {
	for _, block := range diskMap {
		if block == -1 {
			fmt.Print(".")
		} else {
			fmt.Print(block)
		}
	}

	fmt.Println()
}

func findFreeBlock(diskMap []int, from int, minSize int) int {
	for i := from; i < len(diskMap); i++ {
		for j := i; j < len(diskMap); j++ {
			if diskMap[j] != -1 {
				break
			}

			if j-i+1 >= minSize {
				return i
			}
		}
	}

	return -1
}

func reorder(diskMap []int) []int {
	freeBlock := findFreeBlock(diskMap, 0, 1)

	for i := len(diskMap) - 1; i >= 0; i-- {
		if diskMap[i] == -1 {
			continue
		}

		if freeBlock >= i {
			break
		}

		diskMap[freeBlock] = diskMap[i]
		diskMap[i] = -1

		freeBlock = findFreeBlock(diskMap, freeBlock, 1)
		if freeBlock == -1 || freeBlock >= i {
			break
		}
	}

	return diskMap
}

func reorderWholeFiles(diskMap []int) []int {

	for i := len(diskMap) - 1; i >= 0; i-- {
		if diskMap[i] == -1 {
			continue
		}

		length := 1
		for diskMap[i] == diskMap[i-1] {
			length++
			i--
			if i == 0 {
				break
			}
		}

		freeBlock := findFreeBlock(diskMap, 0, length)
		if freeBlock == -1 || freeBlock >= i {
			continue
		}

		for j := 0; j < length; j++ {
			diskMap[freeBlock+j] = diskMap[i+j]
			diskMap[i+j] = -1
		}
	}

	return diskMap
}

func calculateChecksum(diskMap []int) (checksum int) {
	for i, block := range diskMap {
		if block == -1 {
			continue
		}

		checksum += i * block
	}

	return
}

func pt1(diskMap []int) {
	diskMap = reorder(diskMap)
	// print(diskMap)
	fmt.Println("Part 1 Checksum: ", calculateChecksum(diskMap))
}

func pt2(diskMap []int) {
	diskMap = reorderWholeFiles(diskMap)
	// print(diskMap)
	fmt.Println("Part 2 Checksum: ", calculateChecksum(diskMap))
}

func main() {
	diskMap := shared.LoadFileIntRow("./input.txt", "")
	diskMap = expand(diskMap)

	pt1(diskMap)

	diskMap = shared.LoadFileIntRow("./input.txt", "")
	diskMap = expand(diskMap)
	pt2(diskMap)
}
