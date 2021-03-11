From golang:1.14.0
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o main server.go
CMD ["/app/main"]
