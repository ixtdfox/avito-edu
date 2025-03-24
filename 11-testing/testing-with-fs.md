# Тестирование работы с файловой системой в Go

Тестирование кода, взаимодействующего с файловой системой, представляет особые сложности. В этой статье мы рассмотрим:
- Проблемы тестирования файлового ввода-вывода
- Как использовать библиотеку `afero` для мокирования файловой системы

---

## Проблемы тестирования файлового ввода-вывода

### Основные проблемы:

1. **Зависимость от состояния файловой системы**
    - Тесты могут влиять друг на друга, изменяя одни и те же файлы
    - Тесты могут давать разные результаты на разных машинах

2. **Побочные эффекты**
    - Тесты могут создавать/удалять файлы, что нежелательно

3. **Проблемы производительности**
    - Работа с реальной файловой системой медленнее, чем с моками

### Пример 1: Проблемный тест

```go
func CountLines(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	count := 0
	for scanner.Scan() {
		count++
	}
	return count, scanner.Err()
}

func TestCountLines(t *testing.T) {
	// Создаём временный файл
	err := os.WriteFile("test.txt", []byte("line1\nline2\nline3"), 0644)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove("test.txt") // Не всегда срабатывает при панике

	count, err := CountLines("test.txt")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if count != 3 {
		t.Errorf("Expected 3 lines, got %d", count)
	}
}
```

**Проблемы**:
- Тест зависит от файловой системы
- Может не очистить файл при панике
- Не изолирован от других тестов

---

## Как мокировать файловую систему с `afero`

Библиотека `afero` предоставляет абстракцию файловой системы, позволяющую использовать как реальную ФС, так и моки.

### Установка:
```bash
go get github.com/spf13/afero
```

### Пример 2: Рефакторинг с использованием `afero`

```go
import (
	"github.com/spf13/afero"
)

func CountLines(fs afero.Fs, filename string) (int, error) {
	file, err := fs.Open(filename)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	count := 0
	for scanner.Scan() {
		count++
	}
	return count, scanner.Err()
}
```

### Пример 3: Тест с моком файловой системы

```go
func TestCountLinesWithAfero(t *testing.T) {
	// Создаём мок файловой системы
	fs := afero.NewMemMapFs()

	// Создаём файл в памяти
	err := afero.WriteFile(fs, "test.txt", []byte("line1\nline2\nline3"), 0644)
	if err != nil {
		t.Fatal(err)
	}

	count, err := CountLines(fs, "test.txt")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if count != 3 {
		t.Errorf("Expected 3 lines, got %d", count)
	}
}
```

### Пример 4: Тестирование ошибок

```go
func TestCountLines_FileNotExists(t *testing.T) {
	fs := afero.NewMemMapFs()
	_, err := CountLines(fs, "nonexistent.txt")
	if err == nil {
		t.Error("Expected error for nonexistent file")
	}
}
```

---

## Продвинутые сценарии

### Пример 5: Тестирование записи в файл

```go
func WriteData(fs afero.Fs, filename string, data []byte) error {
	file, err := fs.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(data)
	return err
}

func TestWriteData(t *testing.T) {
	fs := afero.NewMemMapFs()
	testData := []byte("test data")

	err := WriteData(fs, "output.txt", testData)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	// Проверяем, что файл создан и содержит правильные данные
	exists, err := afero.Exists(fs, "output.txt")
	if err != nil || !exists {
		t.Error("File was not created")
	}

	content, err := afero.ReadFile(fs, "output.txt")
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	if !bytes.Equal(content, testData) {
		t.Errorf("Expected %s, got %s", testData, content)
	}
}
```

### Пример 6: Тестирование работы с директориями

```go
func ListFiles(fs afero.Fs, dir string) ([]string, error) {
	return afero.ReadDirNames(fs, dir)
}

func TestListFiles(t *testing.T) {
	fs := afero.NewMemMapFs()

	// Создаём структуру директорий и файлов
	fs.MkdirAll("testdir/subdir", 0755)
	afero.WriteFile(fs, "testdir/file1.txt", []byte{}, 0644)
	afero.WriteFile(fs, "testdir/file2.txt", []byte{}, 0644)
	afero.WriteFile(fs, "testdir/subdir/file3.txt", []byte{}, 0644)

	files, err := ListFiles(fs, "testdir")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expected := []string{"file1.txt", "file2.txt", "subdir"}
	if !reflect.DeepEqual(files, expected) {
		t.Errorf("Expected %v, got %v", expected, files)
	}
}
```

---

## Альтернативные подходы

### 1. Использование временных файлов

Если всё же нужно работать с реальной ФС:

```go
func TestWithTempFile(t *testing.T) {
	// Создаём временный файл
	tmpfile, err := os.CreateTemp("", "example")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name()) // Удаляем после теста

	content := []byte("test content")
	if _, err := tmpfile.Write(content); err != nil {
		t.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatal(err)
	}

	// Тестируем на реальной ФС
	fs := afero.NewOsFs()
	count, err := CountLines(fs, tmpfile.Name())
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if count != 1 {
		t.Errorf("Expected 1 line, got %d", count)
	}
}
```

### 2. Использование интерфейсов

Можно создать собственный интерфейс файловой системы:

```go
type FileSystem interface {
	Open(name string) (File, error)
	Create(name string) (File, error)
	// другие методы
}

type File interface {
	io.Reader
	io.Writer
	io.Closer
}
```