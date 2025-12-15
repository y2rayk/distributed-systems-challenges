FROM golang:1.21-alpine

WORKDIR /app

COPY . .

RUN go mod init example.com/webhook-receiver || true
RUN go mod tidy
RUN go build -o /webhook-receiver

EXPOSE 8080

CMD ["/webhook-receiver"]

