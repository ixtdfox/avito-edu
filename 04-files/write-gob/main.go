// Запись в бинарный файл с использованием encoding/gob:
// Этот пример показывает, как работать с бинарным форматом файлов. Мы используем библиотеку encoding/gob для сериализации и десериализации объектов.
package main

import (
	"encoding/gob"
	"fmt"
	"os"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	person := Person{Name: "Alice", Age: 30}

	// Открытие файла для записи в бинарном формате
	file, err := os.Create("person.gob")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	err = encoder.Encode(person)
	if err != nil {
		fmt.Println("Error encoding data:", err)
	} else {
		fmt.Println("Data written successfully in binary format.")
	}
}
