package main

func main() {
	// Проверяем состояние активности
	isActive := true
	if isActive {
		println("hello world")
	}

	// Пример некорректного преобразования в if-условии
	count := 1
	if count == 1 {
		println("неявное преобразование ( if count ) не работает")
	}

	// Пример работы с map
	systemStatus := map[string]string{"systemID": "001", "status": "active"}
	if systemID, found := systemStatus["systemID"]; found {
		println("Найден ключ systemID, его значение: ", systemID)
	} else {
		println("Ключ systemID отсутствует")
	}

	// Проверка существования ключа с противоположной логикой
	if systemID, found := systemStatus["systemID"]; !found {
		println("Ключ systemID не найден")
	} else if systemID == "001" {
		println("systemID = 001")
	} else {
		println("systemID не равен 001")
	}

	// Цикл, который завершится после одного шага
	for {
		println("Цикл без завершения")
		break
	}

	// Работа с срезом
	numbers := []int{10, 20, 30, 40, 50, 60}
	var currentValue int
	position := 0

	// Пример цикла, напоминающего while
	for position < 4 {
		if position < 2 {
			position++
			continue
		}
		currentValue = numbers[position]
		position++
		println("Цикл, похожий на while, индекс:", position, "значение:", currentValue)
	}

	// Пример классического цикла for
	for idx := 0; idx < len(numbers); idx++ {
		println("Цикл C-style", idx, numbers[idx])
	}

	// Пример использования range по индексам
	for index := range numbers {
		println("range по индексу среза", index)
	}

	// Пример использования range с индексом и значением
	for index, value := range numbers {
		println("range по индексу и значению среза", index, value)
	}

	// Итерация по ключам карты
	for key := range systemStatus {
		println("range по ключам карты", key)
	}

	// Итерация по ключам и значениям карты
	for key, value := range systemStatus {
		println("range по ключам и значениям карты", key, value)
	}

	// Итерация по значениям карты
	for _, value := range systemStatus {
		println("range по значениям карты", value)
	}

	// Изменяем значения в map
	systemStatus["systemID"] = "001"
	systemStatus["status"] = "active"

	// Использование switch для проверки значения
	switch systemStatus["systemID"] {
	case "001", "002":
		println("switch - systemID это 001 или 002")
	case "003":
		if systemStatus["status"] == "active" {
			break // Выход из switch
		}
		println("switch - systemID это 003")
		fallthrough
	default:
		println("switch - systemID не из списка")
	}

	// Альтернатива множественным if-else с использованием switch
	switch {
	case systemStatus["systemID"] == "001":
		println("switch2 - это systemID 001")
	case systemStatus["status"] == "inactive":
		println("switch2 - статус неактивен")
	}

	// Прерывание цикла внутри switch
OuterLoop:
	for key, value := range systemStatus {
		println("switch внутри цикла", key, value)
		switch {
		case key == "systemID" && value == "001":
			println("switch - выходим из цикла здесь")
			break OuterLoop
		}
	}
}
