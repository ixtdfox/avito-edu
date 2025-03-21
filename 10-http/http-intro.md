# Введение в HTTP-серверы в Go

## Как работает HTTP-протокол
**HTTP (HyperText Transfer Protocol)** — это протокол передачи данных, используемый для взаимодействия между клиентами (браузерами, API-клиентами) и серверами.

### Основные характеристики HTTP:
- **Структура запроса**:
  ```
  GET /index.html HTTP/1.1
  Host: example.com
  User-Agent: Mozilla/5.0
  ```
- **Структура ответа**:
  ```
  HTTP/1.1 200 OK
  Content-Type: text/html
  
  <html><body>Hello, World!</body></html>
  ```
- **Методы HTTP**:
    - `GET` — Получение данных
    - `POST` — Отправка данных
    - `PUT` — Обновление данных
    - `DELETE` — Удаление данных

## Как работает HTTPS-протокол
**HTTPS (HyperText Transfer Protocol Secure)** — это безопасная версия HTTP, использующая **SSL/TLS** для шифрования данных.

### Основные преимущества HTTPS:
✅ Шифрование данных между клиентом и сервером.  
✅ Аутентификация сервера через сертификаты.  
✅ Защита от атак типа MITM (Man-In-The-Middle).

### Как включить HTTPS в Go
```go
package main
import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Secure Hello, World!")
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServeTLS(":443", "server.crt", "server.key", nil)
}
```

## Как Go позволяет разрабатывать веб-серверы
Go имеет встроенный пакет `net/http`, который делает создание HTTP-серверов простым и удобным.

### Основные компоненты:
1. **`http.ListenAndServe`** — Запуск сервера.
2. **`http.HandleFunc`** — Регистрация обработчиков.
3. **`http.ResponseWriter`** — Формирование ответа.
4. **`http.Request`** — Данные запроса.

Простой веб-сервер:
```go
package main
import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello, World!")
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
```

## Почему встроенный пакет `net/http` удобен и прост в использовании
✅ **Минимум зависимостей** — не требуется сторонних библиотек.  
✅ **Легкость в использовании** — простой API.  
✅ **Гибкость** — поддержка middleware, маршрутизации и работы с JSON.  
✅ **Высокая производительность** — эффективная работа с HTTP-запросами.

Пример обработки JSON-запросов:
```go
package main
import (
    "encoding/json"
    "net/http"
)

type Message struct {
    Text string `json:"text"`
}

func jsonHandler(w http.ResponseWriter, r *http.Request) {
    msg := Message{Text: "Hello, JSON!"}
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(msg)
}

func main() {
    http.HandleFunc("/json", jsonHandler)
    http.ListenAndServe(":8080", nil)
}
```
