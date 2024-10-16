package lib

import (
	"fmt"
	"log"
	"math"
)

// Point представляет координаты точки в лабиринте
type Point struct {
	X, Y int
}

// FieldWrapper представляет матрицу для хранения длины пути
type FieldWrapper struct {
	Data       [][]int
	Rows       int
	Cols       int
	EmptyValue int
}

// NewFieldWrapper создает новый FieldWrapper
func NewFieldWrapper(rows, cols, emptyValue int) FieldWrapper {
	// Создаем срез длиной rows*cols и заполняем его значениями emptyValue
	data := make([][]int, rows)
	for i := range data {
		data[i] = make([]int, cols) // Инициализируем каждый подмассив
		for j := range data[i] {
			data[i][j] = emptyValue // Заполняем значениями emptyValue
		}
	}

	return FieldWrapper{
		Data:       data,
		Rows:       rows,
		Cols:       cols,
		EmptyValue: emptyValue,
	}
}

// Get возвращает значение из матрицы по координатам
func (cw *FieldWrapper) Get(x, y int) int {
	// Проверяем, что индексы находятся в допустимых пределах
	if x < 1 || y < 1 || y > cw.Rows || x > cw.Cols {
		// Обработка ошибки
		fmt.Printf("Get: индекс вне диапазона! x: %d, y: %d\n", x, y)
		return math.MaxInt
	}
	return cw.Data[y-1][x-1] // сдвиг индексов на -1
}

// Set устанавливает значение в матрице по координатам
func (cw *FieldWrapper) Set(x, y, value int) {
	// Проверяем, что индексы находятся в допустимых пределах
	if x < 1 || y < 1 || y > cw.Rows || x > cw.Cols {
		// Обработка ошибки
		log.Printf("Set: индекс вне диапазона! x: %d, y: %d (Rows: %d, Cols: %d)\n", x, y, cw.Rows, cw.Cols)
		return
	}
	cw.Data[y-1][x-1] = value // сдвиг индексов на -1
}

func (cw *FieldWrapper) Print() {
	fmt.Println("FieldWrapper visualization:", cw.Rows, cw.Cols)

	for i := 1; i <= cw.Rows; i++ { // Итерация по строкам
		for j := 1; j <= cw.Cols; j++ { // Итерация по столбцам
			ind := cw.Get(j, i) // Получаем значение с учетом смещения
			if ind == cw.EmptyValue {
				fmt.Printf("%4d ", 0) // Печать значения элемента
			} else {
				fmt.Printf("%4d ", ind) // Печать значения элемента
			}
		}
		fmt.Println() // Переход на новую строку после печати строки
	}
}
