package lib

import (
	"fmt"
	"testing"
)

// Test function for PrintMaze
func TestPrintMaze(t *testing.T) {
	maze := MazeWrapper{
		Vertical: []int{
			1, 1, 0, 1, 1,
			1, 0, 1, 0, 1,
			0, 1, 1, 1, 0,
			1, 0, 0, 1, 1,
			0, 1, 1, 0, 1}, // Vertical walls
		Horizontal: []int{
			1, 0, 1, 0, 1,
			0, 1, 0, 1, 0,
			1, 1, 0, 0, 1,
			0, 0, 1, 1, 0,
			1, 1, 1, 0, 1}, // Horizontal walls
		Rows: 5,
		Cols: 5,
	}

	fmt.Println("Testing PrintMaze function:")
	PrintMaze(maze)
}


func TestPrintMazeWithNumbers(t *testing.T) {
	maze := MazeWrapper{
		Vertical: []int{
			1, 1, 0, 1, 1,
			1, 0, 1, 0, 1,
			0, 1, 1, 1, 0,
			1, 0, 0, 1, 1,
			0, 1, 1, 0, 1}, // Vertical walls
		Horizontal: []int{
			1, 0, 1, 0, 1,
			0, 1, 0, 1, 0,
			1, 1, 0, 0, 1,
			0, 0, 1, 1, 0,
			1, 1, 1, 0, 1}, // Horizontal walls
		Rows: 5,
		Cols: 5,
	}

	fmt.Println("Testing PrintMazeWithNumbers function:")
	PrintMazeWithNumbers(maze)
}

func TestPrintMazeWithSolution(t *testing.T) {
	maze := MazeWrapper{
		Vertical: []int{
			1, 1, 0, 1, 1,
			1, 0, 1, 0, 1,
			0, 1, 1, 1, 0,
			1, 0, 0, 1, 1,
			0, 1, 1, 0, 1}, // Vertical walls
		Horizontal: []int{
			1, 0, 1, 0, 1,
			0, 1, 0, 1, 0,
			1, 1, 0, 0, 1,
			0, 0, 1, 1, 0,
			1, 1, 1, 0, 1}, // Horizontal walls
		Rows: 5,
		Cols: 5,
	}

	from := Point{X: 1, Y: 2}
	to := Point{X: 1, Y: 3}

	pf := NewPathFinder(maze)
	path, err := pf.Solve(maze, from, to)
	if err != nil {
		fmt.Println("Ошибка при решении лабиринта:", err)
	}


	fmt.Println("Testing PrintMazeWithSolution function:")
	PrintMazeWithSolution(maze, path)
}