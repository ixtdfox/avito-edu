package main

import (
	"encoding/json"
	"fmt"
)

// Определение интерфейса
// Интерфейс без методов — это пустой интерфейс, который может содержать значение любого типа
type AnyType interface{}

// Функция, принимающая пустой интерфейс
func printValue(value interface{}) {
	fmt.Println("Value:", value)
}

// Использование type assertion
func checkType(value interface{}) {
	switch v := value.(type) {
	case int:
		fmt.Println("Integer:", v)
	case string:
		fmt.Println("String:", v)
	case bool:
		fmt.Println("Boolean:", v)
	default:
		fmt.Println("Unknown type")
	}
}

func main() {
	// Примеры работы с пустым интерфейсом
	var anything interface{}

	anything = 42
	printValue(anything)

	anything = "Hello, Go!"
	printValue(anything)

	anything = true
	printValue(anything)

	// Использование type assertion для определения типа
	checkType(100)
	checkType("Golang")
	checkType(false)

	// Пример использования пустого интерфейса с JSON
	jsonData := `{"name": "John", "age": 30, "active": true}`
	var result map[string]interface{} // Используем пустой интерфейс для хранения данных
	if err := json.Unmarshal([]byte(jsonData), &result); err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}
	fmt.Println("Parsed JSON:", result)

	// Доступ к значениям через type assertion
	if name, ok := result["name"].(string); ok {
		fmt.Println("Name:", name)
	}
	if age, ok := result["age"].(float64); ok { // JSON числа по умолчанию float64
		fmt.Println("Age:", int(age))
	}
	if active, ok := result["active"].(bool); ok {
		fmt.Println("Active:", active)
	}
}
