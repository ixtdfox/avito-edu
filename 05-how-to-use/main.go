package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"sort"
	"sync"
	"time"
)

// Фильтрация данных: фильтрация четных чисел
func filter(slice []int, predicate func(int) bool) []int {
	var result []int
	for _, value := range slice {
		if predicate(value) {
			result = append(result, value)
		}
	}
	return result
}

// Преобразование данных: умножение каждого элемента на 2
func mapSlice(slice []int, transform func(int) int) []int {
	result := make([]int, len(slice))
	for i, value := range slice {
		result[i] = transform(value)
	}
	return result
}

// Агрегация данных: сумма всех элементов слайса
func reduce(slice []int, accumulator func(int, int) int, initial int) int {
	result := initial
	for _, value := range slice {
		result = accumulator(result, value)
	}
	return result
}

// Сортировка с кастомным компаратором
func sortCustom(slice []int, comparator func(int, int) bool) {
	sort.Slice(slice, func(i, j int) bool {
		return comparator(slice[i], slice[j])
	})
}

// Обработка ошибок: функция-обёртка для обработки ошибок
func withErrorHandler(fn func() error) {
	if err := fn(); err != nil {
		fmt.Println("Error occurred:", err)
	}
}

// Пайплайн обработки данных
func pipeline(slice []int) int {
	filtered := filter(slice, func(x int) bool { return x%2 == 0 })
	transformed := mapSlice(filtered, func(x int) int { return x * 2 })
	return reduce(transformed, func(a, b int) int { return a + b }, 0)
}

// Логирование операций
func withLogging(fn func()) {
	start := time.Now()
	fn()
	fmt.Println("Execution time:", time.Since(start))
}

// Ретри-логика: повторный вызов в случае ошибки
func retry(fn func() error, retries int) error {
	for i := 0; i < retries; i++ {
		if err := fn(); err == nil {
			return nil
		}
		time.Sleep(time.Second)
	}
	return errors.New("operation failed after retries")
}

// Троттлинг: ограничение частоты вызова
func throttle(fn func(), duration time.Duration) func() {
	var lastCall time.Time
	return func() {
		if time.Since(lastCall) > duration {
			lastCall = time.Now()
			fn()
		}
	}
}

// Middleware в веб-приложениях (имитация)
func middleware(fn func()) func() {
	return func() {
		fmt.Println("Middleware: Logging request")
		fn()
	}
}

// Контекст и отмена операций
func processWithTimeout(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	select {
	case <-time.After(2 * time.Second):
		fmt.Println("Process completed")
	case <-ctx.Done():
		fmt.Println("Process cancelled")
	}
}

// Параллельная обработка данных
func parallelProcessing(slice []int, worker func(int) int) []int {
	var wg sync.WaitGroup
	result := make([]int, len(slice))

	for i, v := range slice {
		wg.Add(1)
		go func(i, v int) {
			defer wg.Done()
			result[i] = worker(v)
		}(i, v)
	}

	wg.Wait()
	return result
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Фильтрация четных чисел
	evenNumbers := filter(numbers, func(x int) bool { return x%2 == 0 })
	fmt.Println("Even Numbers:", evenNumbers)

	// Преобразование: умножение каждого числа на 2
	squaredNumbers := mapSlice(numbers, func(x int) int { return x * 2 })
	fmt.Println("Doubled Numbers:", squaredNumbers)

	// Агрегация: сумма всех чисел
	sum := reduce(numbers, func(a, b int) int { return a + b }, 0)
	fmt.Println("Sum of numbers:", sum)

	// Сортировка с кастомным компаратором (по убыванию)
	sortCustom(numbers, func(a, b int) bool { return a > b })
	fmt.Println("Sorted Numbers:", numbers)

	// Обработка ошибок через обёртку
	withErrorHandler(func() error {
		return errors.New("this is a test error")
	})

	// Пайплайн обработки данных
	fmt.Println("Pipeline result:", pipeline(numbers))

	// Логирование времени выполнения
	withLogging(func() {
		time.Sleep(500 * time.Millisecond)
		fmt.Println("Function executed")
	})

	// Ретри-логика
	retryErr := retry(func() error {
		if rand.Float32() < 0.7 {
			return errors.New("random failure")
		}
		return nil
	}, 5)
	fmt.Println("Retry result:", retryErr)

	// Троттлинг вызовов
	throttledFunc := throttle(func() { fmt.Println("Throttled function executed") }, time.Second)
	for i := 0; i < 5; i++ {
		throttledFunc()
		time.Sleep(300 * time.Millisecond)
	}

	// Middleware в веб-приложениях
	wrappedFunction := middleware(func() { fmt.Println("Handling request") })
	wrappedFunction()

	// Контекст и отмена операций
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	var wg sync.WaitGroup
	wg.Add(1)
	go processWithTimeout(ctx, &wg)
	time.Sleep(500 * time.Millisecond)
	cancel()
	wg.Wait()

	// Параллельная обработка данных
	squaredResults := parallelProcessing(numbers, func(x int) int { return x * x })
	fmt.Println("Parallel squared results:", squaredResults)
}
