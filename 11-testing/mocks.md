markdown
Copy
# Использование моков и стабов в тестах Go

## Разница между моками и стабами

**Стабы (Stubs)** - простые заглушки, которые возвращают предопределённые значения:
```go
type UserRepositoryStub struct{}

func (u *UserRepositoryStub) GetUser(id int) (*User, error) {
    return &User{ID: 1, Name: "Test User"}, nil // Всегда возвращает одного и того же пользователя
}
```

**Моки (Mocks)** - более умные заглушки, которые дополнительно проверяют, как они были вызваны:

```go
type UserRepositoryMock struct {
    mock.Mock
}

func (u *UserRepositoryMock) GetUser(id int) (*User, error) {
    args := u.Called(id) // Фиксируем факт вызова
    return args.Get(0).(*User), args.Error(1) // Возвращаем то, что настроили в тесте
}

```
Ключевые различия:
 - Стабы только возвращают данные
 - Моки ещё и проверяют ожидания (был ли вызов, с какими параметрами, сколько раз)

## Как мокировать внешние зависимости

Пример без моков (проблема - реальное обращение к БД):
```go
func TestGetUserEmail(t *testing.T) {
    db := realDatabaseConnection() // Плохо для unit-теста!
    email, err := GetUserEmail(db, 1)
    if err != nil {
        t.Fatal(err)
    }
    if email != "user@example.com" {
        t.Errorf("unexpected email: %s", email)
    }
}
```

Решение с моком:
```go
func TestGetUserEmail(t *testing.T) {
    // 1. Создаём мок
    dbMock := new(DatabaseMock)
    
    // 2. Настраиваем ожидания
    dbMock.On("GetUser", 1).Return(&User{Email: "user@example.com"}, nil)
    
    // 3. Тестируем
    email, err := GetUserEmail(dbMock, 1)
    if err != nil {
        t.Fatal(err)
    }
    if email != "user@example.com" {
        t.Errorf("unexpected email: %s", email)
    }
    
    // 4. Проверяем, что мок был вызван как ожидалось
    dbMock.AssertExpectations(t)
}

```
## Использование testify/mock для упрощения моков
Библиотека testify/mock предоставляет удобный способ создания моков.

Установка:

```shell
go get github.com/stretchr/testify
Пример мока для интерфейса:
```
```go
// Интерфейс, который мы хотим замокать
type Mailer interface {
    Send(email string, body string) error
}

// MockMailer реализует интерфейс Mailer
type MockMailer struct {
    mock.Mock
}

func (m *MockMailer) Send(email string, body string) error {
    args := m.Called(email, body)
    return args.Error(0)
}

func TestNotificationService(t *testing.T) {
    // Создаём мок
    mailer := new(MockMailer)
    
    // Настраиваем ожидание
    mailer.On("Send", "user@example.com", "Hello").Return(nil)
    
    // Тестируем
    service := NewNotificationService(mailer)
    err := service.SendWelcomeEmail("user@example.com")
    assert.NoError(t, err)
    
    // Проверяем вызовы
    mailer.AssertExpectations(t)
}
```

## Использование mockery для генерации моков
mockery автоматически генерирует моки на основе интерфейсов.

Установка:
```shell
go install github.com/vektra/mockery/v2@latest
```
Генерация мока для интерфейса:
```shell
mockery --name=Mailer --output=mocks --case=underscore
```

Пример использования сгенерированного мока:

```go
func TestNotificationService_GeneratedMock(t *testing.T) {
    // Используем сгенерированный мок
    mailer := &mocks.Mailer{}
    
    mailer.On("Send", "user@example.com", "Welcome").Return(nil)
    
    service := NewNotificationService(mailer)
    err := service.SendWelcomeEmail("user@example.com")
    require.NoError(t, err)
    
    mailer.AssertExpectations(t)
}
```

## Практический пример: тестирование сервиса с моком БД

```go
// user_service.go
type UserService struct {
    repo UserRepository
}

func (s *UserService) GetUserName(id int) (string, error) {
    user, err := s.repo.GetUser(id)
    if err != nil {
        return "", err
    }
    return user.Name, nil
}

// user_service_test.go
func TestUserService_GetUserName(t *testing.T) {
    // 1. Создаём мок репозитория
    repoMock := new(UserRepositoryMock)
    
    // 2. Настраиваем ожидаемый вызов
    testUser := &User{ID: 1, Name: "Test User"}
    repoMock.On("GetUser", 1).Return(testUser, nil)
    
    // 3. Создаём сервис с моком
    service := &UserService{repo: repoMock}
    
    // 4. Вызываем метод
    name, err := service.GetUserName(1)
    
    // 5. Проверяем результат
    assert.NoError(t, err)
    assert.Equal(t, "Test User", name)
    
    // 6. Проверяем, что мок был вызван
    repoMock.AssertExpectations(t)
    
    // 7. Проверяем, что не было неожиданных вызовов
    repoMock.AssertNumberOfCalls(t, "GetUser", 1)
}

```

## Когда использовать моки, а когда стабы
**Используйте моки, когда нужно:**
 - Проверить, что метод был вызван с определёнными параметрами
 - Убедиться в определённой последовательности вызовов
 - Проверить количество вызовов метода

**Используйте стабы, когда нужно:**
 - Просто подменить реализацию для изоляции теста
 - Вернуть заранее известные данные
 - Заменить медленные или нестабильные зависимости