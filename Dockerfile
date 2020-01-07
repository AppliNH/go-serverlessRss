FROM golang:latest

ENV GO111MODULE=on

WORKDIR /app

COPY ./go.mod .

EXPOSE 8080

RUN go mod download

COPY . .
CMD ["go", "run", "main.go"]