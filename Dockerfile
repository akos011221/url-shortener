FROM golang:1.23.2

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /url-shortener

EXPOSE 8080

CMD ["/url-shortener"]
