FROM golang:1.23

WORKDIR /

COPY go.mod main.go ./

RUN go mod download

COPY . .

RUN go build main.go

EXPOSE 123

CMD ["./main"]