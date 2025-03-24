# Введение в тестирование в Go

Тестирование — важная часть разработки программного обеспечения. Оно помогает убедиться, что код работает корректно, уменьшает количество ошибок и облегчает поддержку проекта. В Go тестирование встроено в стандартную библиотеку, что делает его простым и удобным.

---

## Почему тестирование важно?

Тестирование позволяет:
- Находить ошибки на ранних этапах.
- Упрощать рефакторинг (если тесты проходят, значит, код работает).
- Документировать поведение функций (тесты показывают, как должен использоваться код).
- Ускорять разработку (автоматические тесты быстрее ручной проверки).

**Пример без тестирования:**
```go
// main.go
package main

func Add(a, b int) int {
    return a + b
}

func main() {
    result := Add(2, 3)
    println(result) // 5
}
```

Если мы случайно изменим Add (например, на return a - b), ошибка останется незамеченной.

## Unit-тесты (тестирование отдельных функций)
Unit-тесты проверяют работу отдельных функций или методов. В Go они пишутся в файлах с суффиксом _test.go.
```go
// math.go
package math

func Add(a, b int) int {
    return a + b
}

```

Тесты:
```go
// math.go
package math

func Add(a, b int) int {
    return a + b
}
```

Запуск теста:
```bash
go test -v
```

##Integration-тесты (проверка взаимодействия компонентов)

```go
// db.go
package db

import "database/sql"

type User struct {
    ID   int
    Name string
}

func GetUser(db *sql.DB, id int) (*User, error) {
    var user User
    err := db.QueryRow("SELECT id, name FROM users WHERE id = ?", id).Scan(&user.ID, &user.Name)
    if err != nil {
        return nil, err
    }
    return &user, nil
}
```

Тесты:
```go
// db.go
package db

import "database/sql"

type User struct {
    ID   int
    Name string
}

func GetUser(db *sql.DB, id int) (*User, error) {
    var user User
    err := db.QueryRow("SELECT id, name FROM users WHERE id = ?", id).Scan(&user.ID, &user.Name)
    if err != nil {
        return nil, err
    }
    return &user, nil
}
```

## End-to-End тесты (проверка системы целиком)
E2E-тесты проверяют работу всей системы, например, HTTP-сервера.

```go
// server.go
package main

import (
    "net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello, World!"))
}

func main() {
    http.HandleFunc("/", HelloHandler)
    http.ListenAndServe(":8080", nil)
}
```

Тесты:

```go
// server_test.go
package main

import (
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestHelloHandler(t *testing.T) {
    req := httptest.NewRequest("GET", "/", nil)
    rr := httptest.NewRecorder()

    HelloHandler(rr, req)

    if rr.Body.String() != "Hello, World!" {
        t.Errorf("Unexpected response: %s", rr.Body.String())
    }
}

func TestServer(t *testing.T) {
    // Запускаем сервер в тесте
    srv := httptest.NewServer(http.HandlerFunc(HelloHandler))
    defer srv.Close()

    // Делаем запрос к серверу
    resp, err := http.Get(srv.URL)
    if err != nil {
        t.Fatal(err)
    }
    defer resp.Body.Close()

    // Проверяем ответ
    if resp.StatusCode != http.StatusOK {
        t.Errorf("Expected status 200, got %d", resp.StatusCode)
    }
}
```



