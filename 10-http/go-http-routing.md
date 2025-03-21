# Обработка маршрутов и запросов в Go

## `http.HandleFunc()` vs `http.NewServeMux()`

### `http.HandleFunc()`
Это самый простой способ обработки HTTP-запросов. Вы регистрируете обработчик для конкретного маршрута и передаёте его в `http.ListenAndServe()`.

#### Пример использования:
```go
package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello, world!")
}

func main() {
    http.HandleFunc("/", handler) // Регистрируем обработчик
    http.ListenAndServe(":8080", nil) // Запускаем сервер
}
```
✅ **Плюсы:** Простота использования, не требует дополнительного кода.
❌ **Минусы:** Все маршруты хранятся в глобальном реестре, что делает код менее гибким.

---

### `http.NewServeMux()`
Этот подход даёт больше контроля над маршрутизацией. `ServeMux` — это мультиплексор (роутер), который управляет маршрутами.

#### Пример использования `http.NewServeMux()`:
```go
package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello from ServeMux!")
}

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", handler) // Регистрация обработчика

    http.ListenAndServe(":8080", mux) // Запуск сервера с роутером
}
```
✅ **Плюсы:** Позволяет лучше структурировать код, использовать вложенные маршруты и middleware.
❌ **Минусы:** Чуть больше кода по сравнению с `http.HandleFunc()`.

---

## Методы HTTP-запросов (GET, POST, PUT, DELETE)

HTTP использует различные методы для взаимодействия с сервером. В Go их можно обрабатывать с помощью `r.Method`.

### Пример обработки различных методов:
```go
package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        fmt.Fprintln(w, "This is a GET request")
    case http.MethodPost:
        fmt.Fprintln(w, "This is a POST request")
    case http.MethodPut:
        fmt.Fprintln(w, "This is a PUT request")
    case http.MethodDelete:
        fmt.Fprintln(w, "This is a DELETE request")
    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

func main() {
    http.HandleFunc("/api", handler)
    http.ListenAndServe(":8080", nil)
}
```

### Краткое описание HTTP-методов:
- **GET** – Получает данные с сервера (например, список пользователей).
- **POST** – Отправляет данные на сервер (например, создаёт нового пользователя).
- **PUT** – Обновляет существующую сущность.
- **DELETE** – Удаляет сущность.

---

## Чтение параметров из запроса

### Чтение Query-параметров (`r.URL.Query()`)

Запрос: `GET /search?query=golang&page=2`

```go
package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    query := r.URL.Query().Get("query")
    page := r.URL.Query().Get("page")
    fmt.Fprintf(w, "Search query: %s, Page: %s", query, page)
}

func main() {
    http.HandleFunc("/search", handler)
    http.ListenAndServe(":8080", nil)
}
```

### Чтение данных из тела запроса (`r.Body`)

При `POST`-запросах данные передаются в теле запроса.

```go
package main

import (
    "encoding/json"
    "fmt"
    "io"
    "net/http"
)

type RequestData struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

func handler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Only POST is allowed", http.StatusMethodNotAllowed)
        return
    }

    body, err := io.ReadAll(r.Body)
    if err != nil {
        http.Error(w, "Cannot read body", http.StatusInternalServerError)
        return
    }
    defer r.Body.Close()

    var data RequestData
    json.Unmarshal(body, &data)
    fmt.Fprintf(w, "Received: %+v", data)
}

func main() {
    http.HandleFunc("/submit", handler)
    http.ListenAndServe(":8080", nil)
}
```

Запрос с `curl`:
```sh
curl -X POST -H "Content-Type: application/json" -d '{"name": "Alice", "age": 30}' http://localhost:8080/submit
```

---

## Итог
✅ `http.HandleFunc()` прост, но `http.NewServeMux()` даёт больше гибкости.
✅ Разные HTTP-методы используются для разных операций.
✅ Query-параметры извлекаются через `r.URL.Query()`.
✅ Данные из `POST`-запросов читаются из `r.Body`.
