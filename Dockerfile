FROM golang:1.14-alpine AS builder

WORKDIR /go/src/app
COPY go.mod go.sum ./

RUN go mod download

COPY . .
RUN go build -o main .

# Next stage
# Use alpine or ubuntu for development
# Use scratch for production
FROM scratch

WORKDIR /app
COPY --from=builder /go/src/app /app

CMD ["./main"]
