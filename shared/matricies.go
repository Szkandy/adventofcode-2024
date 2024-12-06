package shared

type Coord struct {
	X int
	Y int
}

type Matrix struct {
	Rows   [][]string
	Coords []Coord
}

func (c Coord) Add(c2 Coord) Coord {
	return Coord{
		X: c.X + c2.X,
		Y: c.Y + c2.Y,
	}
}

func (m *Matrix) GetValueAt(x, y int) string {
	if x < 0 || y < 0 || y >= len(m.Rows) || x >= len(m.Rows[y]) {
		return ""
	}

	return m.Rows[y][x]
}

func (m *Matrix) GetValue(c Coord) string {
	return m.GetValueAt(c.X, c.Y)
}

func (m *Matrix) SetValue(x, y int, value string) {
	m.Rows[y][x] = value
}

func (m *Matrix) FindValue(value string) (coords []Coord) {
	for y, row := range m.Rows {
		for x, cell := range row {
			if cell == value {
				coords = append(coords, Coord{X: x, Y: y})
			}
		}
	}

	return
}

func (m *Matrix) PrintAreaAroundCenter(center Coord, centerChar string, size int) {
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
