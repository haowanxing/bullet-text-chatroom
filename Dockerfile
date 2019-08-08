FROM golang:1.12.7
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go mod init chatroom
RUN go mod download
RUN go build -o main ./example/main.go
CMD ["/app/main"]