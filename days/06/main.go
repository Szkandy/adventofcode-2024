package main

import (
	"fmt"
	"szkandy/adventofcode-2024/shared"
	"time"
)

type MatrixWalker struct {
	shared.Matrix
	Position   shared.Coord
	Direction  shared.Coord
	Path       []shared.Coord
	Directions []shared.Coord
}

func newMatrixWalker(guardMap shared.Matrix) MatrixWalker {
	position := guardMap.FindValue("^")[0]

	return MatrixWalker{
		Matrix:     guardMap,
		Position:   position,
		Direction:  shared.Coord{X: 0, Y: -1},
		Path:       []shared.Coord{position},
		Directions: []shared.Coord{{X: 0, Y: -1}},
	}
}

func (m *MatrixWalker) Reset() {
	m.Position = m.FindValue("^")[0]
	m.Direction = shared.Coord{X: 0, Y: -1}
	m.Path = []shared.Coord{m.Position}
	m.Directions = []shared.Coord{{X: 0, Y: -1}}
}

func (m *MatrixWalker) GetDistinctPath() []shared.Coord {
	distinctPath := []shared.Coord{}

	for _, coord := range m.Path {
		found := false

		for _, coord2 := range distinctPath {
			if coord == coord2 {
				found = true
				break
			}
		}

		if !found {
			distinctPath = append(distinctPath, coord)
		}
	}

	return distinctPath
}

func (m *MatrixWalker) WasVisited() bool {
	for i, coord := range m.Path {
		if coord == m.Position && m.Directions[i] == m.Direction && i != len(m.Path)-1 {
			return true
		}
	}

	return false
}

func (m *MatrixWalker) GetCharByDirection() string {
	switch m.Direction {
	case shared.Coord{X: 0, Y: -1}:
		return "^"
	case shared.Coord{X: 1, Y: 0}:
		return ">"
	case shared.Coord{X: 0, Y: 1}:
		return "v"
	case shared.Coord{X: -1, Y: 0}:
		return "<"
	}

	return "X"
}

func (m *MatrixWalker) RotateRight() {
	switch m.Direction {
	case shared.Coord{X: 0, Y: -1}:
		m.Direction = shared.Coord{X: 1, Y: 0}
	case shared.Coord{X: 1, Y: 0}:
		m.Direction = shared.Coord{X: 0, Y: 1}
	case shared.Coord{X: 0, Y: 1}:
		m.Direction = shared.Coord{X: -1, Y: 0}
	case shared.Coord{X: -1, Y: 0}:
		m.Direction = shared.Coord{X: 0, Y: -1}
	}
}

func (m *MatrixWalker) AddToPath() {
	m.Path = append(m.Path, m.Position)
	m.Directions = append(m.Directions, m.Direction)
}

func (m *MatrixWalker) MakeStep() {
	nextPosition := m.Position.Add(m.Direction)

	if m.GetValue(nextPosition) == "" {
		m.Position = nextPosition
		return
	}

	for m.GetValue(nextPosition) == "#" {
		m.RotateRight()
		nextPosition = m.Position.Add(m.Direction)
	}

	m.Position = nextPosition
	m.AddToPath()
}

func (m *MatrixWalker) Walk(debug bool) (isLoop bool) {
	for {
		m.MakeStep()

		if debug {
			fmt.Print("\033[H\033[2J")
			m.PrintAreaAroundCenter(m.Position, m.GetCharByDirection(), 10)
			time.Sleep(10 * time.Millisecond)
		}

		if m.WasVisited() {
			return true
		}

		if m.GetValue(m.Position) == "" {
			break
		}
	}

	return false
}

func pt1(walker MatrixWalker, debug bool) {
	walker.Walk(debug)
	fmt.Println("Part 1: ", len(walker.GetDistinctPath()))
}

func pt2(walker MatrixWalker) {
	walker.Walk(false)
	originalPath := walker.GetDistinctPath()

	loops := 0

	for i := range originalPath {
		if i == 0 {
			continue
		}

		fmt.Print("#", i, " Setting obstacle at ", originalPath[i])
		walker.Reset()
		walker.SetValue(originalPath[i].X, originalPath[i].Y, "#")
		isLoop := walker.Walk(false)

		if isLoop {
			loops++
			fmt.Print(" - Loop detected")
		}

		fmt.Println()
		walker.SetValue(originalPath[i].X, originalPath[i].Y, ".")
	}

	fmt.Println("Part 2: ", loops)
}

func main() {
	guardMap := shared.LoadFileStringMatrixStruct("./input.txt")
	walker := newMatrixWalker(guardMap)

	pt1(walker, false)
	pt2(walker)
}
