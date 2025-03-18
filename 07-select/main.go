package main

import (
	"fmt"
	"time"
)

func main() {
	// Пример 1: Объяснение конструкции select
	exampleSelect()

	// Пример 2: Тайм-ауты с использованием time.After и select
	exampleTimeout()

	// Пример 3: Использование select с несколькими каналами и тайм-аутами
	exampleMultipleChannelsWithTimeout()
}

// Пример 1: Объяснение конструкции select
func exampleSelect() {
	fmt.Println("--- Пример 1: Использование select ---")

	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "Сообщение из канала 1"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "Сообщение из канала 2"
	}()

	// select ожидает данные из любого канала
	select {
	case msg1 := <-ch1:
		fmt.Println(msg1)
	case msg2 := <-ch2:
		fmt.Println(msg2)
	}

	// В этом примере select выберет первый канал, который получит данные (в данном случае ch1)
}

// Пример 2: Тайм-ауты с использованием time.After и select
func exampleTimeout() {
	fmt.Println("\n--- Пример 2: Тайм-ауты с time.After ---")

	ch := make(chan string)

	go func() {
		time.Sleep(3 * time.Second)
		ch <- "Данные получены"
	}()

	select {
	case msg := <-ch:
		fmt.Println(msg)
	case <-time.After(2 * time.Second):
		fmt.Println("Тайм-аут: данные не получены вовремя")
	}

	// В этом примере тайм-аут сработает раньше, чем данные придут в канал
}

// Пример 3: Использование select с несколькими каналами и тайм-аутами
func exampleMultipleChannelsWithTimeout() {
	fmt.Println("\n--- Пример 3: Несколько каналов и тайм-ауты ---")

	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "Данные из канала 1"
	}()

	go func() {
		time.Sleep(3 * time.Second)
		ch2 <- "Данные из канала 2"
	}()

	select {
	case msg1 := <-ch1:
		fmt.Println(msg1)
	case msg2 := <-ch2:
		fmt.Println(msg2)
	case <-time.After(2 * time.Second):
		fmt.Println("Тайм-аут: ни один из каналов не ответил вовремя")
	}

	// В этом примере select выберет данные из ch1, так как он ответит первым.
	// Если бы ch1 ответил позже 2 секунд, сработал бы тайм-аут.
}
