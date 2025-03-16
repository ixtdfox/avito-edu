//Чтение файла построчно с использованием bufio.Scanner
//Этот метод позволяет считывать файл построчно, что может быть полезно для обработки больших файлов.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text()) // Чтение строки и вывод
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
