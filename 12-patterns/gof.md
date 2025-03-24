# Паттерны проектирования GoF в Go

## Зачем нужны паттерны GoF и какие проблемы они решают

Паттерны проектирования из книги "Банды четырёх" (GoF) решают несколько ключевых проблем:

1. **Повторное использование проверенных решений** - вместо изобретения велосипедов
2. **Стандартизация подхода** - общий язык для разработчиков
3. **Решения типовых проблем** ООП:
    - Создание объектов
    - Организация взаимодействия между объектами
    - Гибкое изменение поведения

Основные преимущества:
- Ускорение разработки
- Улучшение читаемости кода
- Облегчение поддержки
- Снижение количества ошибок

## Категории паттернов GoF

Паттерны делятся на 3 основные категории:

1. **Порождающие** (Creational) - создание объектов:
    - Одиночка (Singleton)
    - Фабрика (Factory)
    - Абстрактная фабрика (Abstract Factory)
    - Строитель (Builder)
    - Прототип (Prototype)

2. **Структурные** (Structural) - композиция объектов:
    - Адаптер (Adapter)
    - Декоратор (Decorator)
    - Фасад (Facade)
    - Заместитель (Proxy)
    - Мост (Bridge)
    - Компоновщик (Composite)
    - Приспособленец (Flyweight)

3. **Поведенческие** (Behavioral) - взаимодействие объектов:
    - Стратегия (Strategy)
    - Наблюдатель (Observer)
    - Команда (Command)
    - Итератор (Iterator)
    - Состояние (State)
    - Цепочка обязанностей (Chain of Responsibility)
    - Посредник (Mediator)
    - Хранитель (Memento)
    - Шаблонный метод (Template Method)
    - Посетитель (Visitor)
    - Интерпретатор (Interpreter)

## Реализация основных паттернов в Go

### 1. Адаптер (Adapter)

**Проблема**: Несовместимость интерфейсов

**Решение**: Преобразует интерфейс одного класса в интерфейс, ожидаемый клиентом

```go
// Целевой интерфейс
type PaymentGateway interface {
    ProcessPayment(amount float64) error
}

// Старая система (несовместимый интерфейс)
type LegacyPaymentSystem struct{}

func (l *LegacyPaymentSystem) MakePayment(amount float64, currency string) error {
    // реализация
    return nil
}

// Адаптер
type PaymentAdapter struct {
    legacy *LegacyPaymentSystem
}

func (a *PaymentAdapter) ProcessPayment(amount float64) error {
    return a.legacy.MakePayment(amount, "USD")
}

// Использование
func main() {
    legacy := &LegacyPaymentSystem{}
    adapter := &PaymentAdapter{legacy: legacy}
    adapter.ProcessPayment(100.0)
}
```

### 2. Декоратор (Decorator)

**Проблема**: Добавление функциональности без изменения класса

**Решение**: Обёртывание объекта с добавлением поведения

```go
type Coffee interface {
    Cost() float64
    Description() string
}

type SimpleCoffee struct{}

func (c *SimpleCoffee) Cost() float64 { return 1.0 }
func (c *SimpleCoffee) Description() string { return "Simple coffee" }

// Декоратор для молока
type MilkDecorator struct {
    coffee Coffee
}

func (m *MilkDecorator) Cost() float64 {
    return m.coffee.Cost() + 0.5
}

func (m *MilkDecorator) Description() string {
    return m.coffee.Description() + ", milk"
}

// Использование
func main() {
    coffee := &SimpleCoffee{}
    coffeeWithMilk := &MilkDecorator{coffee: coffee}
    
    fmt.Println(coffeeWithMilk.Description()) // "Simple coffee, milk"
    fmt.Println(coffeeWithMilk.Cost())        // 1.5
}
```

### 3. Фасад (Facade)

**Проблема**: Сложная подсистема с множеством компонентов

**Решение**: Упрощённый интерфейс для работы с подсистемой

