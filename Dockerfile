FROM golang:latest

COPY server.go .

RUN go build server.go

CMD ["./server"]