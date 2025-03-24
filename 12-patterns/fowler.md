# Шаблоны корпоративных приложений Мартина Фаулера

## Кто такой Мартин Фаулер

Мартин Фаулер — один из самых влиятельных экспертов в области разработки программного обеспечения, главный научный сотрудник ThoughtWorks. Он автор нескольких классических книг по проектированию ПО, включая:

- "Шаблоны корпоративных приложений" (Patterns of Enterprise Application Architecture)
- "Рефакторинг" (Refactoring)
- "Предметно-ориентированные языки программирования" (Domain-Specific Languages)

Фаулер внес значительный вклад в популяризацию гибких методологий разработки (Agile) и шаблонов проектирования для сложных бизнес-приложений.

## Зачем нужны шаблоны проектирования в корпоративных приложениях

Корпоративные приложения имеют особые требования:
- Сложная бизнес-логика
- Большие объемы данных
- Множество интеграций
- Высокие требования к надежности
- Долгий жизненный цикл

Шаблоны проектирования помогают:
- Организовать сложную логику
- Упростить поддержку
- Обеспечить масштабируемость
- Уменьшить связность компонентов
- Упростить тестирование

### Категории шаблонов корпоративных приложений

1. **Domain Logic Patterns** (Шаблоны предметной области)
2. **Data Source Architectural Patterns** (Шаблоны работы с данными)
3. **Object-Relational Patterns** (Объектно-реляционные шаблоны)
4. **Distribution Patterns** (Шаблоны распределенных систем)

## Domain Logic Patterns (Шаблоны предметной области)

### 1. Domain Model (Модель предметной области)

**Проблема**: Как представить сложную бизнес-логику?

**Решение**: Создать объектную модель, отражающую ключевые концепции предметной области.

```go
type Order struct {
    ID         string
    Customer   Customer
    Items      []OrderItem
    Status     OrderStatus
    CreatedAt  time.Time
}

func (o *Order) CalculateTotal() decimal.Decimal {
    total := decimal.NewFromInt(0)
    for _, item := range o.Items {
        total = total.Add(item.Price.Mul(decimal.NewFromFloat(float64(item.Quantity))))
    }
    return total
}

func (o *Order) Cancel() error {
    if o.Status == Shipped {
        return errors.New("cannot cancel shipped order")
    }
    o.Status = Cancelled
    return nil
}
```

### 2. Transaction Script (Транзакционный сценарий)

**Проблема**: Как организовать простую бизнес-логику?

**Решение**: Каждая операция реализуется как процедура, которая получает данные из БД, выполняет операции и сохраняет результат.

```go
type OrderService struct {
    repo OrderRepository
}

func (s *OrderService) CancelOrder(orderID string) error {
    // 1. Получить данные
    order, err := s.repo.GetByID(orderID)
    if err != nil {
        return err
    }
    
    // 2. Выполнить логику
    if order.Status == Shipped {
        return errors.New("cannot cancel shipped order")
    }
    
    // 3. Сохранить изменения
    order.Status = Cancelled
    return s.repo.Save(order)
}
```

## Data Source Patterns (Шаблоны работы с данными)

### 1. Gateway (Шлюз)

**Table Data Gateway** (Шлюз к таблице):

```go
type ProductGateway interface {
    FindAll() ([]Product, error)
    FindByID(id string) (*Product, error)
    Insert(product *Product) error
    Update(product *Product) error
    Delete(id string) error
}
```

**Row Data Gateway** (Шлюз к строке):

```go
type ProductRowGateway struct {
    ID       string
    Name     string
    Price    decimal.Decimal
    // другие поля
    
    // методы доступа
    Load() error
    Update() error
    Delete() error
}
```

### 2. ActiveRecord (Активная запись)

**Проблема**: Как совместить бизнес-логику и доступ к данным?

**Решение**: Объект содержит и данные, и методы для работы с ними.

```go
type Product struct {
    ID    string
    Name  string
    Price decimal.Decimal
    
    // методы доступа к данным
}

func (p *Product) Save() error {
    if p.ID == "" {
        return db.Insert("products", p)
    }
    return db.Update("products", p)
}

func (p *Product) Delete() error {
    return db.Delete("products", p.ID)
}

func FindProduct(id string) (*Product, error) {
    var p Product
    err := db.Select(&p, "SELECT * FROM products WHERE id = ?", id)
    return &p, err
}
```

### 3. Data Mapper (Преобразователь данных)

**Проблема**: Как отделить бизнес-объекты от структуры БД?

**Решение**: Специальный слой, который переносит данные между объектами и БД.

