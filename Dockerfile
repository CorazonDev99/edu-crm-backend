FROM golang:1.21.1-alpine3.18
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o main ./command/main.go
RUN go mod tidy
RUN go mod vendor
CMD ["./main"]
