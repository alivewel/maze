package lib

import (
	"fmt"
	"testing"
)

// TestGenerate проверяет функцию Generate
func TestGenerate(t *testing.T) {
	tests := []struct {
		name   string
		rows   int
		cols   int
		expect bool
	}{
		{"Valid small maze", 5, 5, true},
		{"Valid medium maze", 10, 10, true},
		{"Valid large maze", 50, 50, true},
		{"Invalid zero rows", 0, 5, false},
		{"Invalid zero cols", 5, 0, false},
		{"Invalid too many rows", 51, 5, false},
		{"Invalid too many cols", 5, 51, false},
		{"Negative rows", -5, 5, false},
		{"Negative cols", 5, -5, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			settings := MazeGenerationSettings{Rows: tt.rows, Cols: tt.cols}
			maze, err := Generate(settings)
			if err != nil {
				fmt.Println(err)
			}
			if (maze.Rows == 0 && maze.Cols == 0) != !tt.expect {
				t.Errorf("expected %v, got %v", tt.expect, maze.Rows != 0 && maze.Cols != 0)
			}
		})
	}
}
