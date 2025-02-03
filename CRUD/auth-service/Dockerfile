FROM golang:1.22.2 AS builder

WORKDIR /app

# Копируем go.mod и go.sum для кэширования зависимостей
COPY go.mod go.sum ./
RUN go mod download

# Копируем остальные файлы и собираем приложение
COPY . .
RUN go build -o auth-service .

# Используем минимальный образ для финального контейнера
FROM ubuntu:22.04

WORKDIR /app
COPY --from=builder /app/auth-service .

CMD ["./auth-service"]