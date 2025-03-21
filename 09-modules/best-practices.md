# Разбор best practices и ошибок в Go

## Как правильно именовать пакеты в Go

### Основные рекомендации:
1. **Используйте осмысленные и короткие имена**
    - ❌ `mypackageutils` → ✅ `utils`
    - ❌ `mymodulehelpers` → ✅ `helpers`

2. **Не дублируйте имя пакета в названиях его содержимого**
   ```go
   // Плохо
   package config
   func ConfigLoad() {}

   // Хорошо
   package config
   func Load() {}
   ```

3. **Следуйте общепринятым названиям**
    - `config`, `logger`, `auth`, `handler`, `repository`

## Использование `internal` пакетов для ограничения видимости

Папка `internal/` запрещает импорт ее содержимого вне основного модуля.

```sh
project/
│── internal/
│   ├── db/
│   │   ├── connection.go
│── cmd/
│── main.go
```

Пример использования:
```go
package db
import "fmt"
func connect() {
    fmt.Println("Connecting to DB...")
}
```
Вне `internal/db`, импорт `project/internal/db` вызовет ошибку.

## Циклические зависимости: как их обнаружить и исправить

**Ошибка:**
- `package A` импортирует `package B`, а `package B` импортирует `package A`

**Как исправить:**
1. Вынести общие зависимости в третий пакет
2. Использовать интерфейсы

```go
package storage

type Storage interface {
    Save(data string)
}
```

## Использование интерфейсов для уменьшения связности

Интерфейсы позволяют зависеть от абстракций, а не конкретных реализаций:

```go
type Repository interface {
    GetUser(id int) string
}
```

Вместо конкретной структуры:
```go
type PostgresRepo struct{}
func (p PostgresRepo) GetUser(id int) string {
    return "User"
}
```

## Неправильное использование глобальных переменных в пакетах

Глобальные переменные могут привести к непредсказуемому поведению:
```go
package config
var DBConnection string
```

### Как исправить:
Использовать `sync.Once` для инициализации:
```go
package config
import "sync"
var (
    dbConnection string
    once sync.Once
)
func GetDBConnection() string {
    once.Do(func() {
        dbConnection = "Initialized"
    })
    return dbConnection
}
```

## Ошибки при работе с видимостью (exported/unexported)

В Go идентификаторы, начинающиеся с заглавной буквы, экспортируются:
```go
// Публичная функция
func PublicFunction() {}
// Приватная функция
func privateFunction() {}
```

### Ошибка:
- Попытка импортировать `privateFunction` в другом пакете

## Игнорирование семантического версионирования в модулях

Всегда фиксируйте версию зависимостей:
```sh
$ go get example.com/mypackage@v1.2.3
```

## Использование устаревших или небезопасных пакетов

Проверяйте зависимости:
```sh
$ go list -m -u all
```
Анализируйте уязвимости:
```sh
$ go list -m -json all | govulncheck
