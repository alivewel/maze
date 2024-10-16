package main

import (
	"fmt"
	"log"

	"maze/internal/lib"
)

func main() {
	// Генерируем матрицу согласно нашему MazeWrapper (имитация считывания из файла)
	// maze := lib.MazeWrapper{
	// 	Vertical: []int{
	// 		1, 1, 1, 1, 1,
	// 		1, 1, 0, 1, 1,
	// 		0, 1, 1, 1, 1,
	// 		1, 0, 0, 1, 1,
	// 		0, 0, 1, 0, 1}, // Инициализация вертикальных стен
	// 	Horizontal: []int{
	// 		0, 0, 0, 0, 0,
	// 		0, 0, 0, 0, 0,
	// 		1, 0, 0, 1, 0,
	// 		0, 1, 0, 0, 0,
	// 		1, 1, 1, 1, 1}, // Инициализация горизонтальных стен
	// 	Rows: 5,
	// 	Cols: 5,
	// }

	fileName := "../../internal/test_data/maze5x5.txt"
	fmt.Println("fileName", fileName)
	maze, err := lib.LoadMazeToFile(fileName)
	if err != nil {
		log.Println("Ошибка при чтении из файла", err)
	}
	fmt.Println("maze", maze)
	from := lib.Point{X: 1, Y: 2}
	to := lib.Point{X: 1, Y: 3}

	pf := lib.NewPathFinder(maze)
	path, err := pf.Solve(maze, from, to)
	if err != nil {
		log.Println("Ошибка при решении лабиринта:", err)
	}

	fmt.Println("Path:", path) // Решение лабиринта

	lib.PrintMaze(maze)                   // печать лабиринта без решения
	lib.PrintMazeWithSolution(maze, path) // печать лабиринта с решением

	// // Генерируем матрицу автоматически
	mazeGenerationSettings := lib.MazeGenerationSettings{Rows: 10, Cols: 10}

	maze2, err := lib.Generate(mazeGenerationSettings)
	if err != nil {
		fmt.Println(err)
	}

	lib.PrintMaze(maze2)

	from2 := lib.Point{X: 1, Y: 1}
	to2 := lib.Point{X: 10, Y: 10}

	pf2 := lib.NewPathFinder(maze2)
	path2, err := pf2.Solve(maze2, from2, to2)
	if err != nil {
		log.Println("Ошибка при решении лабиринта:", err)
	}

	lib.PrintMaze(maze2)                    // печать лабиринта без решения
	lib.PrintMazeWithSolution(maze2, path2) // печать лабиринта с решением
}
