FROM golang:1.23.5-alpine AS builder

# Установка зависимостей для healthcheck
RUN apk add --no-cache curl

COPY . /github.com/RikiTikiTavee17/course/grpc/source/
WORKDIR /github.com/RikiTikiTavee17/course/grpc/source/

RUN go mod download
RUN go build -o ./bin/crud_server cmd/grpc_server/main.go

FROM alpine:latest

# Установка инструментов для healthcheck
RUN apk add --no-cache curl netcat-openbsd

WORKDIR /root/
COPY --from=builder /github.com/RikiTikiTavee17/course/grpc/source/bin/crud_server .

# Добавляем задержку перед запуском для гарантии готовности БД
CMD ["sh", "-c", "echo 'Waiting for DB to be ready...'; sleep 5; ./crud_server"]
