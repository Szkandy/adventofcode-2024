package main

import (
	"fmt"
	"szkandy/adventofcode-2024/shared"
)

type Region struct {
	m     *shared.Matrix[string]
	Value string
	Plots []shared.Point
}

type CornerMask struct {
	SameRegion  []shared.Point
	OtherRegion []shared.Point
}

func (r *Region) IsInRegion(p shared.Point) bool {
	for _, plot := range r.Plots {
		if plot.X == p.X && plot.Y == p.Y {
			return true
		}
	}

	return false
}

func (r *Region) Print() {
	r.m.PrintWithHighlight(r.Plots, shared.Red)
}

func (r *Region) AddAdjacent(p shared.Point) {
	for _, plot := range p.Neighbors() {
		if r.IsInRegion(plot) || r.m.GetValue(plot) != r.Value {
			continue
		}

		r.Plots = append(r.Plots, plot)
		r.AddAdjacent(plot)
	}
}

func (r *Region) Perimeter() int {
	perimeter := 0
	for _, plot := range r.Plots {
		for _, neighbor := range plot.Neighbors() {
			if r.m.GetValue(neighbor) != r.Value {
				perimeter++
			}
		}
	}

	return perimeter
}

func (r *Region) GetCorners() (corners []shared.Point) {
	p := shared.Point{X: 0, Y: 0}
	masks := []CornerMask{
		{OtherRegion: []shared.Point{p.North(), p.West()}},
		{OtherRegion: []shared.Point{p.North(), p.East()}},
		{OtherRegion: []shared.Point{p.South(), p.West()}},
		{OtherRegion: []shared.Point{p.South(), p.East()}},
		{OtherRegion: []shared.Point{p.NorthWest()}, SameRegion: []shared.Point{p.North(), p.West()}},
		{OtherRegion: []shared.Point{p.NorthEast()}, SameRegion: []shared.Point{p.North(), p.East()}},
		{OtherRegion: []shared.Point{p.SouthWest()}, SameRegion: []shared.Point{p.South(), p.West()}},
		{OtherRegion: []shared.Point{p.SouthEast()}, SameRegion: []shared.Point{p.South(), p.East()}},
	}

	for _, plot := range r.Plots {
		for _, mask := range masks {
			sameValid, otherValid := true, true

			for _, neighbor := range mask.SameRegion {
				if !r.IsInRegion(plot.Add(neighbor)) {
					sameValid = false
					break
				}
			}

			for _, neighbor := range mask.OtherRegion {
				if r.IsInRegion(plot.Add(neighbor)) {
					otherValid = false
					break
				}
			}

			if sameValid && otherValid {
				corners = append(corners, plot)
			}
		}
	}

	return corners
}

func MapRegions(m *shared.Matrix[string]) []Region {
	regions := []Region{}
	for _, plot := range m.Coords {
		found := false
		for _, region := range regions {
			if region.IsInRegion(plot) {
				found = true
				break
			}
		}

		if found {
			continue
		}

		region := Region{
			m:     m,
			Value: m.GetValue(plot),
			Plots: []shared.Point{plot},
		}
		region.AddAdjacent(plot)
		regions = append(regions, region)
	}

	return regions
}

func pt1(gardensMap *shared.Matrix[string]) {
	regions := MapRegions(gardensMap)

	price := 0
	for _, region := range regions {
		price += region.Perimeter() * len(region.Plots)
	}

	fmt.Println("Part 1: ", price)
}

func pt2(gardensMap *shared.Matrix[string]) {
	regions := MapRegions(gardensMap)

	price := 0
	for _, region := range regions {
		corners := region.GetCorners()
		price += len(corners) * len(region.Plots)
	}

	fmt.Println("Part 2: ", price)
}

func main() {
	gardensMap := shared.LoadFileStringMatrixStruct("./input.txt")

	pt1(&gardensMap)
	pt2(&gardensMap)
}
