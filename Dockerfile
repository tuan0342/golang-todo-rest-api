FROM golang:1.23
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go mod vendor
RUN CGO_ENABLED=1 go build -o bin/server ./main.go
CMD ./bin/server