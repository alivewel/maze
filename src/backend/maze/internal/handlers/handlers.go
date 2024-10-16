package handlers

import (
	"fmt"
	"log"
	"maze/internal/lib"
	"net/http"

	"github.com/gin-gonic/gin"
)

// MazeRequest представляет параметры для генерации лабиринта
type MazeRequest struct {
	Rows int `json:"rows"`
	Cols int `json:"cols"`
}

// SolveMazeRequest представляет параметры для решения лабиринта
type SolveMazeRequest struct {
	Maze lib.MazeWrapper `json:"maze"`
	From lib.Point       `json:"from"`
	To   lib.Point       `json:"to"`
}

func SolveMazeHandler(c *gin.Context) {
	log.Println("Запрос получен на /solve-maze")
	var req SolveMazeRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("Ошибка декодирования JSON запроса:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный запрос"})
		return
	}

	log.Println("Лабиринт получен:", req.Maze)
	log.Println("Начальная точка:", req.From)
	log.Println("Конечная точка:", req.To)

	pf := lib.NewPathFinder(req.Maze)

	path, err := pf.Solve(req.Maze, req.From, req.To)
	if err != nil {
		log.Println("Ошибка при решении лабиринта:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(path) == 0 {
		log.Println("Путь не найден")
		c.JSON(http.StatusNotFound, gin.H{"error": "Путь не найден"})
		return
	}

	log.Println("Путь найден:", path)
	c.JSON(http.StatusOK, gin.H{"path": path})
}

func GenerateMazeHandler(c *gin.Context) {
	log.Println("Запрос получен на /generate-maze")
	var req MazeRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("Ошибка декодирования JSON запроса:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный запрос"})
		return
	}

	log.Printf("Генерация лабиринта с %d строками и %d столбцами", req.Rows, req.Cols)

	maze, err := lib.Generate(lib.MazeGenerationSettings{
		Rows: req.Rows,
		Cols: req.Cols,
	})
	if err != nil {
		log.Println("Ошибка при генерации лабиринта:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fileName := fmt.Sprintf("maze_%dx%d.txt", req.Rows, req.Cols)

	if err := lib.SaveMazeToFile(maze, fileName); err != nil {
		log.Println("Ошибка сохранения лабиринта:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось сохранить лабиринт"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"vertical":   maze.Vertical,
		"horizontal": maze.Horizontal,
		"rows":       maze.Rows,
		"cols":       maze.Cols,
	})
}
