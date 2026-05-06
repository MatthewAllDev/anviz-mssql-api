FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main .

FROM alpine:latest

RUN apk --no-cache add ca-certificates \
	&& addgroup -S app \
	&& adduser -S -G app app

WORKDIR /app

COPY --from=builder --chown=app:app /app/main .

USER app

CMD ["./main"]
