FROM golang:1.23
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go mod tidy
RUN go mod vendor
RUN CGO_ENABLED=0 go build -mod=vendor -o bin/server main.go

FROM alpine:latest 
WORKDIR /root/
COPY --from=builder /app/server .
EXPOSE 9090 
CMD ["./server"]