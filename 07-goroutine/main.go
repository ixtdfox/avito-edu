package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello from goroutine!")
}

// Пример 2: Горутина с анонимной функцией
// Запуск анонимной функции в горутине прямо в месте вызова.
func example2() {
	// Используем анонимную функцию внутри горутины
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("A goroutine is running", i)
			time.Sleep(100 * time.Millisecond)
		}
	}()

	// Ожидаем завершения горутины с помощью time.Sleep
	time.Sleep(1 * time.Second)
}

// Пример 3: Горутина с параметрами
// Передача параметров в функцию, запускаемую в горутине.
func example3() {
	// Запускаем горутину, передав параметры в функцию
	go printNumbers(5)

	// Задержка для завершения работы горутины
	time.Sleep(1 * time.Second)
}

func printNumbers(count int) {
	for i := 0; i < count; i++ {
		fmt.Println(i)
		time.Sleep(200 * time.Millisecond)
	}
}

// Пример 4: Несколько горутин
// Запуск нескольких горутин параллельно.
func example4() {
	// Запускаем несколько горутин с разными функциями
	go sayHello()
	go printNumbers(3)

	// Задержка для завершения всех горутин
	time.Sleep(1 * time.Second)
}

func main() {
	go sayHello()
	time.Sleep(1 * time.Second)

	example2()

	example3()

	example4()

	fmt.Println("Main function is finished.")
}
