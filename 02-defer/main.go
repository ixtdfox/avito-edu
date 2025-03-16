package main

import (
	"fmt"
	"os"
)

// Пример использования defer для отсроченного выполнения функции
func deferredExample() {
	defer fmt.Println("This is printed last")
	fmt.Println("This is printed first")
}

// Использование defer для закрытия файла
func fileHandling() {
	// Открываем файл
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	// Используем defer, чтобы закрыть файл по завершении функции
	defer file.Close()

	// Читаем из файла (для примера)
	buffer := make([]byte, 100)
	_, err = file.Read(buffer)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Файл будет закрыт после того, как функция завершит выполнение,
	// даже если произойдет ошибка или функция завершится раньше.
	fmt.Println("File is being read")
}

// Использование defer в контексте работы с базой данных (имитация)
func dbTransaction() {
	// Открываем соединение с базой данных
	fmt.Println("Connecting to database...")

	// Используем defer, чтобы симулировать закрытие соединения с базой данных
	defer fmt.Println("Database connection closed")

	// Начинаем транзакцию
	fmt.Println("Starting transaction...")

	// Завершаем транзакцию
	fmt.Println("Committing transaction...")
}

// Пример с несколькими defer, чтобы показать порядок их вызова
func multipleDefer() {
	defer fmt.Println("Deferred 1")
	defer fmt.Println("Deferred 2")
	defer fmt.Println("Deferred 3")

	// Важно: defer вызывает функции в обратном порядке
	fmt.Println("Main function")
}

func main() {
	// Пример 1: deferredExample
	deferredExample()

	// Пример 2: fileHandling (закрытие файла с defer)
	fileHandling()

	// Пример 3: dbTransaction
	dbTransaction()

	// Пример 4: multipleDefer
	multipleDefer()
}
