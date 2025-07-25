FROM golang:1.24.4-alpine3.22 AS builder

WORKDIR /builder

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main cmd/server/main.go

FROM alpine:latest

WORKDIR /usr/src/

COPY --from=builder /builder/main .

EXPOSE 8000

CMD ["./main"]