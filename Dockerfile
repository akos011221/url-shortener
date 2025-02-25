FROM golang:1.23.2

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /url-shortener

ENV ENV="production"
ENV SERVER_ADDRESS=:8080
ENV DATABASE_URL=""

EXPOSE 8080

CMD ["/url-shortener"]
