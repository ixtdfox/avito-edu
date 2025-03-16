package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address,omitempty"` // `omitempty` означает, что поле не будет записано в JSON, если оно пустое
}

func main() {
	person := Person{
		Name:    "Alice",
		Age:     30,
		Address: "123 Main St",
	}

	// Преобразование структуры в JSON
	jsonData, err := json.Marshal(person)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Marshalled JSON:", string(jsonData)) // Выводим JSON как строку
}
