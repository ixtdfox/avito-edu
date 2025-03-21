# Структура проекта и работа с пакетами в Go

## Flat Structure: плюсы и минусы

### Что такое Flat Structure?
Flat Structure — это организация кода, при которой все файлы проекта находятся в одном уровне или минимально разнесены по директориям.

### Преимущества Flat Structure:
✅ Простота — легко ориентироваться в небольших проектах.  
✅ Быстрое добавление новых файлов без необходимости пересмотра структуры.  
✅ Удобство для скриптов и небольших утилит.

### Недостатки Flat Structure:
❌ Плохая масштабируемость — сложнее поддерживать большие проекты.  
❌ Возможность пересечения имен файлов.  
❌ Нарушение принципа разделения ответственности (код разных слоев смешан).

### Пример Flat Structure:
```
myapp/
├── main.go
├── config.go
├── database.go
├── handler.go
├── service.go
├── repository.go
```

Такой подход может быть удобен для маленьких программ, но быстро становится неуправляемым в масштабных проектах.

---

## Standard Project Layout

**Standard Project Layout** — это устоявшаяся структура проектов Go, рекомендованная сообществом и используемая в продакшн-проектах.

### Преимущества Standard Project Layout:
✅ Четкое разделение ответственности.  
✅ Масштабируемость — легко добавлять новые компоненты.  
✅ Удобство тестирования и поддержки кода.  
✅ Следование best practices, понятных другим разработчикам.

### Пример структуры:
```
myapp/
├── cmd/             # Точки входа в приложение
│   ├── app/         # Основное приложение
│   │   ├── main.go
│   ├── worker/      # Вторичный сервис (например, background worker)
│   │   ├── main.go
├── internal/        # Внутренние пакеты (нельзя импортировать извне)
│   ├── config/      # Конфигурация
│   ├── database/    # Логика работы с базой данных
│   ├── service/     # Бизнес-логика
│   ├── repository/  # Репозитории (работа с БД)
│   ├── handler/     # HTTP-хендлеры
├── pkg/             # Общие пакеты, которые можно переиспользовать
├── api/             # API-спецификации (Swagger, Protobuf)
├── configs/         # Конфигурационные файлы
├── scripts/         # Скрипты для деплоя, миграций и т. д.
├── go.mod           # Модульный файл Go
├── go.sum           # Контрольные суммы зависимостей
```

---

## Разделение бизнес-логики и слоев

### Основные слои в проекте

1. **Handler (Контроллеры)** — обработка входящих HTTP-запросов.
2. **Service (Сервисный слой)** — бизнес-логика.
3. **Repository (Репозиторий)** — взаимодействие с базой данных.

### Почему важно разделять слои?
- Повышает читаемость кода.
- Уменьшает связанность компонентов.
- Упрощает тестирование.

### Пример кода с разделением слоев

#### `handler/user.go` (Контроллер)
```go
package handler

import (
    "net/http"
    "myapp/internal/service"
    "github.com/gin-gonic/gin"
)

type UserHandler struct {
    service service.UserService
}

func NewUserHandler(s service.UserService) *UserHandler {
    return &UserHandler{service: s}
}

func (h *UserHandler) GetUser(c *gin.Context) {
    id := c.Param("id")
    user, err := h.service.GetUser(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }
    c.JSON(http.StatusOK, user)
}
```

#### `service/user.go` (Сервисный слой)
```go
package service

import "myapp/internal/repository"

type UserService struct {
    repo repository.UserRepository
}

func NewUserService(r repository.UserRepository) *UserService {
    return &UserService{repo: r}
}

func (s *UserService) GetUser(id string) (string, error) {
    return s.repo.FindById(id)
}
```

#### `repository/user.go` (Репозиторий)
```go
package repository

type UserRepository struct {}

func NewUserRepository() *UserRepository {
    return &UserRepository{}
}

func (r *UserRepository) FindById(id string) (string, error) {
    return "John Doe", nil
}
```

Этот подход делает код чистым, модульным и легко тестируемым.

---

## Итоги
- Flat Structure удобна для маленьких проектов, но не масштабируется.
- Standard Project Layout — лучший вариант для продакшн-проектов.
- Разделение слоев (Handler, Service, Repository) улучшает читаемость и тестируемость кода.

Следуя этим принципам, ваш Go-проект будет легко поддерживать и расширять. 🚀
