FROM golang:1.19

WORKDIR /app

COPY mac-app/go.mod .

RUN go mod download

COPY mac-app/. .

RUN CGO_ENABLED=0 GOOS=linux go build -o app

EXPOSE 3000

CMD ["./app"]

