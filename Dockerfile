FROM golang:1.16
MAINTAINER Ebubekir Tabak <ebubekir.tabak@yahoo.com>


COPY ./ /app

WORKDIR /app

RUN go mod download

COPY *.go ./

RUN go build -o main .

EXPOSE 3000

CMD ["./main"]
