FROM golang:latest

WORKDIR /src/app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ENV GIN_MODE=release

RUN CGO_ENABLE=0 GOOS=linux go build -o /docker-tamer-api

EXPOSE 8080

CMD ["/docker-tamer-api"]
