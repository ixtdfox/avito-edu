package main

import (
	"fmt"
	"sync"
	"time"
)

// Пример использования мьютекса для защиты общих данных
func exampleMutex() {
	var (
		counter int
		mu      sync.Mutex
		wg      sync.WaitGroup
	)

	// Запускаем 10 горутин, которые увеличивают счетчик
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()   // Блокируем доступ к общим данным
			counter++   // Увеличиваем счетчик
			mu.Unlock() // Разблокируем доступ
		}()
	}

	wg.Wait() // Ожидаем завершения всех горутин
	fmt.Println("Counter:", counter)
}

// Типичные ошибки: забытый Unlock, повторный Lock в одной горутине
func exampleMutexErrors() {
	var mu sync.Mutex

	// Забытый Unlock
	mu.Lock()
	// mu.Unlock() // Если забыть разблокировать, это приведет к deadlock

	// Повторный Lock в одной горутине
	mu.Lock() // Это приведет к deadlock, так как мьютекс уже заблокирован
	mu.Unlock()
}

// Пример кода с гонкой данных и его исправление с помощью мьютекса
func exampleDataRace() {
	var (
		counter int
		wg      sync.WaitGroup
	)

	// Гонка данных: несколько горутин пытаются изменить переменную одновременно
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter++ // Гонка данных!
		}()
	}

	wg.Wait()
	fmt.Println("Counter with data race:", counter)

	// Исправление с помощью мьютекса
	var mu sync.Mutex
	counter = 0 // Сбрасываем счетчик

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}

	wg.Wait()
	fmt.Println("Counter with mutex:", counter)
}

// Когда использовать RWMutex вместо Mutex
func exampleRWMutex() {
	var (
		data map[string]string = make(map[string]string)
		mu   sync.RWMutex
		wg   sync.WaitGroup
	)

	// Запись данных
	wg.Add(1)
	go func() {
		defer wg.Done()
		mu.Lock() // Блокируем для записи
		data["key"] = "value"
		mu.Unlock()
	}()

	// Чтение данных
	wg.Add(1)
	go func() {
		defer wg.Done()
		mu.RLock() // Блокируем для чтения
		value := data["key"]
		mu.RUnlock()
		fmt.Println("Read value:", value)
	}()

	wg.Wait()
}

// Пример использования sync.WaitGroup
func exampleWaitGroup() {
	var wg sync.WaitGroup

	// Запускаем несколько горутин
	for i := 0; i < 5; i++ {
		wg.Add(1) // Увеличиваем счетчик WaitGroup
		go func(id int) {
			defer wg.Done() // Уменьшаем счетчик при завершении горутины
			fmt.Printf("Goroutine %d is working\n", id)
			time.Sleep(time.Second)
		}(i)
	}

	wg.Wait() // Ожидаем завершения всех горутин
	fmt.Println("All goroutines are done")
}

// Ошибки: забытый Add или Done, неправильное использование Wait
func exampleWaitGroupErrors() {
	var wg sync.WaitGroup

	// Забытый Add
	// wg.Add(1) // Если забыть добавить, Wait завершится сразу
	go func() {
		defer wg.Done()
		fmt.Println("Working...")
	}()

	wg.Wait()

	// Забытый Done
	wg.Add(1)
	go func() {
		// defer wg.Done() // Если забыть Done, Wait будет ждать вечно
		fmt.Println("Working...")
	}()

	// Неправильное использование Wait
	wg.Wait() // Это приведет к deadlock, если Done не будет вызван
}

func main() {
	fmt.Println("--- Example Mutex ---")
	exampleMutex()

	fmt.Println("\n--- Example Mutex Errors ---")
	// exampleMutexErrors() // Раскомментируйте, чтобы увидеть ошибки

	fmt.Println("\n--- Example Data Race ---")
	exampleDataRace()

	fmt.Println("\n--- Example RWMutex ---")
	exampleRWMutex()

	fmt.Println("\n--- Example WaitGroup ---")
	exampleWaitGroup()

	fmt.Println("\n--- Example WaitGroup Errors ---")
	// exampleWaitGroupErrors() // Раскомментируйте, чтобы увидеть ошибки
}
