# syntax=docker/dockerfile:1

FROM golang:1.21

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./api ./service ./models ./storage ./utils ./config ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /url-shortener

EXPOSE 8080

CMD ["/url-shortener"]
