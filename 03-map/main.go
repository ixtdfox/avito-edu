package main

import (
	"fmt"
)

func main() {
	// Что такое карты?
	// Карта (map) — это структура данных, которая хранит пары "ключ-значение".
	// Каждый ключ в карте уникален, и его можно использовать для доступа к соответствующему значению.

	// Инициализация карты с помощью make()
	// Используем make() для создания карты с типом ключа string и типом значения int
	ageMap := make(map[string]int)

	// Добавление элементов в карту
	ageMap["Alice"] = 25
	ageMap["Bob"] = 30
	ageMap["Charlie"] = 35

	// Выводим элементы карты
	fmt.Println("ageMap:", ageMap)

	// Как карты отличаются от слайсов и массивов?
	// В отличие от слайсов и массивов, карты не имеют фиксированной длины, и элементы не упорядочены.
	// В слайсах и массивах индексы всегда числовые, а в картах ключи могут быть любыми типами данных (например, строки, числа, и т.д.).

	// Пример поиска элемента по ключу
	age := ageMap["Alice"]
	fmt.Println("Age of Alice:", age)

	// Пример удаления элемента из карты
	delete(ageMap, "Bob")
	fmt.Println("ageMap after deleting Bob:", ageMap)

	// Проверка существования ключа в карте
	// При поиске по ключу можно одновременно получить и булевое значение, которое указывает, существует ли такой ключ.
	value, exists := ageMap["Charlie"]
	if exists {
		fmt.Println("Charlie is", value, "years old")
	} else {
		fmt.Println("Charlie not found")
	}

	// Пример, когда ключ не существует
	_, exists = ageMap["Bob"]
	if !exists {
		fmt.Println("Bob is not in the map")
	}

	// Пример инициализации карты с заданными значениями
	// Также можно инициализировать карту с использованием литерала
	productPrices := map[string]float64{
		"apple":  0.99,
		"banana": 1.29,
		"orange": 1.49,
	}
	fmt.Println("Product prices:", productPrices)

	// Перебор элементов карты
	// Можно использовать цикл range для перебора ключей и значений в карте.
	for product, price := range productPrices {
		fmt.Printf("The price of %s is %.2f\n", product, price)
	}

	// Использование карты для подсчета частоты элементов
	// Например, подсчет частоты появления символов в строке
	text := "hello"
	charFrequency := make(map[rune]int)
	for _, char := range text {
		charFrequency[char]++
	}
	fmt.Println("Character frequencies in 'hello':", charFrequency)

	// Пример с пустой картой и проверкой существования ключа
	var emptyMap map[string]int
	if emptyMap == nil {
		fmt.Println("The map is nil")
	} else {
		fmt.Println("The map is not nil")
	}
}
