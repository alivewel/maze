package lib

import "fmt"

func PrintMaze(maze MazeWrapper) {
	fmt.Println("Maze visualization:")
	for i := 0; i < maze.Rows; i++ {
		// Верхние стены
		for j := 0; j < maze.Cols; j++ {
			fmt.Print("+")
			if i == 0 {
				fmt.Print("-----") // Верхняя граница
			} else {
				if maze.At(i-1, j, false) == 1 {
					fmt.Print("-----") // Горизонтальная стена
				} else {
					fmt.Print("     ")
				}
			}
		}
		fmt.Println("+")

		// Вертикальные стены и пространство между ними
		for j := 0; j < maze.Cols; j++ {
			if j == 0 || maze.At(i, j-1, true) == 1 {
				fmt.Print("|")
			} else {
				fmt.Print(" ")
			}
			fmt.Print("     ")
		}
		fmt.Println("|")
	}

	// Нижние стены
	for j := 0; j < maze.Cols; j++ {
		fmt.Print("+-----")
	}
	fmt.Println("+")
}

func PrintMazeWithNumbers(maze MazeWrapper) {
	fmt.Println("Maze visualization:")
	for i := 0; i < maze.Rows; i++ {
		// Верхние стены
		for j := 0; j < maze.Cols; j++ {
			fmt.Print("+")
			if i == 0 {
				fmt.Print("-----") // Верхняя граница
			} else {
				if maze.At(i-1, j, false) == 1 {
					fmt.Print("-----") // Горизонтальная стена
				} else {
					fmt.Printf("  %d  ", maze.At(i-1, j, false)) // Печать значения ячейки вместо стены
				}
			}
		}
		fmt.Println("+")

		// Вертикальные стены и пространство между ними
		for j := 0; j <= maze.Cols; j++ {
			if j == 0 || maze.At(i, j-1, true) == 1 {
				if j == 0 {
					fmt.Print("|")
				} else {
					fmt.Printf("%d", maze.At(i, j-1, true)) // Печать значения ячейки вместо стены
				}
			} else {
				fmt.Printf("%d", maze.At(i, j-1, true)) // Печать значения ячейки вместо стены
			}
			fmt.Print("     ")
		}
		fmt.Println()
	}

	// Нижние стены
	for j := 0; j < maze.Cols; j++ {
		fmt.Print("+")
		fmt.Printf("  %d  ", maze.At(maze.Rows-1, j, false)) // Печать значения ячейки вместо нижней стены
	}
	fmt.Println("+")
}

func PrintMazeWithSolution(maze MazeWrapper, path []Point) {
	fmt.Println("Maze visualization:")

	// Создаем карту для быстрого поиска, находится ли точка на пути
	pathMap := make(map[Point]bool)
	for _, p := range path {
		pathMap[p] = true
	}

	for i := 0; i < maze.Rows; i++ {
		// Верхние стены
		for j := 0; j < maze.Cols; j++ {
			fmt.Print("+")
			if i == 0 {
				fmt.Print("-----") // Верхняя граница
			} else {
				if maze.At(i-1, j, false) == 1 {
					fmt.Print("-----") // Горизонтальная стена
				} else {
					fmt.Print("     ")
				}
			}
		}
		fmt.Println("+")

		// Вертикальные стены и пространство между ними
		for j := 0; j < maze.Cols; j++ {
			if j == 0 || maze.At(i, j-1, true) == 1 {
				fmt.Print("|")
			} else {
				fmt.Print(" ")
			}

			// Проверяем, принадлежит ли текущая позиция пути
			if pathMap[Point{i + 1, j + 1}] {
				fmt.Print("  *  ") // Отображаем путь как '*'
			} else {
				fmt.Print("     ")
			}
		}
		fmt.Println("|")
	}

	// Нижние стены
	for j := 0; j < maze.Cols; j++ {
		fmt.Print("+-----")
	}
	fmt.Println("+")
}
