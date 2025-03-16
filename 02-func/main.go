package main

import (
	"errors"
	"fmt"
)

// Простая функция без параметров и возврата значения
func greet() {
	fmt.Println("Hello, World!")
}

// Функция с параметрами, возвращающая значение
func add(a int, b int) int {
	return a + b
}

// Функция с несколькими возвращаемыми значениями
func divide(a, b int) (int, int, error) {
	if b == 0 {
		return 0, 0, errors.New("division by zero")
	}
	return a / b, a % b, nil
}

// Функция с именованными возвращаемыми значениями
func calculate(a, b int) (sum, difference int) {
	sum = a + b
	difference = a - b
	return
}

// Функция с переменным числом аргументов
func average(numbers ...int) float64 {
	var sum int
	for _, num := range numbers {
		sum += num
	}
	return float64(sum) / float64(len(numbers))
}

// Функция, вызывающая другую функцию
func performOperation(a, b int) {
	result := add(a, b)
	fmt.Println("Result of addition:", result)
}

// Функция с замыканием
func multiplyBy(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

// Функция как переменная
var square = func(x int) int {
	return x * x
}

func main() {
	// Вызов простой функции без параметров
	greet()

	// Вызов функции с параметрами и возвращаемым значением
	result := add(10, 5)
	fmt.Println("Addition result:", result)

	// Вызов функции с несколькими значениями, включая обработку ошибок
	quotient, remainder, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Quotient:", quotient, "Remainder:", remainder)
	}

	// Функция с именованными возвращаемыми значениями
	total, diff := calculate(10, 3)
	fmt.Println("Sum:", total, "Difference:", diff)

	// Вызов функции с переменным числом аргументов
	avg := average(1, 2, 3, 4, 5)
	fmt.Println("Average:", avg)

	// Вызов функции, которая вызывает другую
	performOperation(3, 7)

	// Вызов функции с замыканием
	multiply := multiplyBy(3)
	fmt.Println("Multiplying by 3:", multiply(10))

	// Вызов функции как переменной
	fmt.Println("Square of 4:", square(4))
}
