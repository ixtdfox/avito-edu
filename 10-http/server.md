# Организация кода и запуск сервера в Go

## Как структурировать код сервера

Хорошая организация кода упрощает поддержку, масштабируемость и тестирование веб-приложений. В небольших проектах можно использовать минималистичный подход, но для более сложных решений рекомендуется четкая архитектура.

### Базовая структура проекта:
```
project/
├── cmd/             # Главные исполняемые файлы
│   ├── server/      # Основной сервер
│   │   ├── main.go
├── internal/        # Внутренние модули приложения
│   ├── handlers/    # HTTP-обработчики
│   ├── routes/      # Определение маршрутов
│   ├── services/    # Бизнес-логика
│   ├── repository/  # Доступ к данным (БД, файлы и т.д.)
│   ├── config/      # Настройки приложения
├── pkg/             # Общие пакеты, которые можно переиспользовать
├── go.mod           # Файл модуля Go
├── go.sum           # Контрольные суммы зависимостей
```

Такой подход помогает разделить код по функциональным зонам, улучшая читаемость и поддержку.

## Разделение API на пакеты

### Пакет `handlers`
Этот пакет содержит HTTP-обработчики, которые принимают запросы и формируют ответы.

**Пример файла `internal/handlers/user.go`**:
```go
package handlers

import (
    "encoding/json"
    "net/http"
    "project/internal/services"
)

type UserHandler struct {
    Service *services.UserService
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
    id := r.URL.Query().Get("id")
    user, err := h.Service.GetUserByID(id)
    if err != nil {
        http.Error(w, "User not found", http.StatusNotFound)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(user)
}
```

### Пакет `services`
Этот пакет содержит бизнес-логику и работу с данными.

**Пример файла `internal/services/user.go`**:
```go
package services

type UserService struct {}

func (s *UserService) GetUserByID(id string) (map[string]string, error) {
    // В реальном приложении здесь будет запрос в базу данных
    if id == "1" {
        return map[string]string{"id": "1", "name": "John Doe"}, nil
    }
    return nil, fmt.Errorf("user not found")
}
```

### Пакет `routes`
Этот пакет организует маршруты и подключает обработчики.

**Пример файла `internal/routes/router.go`**:
```go
package routes

import (
    "net/http"
    "project/internal/handlers"
)

func NewRouter(userHandler *handlers.UserHandler) *http.ServeMux {
    mux := http.NewServeMux()
    mux.HandleFunc("/user", userHandler.GetUser)
    return mux
}
```

## Использование `http.NewServeMux()` для организации маршрутов
Встроенный `http.NewServeMux()` позволяет эффективно управлять маршрутами и упрощает поддержку кода.

### Запуск сервера с маршрутизатором
**Пример файла `cmd/server/main.go`**:
```go
package main

import (
    "fmt"
    "net/http"
    "project/internal/handlers"
    "project/internal/routes"
    "project/internal/services"
)

func main() {
    userService := &services.UserService{}
    userHandler := &handlers.UserHandler{Service: userService}
    router := routes.NewRouter(userHandler)

    fmt.Println("Starting server on :8080")
    err := http.ListenAndServe(":8080", router)
    if err != nil {
        fmt.Println("Error starting server:", err)
    }
}
```

## Итог
- **Структурирование кода** улучшает читаемость и поддержку.
- **Разделение API на пакеты** помогает избежать нагромождения кода в одном файле.
- **Использование `http.NewServeMux()`** делает маршрутизацию гибкой и удобной.