FROM golang:1.21.6-alpine

WORKDIR /app

# Копирование go.mod и go.sum, если он есть
COPY go.mod ./
COPY go.sum* ./

# Загрузка зависимостей
RUN go mod download

# Копирование исходного кода проекта
COPY . .

# Сборка приложения
RUN go build -o products-service

EXPOSE 8082

CMD ["./products-service"]
