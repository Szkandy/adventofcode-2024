package main

import (
	"fmt"
	"szkandy/adventofcode-2024/shared"
)

type Letter struct {
	x int
	y int
}

type Word struct {
	letters []Letter
}

func getLetter(i, j int, puzzle [][]string) string {
	if i < 0 || j < 0 || i >= len(puzzle) || j >= len(puzzle[i]) {
		return ""
	}

	return puzzle[i][j]
}

func (w Word) isXmas(puzzle [][]string) bool {
	s := ""

	for i := range w.letters {
		s += getLetter(w.letters[i].x, w.letters[i].y, puzzle)
	}

	return s == "XMAS" || s == "SAMX"
}

func findWords(i, j int, puzzle [][]string) []Word {
	words := []Word{}

	horizontal := Word{letters: []Letter{}}
	vertical := Word{letters: []Letter{}}
	diagonalRight := Word{letters: []Letter{}}
	diagonalLeft := Word{letters: []Letter{}}

	for k := 0; k < 4; k++ {
		horizontal.letters = append(horizontal.letters, Letter{x: i, y: j + k})
		vertical.letters = append(vertical.letters, Letter{x: i + k, y: j})
		diagonalRight.letters = append(diagonalRight.letters, Letter{x: i + k, y: j + k})
		diagonalLeft.letters = append(diagonalLeft.letters, Letter{x: i + k, y: j - k})
	}

	if horizontal.isXmas(puzzle) {
		words = append(words, horizontal)
	}

	if vertical.isXmas(puzzle) {
		words = append(words, vertical)
	}

	if diagonalLeft.isXmas(puzzle) {
		words = append(words, diagonalLeft)
	}

	if diagonalRight.isXmas(puzzle) {
		words = append(words, diagonalRight)
	}

	return words
}

func isXMas(i, j int, puzzle [][]string) bool {
	upperLeft := getLetter(i-1, j-1, puzzle)
	upperRight := getLetter(i-1, j+1, puzzle)
	lowerLeft := getLetter(i+1, j-1, puzzle)
	lowerRight := getLetter(i+1, j+1, puzzle)
	center := getLetter(i, j, puzzle)

	return center == "A" &&
		((upperLeft == "M" && lowerRight == "S") || (upperLeft == "S" && lowerRight == "M")) &&
		((upperRight == "M" && lowerLeft == "S") || (upperRight == "S" && lowerLeft == "M"))
}

func pt1(puzzle [][]string) {
	words := []Word{}

	for i := range puzzle {
		for j := range puzzle[i] {
			words = append(words, findWords(i, j, puzzle)...)
		}
	}

	fmt.Println("Part 1 Count: ", len(words))
}

func pt2(puzzle [][]string) {
	count := 0

	for i := range puzzle {
		for j := range puzzle[i] {
			if isXMas(i, j, puzzle) {
				count++
			}
		}
	}

	fmt.Println("Part 2 Count: ", count)
}

func main() {
	puzzle := shared.LoadFileStringMatrix("./input.txt")

	pt1(puzzle)
	pt2(puzzle)
}
