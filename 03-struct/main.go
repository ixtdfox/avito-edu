package main

import "fmt"

// Определение структуры
type Person struct {
	FirstName string
	LastName  string
	Age       int
}

type Address struct {
	Street  string
	City    string
	ZipCode int
}

// Структура с вложенной структурой
type Employee struct {
	Person   // Вложенная структура (неявное использование Person)
	Position string
	Salary   int
}

// Методы для структур

// Метод для структуры Person
func (p Person) FullName() string {
	return p.FirstName + " " + p.LastName
}

// Метод для структуры Employee (использование указателя для изменения значений)
func (e *Employee) UpdateSalary(newSalary int) {
	e.Salary = newSalary
}

// Конструктор для Person, возвращающий указатель на структуру
func NewPerson(firstName, lastName string, age int) *Person {
	return &Person{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}
}

// Конструктор для Employee, возвращающий указатель на структуру
func NewEmployee(firstName, lastName string, age int, position string, salary int) *Employee {
	return &Employee{
		Person: Person{
			FirstName: firstName,
			LastName:  lastName,
			Age:       age,
		},
		Position: position,
		Salary:   salary,
	}
}

func main() {
	// 1. Инициализация структуры с именованными полями
	p1 := Person{
		FirstName: "John",
		LastName:  "Doe",
		Age:       30,
	}

	// 2. Доступ к полям структуры
	fmt.Println("First Name:", p1.FirstName)
	fmt.Println("Last Name:", p1.LastName)
	fmt.Println("Age:", p1.Age)

	// 3. Изменение полей структуры
	p1.Age = 31
	fmt.Println("Updated Age:", p1.Age)

	// 4. Передача структуры в функцию
	printPerson(p1)

	// 5. Инициализация структуры с вложенной структурой
	e1 := Employee{
		Person: Person{
			FirstName: "Alice",
			LastName:  "Smith",
			Age:       28,
		},
		Position: "Software Engineer",
		Salary:   70000,
	}

	// Доступ к полям вложенной структуры
	fmt.Println("Employee Full Name:", e1.FullName())
	fmt.Println("Employee Position:", e1.Position)

	// 6. Использование методов для изменения данных структуры
	fmt.Println("Old Salary:", e1.Salary)
	e1.UpdateSalary(75000)
	fmt.Println("Updated Salary:", e1.Salary)

	// 7. Вложенные структуры с доступом к вложенным полям
	address := Address{
		Street:  "123 Main St",
		City:    "Springfield",
		ZipCode: 12345,
	}

	// Создаем новый объект Person с вложенным адресом
	p2 := struct {
		Person
		Address
	}{
		Person: Person{
			FirstName: "Bob",
			LastName:  "Johnson",
			Age:       35,
		},
		Address: address,
	}

	// Доступ к полям вложенной структуры
	fmt.Println("Person with Address:", p2.FullName(), "from", p2.Street, p2.City, p2.ZipCode)

	// 8. Инициализация с использованием new (Конструктор)
	p3 := new(Person)
	p3.FirstName = "Charlie"
	p3.LastName = "Brown"
	p3.Age = 40
	fmt.Println("Person from new constructor:", p3.FullName(), "Age:", p3.Age)

	// 9. Инициализация без именованных полей (значения идут в порядке определения полей)
	p4 := Person{"David", "White", 25}
	fmt.Println("Person with positional initialization:", p4.FullName(), "Age:", p4.Age)

	// 10. Инициализация Employee с помощью конструктора
	e2 := NewEmployee("Eve", "Green", 32, "Data Scientist", 85000)
	fmt.Println("Employee from constructor:", e2.FullName(), "Position:", e2.Position, "Salary:", e2.Salary)
}

// Функция, принимающая структуру в качестве аргумента
func printPerson(p Person) {
	fmt.Println("Person Details:", p.FirstName, p.LastName, p.Age)
}
