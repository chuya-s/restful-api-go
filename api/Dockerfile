FROM golang:1.15
RUN mkdir /app
WORKDIR /app/api
COPY . /app/
RUN go build -a -o main ./...
CMD ["/app/api/main"]
