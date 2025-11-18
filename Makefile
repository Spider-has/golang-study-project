include .env.local
export

# Цель: сборка бинарника
build:
	templ generate ./src/web/templates/
	go build -o bin/server src/cmd/server/main.go
	@echo "Сборка завершена. файл запуска: bin/server"

# Цель: запуск сервера
run: build
	./bin/server

# Цель: запуск тестов
test:
	go test ./src/...

# Цель: очистка (удаление бинарников)
clean:
	rm -f bin/server
	@echo "Очистка завершена."

# Цель: всё вместе — сборка, запуск тестов, очистка
all: build test clean