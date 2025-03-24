# Dependency Injection (DI) и Inversion of Control (IoC) в Go

## Что такое Dependency Injection (DI)

Dependency Injection (внедрение зависимостей) — это паттерн проектирования, при котором зависимости объекта не создаются внутри него, а передаются извне. Это делает код более гибким, тестируемым и поддерживаемым.

**Основные виды DI:**
1. **Constructor Injection** - зависимости передаются через конструктор
2. **Method Injection** - зависимости передаются через метод
3. **Property Injection** - зависимости устанавливаются через свойства (поля структуры)

### Пример Constructor Injection в Go:

```go
type Database interface {
    Query(query string) ([]byte, error)
}

type MySQLDatabase struct{}

func (m *MySQLDatabase) Query(query string) ([]byte, error) {
    // реализация запроса к MySQL
    return []byte("result"), nil
}

type Service struct {
    db Database
}

// Конструктор с внедрением зависимости
func NewService(db Database) *Service {
    return &Service{db: db}
}

func (s *Service) GetData() ([]byte, error) {
    return s.db.Query("SELECT * FROM data")
}

// Использование
func main() {
    db := &MySQLDatabase{}
    service := NewService(db)
    data, _ := service.GetData()
    fmt.Println(string(data))
}
```

## Почему в Go нет встроенного DI и как его реализовать

### Причины отсутствия встроенного DI в Go:
1. **Философия простоты** - Go предпочитает явные решения магическим
2. **Отсутствие классов** - DI-контейнеры часто используют рефлексию, что противоречит идеологии Go
3. **Композиция вместо наследования** - Go поощряет явную композицию структур

### Способы реализации DI в Go:

#### 1. Ручное внедрение (наиболее идиоматичный способ)

```go
type Logger interface {
    Log(message string)
}

type ConsoleLogger struct{}

func (c *ConsoleLogger) Log(message string) {
    fmt.Println(message)
}

type App struct {
    logger Logger
}

func NewApp(logger Logger) *App {
    return &App{logger: logger}
}

func main() {
    logger := &ConsoleLogger{}
    app := NewApp(logger)
    app.logger.Log("Application started")
}
```

#### 2. Использование DI-контейнеров (библиотеки)

Популярные библиотеки:
- [google/wire](https://github.com/google/wire)
- [uber-go/dig](https://github.com/uber-go/dig)
- [facebookgo/inject](https://github.com/facebookgo/inject)

Пример с **google/wire**:

```go
// +build wireinject

package main

import "github.com/google/wire"

type Message string

func NewMessage() Message {
    return "Hello, Wire!"
}

func NewGreeter(m Message) *Greeter {
    return &Greeter{Message: m}
}

type Greeter struct {
    Message Message
}

func (g *Greeter) Greet() string {
    return string(g.Message)
}

func InitializeGreeter() *Greeter {
    wire.Build(NewGreeter, NewMessage)
    return &Greeter{}
}

func main() {
    g := InitializeGreeter()
    fmt.Println(g.Greet())
}
```

#### 3. Функциональные опции (Functional Options)

```go
type Server struct {
    host    string
    port    int
    timeout time.Duration
}

type Option func(*Server)

func WithHost(host string) Option {
    return func(s *Server) {
        s.host = host
    }
}

func WithPort(port int) Option {
    return func(s *Server) {
        s.port = port
    }
}

func NewServer(opts ...Option) *Server {
    s := &Server{
        host:    "localhost",
        port:    8080,
        timeout: 30 * time.Second,
    }
    
    for _, opt := range opts {
        opt(s)
    }
    
    return s
}

func main() {
    server := NewServer(
        WithHost("example.com"),
        WithPort(9000),
    )
    fmt.Printf("%+v\n", server)
}
```

## Понятие Inversion of Control (IoC)

Inversion of Control (Инверсия управления) — более общий принцип, частью которого является DI. IoC означает, что управление программой передается внешнему фреймворку или контейнеру, а не реализуется непосредственно в коде.

### Разница между DI и IoC:
- **DI** — это техника реализации (как)
- **IoC** — это принцип проектирования (что)

### Пример IoC в Go:

Без IoC:
```go
type App struct {
    logger *Logger
}

func NewApp() *App {
    return &App{
        logger: &Logger{}, // создание зависимости внутри
    }
}
```

С IoC:
```go
type App struct {
    logger Logger
}

func NewApp(logger Logger) *App { // зависимость передается извне
    return &App{
        logger: logger,
    }
}
```

### Реализация IoC контейнера (упрощенный вариант):

```go
type Container struct {
    services map[string]interface{}
}

func NewContainer() *Container {
    return &Container{
        services: make(map[string]interface{}),
    }
}

func (c *Container) Register(name string, service interface{}) {
    c.services[name] = service
}

func (c *Container) Get(name string) (interface{}, error) {
    service, ok := c.services[name]
    if !ok {
        return nil, fmt.Errorf("service %s not found", name)
    }
    return service, nil
}

func main() {
    container := NewContainer()
    container.Register("logger", &ConsoleLogger{})
    
    logger, _ := container.Get("logger")
    app := NewApp(logger.(Logger))
    app.Run()
}
```

## Практические рекомендации по DI в Go

1. **Используйте интерфейсы** для зависимостей
2. **Предпочитайте явное внедрение** через конструкторы
3. **Избегайте глобальных состояний** - они усложняют тестирование
4. **Для сложных приложений** рассмотрите DI-библиотеки
5. **Не злоупотребляйте рефлексией** - она снижает читаемость кода

Пример тестируемого кода с DI:

```go
type UserRepository interface {
    GetUser(id int) (*User, error)
}

type RealUserRepository struct {
    db *sql.DB
}

func (r *RealUserRepository) GetUser(id int) (*User, error) {
    // запрос к реальной БД
}

type MockUserRepository struct{}

func (m *MockUserRepository) GetUser(id int) (*User, error) {
    return &User{ID: id, Name: "Test User"}, nil
}

type UserService struct {
    repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
    return &UserService{repo: repo}
}

func TestUserService(t *testing.T) {
    mockRepo := &MockUserRepository{}
    service := NewUserService(mockRepo)
    
    user, err := service.repo.GetUser(1)
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }
    
    if user.Name != "Test User" {
        t.Errorf("expected Test User, got %s", user.Name)
    }
}
```

## Заключение

Хотя в Go нет встроенного DI, принципы внедрения зависимостей и инверсии управления могут и должны применяться:
- Для повышения тестируемости кода
- Уменьшения связанности компонентов
- Упрощения поддержки и модификации
- Реализации гибкой архитектуры

Выбор конкретного подхода (ручное внедрение, DI-библиотеки или функциональные опции) зависит от сложности проекта и предпочтений команды.