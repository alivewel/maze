package lib

import (
	"errors"
	"math"
)

// PathFinder представляет алгоритм поиска пути
type PathFinder struct {
	LengthMap  FieldWrapper
	Wave       []Point
	OldWave    []Point
	WaveStep   int
	EmptyValue int
}

// NewPathFinder создает новый PathFinder
// Rows - строки, Cols - столбцы
func NewPathFinder(maze MazeWrapper) *PathFinder {
	return &PathFinder{
		LengthMap:  NewFieldWrapper(maze.Rows, maze.Cols, math.MaxInt),
		Wave:       []Point{},
		OldWave:    []Point{},
		WaveStep:   0,
		EmptyValue: math.MaxInt,
	}
}

// Solve ищет путь от начальной точки до конечной в лабиринте
func (pf *PathFinder) Solve(maze MazeWrapper, from, to Point) ([]Point, error) {
	if valid, err := pf.IsValid(maze, from, to); !valid {
		return nil, err
	}

	// меняем местами точки нахождения пути
	reverPath := true
	if from.X > to.X || from.Y > to.Y {
		temp := from
		from = to
		to = temp
		reverPath = false
	}

	pf.InitializeStartState(maze, from)
	for len(pf.OldWave) > 0 {
		if pf.StepWave(maze, to) {
			break
		}
	}
	pf.LengthMap.Set(from.Y, from.X, 0)

	return pf.MakePath(maze, to, reverPath), nil
}

// InitializeStartState инициализирует начальное состояние для поиска пути
func (pf *PathFinder) InitializeStartState(maze MazeWrapper, from Point) {
	pf.Wave = nil
	pf.WaveStep = 0
	pf.OldWave = []Point{from}
	pf.LengthMap = NewFieldWrapper(maze.Rows, maze.Cols, pf.EmptyValue)
	pf.LengthMap.Set(from.X, from.Y, pf.WaveStep)
}

// IsValid проверяет, находятся ли точки внутри границ лабиринта и возвращает ошибку, если что-то не так
func (pf *PathFinder) IsValid(maze MazeWrapper, from, to Point) (bool, error) {
	// Проверяем, что лабиринт корректен
	if !maze.IsGood() {
		return false, errors.New("лабиринт некорректен")
	}
	// Проверяем, что начальная точка находится в допустимых пределах
	if from.X < 1 || from.X >= maze.Rows+1 || from.Y < 1 || from.Y >= maze.Cols+1 {
		return false, errors.New("начальная точка вне границ лабиринта")
	}
	// Проверяем, что конечная точка находится в допустимых пределах
	if to.X < 1 || to.X >= maze.Rows+1 || to.Y < 1 || to.Y >= maze.Cols+1 {
		return false, errors.New("конечная точка вне границ лабиринта")
	}
	return true, nil
}

// StepWave выполняет один шаг алгоритма BFS и возвращает true, если достигнута конечная точка
func (pf *PathFinder) StepWave(maze MazeWrapper, to Point) bool {
	pf.WaveStep++
	for _, p := range pf.OldWave {
		neighbors := []struct {
			x, y, value int
		}{
			{p.X + 1, p.Y, maze.At(p.X-1, p.Y-1, false)}, // Проверяем стену снизу
			{p.X - 1, p.Y, maze.At(p.X-2, p.Y-1, false)}, // Проверяем стену сверху
			{p.X, p.Y + 1, maze.At(p.X-1, p.Y-1, true)},  // Проверяем стену справа
			{p.X, p.Y - 1, maze.At(p.X-1, p.Y-2, true)},  // Проверяем стену слева
		}
		for _, n := range neighbors {
			// Проверка на допустимость координат
			if n.x > 0 && n.x <= maze.Rows && n.y > 0 && n.y <= maze.Cols {
				if n.value == 0 && pf.LengthMap.Get(n.y, n.x) == pf.EmptyValue {
					pf.Wave = append(pf.Wave, Point{n.x, n.y})
					pf.LengthMap.Set(n.y, n.x, pf.WaveStep)
					if n.x == to.X && n.y == to.Y {
						return true
					}
				}
			}
		}
	}

	pf.OldWave = pf.Wave
	pf.Wave = nil
	return false
}

// MakePath восстанавливает путь из конечной точки в начальную
func (pf *PathFinder) MakePath(maze MazeWrapper, to Point, reverPath bool) []Point {
	path := []Point{to}
	row, col := to.X, to.Y
	// pf.LengthMap.Print()
	for pf.LengthMap.Get(col, row) != 0 { // Используем (col, row)
		currentLen := pf.LengthMap.Get(col, row) // Используем (col, row)

		// Проверяем движение влево
		if col > 1 && pf.LengthMap.Get(col-1, row) == currentLen-1 && maze.At(row-1, col-2, true) == 0 {
			col-- // Движение влево
		} else if col < maze.Cols && pf.LengthMap.Get(col+1, row) == currentLen-1 && maze.At(row-1, col-1, true) == 0 {
			col++ // Движение вправо
		} else if row < maze.Rows && pf.LengthMap.Get(col, row+1) == currentLen-1 && maze.At(row-1, col-1, false) == 0 {
			row++ // Движение вниз
		} else if row > 1 && pf.LengthMap.Get(col, row-1) == currentLen-1 && maze.At(row-2, col-1, false) == 0 {
			row-- // Движение вверх
		} else {
			return nil // Если путь не найден, возвращаем nil
		}

		path = append(path, Point{row, col})
	}
	if reverPath {
		reversePath(path)
	}
	return path
}

// reversePath переворачивает путь
func reversePath(path []Point) {
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}
}
