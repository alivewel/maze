package lib

// MazeWrapper представляет лабиринт с его границами и размерами
type MazeWrapper struct {
	Vertical   []int
	Horizontal []int
	Rows       int
	Cols       int
}

// IsGood проверяет, является ли лабиринт корректным
func (m *MazeWrapper) IsGood() bool {
	return len(m.Vertical) == m.Rows*m.Cols && len(m.Horizontal) == m.Rows*m.Cols
}

// At возвращает значение ячейки в вертикальной или горизонтальной линии
func (m *MazeWrapper) At(x, y int, vertical bool) int {
	index := x*m.Cols + y

	if index < 0 || x < 0 || y < 0 || x >= m.Rows || y >= m.Cols {
		return 1 // Считаем, что за пределами лабиринта есть стена
	}
	if vertical {
		return m.Vertical[index]
	}
	return m.Horizontal[index]
}
