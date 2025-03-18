package main

import (
	"fmt"
	"sync"
	"time"
)

// Пример 1: Проблемы с гонками данных (data race) и их решение с помощью каналов
func exampleDataRace() {
	var counter int
	var wg sync.WaitGroup
	ch := make(chan int, 1) // Буферизованный канал для синхронизации

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ch <- 1 // Блокируем доступ к counter
			counter++
			<-ch // Освобождаем доступ к counter
		}()
	}

	wg.Wait()
	fmt.Println("Counter:", counter)
}

// Пример 2: Проблемы с паниками при обращении к закрытым каналам
func exampleClosedChannelPanic() {
	ch := make(chan int)

	go func() {
		ch <- 1
		close(ch)
	}()

	// Первое чтение из канала
	fmt.Println(<-ch)

	// Второе чтение из закрытого канала (не вызовет панику, вернет нулевое значение)
	fmt.Println(<-ch)

	// Попытка записи в закрытый канал вызовет панику
	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered from panic:", r)
			}
		}()
		ch <- 2 // Паника: попытка записи в закрытый канал
	}()

	time.Sleep(time.Second) // Даем время для завершения горутины
}

// Пример 3: Как горутины могут вызывать паники в основном потоке, если их не обрабатывать правильно
func exampleGoroutinePanic() {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in goroutine:", r)
			}
		}()
		panic("goroutine panic")
	}()

	time.Sleep(time.Second) // Даем время для завершения горутины
	fmt.Println("Main goroutine continues")
}

// Пример 4: Лучшие практики синхронизации горутин через sync.WaitGroup, каналы и другие механизмы
func exampleSyncBestPractices() {
	var wg sync.WaitGroup
	ch := make(chan int, 5) // Буферизованный канал для ограничения количества одновременно работающих горутин

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			ch <- i                 // Блокируем, если канал заполнен
			defer func() { <-ch }() // Освобождаем слот в канале после завершения работы

			fmt.Println("Processing task", i)
			time.Sleep(time.Second) // Имитация работы
		}(i)
	}

	wg.Wait()
	fmt.Println("All tasks processed")
}

func main() {
	fmt.Println("Example 1: Data Race and Solution with Channels")
	exampleDataRace()

	fmt.Println("\nExample 2: Panic with Closed Channels")
	exampleClosedChannelPanic()

	fmt.Println("\nExample 3: Goroutine Panic in Main Thread")
	exampleGoroutinePanic()

	fmt.Println("\nExample 4: Best Practices for Goroutine Synchronization")
	exampleSyncBestPractices()
}
