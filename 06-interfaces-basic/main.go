package main

import (
	"fmt"
)

// Определение интерфейса
// В Go интерфейсы определяют поведение, а не структуру

type Speaker interface {
	Speak() string
}

// Структуры, реализующие интерфейс Speaker неявно

type Dog struct{}

func (d Dog) Speak() string {
	return "Woof!"
}

type Cat struct{}

func (c Cat) Speak() string {
	return "Meow!"
}

// Динамическая типизация: использование пустого интерфейса
func PrintAnything(value interface{}) {
	fmt.Println("Value:", value)
}

// Пример принципа разделения интерфейсов (ISP - Interface Segregation Principle)
// Разделяем функциональность на отдельные маленькие интерфейсы

type Printer interface {
	Print()
}

type Scanner interface {
	Scan()
}

// Многофункциональное устройство реализует оба интерфейса

type MultiFunctionDevice struct{}

func (m MultiFunctionDevice) Print() {
	fmt.Println("Printing document...")
}

func (m MultiFunctionDevice) Scan() {
	fmt.Println("Scanning document...")
}

func main() {
	// Использование интерфейсов
	var s Speaker

	dog := Dog{}
	cat := Cat{}

	s = dog
	fmt.Println("Dog says:", s.Speak())

	s = cat
	fmt.Println("Cat says:", s.Speak())

	// Динамическая типизация
	PrintAnything(42)
	PrintAnything("Hello")

	// Использование принципа ISP
	var printer Printer = MultiFunctionDevice{}
	var scanner Scanner = MultiFunctionDevice{}

	printer.Print()
	scanner.Scan()
}
