package shared

type Point struct {
	X int
	Y int
}

type Matrix struct {
	Rows   [][]string
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

func (m *Matrix) GetValueAt(x, y int) string {
	if x < 0 || y < 0 || y >= len(m.Rows) || x >= len(m.Rows[y]) {
		return ""
	}

	return m.Rows[y][x]
}

func (m *Matrix) GetValue(c Point) string {
	return m.GetValueAt(c.X, c.Y)
}

func (m *Matrix) SetValue(x, y int, value string) {
	m.Rows[y][x] = value
}

func (m *Matrix) FindValue(value string) (coords []Point) {
	for y, row := range m.Rows {
		for x, cell := range row {
			if cell == value {
				coords = append(coords, Point{X: x, Y: y})
			}
		}
	}

	return
}

func (m *Matrix) PrintAreaAroundCenter(center Point, centerChar string, size int) {
	for y := center.Y - size; y <= center.Y+size; y++ {
		for x := center.X - size; x <= center.X+size; x++ {
			if x == center.X && y == center.Y {
				print(centerChar)
			} else {
				print(m.GetValueAt(x, y))
			}
		}
		println()
	}
}

func (m *Matrix) Print() {
	for _, row := range m.Rows {
		for _, cell := range row {
			print(cell)
		}
		println()
	}
}

func (m *Matrix) ApplyPoints(points []Point, value string) {
	for _, p := range points {
		m.SetValue(p.X, p.Y, value)
	}
}
