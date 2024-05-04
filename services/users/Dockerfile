FROM golang:1.21.6-alpine

WORKDIR /app
COPY . .

RUN go build -o users-service

CMD ["./users-service"]
