# Заготовка библиотеки с internal-пакетом

Book и Library показывают экспорт полей, конструкторы и методы. Member ещё не определён; ряд функций содержит пустой return. cmd/library/main.go объявляет package library и функцию library, поэтому исполняемой точки входа пока нет.

Учебная папка `OOP/library`. Код сохранён в исходном состоянии.

## Структура

- `.gitignore`
- `cmd/library/main.go`
- `go.mod`
- `internal/library/book.go`
- `internal/library/library.go`
- `internal/library/member.go`

## Требования и команды

Проверено на Windows с Go 1.27.1 (windows/amd64); `go.mod` требует Go 1.27.1. Команды выполняются из корня репозитория.

Исполняемой точки входа нет. Это пакет или незавершённая заготовка; `go run .` для него не применяется.

```sh
go test ./...
go vet ./...
```

## Результаты проверки

- `go test ./...`: ошибка (код 1).
- `go vet ./...`: ошибка (код 1).

Файлов `*_test.go` нет. Успех `go test` здесь означает проверку сборки, а не прохождение набора тестов.

Диагностика `go test ./...` (фактический результат):

```text
?   	library/cmd/library	[no test files]
FAIL	library/internal/library [build failed]
FAIL
# library/internal/library
internal\library\library.go:8:15: undefined: Member
internal\library\library.go:12:2: not enough return values
	have ()
	want (*Library)
internal\library\library.go:18:2: not enough return values
	have ()
	want (*Book, bool)
internal\library\library.go:21:2: not enough return values
	have ()
	want (error)
internal\library\library.go:24:48: undefined: Member
internal\library\library.go:25:2: not enough return values
	have ()
	want (unknown type, bool)
internal\library\book.go:13:2: not enough return values
	have ()
	want (*Book, error)
```

Проверки выполнены 22 ноября 2026 года. Сборка и тесты не исправляют существующие учебные ошибки.
