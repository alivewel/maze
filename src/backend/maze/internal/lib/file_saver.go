package lib

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

// Функция для создания директории и сохранения файла с лабиринтом
func SaveMazeToFile(maze MazeWrapper, fileName string) error {
	tempDir, err := filepath.Abs("temp_template")
	if err != nil {
		return fmt.Errorf("error getting absolute path: %v", err)
	}

	log.Printf("Absolute path for temp_template: %s", tempDir)

	if _, err := os.Stat(tempDir); os.IsNotExist(err) {
		log.Println("Creating temp_template directory...")
		err = os.Mkdir(tempDir, os.ModePerm)
		if err != nil {
			return fmt.Errorf("error creating temp_template directory: %v", err)
		}
	}

	timestamp := time.Now().Format("20060102_150405")
	uniqueFileName := fmt.Sprintf("%s_%s.txt", fileName[:len(fileName)-4], timestamp)
	filePath := filepath.Join(tempDir, uniqueFileName)
	log.Printf("Creating file: %s", filePath)

	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("error creating maze file: %v", err)
	}
	defer file.Close()

	file.WriteString(fmt.Sprintf("%d %d\n", maze.Rows, maze.Cols))

	for i := 0; i < maze.Rows; i++ {
		for j := 0; j < maze.Cols; j++ {
			file.WriteString(fmt.Sprintf("%d ", maze.Vertical[i*maze.Cols+j]))
		}
		file.WriteString("\n")
	}

	file.WriteString("\n")

	for i := 0; i < maze.Rows; i++ {
		for j := 0; j < maze.Cols; j++ {
			file.WriteString(fmt.Sprintf("%d ", maze.Horizontal[i*maze.Cols+j]))
		}
		file.WriteString("\n")
	}

	log.Printf("Maze saved to %s", filePath)
	return nil
}
