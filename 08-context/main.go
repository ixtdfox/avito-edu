package main

import (
	"context"
	"fmt"
	"time"
)

// Пример 1: Что такое контекст и зачем он нужен.
func exampleContextUsage() {
	// Создаем контекст с таймаутом 2 секунды.
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel() // Освобождаем ресурсы контекста после завершения.

	// Запускаем горутину, которая будет выполнять какую-то работу.
	go func(ctx context.Context) {
		select {
		case <-time.After(3 * time.Second):
			fmt.Println("Работа завершена")
		case <-ctx.Done():
			fmt.Println("Контекст отменен:", ctx.Err())
		}
	}(ctx)

	// Ждем завершения горутины.
	time.Sleep(4 * time.Second)
}

// Пример 2: Как передавать контекст между горутинами.
func exampleContextPassing() {
	// Создаем контекст с возможностью отмены.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Запускаем горутину, которая будет ждать отмены контекста.
	go func(ctx context.Context) {
		select {
		case <-ctx.Done():
			fmt.Println("Горутина 1: Контекст отменен")
		}
	}(ctx)

	// Передаем контекст в другую горутину.
	go func(ctx context.Context) {
		select {
		case <-ctx.Done():
			fmt.Println("Горутина 2: Контекст отменен")
		}
	}(ctx)

	// Отменяем контекст через 1 секунду.
	time.Sleep(1 * time.Second)
	cancel()

	// Даем время горутинам завершиться.
	time.Sleep(1 * time.Second)
}

// Пример 3: Утечка горутин из-за отсутствия отмены через контекст.
func exampleGoroutineLeak() {
	// Создаем канал для сигнала завершения.
	done := make(chan struct{})

	// Запускаем горутину, которая будет работать бесконечно.
	go func() {
		for {
			select {
			case <-done:
				fmt.Println("Горутина завершена")
				return
			default:
				// Имитация работы.
				time.Sleep(500 * time.Millisecond)
				fmt.Println("Горутина работает...")
			}
		}
	}()

	// Ждем некоторое время, а затем завершаем работу.
	time.Sleep(2 * time.Second)
	close(done)

	// Даем время горутине завершиться.
	time.Sleep(1 * time.Second)
	fmt.Println("Основная горутина завершена")
}

func main() {
	fmt.Println("Пример 1: Что такое контекст и зачем он нужен.")
	exampleContextUsage()

	fmt.Println("\nПример 2: Как передавать контекст между горутинами.")
	exampleContextPassing()

	fmt.Println("\nПример 3: Утечка горутин из-за отсутствия отмены через контекст.")
	exampleGoroutineLeak()
}
