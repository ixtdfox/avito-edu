package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}

func main() {
	// Пример JSON данных
	jsonData := `{"name":"Alice","age":30,"address":"123 Main St"}`

	var person Person

	// Преобразование JSON обратно в структуру
	err := json.Unmarshal([]byte(jsonData), &person)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Unmarshalled struct:", person) // Выводим структуру
}
