package lib

import (
	// "io/ioutil"
	// "io"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// Test function for SaveMazeToFile
func TestSaveMazeToFile(t *testing.T) {
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

	fileName := "test_maze.txt"

	// Call the function to save the maze to a file
	err := SaveMazeToFile(maze, fileName)
	if err != nil {
		t.Fatalf("Failed to save maze to file: %v", err)
	}

	// Check if the file was created
	tempDir, _ := filepath.Abs("temp_template")
	files, err := os.ReadDir(tempDir)
	if err != nil {
		t.Fatalf("Failed to read directory: %v", err)
	}

	var savedFilePath string
	for _, file := range files {
		if strings.HasPrefix(file.Name(), strings.TrimSuffix(fileName, ".txt")) {
			savedFilePath = filepath.Join(tempDir, file.Name())
			break
		}
	}

	if savedFilePath == "" {
		t.Fatalf("Maze file was not created")
	}

	// Clean up: remove the created file
	defer os.Remove(savedFilePath)

	// Optionally, read the file and verify its contents
	content, err := os.ReadFile(savedFilePath)
	if err != nil {
		t.Fatalf("Failed to read saved maze file: %v", err)
	}

	expectedContent := "5 5\n1 1 0 1 1 \n1 0 1 0 1 \n0 1 1 1 0 \n1 0 0 1 1 \n0 1 1 0 1 \n\n1 0 1 0 1 \n0 1 0 1 0 \n1 1 0 0 1 \n0 0 1 1 0 \n1 1 1 0 1 \n"
	if string(content) != expectedContent {
		t.Errorf("Maze file content does not match expected content.\nExpected:\n%s\nGot:\n%s", expectedContent, content)
	}
}
