package shared

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Magenta = "\033[35m"
var Cyan = "\033[36m"
var Gray = "\033[37m"
var White = "\033[97m"

func (m *Matrix[T]) GetValueAt(x, y int) T {
	if x < 0 || y < 0 || y >= len(m.Rows) || x >= len(m.Rows[y]) {
		var zero T
		switch any(zero).(type) {
		case string:
			return any("").(T)
		case int:
			return any(-1).(T)
		}
	}

	return m.Rows[y][x]
}

func (m *Matrix[T]) GetValue(c Point) T {
	return m.GetValueAt(c.X, c.Y)
}

func (m *Matrix[T]) SetValue(x, y int, value T) {
	m.Rows[y][x] = value
}

func (m *Matrix[T]) FindValue(value T) (coords []Point) {
	for y, row := range m.Rows {
		for x, cell := range row {
			if cell == value {
				coords = append(coords, Point{X: x, Y: y})
			}
		}
	}

	return
}

func (m *Matrix[T]) PrintAreaAroundCenter(center Point, centerChar string, size int) {
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

func (m *Matrix[T]) Print() {
	for _, row := range m.Rows {
		for _, cell := range row {
			print(cell)
		}
		println()
	}
}

func (m *Matrix[T]) ApplyPoints(points []Point, value T) {
	for _, p := range points {
		m.SetValue(p.X, p.Y, value)
	}
}

func (m *Matrix[T]) PrintWithHighlight(points []Point, color string) {
	for y, row := range m.Rows {
		for x, cell := range row {
			containsPoint := false
			for _, point := range points {
				if point.X == x && point.Y == y {
					containsPoint = true
					break
				}
			}

			if containsPoint {
				print(color)
				print(cell)
				print(Reset)
			} else {
				print(cell)
			}
		}
		println()
	}
}
