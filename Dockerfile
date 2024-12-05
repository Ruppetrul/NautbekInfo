FROM golang:1.23

WORKDIR /app

COPY go.mod main.go ./

RUN go mod download

COPY . .

RUN go build -o main

EXPOSE 123

CMD ["./main"]