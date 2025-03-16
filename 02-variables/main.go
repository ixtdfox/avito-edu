package main

import "fmt"

func main() {
	// Числовые переменные
	var wholeNumber int = 25
	var inferredNumber = -7
	var largeNumber int64 = 1<<40 - 1
	var positiveNumber uint = 123456
	var maxUnsigned uint64 = 1<<63 - 1
	fmt.Println("Numbers:", wholeNumber, inferredNumber, largeNumber, positiveNumber, maxUnsigned)

	// Числа с плавающей точкой
	var pi float64 = 3.1415926535
	fmt.Println("Pi value:", pi)

	// Логические переменные
	var flag bool = false
	fmt.Println("Boolean flag:", flag)

	// Строки
	var greeting string = "Hi there!\n"
	var name = "Gopher"
	fmt.Println(greeting, name)

	// Бинарные данные
	var specialByte byte = '\x41' // ASCII символ 'A'
	fmt.Println("Byte value:", specialByte)

	// Краткая запись объявления переменных
	magicNumber := 7
	fmt.Println("Lucky number:", magicNumber)

	// Преобразование типов
	fmt.Println("Float to Integer:", int(pi))
	fmt.Println("Integer to Character:", string(97)) // ASCII 'a'

	// Комплексные числа
	complexNum := 1 + 4i
	fmt.Println("Complex number:", complexNum)

	// Операции со строками
	firstName := "John"
	lastName := "Doe"
	fullName := firstName + " " + lastName
	fmt.Println("Full name and its length:", fullName, len(fullName))

	escapeExample := `Line1\nLine2`
	fmt.Println("Raw string:", escapeExample)

	// Значения по умолчанию
	var zeroInt int
	var zeroFloat float64
	var emptyString string
	var zeroBool bool
	fmt.Println("Default values:", zeroInt, zeroFloat, emptyString, zeroBool)

	// Объявление нескольких переменных
	var first, second = "Alpha", "Beta"
	fmt.Println(first, second)

	var (
		num1 int    = 99
		text string = "Example"
		val         = 42
	)
	fmt.Println(num1, text, val)
}
