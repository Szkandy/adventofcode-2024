package shared

type Point struct {
	X int
	Y int
}

type Matrix[T string | int] struct {
	Rows   [][]T
	Coords []Point
}

func (c Point) Add(c2 Point) Point {
	return Point{
		X: c.X + c2.X,
		Y: c.Y + c2.Y,
	}
}

func (c Point) Subtract(c2 Point) Point {
	return Point{
		X: c.X - c2.X,
		Y: c.Y - c2.Y,
	}
}

func (c Point) Invert() Point {
	return Point{
		X: -c.X,
		Y: -c.Y,
	}
}

func (c Point) MultiplyScalar(scalar int) Point {
	return Point{
		X: c.X * scalar,
		Y: c.Y * scalar,
	}
}

func (c Point) North() Point {
	return Point{
		X: c.X,
		Y: c.Y - 1,
	}
}

func (c Point) South() Point {
	return Point{
		X: c.X,
		Y: c.Y + 1,
	}
}

func (c Point) East() Point {
	return Point{
		X: c.X + 1,
		Y: c.Y,
	}
}

func (c Point) West() Point {
	return Point{
		X: c.X - 1,
		Y: c.Y,
	}
}

func (c Point) NorthEast() Point {
	return c.North().East()
}

func (c Point) NorthWest() Point {
	return c.North().West()
}

func (c Point) SouthEast() Point {
	return c.South().East()
}

func (c Point) SouthWest() Point {
	return c.South().West()
}

func (c Point) Neighbors() []Point {
	return []Point{
		c.North(),
		c.South(),
		c.East(),
		c.West(),
	}
}
