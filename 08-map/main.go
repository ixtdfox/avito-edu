package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// Пример 1: Что такое sync.Map и когда его использовать
	example1()

	// Пример 2: Пример использования sync.Map для конкурентного доступа к данным
	example2()

	// Пример 3: Сравнение с обычным map и мьютексами
	example3()

	// Пример 4: Ограничения и подводные камни sync.Map
	example4()
}

// Пример 1: Что такое sync.Map и когда его использовать
func example1() {
	fmt.Println("Пример 1: Что такое sync.Map и когда его использовать")

	// sync.Map — это потокобезопасная реализация map в Go.
	// Он предназначен для использования в случаях, когда у вас есть множество горутин,
	// которые одновременно читают, записывают и удаляют данные из map.
	// В отличие от обычного map, sync.Map не требует внешней синхронизации с помощью мьютексов.

	var sm sync.Map

	// Запись данных
	sm.Store("key1", "value1")
	sm.Store("key2", "value2")

	// Чтение данных
	if value, ok := sm.Load("key1"); ok {
		fmt.Println("key1:", value)
	}

	// Удаление данных
	sm.Delete("key2")

	// Проверка наличия ключа
	if _, ok := sm.Load("key2"); !ok {
		fmt.Println("key2 удален")
	}
}

// Пример 2: Пример использования sync.Map для конкурентного доступа к данным
func example2() {
	fmt.Println("\nПример 2: Пример использования sync.Map для конкурентного доступа к данным")

	var sm sync.Map
	var wg sync.WaitGroup

	// Запускаем несколько горутин для конкурентного доступа к sync.Map
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			sm.Store(fmt.Sprintf("key%d", i), fmt.Sprintf("value%d", i))
		}(i)
	}

	wg.Wait()

	// Чтение данных из sync.Map
	sm.Range(func(key, value interface{}) bool {
		fmt.Printf("%s: %s\n", key, value)
		return true
	})
}

// Пример 3: Сравнение с обычным map и мьютексами
func example3() {
	fmt.Println("\nПример 3: Сравнение с обычным map и мьютексами")

	// Обычный map с мьютексом
	var regularMap = make(map[string]string)
	var mutex sync.Mutex

	// sync.Map
	var sm sync.Map

	// Запись данных в обычный map с мьютексом
	start := time.Now()
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			mutex.Lock()
			regularMap[fmt.Sprintf("key%d", i)] = fmt.Sprintf("value%d", i)
			mutex.Unlock()
		}(i)
	}
	wg.Wait()
	fmt.Println("Обычный map с мьютексом:", time.Since(start))

	// Запись данных в sync.Map
	start = time.Now()
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			sm.Store(fmt.Sprintf("key%d", i), fmt.Sprintf("value%d", i))
		}(i)
	}
	wg.Wait()
	fmt.Println("sync.Map:", time.Since(start))
}

// Пример 4: Ограничения и подводные камни sync.Map
func example4() {
	fmt.Println("\nПример 4: Ограничения и подводные камни sync.Map")

	// 1. sync.Map не типизирован, поэтому вам нужно приводить типы вручную.
	var sm sync.Map
	sm.Store("key1", 123)

	if value, ok := sm.Load("key1"); ok {
		if intValue, ok := value.(int); ok {
			fmt.Println("key1:", intValue)
		} else {
			fmt.Println("Ошибка при приведении типа")
		}
	}

	// 2. sync.Map может быть менее эффективным, чем обычный map с мьютексом,
	// если у вас небольшое количество горутин или если операции записи редки.

	// 3. sync.Map не поддерживает некоторые операции, такие как получение длины map.
	// Для этого нужно использовать Range и подсчитывать элементы вручную.

	// 4. sync.Map может потреблять больше памяти, чем обычный map, из-за внутренней структуры,
	// оптимизированной для конкурентного доступа.

	// 5. sync.Map не подходит для случаев, когда вам нужно часто обновлять одни и те же ключи,
	// так как это может привести к contention (конкуренции за ресурсы).
}
