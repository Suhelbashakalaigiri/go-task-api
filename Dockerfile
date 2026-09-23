FROM golang:1.27-alpine AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -o go-task-api .

FROM alpine:3.22

WORKDIR /app

COPY --from=builder /app/go-task-api .

EXPOSE 8080

CMD ["./go-task-api"]