FROM golang:1.21 AS builder

WORKDIR /app

COPY . .

RUN go mod tidy

EXPOSE 8080

WORKDIR /app

CMD ["tail", "-f", "/dev/null"]
