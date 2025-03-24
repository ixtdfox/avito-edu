# Тестирование работы с временем в Go

Работа с временем в тестах представляет особую сложность, так как функции вроде `time.Now()` и `time.Sleep()` делают тесты недетерминированными. Рассмотрим, как правильно тестировать временнозависимый код.

## Проблемы тестирования временнозависимого кода

### Почему `time.Now()` делает тесты нестабильными

Основные проблемы:
- Реальное время постоянно меняется
- Тесты могут давать разные результаты в разное время суток
- Проверки на конкретные временные значения хрупкие

**Плохой пример:**
```go
func IsMorning() bool {
    hour := time.Now().Hour()
    return hour >= 5 && hour < 12
}

func TestIsMorning(t *testing.T) {
    if !IsMorning() {
        t.Error("Expected morning time")
    }
}
```
Этот тест будет падать ночью и вечером!

## Как тестировать функции, зависящие от времени

### 1. Использование интерфейсов для инъекции зависимостей

Лучший подход - сделать время явной зависимостью:

```go
type Clock interface {
    Now() time.Time
}

type RealClock struct{}

func (RealClock) Now() time.Time {
    return time.Now()
}

func IsMorning(clock Clock) bool {
    hour := clock.Now().Hour()
    return hour >= 5 && hour < 12
}
```

Теперь можно тестировать с mock-часами:

```go
type MockClock struct {
    fixedTime time.Time
}

func (m MockClock) Now() time.Time {
    return m.fixedTime
}

func TestIsMorning(t *testing.T) {
    tests := []struct {
        time time.Time
        want bool
    }{
        {time.Date(2023, 1, 1, 6, 0, 0, 0, time.UTC), true},   // утро
        {time.Date(2023, 1, 1, 15, 0, 0, 0, time.UTC), false},  // день
    }

    for _, tt := range tests {
        clock := MockClock{fixedTime: tt.time}
        got := IsMorning(clock)
        if got != tt.want {
            t.Errorf("IsMorning(%v) = %v, want %v", tt.time, got, tt.want)
        }
    }
}
```

### 2. Тестирование time.Sleep

Для тестирования функций с задержками:

```go
type Sleeper interface {
    Sleep(time.Duration)
}

type RealSleeper struct{}

func (RealSleeper) Sleep(d time.Duration) {
    time.Sleep(d)
}

func Countdown(sleeper Sleeper, n int) {
    for i := n; i > 0; i-- {
        sleeper.Sleep(1 * time.Second)
    }
}
```

Тест с mock-задержкой:

```go
type SpySleeper struct {
    calls int
    durations []time.Duration
}

func (s *SpySleeper) Sleep(d time.Duration) {
    s.calls++
    s.durations = append(s.durations, d)
}

func TestCountdown(t *testing.T) {
    sleeper := &SpySleeper{}
    Countdown(sleeper, 3)
    
    if sleeper.calls != 3 {
        t.Errorf("expected 3 sleeps, got %d", sleeper.calls)
    }
    
    expected := []time.Duration{1*time.Second, 1*time.Second, 1*time.Second}
    if !reflect.DeepEqual(sleeper.durations, expected) {
        t.Errorf("expected sleeps %v, got %v", expected, sleeper.durations)
    }
}
```

## Готовые решения для работы со временем

### Пакет clockwork

Библиотека [clockwork](https://github.com/jonboulle/clockwork) предоставляет готовую реализацию:

```go
import "github.com/jonboulle/clockwork"

func TestWithClockwork(t *testing.T) {
    // Создаем fake clock
    clock := clockwork.NewFakeClock()
    
    // Устанавливаем конкретное время
    clock.Set(time.Date(2023, 1, 1, 10, 0, 0, 0, time.UTC))
    
    // Проверяем утреннее время
    if !IsMorning(clock) {
        t.Error("Expected morning time")
    }
    
    // Можем "переводить" время вперед
    clock.Advance(2 * time.Hour)
    if IsMorning(clock) {
        t.Error("Expected not morning after advance")
    }
}
```

### Пакет testify/suite с временными моками

Для комплексных тестов:

```go
import (
    "testing"
    "github.com/stretchr/testify/suite"
    "github.com/stretchr/testify/mock"
)

type TimeMock struct {
    mock.Mock
}

func (m *TimeMock) Now() time.Time {
    args := m.Called()
    return args.Get(0).(time.Time)
}

type TimeTestSuite struct {
    suite.Suite
    timeMock *TimeMock
}

func (s *TimeTestSuite) SetupTest() {
    s.timeMock = new(TimeMock)
}

func (s *TimeTestSuite) TestIsMorning() {
    morningTime := time.Date(2023, 1, 1, 8, 0, 0, 0, time.UTC)
    s.timeMock.On("Now").Return(morningTime)
    
    s.True(IsMorning(s.timeMock))
    s.timeMock.AssertExpectations(s.T())
}

func TestTimeSuite(t *testing.T) {
    suite.Run(t, new(TimeTestSuite))
}
```

## Рекомендации по тестированию временнозависимого кода

1. **Избегайте прямых вызовов `time.Now()`** в бизнес-логике
2. **Используйте интерфейсы** для работы со временем
3. **Для простых случаев** достаточно мок-структур
4. **Для сложных сценариев** используйте готовые библиотеки вроде clockwork
5. **Тестируйте граничные случаи**:
    - Переход через полночь
    - Летнее время
    - Разные часовые пояса

Пример теста граничного случая:

```go
func TestMidnightTransition(t *testing.T) {
    clock := MockClock{
        fixedTime: time.Date(2023, 1, 1, 23, 59, 59, 0, time.UTC),
    }
    
    // Проверяем поведение перед полуночью
    if IsMorning(clock) {
        t.Error("23:59 should not be morning")
    }
    
    // Переводим время на 1 секунду вперед
    clock.fixedTime = clock.fixedTime.Add(1 * time.Second)
    
    // Теперь полночь - еще не утро
    if IsMorning(clock) {
        t.Error("00:00 should not be morning")
    }
    
    // Переводим на 5 часов вперед
    clock.fixedTime = clock.fixedTime.Add(5 * time.Hour)
    
    // 5 утра - уже утро
    if !IsMorning(clock) {
        t.Error("05:00 should be morning")
    }
}
```

Правильный подход к тестированию времени делает ваши тесты:
- **Надёжными** - не зависят от реального времени
- **Детерминированными** - всегда одинаковый результат
- **Поддерживаемыми** - легко изменять и расширять