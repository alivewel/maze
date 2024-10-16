package lib

import (
	"testing"
)

func TestLoadMazeToFile_MultipleFiles(t *testing.T) {
	// Определение тестовых случаев
	tests := []struct {
		fileName     string
		expectedMaze MazeWrapper
	}{
		{
			fileName: "../test_data/maze4x4.txt",
			expectedMaze: MazeWrapper{
				Vertical:   []int{0, 0, 0, 1, 1, 0, 1, 1, 0, 1, 0, 1, 0, 0, 0, 1},
				Horizontal: []int{1, 0, 1, 0, 0, 0, 1, 0, 1, 1, 0, 1, 1, 1, 1, 1},
				Rows:       4,
				Cols:       4,
			},
		},
		{
			fileName: "../test_data/maze5x5.txt",
			expectedMaze: MazeWrapper{
				Vertical:   []int{1, 1, 1, 1, 1, 1, 1, 0, 1, 1, 0, 1, 1, 1, 1, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1},
				Horizontal: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0, 0, 0, 1, 1, 1, 1, 1},
				Rows:       5,
				Cols:       5,
			},
		},
	}

	// Проход по каждому тестовому случаю
	for _, tt := range tests {
		t.Run(tt.fileName, func(t *testing.T) {
			maze, err := LoadMazeToFile(tt.fileName)
			if err != nil {
				t.Fatalf("Ошибка при загрузке лабиринта из файла %s: %v", tt.fileName, err)
			}

			// Проверка, что загруженный лабиринт соответствует ожидаемому
			if maze.Rows != tt.expectedMaze.Rows || maze.Cols != tt.expectedMaze.Cols {
				t.Errorf("Ожидались размеры лабиринта %dx%d, но получены %dx%d", tt.expectedMaze.Rows, tt.expectedMaze.Cols, maze.Rows, maze.Cols)
			}

			for i, v := range maze.Vertical {
				if v != tt.expectedMaze.Vertical[i] {
					t.Errorf("Ожидалась вертикальная стена на индексе %d равная %d, но получена %d", i, tt.expectedMaze.Vertical[i], v)
				}
			}

			for i, h := range maze.Horizontal {
				if h != tt.expectedMaze.Horizontal[i] {
					t.Errorf("Ожидалась горизонтальная стена на индексе %d равная %d, но получена %d", i, tt.expectedMaze.Horizontal[i], h)
				}
			}
		})
	}
}
