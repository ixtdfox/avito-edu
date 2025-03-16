package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString("Hello, Go!")
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}
}
