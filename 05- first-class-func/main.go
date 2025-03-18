package main

import (
	"fmt"
	"math"
)

// Простая функция, которая выполняет сложение двух чисел
func add(a int, b int) int {
	return a + b
}

// Функция, возвращающая два значения (результат деления и остаток)
func divide(a, b int) (int, int) {
	return a / b, a % b
}

// Функция с именованными возвращаемыми значениями
func rectangleProperties(length, width float64) (area, perimeter float64) {
	area = length * width
	perimeter = 2 * (length + width)
	return
}

// Функция как значение
var multiply = func(a, b int) int {
	return a * b
}

// Функция, принимающая другую функцию в качестве аргумента
func applyOperation(a, b int, operation func(int, int) int) int {
	return operation(a, b)
}

// Функция, возвращающая другую функцию (замыкание)
func powerFunction(exponent float64) func(float64) float64 {
	return func(base float64) float64 {
		return math.Pow(base, exponent)
	}
}

// Функция высшего порядка: принимает функцию и применяет её к каждому элементу среза
func mapSlice(slice []int, operation func(int) int) []int {
	result := make([]int, len(slice))
	for i, v := range slice {
		result[i] = operation(v)
	}
	return result
}

func main() {
	// Использование функций как строительных блоков
	fmt.Println("Sum:", add(3, 5))

	// Использование функции с несколькими возвращаемыми значениями
	quotient, remainder := divide(10, 3)
	fmt.Println("Quotient:", quotient, "Remainder:", remainder)

	// Использование функции с именованными возвращаемыми значениями
	area, perimeter := rectangleProperties(5, 10)
	fmt.Println("Area:", area, "Perimeter:", perimeter)

	// Использование функции как значения
	fmt.Println("Multiply:", multiply(4, 5))

	// Передача функции как аргумента
	result := applyOperation(6, 2, add)
	fmt.Println("Apply Operation (Addition):", result)

	// Использование функции, которая возвращает другую функцию
	square := powerFunction(2)
	fmt.Println("Square of 4:", square(4))

	// Использование функции высшего порядка с mapSlice
	numbers := []int{1, 2, 3, 4, 5}
	squaredNumbers := mapSlice(numbers, func(x int) int { return x * x })
	fmt.Println("Squared Numbers:", squaredNumbers)
}
