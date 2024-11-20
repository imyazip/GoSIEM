# Используем официальный образ golang
FROM golang:1.23-alpine AS builder

# Устанавливаем рабочую директорию для сборки
WORKDIR /build

# Копируем go.mod и go.sum для загрузки зависимостей
COPY go.mod go.sum ./

# Загружаем зависимости
RUN go mod tidy

# Копируем все остальные исходники
COPY . .

# Компилируем приложение
RUN go build -o /auth-server ./cmd/server

# Используем минимальный образ для выполнения
FROM alpine:latest

# Устанавливаем рабочую директорию
WORKDIR /root/

# Копируем скомпилированный бинарник из предыдущего этапа
COPY --from=builder /auth-server .

# Копируем конфигурационный файл в контейнер
COPY server-config.yaml /root/

# Открываем порт для gRPC
EXPOSE 50051

# Запускаем сервер
CMD ["./auth-server"]
