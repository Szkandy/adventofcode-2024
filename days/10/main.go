package main

import (
	"fmt"
	"szkandy/adventofcode-2024/shared"
)

type Trail struct {
	m        *shared.Matrix[int]
	Start    shared.Point
	Position shared.Point
	End      shared.Point
	Path     []shared.Point
}

func newTrail(m *shared.Matrix[int], start shared.Point) Trail {
	return Trail{
		m:        m,
		Start:    start,
		Position: start,
		Path:     []shared.Point{start},
	}
}

func (t *Trail) Print() {
	fmt.Println("Start: ", t.Start, " End: ", t.End)
	t.m.PrintWithHighlight(t.Path, shared.Red)
}

func (t *Trail) copy(step shared.Point) Trail {
	newTrail := newTrail(t.m, t.Start)
	newTrail.Position = step

	copiedPath := make([]shared.Point, len(t.Path))
	copy(copiedPath, t.Path)
	newTrail.Path = append(copiedPath, step)

	return newTrail
}

func (t *Trail) GetPossibleSteps() []shared.Point {
	value := t.m.GetValue(t.Position)
	points := []shared.Point{}

	if t.m.GetValue(t.Position.North()) == value+1 {
		points = append(points, t.Position.North())
	}

	if t.m.GetValue(t.Position.South()) == value+1 {
		points = append(points, t.Position.South())
	}

	if t.m.GetValue(t.Position.East()) == value+1 {
		points = append(points, t.Position.East())
	}

	if t.m.GetValue(t.Position.West()) == value+1 {
		points = append(points, t.Position.West())
	}

	return points
}

func (t *Trail) FindPath() []Trail {
	possibleSteps := t.GetPossibleSteps()
	if len(possibleSteps) == 0 {
		return nil
	}

	trails := []Trail{}

	for _, step := range possibleSteps {
		newTrail := t.copy(step)

		if t.m.GetValue(step) == 9 {
			newTrail.End = step
			trails = append(trails, newTrail)
			continue
		}

		if path := newTrail.FindPath(); path != nil {
			trails = append(trails, path...)
		}
	}

	return trails
}

func getScore(trails []Trail) int {
	score := 0
	distinctStartToEnd := map[shared.Point][]shared.Point{}

	for _, trail := range trails {
		endIsThere := false
		for _, end := range distinctStartToEnd[trail.Start] {
			if end == trail.End {
				endIsThere = true
				break
			}
		}

		if !endIsThere {
			distinctStartToEnd[trail.Start] = append(distinctStartToEnd[trail.Start], trail.End)
			score++
		}
	}

	return score
}

func pt1(topoMap shared.Matrix[int]) {
	starts := topoMap.FindValue(0)

	trails := []Trail{}
	for _, start := range starts {

		trail := newTrail(&topoMap, start)

		if path := trail.FindPath(); path != nil {
			trails = append(trails, path...)
		}
	}

	fmt.Println("Part 1: ", getScore(trails))
}

func pt2(topoMap shared.Matrix[int]) {
	starts := topoMap.FindValue(0)

	trails := []Trail{}
	for _, start := range starts {

		trail := newTrail(&topoMap, start)

		if path := trail.FindPath(); path != nil {
			trails = append(trails, path...)
		}
	}

	fmt.Println("Part 2: ", len(trails))
}

func main() {
	topoMap := shared.LoadFileIntMatrixStruct("./input.txt")
	pt1(topoMap)
	pt2(topoMap)
}
