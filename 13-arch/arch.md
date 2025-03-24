# Архитектурные паттерны в разработке ПО

## Многослойная архитектура (Layered Architecture)

**Классическая трехслойная структура**:
1. **Presentation Layer** (Уровень представления)
    - Обработка HTTP-запросов
    - Валидация входных данных
    - Формирование ответов

2. **Business Logic Layer** (Уровень бизнес-логики)
    - Ядро приложения
    - Основные алгоритмы и правила
    - Не зависит от способа представления и хранения данных

3. **Data Access Layer** (Уровень доступа к данным)
    - Работа с базами данных
    - Внешние API и сервисы
    - Кеширование

### Пример в Go:
```go
// Data Access Layer
type UserRepository interface {
    FindByID(id int) (*User, error)
}

type MySQLUserRepository struct {
    db *sql.DB
}

// Business Layer
type UserService struct {
    repo UserRepository
}

func (s *UserService) GetUserProfile(id int) (*Profile, error) {
    user, err := s.repo.FindByID(id)
    // бизнес-логика
}

// Presentation Layer
type UserHandler struct {
    service *UserService
}

func (h *UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    id, _ := strconv.Atoi(r.URL.Query().Get("id"))
    profile, _ := h.service.GetUserProfile(id)
    json.NewEncoder(w).Encode(profile)
}
```

**Преимущества**:
- Простота понимания
- Четкое разделение ответственности
- Легкость в тестировании

**Недостатки**:
- Риск "утечки" логики между слоями
- Сложности с добавлением новых способов взаимодействия

## Гексагональная архитектура (Hexagonal Architecture / Ports & Adapters)

**Ключевые концепции**:
- **Ядро** (Core) - бизнес-логика
- **Порты** (Ports) - интерфейсы для взаимодействия
- **Адаптеры** (Adapters) - реализация портов

### Пример в Go:
```go
// Порт (интерфейс)
type UserStore interface {
    Save(user *User) error
    FindByID(id string) (*User, error)
}

// Ядро
type UserService struct {
    store UserStore
}

func (s *UserService) Register(user *User) error {
    // валидация и бизнес-правила
    return s.store.Save(user)
}

// Адаптер для MySQL
type MySQLUserAdapter struct {
    db *sql.DB
}

func (a *MySQLUserAdapter) Save(user *User) error {
    // SQL-запрос
}

// Адаптер для gRPC
type UserGRPCHandler struct {
    service *UserService
}

func (h *UserGRPCHandler) Register(ctx context.Context, req *pb.RegisterReq) (*pb.RegisterResp, error) {
    user := convertToDomain(req)
    err := h.service.Register(user)
    // преобразование ответа
}
```

**Преимущества**:
- Полная изоляция бизнес-логики
- Легкая замена адаптеров
- Поддержка множества интерфейсов

**Недостатки**:
- Больше шаблонного кода
- Сложнее для простых приложений

## Onion Architecture

**Концентрические слои**:
1. **Domain Model** (Ядро)
2. **Domain Services**
3. **Application Services**
4. **Infrastructure** (внешние сервисы, UI, DB)

**Правила**:
- Внутренние слои ничего не знают о внешних
- Зависимости направлены только внутрь
- Интерфейсы определяются внутри, реализуются снаружи

### Пример в Go:
```go
// Domain Layer
type User struct {
    ID       string
    Name     string
    Email    string
}

// Domain Service Interface
type UserRepository interface {
    Save(user *User) error
}

// Application Service
type UserRegistration struct {
    repo    UserRepository
    notifier NotificationService
}

func (u *UserRegistration) Register(name, email string) error {
    user := &User{Name: name, Email: email}
    // валидация
    return u.repo.Save(user)
}

// Infrastructure Layer
type SQLUserRepository struct {
    db *sql.DB
}

func (r *SQLUserRepository) Save(user *User) error {
    // SQL-запрос
}
```

**Преимущества**:
- Максимальная изоляция домена
- Легкость замены инфраструктуры
- Удобство тестирования

**Недостатки**:
- Сложность начальной настройки
- Избыточность для простых проектов

## Clean Architecture

**Принципы**:
1. Независимость от фреймворков
2. Тестируемость
3. Независимость от UI
4. Независимость от БД
5. Независимость от внешних сервисов

**Слои**:
- **Entities** (Бизнес-правила)
- **Use Cases** (Сценарии применения)
- **Interface Adapters**
- **Frameworks & Drivers**

### Пример в Go:
```go
// Entity
type User struct {
    ID    string
    Name  string
    Email string
}

// Use Case
type UserRegister interface {
    Register(user *User) error
}

type UserInteractor struct {
    userRepo UserRepository
}

func (i *UserInteractor) Register(user *User) error {
    // бизнес-правила
    return i.userRepo.Store(user)
}

// Interface
type UserRepository interface {
    Store(user *User) error
}

// Adapter
type UserHTTPHandler struct {
    interactor UserRegister
}

func (h *UserHTTPHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    var user User
    json.NewDecoder(r.Body).Decode(&user)
    h.interactor.Register(&user)
    // ответ
}

// Infrastructure
type PostgreSQLUserRepo struct {
    db *sql.DB
}

func (r *PostgreSQLUserRepo) Store(user *User) error {
    // SQL-запрос
}
```

**Преимущества**:
- Полная независимость бизнес-логики
- Легкость тестирования
- Гибкость в выборе технологий
- Долгосрочная поддерживаемость

**Недостатки**:
- Высокий порог входа
- Избыточность для простых CRUD-приложений
- Большое количество промежуточных слоев

## Сравнение архитектур

| Критерий               | Многослойная | Гексагональная | Onion | Clean |
|------------------------|--------------|----------------|-------|-------|
| Сложность реализации   | Низкая       | Средняя        | Высокая | Очень высокая |
| Гибкость               | Ограниченная | Высокая        | Очень высокая | Максимальная |
| Тестируемость          | Хорошая      | Отличная       | Отличная | Идеальная |
| Скорость разработки    | Средняя      | Средняя        | Низкая | Очень низкая |

## Практические рекомендации для Go

1. **Начинайте с простого**: Для небольших проектов достаточно слоистой архитектуры
2. **Используйте интерфейсы**: Это основа для любой чистой архитектуры
3. **Избегайте зависимостей**: Особенно от глобальных переменных
4. **Dependency Injection**: Позволяет легко менять реализации
5. **Тестируемость**: Архитектура должна позволять легко подменять адаптеры

Пример организации кода для Clean Architecture:
```
project/
├── cmd/              # Точки входа
├── internal/
│   ├── entity/       # Сущности
│   ├── usecase/      # Сценарии использования
│   ├── repository/   # Интерфейсы репозиториев
│   └── delivery/     # Доставка (HTTP, gRPC и т.д.)
└── pkg/              
    └── postgres/     # Реализация репозитория
```

Выбор архитектуры зависит от:
- Сложности предметной области
- Требований к гибкости
- Ожидаемого срока жизни проекта
- Опыта команды

Для большинства Go-проектов оптимальны:
- Упрощенная слоистая архитектура (для простых сервисов)
- Гексагональная архитектура (для сервисов со сложной логикой)
- Clean Architecture (для долгоживущих стратегических проектов)