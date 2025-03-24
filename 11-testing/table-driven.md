# Table-Driven тесты в Go

## Что такое table-driven тестирование?

Table-Driven тестирование — это подход, при котором тестовые случаи (входные данные и ожидаемые результаты) описываются в виде таблицы (обычно слайса структур), а затем один и тот же тест прогоняется для всех этих случаев в цикле.

**Пример простого теста без table-driven:**
```go
func TestAdd(t *testing.T) {
    if Add(1, 2) != 3 {
        t.Error("1 + 2 != 3")
    }
    if Add(0, 0) != 0 {
        t.Error("0 + 0 != 0")
    }
    if Add(-1, 1) != 0 {
        t.Error("-1 + 1 != 0")
    }
}
```
Тот же тест, но с table-driven подходом:

```go
func TestAdd(t *testing.T) {
    tests := []struct {
        a, b     int
        expected int
    }{
        {1, 2, 3},
        {0, 0, 0},
        {-1, 1, 0},
        {100, -50, 50},
    }

    for _, test := range tests {
        result := Add(test.a, test.b)
        if result != test.expected {
            t.Errorf("Add(%d, %d) = %d, expected %d", 
                test.a, test.b, result, test.expected)
        }
    }
}
```

## Почему table-driven удобнее простого набора if?
### 1. Компактность и читаемость
Все тест-кейсы собраны в одном месте, их легко добавлять/изменять.

**Добавление нового случая:**
```go
tests := []struct {
    a, b     int
    expected int
}{
    {1, 2, 3},
    {0, 0, 0},
    {-1, 1, 0},
    // Новый случай:
    {10, -5, 5}, // Добавили одной строкой
}
```

### 2. Единообразие ошибок
Все ошибки выводятся в одинаковом формате, что упрощает анализ.

**Сравнение выводов:**

```go
// Простой тест:
add_test.go:10: 1 + 2 != 3
add_test.go:13: 0 + 0 != 0

// Table-Driven:
add_test.go:15: Add(1, 2) = 4, expected 3
add_test.go:15: Add(0, 0) = 1, expected 0
```

### 3. Возможность добавить описание тестов
Можно расширить структуру полем name для пояснения:

```go
func TestMultiply(t *testing.T) {
    tests := []struct {
        name     string
        a, b     int
        expected int
    }{
        {"Positive numbers", 2, 3, 6},
        {"With zero", 5, 0, 0},
        {"Negative", -2, 4, -8},
    }

    for _, test := range tests {
        t.Run(test.name, func(t *testing.T) {
            result := Multiply(test.a, test.b)
            if result != test.expected {
                t.Errorf("got %d, want %d", result, test.expected)
            }
        })
    }
}
```

### 4. Легко тестировать edge-кейсы
Можно систематически проверять граничные условия:

```go
func TestSafeDivide(t *testing.T) {
    tests := []struct {
        a, b     int
        expected int
        hasError bool
    }{
        {4, 2, 2, false},
        {1, 0, 0, true},  // Деление на ноль
        {-9, 3, -3, false},
    }

    for _, test := range tests {
        result, err := SafeDivide(test.a, test.b)
        if test.hasError {
            if err == nil {
                t.Errorf("Expected error for %d/%d", test.a, test.b)
            }
        } else {
            if result != test.expected {
                t.Errorf("%d/%d: got %d, want %d", 
                    test.a, test.b, result, test.expected)
            }
        }
    }
}
```

### 5. Поддержка субтестов (subtests)
Каждый случай можно запускать отдельно через t.Run():

```go
func TestParseDate(t *testing.T) {
    tests := []struct {
        name     string
        input    string
        expected time.Time
        hasError bool
    }{
        {"Valid date", "2023-01-15", time.Date(2023, 1, 15, 0, 0, 0, 0, time.UTC), false},
        {"Invalid format", "15-01-2023", time.Time{}, true},
        {"Empty string", "", time.Time{}, true},
    }

    for _, test := range tests {
        t.Run(test.name, func(t *testing.T) {
            result, err := ParseDate(test.input)
            if test.hasError {
                if err == nil {
                    t.Error("Expected error, got nil")
                }
            } else {
                if !result.Equal(test.expected) {
                    t.Errorf("got %v, want %v", result, test.expected)
                }
            }
        })
    }
}
```
**Преимущество:** При запуске go test -v вы увидите каждый подтест отдельно:

### 6. Легко пропускать определенные тесты
Можно добавлять флаги для сложных или долгих тестов:

```go
func TestComplexCalculation(t *testing.T) {
    tests := []struct {
        name     string
        input    int
        expected float64
        skip     bool
    }{
        {"Simple case", 1, 1.0, false},
        {"Edge case", 1000, 123.45, true}, // Пропускаем в обычных прогонах
    }

    for _, test := range tests {
        t.Run(test.name, func(t *testing.T) {
            if test.skip && !testing.Short() {
                t.Skip("Skipping in short mode")
            }
            result := ComplexCalculation(test.input)
            if math.Abs(result-test.expected) > 0.01 {
                t.Errorf("got %.2f, want %.2f", result, test.expected)
            }
        })
    }
}
```

Запуск без пропусков:
```shell
go test -v
```

Запуск только быстрых тестов:
```shell
go test -v -short
```