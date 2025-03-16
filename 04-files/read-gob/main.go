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
	// Открытие бинарного файла для чтения
	file, err := os.Open("person.gob")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var person Person
	decoder := gob.NewDecoder(file)
	err = decoder.Decode(&person)
	if err != nil {
		fmt.Println("Error decoding data:", err)
	} else {
		fmt.Println("Read data:", person)
	}
}
