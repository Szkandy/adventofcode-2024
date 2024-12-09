package main

import (
	"fmt"
	"szkandy/adventofcode-2024/shared"
)

func findAllAntennasGroups(m shared.Matrix) map[string][]shared.Point {
	antennas := make(map[string][]shared.Point)

	for y, row := range m.Rows {
		for x, cell := range row {
			if cell == "." {
				continue
			}

			antennas[cell] = append(antennas[cell], shared.Point{X: x, Y: y})
		}
	}

	return antennas
}

func getAntiNodes(m shared.Matrix, p1 shared.Point, p2 shared.Point, unlimitedDistance bool) []shared.Point {
	antiNodes := make([]shared.Point, 0)
	distance := p1.Subtract(p2)

	if !unlimitedDistance {
		distance = distance.MultiplyScalar(2)
	}

	for i := 1; ; i++ {
		p1Antinode := p1.Add(distance.Invert().MultiplyScalar(i))
		if m.GetValue(p1Antinode) != "" {
			antiNodes = append(antiNodes, p1Antinode)
		}

		if !unlimitedDistance || m.GetValue(p1Antinode) == "" {
			break
		}
	}

	for i := 1; ; i++ {
		p2Antinode := p2.Add(distance.MultiplyScalar(i))
		if m.GetValue(p2Antinode) != "" {
			antiNodes = append(antiNodes, p2Antinode)
		}

		if !unlimitedDistance || m.GetValue(p2Antinode) == "" {
			break
		}
	}

	return antiNodes
}

func getAntiNodesForGroup(m shared.Matrix, group []shared.Point, unlimitedDistance bool) []shared.Point {
	antiNodes := make([]shared.Point, 0)

	for i, p1 := range group {
		for j, p2 := range group {
			if i == j || i > j {
				continue
			}

			antiNodes = append(antiNodes, getAntiNodes(m, p1, p2, unlimitedDistance)...)
		}
	}

	return antiNodes
}

func getDistinctPoints(points []shared.Point) []shared.Point {
	distinctPoints := make([]shared.Point, 0)
	seenPoints := make(map[shared.Point]bool)

	for _, point := range points {
		if _, seen := seenPoints[point]; !seen {
			distinctPoints = append(distinctPoints, point)
			seenPoints[point] = true
		}
	}

	return distinctPoints
}

func pt1(m shared.Matrix) {
	groups := findAllAntennasGroups(m)

	antiNodes := make([]shared.Point, 0)
	for _, group := range groups {
		antiNodes = append(antiNodes, getAntiNodesForGroup(m, group, false)...)
	}

	antiNodes = getDistinctPoints(antiNodes)

	m.ApplyPoints(antiNodes, "#")
	m.Print()
	fmt.Println("Part 1: ", len(antiNodes))
}

func pt2(m shared.Matrix) {
	groups := findAllAntennasGroups(m)

	antiNodes := make([]shared.Point, 0)
	for _, group := range groups {
		antiNodes = append(antiNodes, getAntiNodesForGroup(m, group, true)...)
	}

	antiNodes = getDistinctPoints(antiNodes)

	m.ApplyPoints(antiNodes, "#")
	m.Print()
	fmt.Println("Part 2: ", len(antiNodes))
}

func main() {
	antennasMap := shared.LoadFileStringMatrixStruct("./input.txt")
	pt1(antennasMap)

	antennasMap = shared.LoadFileStringMatrixStruct("./input.txt")
	pt2(antennasMap)
}
