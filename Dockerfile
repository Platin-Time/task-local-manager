# ---------- Build stage ----------
FROM golang:1.27-alpine AS builder

WORKDIR /app

# Кэшируем зависимости отдельно от кода
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходники и собираем
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-s -w" -o /app/bin/server ./cmd/main.go

# ---------- Runtime stage ----------
FROM alpine:3.20

RUN apk add --no-cache ca-certificates tzdata && \
    addgroup -S app && adduser -S app -G app

ENV TZ=UTC

WORKDIR /app

COPY --from=builder /app/bin/server /app/server

USER app

EXPOSE 8000

ENTRYPOINT ["/app/server"]