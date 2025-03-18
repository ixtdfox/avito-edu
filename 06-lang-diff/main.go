package main

import (
	"fmt"
)

// В ООП-языках (Java, C#) наследование используется для иерархий классов.
// В Go вместо наследования используется композиция и интерфейсы.

// Базовая структура (аналог базового класса в ООП)
type Animal struct {
	Name string
}

func (a Animal) Speak() {
	fmt.Println("Some generic animal sound")
}

// Наследование в ООП (Java, C#) выглядело бы как: class Dog extends Animal {}
// В Go вместо этого используется встраивание (композиция)
type Dog struct {
	Animal // Встраиваем структуру Animal
	Breed  string
}

func (d Dog) Speak() {
	fmt.Println("Woof!")
}

// Интерфейсы в Go (в отличие от классов и наследования)
type Speaker interface {
	Speak()
}

type Walker interface {
	Walk()
}

// Структура может реализовывать несколько интерфейсов
type Person struct {
	Name string
}

func (p Person) Speak() {
	fmt.Println("Hello, my name is", p.Name)
}

func (p Person) Walk() {
	fmt.Println(p.Name, "is walking")
}

func main() {
	// Демонстрация композиции
	dog := Dog{Animal: Animal{Name: "Buddy"}, Breed: "Golden Retriever"}
	dog.Speak()

	// Использование интерфейса
	var speaker Speaker
	speaker = Person{Name: "Alice"}
	speaker.Speak()

	// Использование нескольких интерфейсов
	var walker Walker = Person{Name: "Bob"}
	walker.Walk()
}