```go
type ProductMapper struct {
    db *sql.DB
}

func (m *ProductMapper) Find(id string) (*Product, error) {
    row := m.db.QueryRow("SELECT * FROM products WHERE id = ?", id)
    
    var p Product
    err := row.Scan(&p.ID, &p.Name, &p.Price)
    if err != nil {
        return nil, err
    }
    
    return &p, nil
}

func (m *ProductMapper) Insert(p *Product) error {
    _, err := m.db.Exec(
        "INSERT INTO products (id, name, price) VALUES (?, ?, ?)",
        p.ID, p.Name, p.Price,
    )
    return err
}
```

### 4. Repository (Репозиторий)

**Проблема**: Как абстрагировать доступ к данным?

**Решение**: Коллекция объектов в памяти, которая имитирует работу с коллекцией.

```go
type ProductRepository interface {
    FindAll() ([]Product, error)
    FindByID(id string) (*Product, error)
    Add(product *Product) error
    Update(product *Product) error
    Remove(id string) error
}

type MySQLProductRepository struct {
    db *sql.DB
}

func (r *MySQLProductRepository) FindByID(id string) (*Product, error) {
    // реализация поиска в MySQL
}
```

## Object-Relational Patterns (Объектно-реляционные шаблоны)

### 1. Identity Map (Карта идентичностей)

**Проблема**: Как избежать дублирования объектов при загрузке одних и тех же данных?

**Решение**: Хранить загруженные объекты в карте.

```go
type IdentityMap struct {
    products map[string]*Product
    mu       sync.RWMutex
}

func (im *IdentityMap) Get(id string) (*Product, error) {
    im.mu.RLock()
    defer im.mu.RUnlock()
    
    if p, ok := im.products[id]; ok {
        return p, nil
    }
    return nil, ErrNotFound
}

func (im *IdentityMap) Add(p *Product) {
    im.mu.Lock()
    defer im.mu.Unlock()
    
    im.products[p.ID] = p
}
```

### 2. Unit of Work (Единица работы)

**Проблема**: Как отслеживать изменения объектов для эффективного сохранения?

**Решение**: Регистрировать все изменения и выполнять их одной транзакцией.

```go
type UnitOfWork struct {
    newObjects     []interface{}
    dirtyObjects   []interface{}
    removedObjects []interface{}
}

func (uow *UnitOfWork) RegisterNew(obj interface{}) {
    uow.newObjects = append(uow.newObjects, obj)
}

func (uow *UnitOfWork) Commit() error {
    // Выполнить все операции в транзакции
    tx := db.Begin()
    
    for _, obj := range uow.newObjects {
        if err := mapper.Insert(tx, obj); err != nil {
            tx.Rollback()
            return err
        }
    }
    
    // аналогично для измененных и удаленных объектов
    
    return tx.Commit()
}
```

### 3. Lazy Load (Ленивая загрузка)

**Проблема**: Как избежать загрузки связанных объектов, пока они не нужны?

**Решение**: Использовать прокси-объекты, которые загружают данные при первом обращении.

```go
type LazyCustomer struct {
    customer *Customer
    loader   func() (*Customer, error)
}

func (lc *LazyCustomer) Get() (*Customer, error) {
    if lc.customer == nil {
        customer, err := lc.loader()
        if err != nil {
            return nil, err
        }
        lc.customer = customer
    }
    return lc.customer, nil
}

// Использование
lazyCustomer := &LazyCustomer{
    loader: func() (*Customer, error) {
        return customerRepo.FindByID(order.CustomerID)
    },
}
```

## Distribution Patterns (Шаблоны распределенных систем)

### 1. DTO (Data Transfer Object)

**Проблема**: Как передавать данные между процессами/сервисами?

**Решение**: Создавать специальные объекты для передачи данных.

```go
type OrderDTO struct {
    ID         string          `json:"id"`
    CustomerID string          `json:"customerId"`
    Items      []OrderItemDTO  `json:"items"`
    Total      decimal.Decimal `json:"total"`
    Status     string          `json:"status"`
}

func ToOrderDTO(order *Order) OrderDTO {
    items := make([]OrderItemDTO, len(order.Items))
    for i, item := range order.Items {
        items[i] = OrderItemDTO{
            ProductID: item.ProductID,
            Quantity:  item.Quantity,
            Price:     item.Price,
        }
    }
    
    return OrderDTO{
        ID:         order.ID,
        CustomerID: order.CustomerID,
        Items:      items,
        Total:      order.CalculateTotal(),
        Status:     order.Status.String(),
    }
}
```