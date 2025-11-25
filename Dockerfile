FROM golang:1.25 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go mod tidy
RUN go mod vendor
RUN CGO_ENABLED=0 go build -mod=vendor -o bin/server main.go

FROM alpine:latest 
WORKDIR /root/
COPY --from=builder /app/bin/server .
EXPOSE 9090 
CMD ["./server"]