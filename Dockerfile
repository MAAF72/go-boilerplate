FROM golang:1.18.5

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o binary

CMD ["/app/binary"]