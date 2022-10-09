FROM golang:1.19
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go work init api model mysql 
WORKDIR /app/api
RUN go mod download
WORKDIR /app/model
RUN go mod download
WORKDIR /app/mysql
RUN go mod download
RUN go build -o main ./...
CMD ["/app/main"]