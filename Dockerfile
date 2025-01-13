FROM golang:1.20 AS builder

WORKDIR /app

COPY . .

RUN go mod tidy

EXPOSE 8080

WORKDIR /app

CMD ["tail", "-f", "/dev/null"]
