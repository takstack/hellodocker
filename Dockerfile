FROM golang:1.15.2-alpine3.12
RUN apk add git
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go get github.com/go-sql-driver/mysql
RUN go build -o hd .
CMD ["/app/hd"]