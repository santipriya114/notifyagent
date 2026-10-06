FROM golang:1.26.5

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o notification-agent

EXPOSE 8085

CMD ["./notification-agent"]