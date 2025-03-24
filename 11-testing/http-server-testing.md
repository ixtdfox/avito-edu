# Тестирование HTTP-серверов и клиентов в Go

В Go стандартная библиотека предоставляет мощные инструменты для тестирования HTTP-серверов и клиентов. В этой статье разберём, как тестировать обработчики API и как использовать пакет `httptest` для мокирования HTTP-запросов.

---

## Как тестировать обработчики API

Обработчики (handlers) в Go — это функции, которые принимают `http.ResponseWriter` и `http.Request`, обрабатывают запрос и возвращают ответ. Чтобы протестировать их, можно:
1. Создать тестовый HTTP-запрос.
2. Передать его в обработчик.
3. Проверить ответ.

### Пример 1: Простой обработчик

Допустим, у нас есть такой обработчик:

```go
package main

import (
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, World!"))
}
```

Тест для него:

```go
package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/hello", nil)
	rec := httptest.NewRecorder()

	HelloHandler(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, rec.Code)
	}

	expectedBody := "Hello, World!"
	if rec.Body.String() != expectedBody {
		t.Errorf("Expected body '%s', got '%s'", expectedBody, rec.Body.String())
	}
}
```

### Пример 2: Обработчик с параметрами URL

Допустим, обработчик принимает параметр `name`:

```go
func GreetHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Name is required"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, " + name + "!"))
}
```

Тест:

```go
func TestGreetHandler(t *testing.T) {
	tests := []struct {
		name         string
		url          string
		expectedCode int
		expectedBody string
	}{
		{"Valid name", "/greet?name=Alice", http.StatusOK, "Hello, Alice!"},
		{"Empty name", "/greet", http.StatusBadRequest, "Name is required"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", tt.url, nil)
			rec := httptest.NewRecorder()

			GreetHandler(rec, req)

			if rec.Code != tt.expectedCode {
				t.Errorf("Expected status %d, got %d", tt.expectedCode, rec.Code)
			}

			if rec.Body.String() != tt.expectedBody {
				t.Errorf("Expected body '%s', got '%s'", tt.expectedBody, rec.Body.String())
			}
		})
	}
}
```

---

## Использование `httptest` для мокирования HTTP-запросов

Часто HTTP-клиенты взаимодействуют с внешними API, и в тестах нужно эмулировать их поведение. Пакет `net/http/httptest` позволяет создать фейковый сервер (`httptest.Server`), который можно использовать вместо реального API.

### Пример 3: Тестирование HTTP-клиента

Допустим, у нас есть клиент, который делает запрос к API:

```go
func GetUserInfo(apiURL string) (string, error) {
	resp, err := http.Get(apiURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("HTTP error: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
```

Тест с мок-сервером:

```go
func TestGetUserInfo(t *testing.T) {
	// Создаём тестовый сервер
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"name": "Alice", "age": 30}`))
	}))
	defer server.Close() // Важно закрыть сервер после теста

	// Вызываем функцию, передавая URL тестового сервера
	result, err := GetUserInfo(server.URL)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expected := `{"name": "Alice", "age": 30}`
	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}
```

### Пример 4: Тестирование ошибок сервера

Проверим, как клиент обрабатывает ошибки:

```go
func TestGetUserInfo_ServerError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	_, err := GetUserInfo(server.URL)
	if err == nil {
		t.Error("Expected error, got nil")
	}

	expectedErr := "HTTP error: 500"
	if err.Error() != expectedErr {
		t.Errorf("Expected error '%s', got '%v'", expectedErr, err)
	}
}
```

### Пример 5: Мокирование задержки ответа

Иногда API отвечает медленно, и нужно проверить, как клиент обрабатывает таймауты:

```go
func TestGetUserInfo_Timeout(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(2 * time.Second) // Имитируем долгий ответ
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	// Создаём клиент с таймаутом
	client := &http.Client{
		Timeout: 1 * time.Second,
	}

	// Переопределяем клиент (в реальном коде лучше использовать интерфейсы)
	oldClient := http.DefaultClient
	defer func() { http.DefaultClient = oldClient }()
	http.DefaultClient = client

	_, err := GetUserInfo(server.URL)
	if err == nil {
		t.Error("Expected timeout error, got nil")
	}
}
```