```go
// Сложная подсистема
type CPU struct{}
func (c *CPU) Execute() { fmt.Println("CPU execute") }

type Memory struct{}
func (m *Memory) Load() { fmt.Println("Memory load") }

type HardDrive struct{}
func (h *HardDrive) Read() { fmt.Println("HardDrive read") }

// Фасад
type ComputerFacade struct {
    cpu       *CPU
    memory    *Memory
    hardDrive *HardDrive
}

func NewComputerFacade() *ComputerFacade {
    return &ComputerFacade{
        cpu:       &CPU{},
        memory:    &Memory{},
        hardDrive: &HardDrive{},
    }
}

func (c *ComputerFacade) Start() {
    c.cpu.Execute()
    c.memory.Load()
    c.hardDrive.Read()
}

// Использование
func main() {
    computer := NewComputerFacade()
    computer.Start()
}
```

### 4. Одиночка (Singleton)

**Проблема**: Гарантия единственного экземпляра класса

**Решение**: Контроль создания экземпляра

```go
type Database struct {
    connection string
}

var (
    instance *Database
    once     sync.Once
)

func GetDatabase() *Database {
    once.Do(func() {
        instance = &Database{connection: "mysql:3306"}
    })
    return instance
}

// Использование
func main() {
    db1 := GetDatabase()
    db2 := GetDatabase()
    
    fmt.Println(db1 == db2) // true
}
```

### 5. Фабрика (Factory)

**Проблема**: Создание объектов без указания конкретного класса

**Решение**: Вынесение логики создания в отдельный метод

```go
type PaymentMethod interface {
    Pay(amount float64) string
}

type CreditCard struct{}
func (c *CreditCard) Pay(amount float64) string {
    return fmt.Sprintf("Paid %.2f via Credit Card", amount)
}

type PayPal struct{}
func (p *PayPal) Pay(amount float64) string {
    return fmt.Sprintf("Paid %.2f via PayPal", amount)
}

func GetPaymentMethod(method string) (PaymentMethod, error) {
    switch method {
    case "credit":
        return &CreditCard{}, nil
    case "paypal":
        return &PayPal{}, nil
    default:
        return nil, fmt.Errorf("unknown payment method")
    }
}

// Использование
func main() {
    payment, _ := GetPaymentMethod("credit")
    fmt.Println(payment.Pay(100.0))
}
```

### 6. Стратегия (Strategy)

**Проблема**: Вариативное поведение алгоритма

**Решение**: Инкапсуляция алгоритмов в отдельные классы

```go
type SortStrategy interface {
    Sort([]int) []int
}

type BubbleSort struct{}
func (b *BubbleSort) Sort(data []int) []int {
    // реализация пузырьковой сортировки
    return data
}

type QuickSort struct{}
func (q *QuickSort) Sort(data []int) []int {
    // реализация быстрой сортировки
    return data
}

type Sorter struct {
    strategy SortStrategy
}

func (s *Sorter) SetStrategy(strategy SortStrategy) {
    s.strategy = strategy
}

func (s *Sorter) Sort(data []int) []int {
    return s.strategy.Sort(data)
}

// Использование
func main() {
    data := []int{3, 1, 4, 1, 5, 9, 2, 6}
    
    sorter := &Sorter{}
    
    sorter.SetStrategy(&BubbleSort{})
    result := sorter.Sort(data)
    
    sorter.SetStrategy(&QuickSort{})
    result = sorter.Sort(data)
}
```

## Особенности реализации в Go

В отличие от классических ООП-языков, в Go паттерны реализуются с учётом:
- Отсутствия классического наследования
- Наличия интерфейсов
- Встроенной поддержки композиции
- Возможности возвращать несколько значений

**Советы по реализации**:
1. Используйте интерфейсы для определения стратегий
2. Применяйте композицию вместо наследования
3. Используйте замыкания для упрощённых стратегий
4. Для фабрик возвращайте интерфейсы, а не конкретные типы

Паттерны в Go часто выглядят проще, чем в классических ООП-языках, но решают те же самые проблемы проектирования.