FROM golang:1.26.3-alpine3.23 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o estante-viva ./cmd/app/main.go

FROM alpine:3.23
WORKDIR /app
COPY --from=builder /app/estante-viva .
CMD [ "./estante-viva" ]
EXPOSE 8080
