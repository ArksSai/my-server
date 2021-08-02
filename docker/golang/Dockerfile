FROM golang:alpine

WORKDIR /app
COPY . /app

RUN go build -o main *.go

CMD ["/app/main"]
