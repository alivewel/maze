package lib

import (
	"fmt"
	"testing"
)

// Тесты для PathFinder
func TestPathFinder_Solve(t *testing.T) {
	maze := MazeWrapper{
		Vertical: []int{
			1, 1, 1, 1, 1,
			1, 1, 0, 1, 1,
			0, 1, 1, 1, 1,
			1, 0, 0, 1, 1,
			0, 0, 1, 0, 1}, // Инициализация вертикальных стен
		Horizontal: []int{
			0, 0, 0, 0, 0,
			0, 0, 0, 0, 0,
			1, 0, 0, 1, 0,
			0, 1, 0, 0, 0,
			1, 1, 1, 1, 1}, // Инициализация горизонтальных стен
		Rows: 5,
		Cols: 5,
	}

	tests := []struct {
		name   string
		maze   MazeWrapper
		from   Point
		to     Point
		expect []Point
	}{
		{
			name:   "Simple path",
			maze:   maze,
			from:   Point{X: 1, Y: 1},
			to:     Point{X: 4, Y: 4},
			expect: []Point{{1, 1}, {2, 1}, {3, 1}, {3, 2}, {4, 2}, {4, 3}, {4, 4}},
		},
		{
			name:   "Path from (1, 2) to (1, 3)",
			maze:   maze,
			from:   Point{X: 1, Y: 2},
			to:     Point{X: 1, Y: 3},
			expect: []Point{{1, 2}, {2, 2}, {3, 2}, {4, 2}, {4, 3}, {3, 3}, {2, 3}, {1, 3}},
		},
		{
			name:   "Path from (1, 4) to (5, 2)",
			maze:   maze,
			from:   Point{X: 1, Y: 4},
			to:     Point{X: 5, Y: 2},
			expect: []Point{{1, 4}, {2, 4}, {2, 3}, {3, 3}, {4, 3}, {5, 3}, {5, 2}},
		},
		{
			name:   "Path from (4, 2) to (2, 3)",
			maze:   maze,
			from:   Point{X: 4, Y: 2},
			to:     Point{X: 2, Y: 3},
			expect: []Point{{4, 2}, {4, 3}, {3, 3}, {2, 3}},
		},
		{
			name:   "Path from (2, 5) to (2, 3)",
			maze:   maze,
			from:   Point{X: 2, Y: 5},
			to:     Point{X: 2, Y: 3},
			expect: []Point{{2, 5}, {3, 5}, {4, 5}, {5, 5}, {5, 4}, {4, 4}, {4, 3}, {3, 3}, {2, 3}},
		},
		{
			name:   "Path from (5, 5) to (3, 3)",
			maze:   maze,
			from:   Point{X: 5, Y: 5},
			to:     Point{X: 3, Y: 3},
			expect: []Point{{5, 5}, {5, 4}, {4, 4}, {4, 3}, {3, 3}},
		},
		{
			name:   "Path from (3, 3) to (3, 3)",
			maze:   maze,
			from:   Point{X: 3, Y: 3},
			to:     Point{X: 3, Y: 3},
			expect: []Point{{3, 3}},
		},
		{
			name:   "Path from (1, 1) to (5, 5)",
			maze:   maze,
			from:   Point{X: 1, Y: 1},
			to:     Point{X: 5, Y: 5},
			expect: []Point{{1, 1}, {2, 1}, {3, 1}, {3, 2}, {4, 2}, {4, 3}, {4, 4}, {5, 4}, {5, 5}},
		},
		{
			name:   "Path from (5, 5) to (1, 1)",
			maze:   maze,
			from:   Point{X: 5, Y: 5},
			to:     Point{X: 1, Y: 1},
			expect: []Point{{5, 5}, {5, 4}, {4, 4}, {4, 3}, {4, 2}, {3, 2}, {3, 1}, {2, 1}, {1, 1}},
		},
		{
			name:   "Start point out of bounds",
			maze:   maze,
			from:   Point{X: 0, Y: 1},
			to:     Point{X: 4, Y: 4},
			expect: nil,
		},
		{
			name:   "End point out of bounds (4, 0)",
			maze:   maze,
			from:   Point{X: 1, Y: 1},
			to:     Point{X: 4, Y: 0},
			expect: nil,
		},
		{
			name:   "End point out of bounds (6, 0)",
			maze:   maze,
			from:   Point{X: 1, Y: 1},
			to:     Point{X: 6, Y: 0},
			expect: nil,
		},
		{
			name:   "End point out of bounds (6, 1)",
			maze:   maze,
			from:   Point{X: 1, Y: 1},
			to:     Point{X: 6, Y: 1},
			expect: nil,
		},
		{
			name:   "End point out of bounds (6, 6)",
			maze:   maze,
			from:   Point{X: 1, Y: 1},
			to:     Point{X: 6, Y: 6},
			expect: nil,
		},
	}

	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pf := NewPathFinder(tt.maze)
			path, err := pf.Solve(tt.maze, tt.from, tt.to)
			if err != nil {
				fmt.Println(err)
			}
			if i == 0 {
				pf.LengthMap.Print()
			}

			if len(path) != len(tt.expect) {
				t.Errorf("expected path length %d, got %d", len(tt.expect), len(path))
			}

			for i, p := range path {
				if p != tt.expect[i] {
					t.Errorf("expected point %v, got %v", tt.expect[i], p)
				}
			}
		})
	}
}
