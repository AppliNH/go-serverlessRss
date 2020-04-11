FROM golang:latest

ENV GO111MODULE=on

WORKDIR /app

COPY ./go.mod .

RUN go mod download

COPY . .

# Build the Go app
RUN go build -o go-serverlessRss .

# Expose port 80 to the outside world
# EXPOSE 8080

CMD ["./go-serverlessRss"]