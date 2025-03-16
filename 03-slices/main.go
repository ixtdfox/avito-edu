package main

import "fmt"

// Пример 1: Что такое слайсы
func sliceExample() {
	// Слайс - это динамическая структура данных, которая может изменять свой размер
	s := []int{1, 2, 3, 4, 5}
	fmt.Println("Слайс:", s)
}

// Пример 2: Разница между массивами и слайсами
func arrayVsSlice() {
	// Массив фиксированной длины
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Массив:", arr)

	// Слайс динамической длины
	slice := arr[:3] // Слайс от массива, включает первые три элемента
	fmt.Println("Слайс из массива:", slice)
}

// Пример 3: Преимущества слайсов перед массивами
func sliceAdvantages() {
	// Массив
	arr := [3]int{10, 20, 30}
	fmt.Println("Массив:", arr)

	// Слайс, который может изменяться в размере
	slice := []int{10, 20, 30}
	slice = append(slice, 40) // Добавляем новый элемент
	fmt.Println("Слайс после добавления:", slice)
}

// Пример 4: Операции с слайсами: добавление, изменение, инициализация через make()
func sliceOperations() {
	// Инициализация слайса через make
	slice := make([]int, 5) // создаем слайс длины 5 с нулевыми значениями
	fmt.Println("Слайс после make:", slice)

	// Изменение значений в слайсе
	slice[0] = 100
	slice[1] = 200
	fmt.Println("Слайс после изменений:", slice)

	// Добавление элементов с помощью append
	slice = append(slice, 300, 400)
	fmt.Println("Слайс после добавления элементов:", slice)
}

// Пример 5: Расширение слайсов, что такое capacity
func sliceCapacity() {
	// Создание слайса с использованием make и проверка его емкости
	slice := make([]int, 3, 5) // слайс длины 3 и емкостью 5
	fmt.Println("Слайс после make:", slice)
	fmt.Println("Длина слайса:", len(slice))
	fmt.Println("Емкость слайса:", cap(slice))

	// Расширение слайса
	slice = append(slice, 10, 20)
	fmt.Println("Слайс после добавления элементов:", slice)
	fmt.Println("Длина слайса:", len(slice))
	fmt.Println("Емкость слайса:", cap(slice))
}

// Пример 6: Slice оператор slice[low:hight]
func sliceOperator() {
	// Создание слайса
	slice := []int{10, 20, 30, 40, 50, 60}

	// Использование оператора slice
	subSlice1 := slice[1:4] // Слайс от индекса 1 до 3
	subSlice2 := slice[:3]  // Слайс от начала до индекса 2 (не включая 3)
	subSlice3 := slice[3:]  // Слайс с индекса 3 до конца

	fmt.Println("Слайс с 1 по 3 элементы:", subSlice1)
	fmt.Println("Слайс с начала до 2 элемента:", subSlice2)
	fmt.Println("Слайс с 3 элемента до конца:", subSlice3)
}

func main() {
	// Пример 1: Что такое слайсы
	sliceExample()

	// Пример 2: Разница между массивами и слайсами
	arrayVsSlice()

	// Пример 3: Преимущества слайсов перед массивами
	sliceAdvantages()

	// Пример 4: Операции с слайсами
	sliceOperations()

	// Пример 5: Расширение слайсов и емкость
	sliceCapacity()

	// Пример 6: Slice оператор
	sliceOperator()
}
