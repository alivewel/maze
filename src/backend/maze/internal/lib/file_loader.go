package lib

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func LoadMazeToFile(fileName string) (MazeWrapper, error) {
	content, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Ошибка при чтении файла:", err)
		return MazeWrapper{}, err
	}

	lines := strings.Split(string(content), "\n")

	// Первая строка - размеры лабиринта
	dimensions := strings.Fields(lines[0])
	if len(dimensions) != 2 {
		return MazeWrapper{}, fmt.Errorf("неверный формат размеров")
	}

	rows, err := strconv.Atoi(dimensions[0])
	if err != nil {
		return MazeWrapper{}, err
	}

	cols, err := strconv.Atoi(dimensions[1])
	if err != nil {
		return MazeWrapper{}, err
	}

	vertical := make([]int, 0, rows*cols)
	horizontal := make([]int, 0, rows*cols)

	// Остальные строки - данные вертикальных и горизонтальных линий
	for i := 1; i <= rows; i++ {
		values := strings.Fields(lines[i])
		if len(values) != cols {
			return MazeWrapper{}, fmt.Errorf("неверное количество значений в строке %d", i)
		}

		for _, v := range values {
			num, err := strconv.Atoi(v)
			if err != nil {
				return MazeWrapper{}, err
			}
			vertical = append(vertical, num)
		}
	}

	for i := rows + 2; i <= 2*rows+1; i++ {
		values := strings.Fields(lines[i])
		if len(values) != cols {
			return MazeWrapper{}, fmt.Errorf("неверное количество значений в строке %d", i)
		}

		for _, v := range values {
			num, err := strconv.Atoi(v)
			if err != nil {
				return MazeWrapper{}, err
			}
			horizontal = append(horizontal, num)
		}
	}

	return MazeWrapper{
		Vertical:   vertical,
		Horizontal: horizontal,
		Rows:       rows,
		Cols:       cols,
	}, nil
}
