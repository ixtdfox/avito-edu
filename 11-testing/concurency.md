# Написание тестов для конкурентного кода в Go

Конкурентный код в Go (с использованием горутин и каналов) требует особого подхода к тестированию. В этой статье разберём:
- Как тестировать горутины
- Как выявлять гонки данных (data race)
- Использование `sync.WaitGroup` и `sync.Mutex` в тестах

---

## Особенности тестирования горутин

Горутины выполняются асинхронно, поэтому тесты должны дожидаться их завершения. Иначе тест может завершиться раньше, чем горутина выполнит свою работу.

### Пример 1: Тест без ожидания горутины (проблемный вариант)

```go
func ProcessAsync(data int, resultChan chan int) {
	go func() {
		resultChan <- data * 2
	}()
}

func TestProcessAsync_Bad(t *testing.T) {
	resultChan := make(chan int)
	ProcessAsync(5, resultChan)

	// Тест завершится, не дожидаясь результата!
	if result := <-resultChan; result != 10 {
		t.Errorf("Expected 10, got %d", result)
	}
}
```
**Проблема**: В реальном коде горутина может не успеть выполниться.

### Пример 2: Правильный тест с ожиданием

```go
func TestProcessAsync_Good(t *testing.T) {
	resultChan := make(chan int, 1) // Буферизованный канал
	ProcessAsync(5, resultChan)

	select {
	case result := <-resultChan:
		if result != 10 {
			t.Errorf("Expected 10, got %d", result)
		}
	case <-time.After(1 * time.Second): // Таймаут на случай зависания
		t.Error("Timeout waiting for result")
	}
}
```

---

## Как выявлять гонки данных (data race)

Гонки данных возникают, когда несколько горутин одновременно обращаются к одной переменной, и хотя бы одна из них изменяет её.

### Пример 3: Код с гонкой данных

```go
var counter int

func Increment() {
	counter++
}

func TestIncrement_Race(t *testing.T) {
	for i := 0; i < 1000; i++ {
		go Increment()
	}

	time.Sleep(1 * time.Second) // Наивная попытка дождаться завершения
	t.Logf("Counter: %d", counter) // Результат непредсказуем!
}
```
**Проблема**: Тест может выводить разные значения `counter` из-за гонки.

### Способ 1: Запуск теста с детектором гонок

Добавьте флаг `-race` при запуске тестов:
```bash
go test -race ./...
```
Детектор сообщит о проблеме:
```
WARNING: DATA RACE
Write at 0x00000123456 by goroutine 7:
  main.Increment()
```

### Способ 2: Использование `sync.Mutex`

```go
var (
	counter int
	mu      sync.Mutex
)

func IncrementSafe() {
	mu.Lock()
	defer mu.Unlock()
	counter++
}

func TestIncrementSafe(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			IncrementSafe()
		}()
	}
	wg.Wait()

	if counter != 1000 {
		t.Errorf("Expected 1000, got %d", counter)
	}
}
```

---

## Использование `sync.WaitGroup` и `sync.Mutex` в тестах

### Пример 4: `WaitGroup` для ожидания горутин

```go
func ProcessBatch(data []int, resultChan chan int) {
	var wg sync.WaitGroup
	for _, num := range data {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			resultChan <- n * 2
		}(num)
	}
	wg.Wait()
	close(resultChan)
}

func TestProcessBatch(t *testing.T) {
	data := []int{1, 2, 3, 4, 5}
	resultChan := make(chan int, len(data))

	ProcessBatch(data, resultChan)

	var results []int
	for res := range resultChan {
		results = append(results, res)
	}

	expected := []int{2, 4, 6, 8, 10}
	if !reflect.DeepEqual(results, expected) {
		t.Errorf("Expected %v, got %v", expected, results)
	}
}
```

### Пример 5: Тестирование с `Mutex`

Допустим, у нас есть кэш с конкурентным доступом:

```go
type Cache struct {
	mu    sync.Mutex
	items map[string]string
}

func (c *Cache) Set(key, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.items[key] = value
}

func (c *Cache) Get(key string) (string, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	val, ok := c.items[key]
	return val, ok
}
```

Тест:

```go
func TestCache_ConcurrentAccess(t *testing.T) {
	cache := &Cache{items: make(map[string]string)}

	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d", n)
			cache.Set(key, "value")
		}(i)
	}

	wg.Wait()

	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("key%d", i)
		if val, ok := cache.Get(key); !ok || val != "value" {
			t.Errorf("Key %s not found or incorrect value", key)
		}
	}
}
```

---

## Полезные советы

1. **Всегда используйте `-race`**
   ```bash
   go test -race ./...
   ```

2. **Избегайте `time.Sleep` в тестах**  
   Вместо этого используйте `WaitGroup` или каналы.

3. **Тестируйте разные сценарии**
    - Конкурентные запись и чтение
    - Очень большая нагрузка (1000+ горутин)
    - Ошибки и паники в горутинах

4. **Пример теста с паникой**
   ```go
   func TestGoroutinePanic(t *testing.T) {
       var wg sync.WaitGroup
       wg.Add(1)

       go func() {
           defer func() {
               if r := recover(); r == nil {
                   t.Error("Expected panic, got none")
               }
               wg.Done()
           }()
           panic("expected")
       }()

       wg.Wait()
   }
   ```