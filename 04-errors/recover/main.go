package main

import "fmt"

func safeDivide(a, b int) (result int, err error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
			err = fmt.Errorf("error: %v", r)
		}
	}()
	if b == 0 {
		panic("division by zero")
	}
	return a / b, nil
}

func main() {
	result, err := safeDivide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
}
