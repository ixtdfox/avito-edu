package main

import (
	"fmt"
)

func main() {
	// Объявление и вызов анонимной функции сразу после создания
	func() {
		fmt.Println("Hello from an anonymous function!")
	}()

	// Анонимная функция, принимающая параметры
	func(msg string) {
		fmt.Println("Message:", msg)
	}("This is a test message")

	// Присвоение анонимной функции переменной и последующий вызов
	multiply := func(a, b int) int {
		return a * b
	}
	fmt.Println("Multiplication result:", multiply(4, 5))

	// Замыкание, которое сохраняет состояние
	counter := func() func() int {
		count := 0
		return func() int {
			count++
			return count
		}
	}()

	fmt.Println("Counter 1:", counter())
	fmt.Println("Counter 2:", counter())
	fmt.Println("Counter 3:", counter())

	// Передача анонимной функции в качестве аргумента другой функции
	applyFunc := func(a, b int, operation func(int, int) int) int {
		return operation(a, b)
	}

	result := applyFunc(10, 2, func(x, y int) int {
		return x - y
	})
	fmt.Println("Subtraction result:", result)

	// Пример использования замыканий для генерации последовательности чисел
	sequenceGenerator := func(start int) func() int {
		return func() int {
			start++
			return start
		}
	}

	nextNumber := sequenceGenerator(100)
	fmt.Println("Next number:", nextNumber())
	fmt.Println("Next number:", nextNumber())
	fmt.Println("Next number:", nextNumber())
}
