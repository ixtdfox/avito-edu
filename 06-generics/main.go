package main

import (
	"fmt"
)

// Обобщенная функция для работы с любым типом данных
func PrintValue[T any](value T) {
	fmt.Println("Value:", value)
}

// Обобщенная функция для поиска максимального значения
func Max[T int | float64](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// Обобщенная структура
type Box[T any] struct {
	content T
}

// Метод структуры, работающий с дженериками
func (b Box[T]) GetContent() T {
	return b.content
}

func main() {
	// Использование обобщенной функции
	PrintValue(42)
	PrintValue("Hello, Generics!")

	// Использование функции Max
	fmt.Println("Max of 10 and 20:", Max(10, 20))
	fmt.Println("Max of 5.5 and 2.3:", Max(5.5, 2.3))

	// Использование структуры с дженериками
	intBox := Box[int]{content: 100}
	stringBox := Box[string]{content: "Generic Box"}

	fmt.Println("Box content (int):", intBox.GetContent())
	fmt.Println("Box content (string):", stringBox.GetContent())
}
