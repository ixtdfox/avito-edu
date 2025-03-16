package main

import "fmt"

func divide(a, b int) int {
	if b == 0 {
		panic("division by zero") // вызов паники при делении на ноль
	}
	return a / b
}

func main() {
	fmt.Println(divide(10, 0)) // вызывает панику
}
