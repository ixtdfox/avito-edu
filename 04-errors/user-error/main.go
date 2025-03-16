package main

import (
	"errors"
	"fmt"
)

func checkValue(val int) error {
	if val < 0 {
		return errors.New("value cannot be negative")
	}
	return nil
}

func main() {
	err := checkValue(-5)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Value is valid.")
	}
}
