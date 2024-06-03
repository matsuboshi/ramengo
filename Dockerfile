FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o api cmd/api/main.go

FROM scratch
WORKDIR /app
COPY --from=builder /app/api .
EXPOSE 8080
CMD ["/api"]
