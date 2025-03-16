// Простой пример чтения файла
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	data, err := os.ReadFile("example.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data)) // выводим содержимое файла
}
