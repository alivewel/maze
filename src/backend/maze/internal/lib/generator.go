package lib

import (
	"errors"
	"log"
	"math/rand"
	"time"
)

// MazeGenerationSettings содержит настройки для генерации лабиринта
type MazeGenerationSettings struct {
	Rows int
	Cols int
}

// isGood проверяет корректность данных для генерации лабиринта и возвращает ошибку, если что-то не так
func isGood(data []int, rows, cols int) (bool, error) {
	if rows == 0 || rows > 50 || cols == 0 || cols > 50 {
		return false, errors.New("некорректное количество строк или столбцов")
	}
	if rows*cols != len(data) {
		return false, errors.New("несоответствие количества элементов и размеров лабиринта")
	}

	for _, val := range data {
		if val != 0 && val != 1 {
			return false, errors.New("данные содержат недопустимые значения")
		}
	}

	return true, nil
}

// Generate генерирует лабиринт с использованием алгоритма Эллера
func Generate(s MazeGenerationSettings) (MazeWrapper, error) {
	if s.Rows < 0 || s.Cols < 0 {
		errStr := "отрицательные значения столбцов или строк"
		log.Println(errStr)
		return MazeWrapper{}, errors.New(errStr)
	}

	vertical := make([]int, s.Rows*s.Cols)
	horizontal := make([]int, s.Rows*s.Cols)

	// Проверка вертикальных стенок
	if valid, err := isGood(vertical, s.Rows, s.Cols); !valid {
		log.Printf("Некорректные значения вертикальных стенок: %v", err)
		return MazeWrapper{}, err
	}

	// Проверка горизонтальных стенок
	if valid, err := isGood(horizontal, s.Rows, s.Cols); !valid {
		log.Printf("Некорректные значения горизонтальных стенок: %v", err)
		return MazeWrapper{}, err
	}

	random := make([]int, s.Rows*s.Cols*3)
	// Создание локального генератора случайных чисел
	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))
	for i := range random {
		random[i] = r.Intn(2)
	}

	randIndex := 0
	counter := 1

	line := make([]int, s.Cols)

	merge := func(i int) {
		mergedItem := line[i+1]
		for col := range line {
			if line[col] == mergedItem {
				line[col] = line[i]
			}
		}
	}

	for row := 0; row < s.Rows-1; row++ {
		for i := range line {
			if line[i] == 0 {
				line[i] = counter
				counter++
			}
		}

		for col := 0; col < s.Cols-1; col++ {
			if line[col] == line[col+1] {
				vertical[row*s.Cols+col] = 1
			}
		}

		for col := 0; col < s.Cols-1; col++ {
			choice := random[randIndex]
			randIndex++
			if choice == 1 || line[col] == line[col+1] {
				vertical[row*s.Cols+col] = 1
			} else {
				merge(col)
			}
		}
		vertical[row*s.Cols+s.Cols-1] = 1

		for col := 0; col < s.Cols; col++ {
			choice := random[randIndex]
			randIndex++
			if choice == 1 {
				count := 0
				for c := 0; c < s.Cols; c++ {
					if line[c] == line[col] && horizontal[row*s.Cols+c] == 0 {
						count++
					}
				}
				if count != 1 {
					horizontal[row*s.Cols+col] = 1
				}
			}
		}

		for col := range line {
			if horizontal[row*s.Cols+col] == 1 {
				line[col] = 0
			}
		}
	}

	for i := range line {
		if line[i] == 0 {
			line[i] = counter
			counter++
		}
	}

	for col := 0; col < s.Cols-1; col++ {
		choice := random[randIndex]
		randIndex++
		if choice == 1 || line[col] == line[col+1] {
			vertical[(s.Rows-1)*s.Cols+col] = 1
		} else {
			merge(col)
		}
	}
	vertical[(s.Rows-1)*s.Cols+s.Cols-1] = 1

	for col := 0; col < s.Cols-1; col++ {
		if line[col] != line[col+1] {
			vertical[(s.Rows-1)*s.Cols+col] = 0
			merge(col)
		}
		horizontal[(s.Rows-1)*s.Cols+col] = 1
	}
	vertical[(s.Rows-1)*s.Cols+s.Cols-1] = 1

	return MazeWrapper{
		Vertical:   vertical,
		Horizontal: horizontal,
		Rows:       s.Rows,
		Cols:       s.Cols,
	}, nil
}
