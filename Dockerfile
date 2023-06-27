FROM golang:1.19

WORKDIR /app

COPY go.mod .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o app

EXPOSE 3000

CMD ["./app"]

