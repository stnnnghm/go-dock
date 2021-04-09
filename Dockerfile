FROM golang:latest

COPY main.go .

RUN go build main.go

CMD ["./main"]