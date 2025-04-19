FROM golang:1.21-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /app/api-server ./cmd/api-server

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/api-server .

EXPOSE 8080

CMD ["./api-server"] 