package main

const number = 5
const explicitNumber int16 = 42
const nickname = "GoDev"

const (
	featureA = 10
	featureB = 20
)

const (
	alpha = iota
	beta
	_     // пропуск значения
	delta // автоматически 3
)

const (
	_               = iota // пропускаем 0
	Kilobyte uint64 = 1 << (10 * iota)
	Megabyte
	Gigabyte
	Terabyte
	Petabyte
	Exabyte
)

func main() {
	euler := 2.718

	// Компилятор определяет тип константы автоматически
	println(euler + number)

	// Ошибка типов при сложении
	// println(euler + explicitNumber)
	// invalid operation: mismatched types float64 and int16

	println(Kilobyte, Megabyte, Gigabyte, Terabyte, Petabyte, Exabyte)
}
