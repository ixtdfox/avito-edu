# Работа с JSON и хендлерами в Go

В языке Go для работы с JSON используется стандартный пакет `encoding/json`. Он позволяет легко кодировать (маршалить) и декодировать (анмаршалить) данные в формате JSON.

---

## Кодирование и декодирование JSON (encoding/json)

### Кодирование (маршаллинг) JSON
Для преобразования структуры Go в JSON используется `json.Marshal`.

#### Пример:
```go
package main

import (
    "encoding/json"
    "fmt"
)

type User struct {
    Name  string `json:"name"`
    Age   int    `json:"age"`
    Email string `json:"email"`
}

func main() {
    user := User{Name: "John Doe", Age: 30, Email: "john@example.com"}
    jsonData, err := json.Marshal(user)
    if err != nil {
        fmt.Println("Error encoding JSON:", err)
        return
    }
    fmt.Println(string(jsonData))
}
```

### Декодирование (анмаршаллинг) JSON
Для преобразования JSON в структуру Go используется `json.Unmarshal`.

#### Пример:
```go
package main

import (
    "encoding/json"
    "fmt"
)

type User struct {
    Name  string `json:"name"`
    Age   int    `json:"age"`
    Email string `json:"email"`
}

func main() {
    jsonString := `{"name":"Alice","age":25,"email":"alice@example.com"}`
    var user User
    err := json.Unmarshal([]byte(jsonString), &user)
    if err != nil {
        fmt.Println("Error decoding JSON:", err)
        return
    }
    fmt.Println(user)
}
```

---

## `http.Request.Body` для обработки входящих данных

Когда клиент отправляет JSON-данные на сервер, их можно прочитать из `http.Request.Body` и декодировать с помощью `json.Decoder`.

### Пример обработки JSON-запроса:
```go
package main

import (
    "encoding/json"
    "fmt"
    "io"
    "net/http"
)

type User struct {
    Name  string `json:"name"`
    Age   int    `json:"age"`
    Email string `json:"email"`
}

func userHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    var user User
    body, err := io.ReadAll(r.Body)
    if err != nil {
        http.Error(w, "Error reading request body", http.StatusInternalServerError)
        return
    }
    defer r.Body.Close()

    err = json.Unmarshal(body, &user)
    if err != nil {
        http.Error(w, "Invalid JSON format", http.StatusBadRequest)
        return
    }
    fmt.Fprintf(w, "Received user: %+v", user)
}

func main() {
    http.HandleFunc("/user", userHandler)
    http.ListenAndServe(":8080", nil)
}
```

Этот сервер принимает JSON-запрос `POST /user` и возвращает данные пользователю.

---

## `http.ResponseWriter` для отправки JSON-ответов

При работе с API важно правильно форматировать JSON-ответы и указывать корректные заголовки.

### Пример JSON-ответа:
```go
package main

import (
    "encoding/json"
    "net/http"
)

type Response struct {
    Message string `json:"message"`
}

func jsonResponseHandler(w http.ResponseWriter, r *http.Request) {
    response := Response{Message: "Hello, JSON!"}
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(response)
}

func main() {
    http.HandleFunc("/json", jsonResponseHandler)
    http.ListenAndServe(":8080", nil)
}
```

### Что здесь важно:
- Устанавливаем заголовок `Content-Type: application/json`.
- Используем `json.NewEncoder(w).Encode(response)`, чтобы избежать лишнего создания буферов.
- Устанавливаем код ответа с `w.WriteHeader(http.StatusOK)` (по умолчанию 200 OK).

---

## Итог

- `encoding/json` позволяет легко кодировать и декодировать JSON в Go.
- `http.Request.Body` используется для чтения входящих JSON-данных.
- `http.ResponseWriter` отвечает за отправку JSON-ответов клиентам.
- `json.NewEncoder` и `json.NewDecoder` — удобные методы для работы с JSON.

