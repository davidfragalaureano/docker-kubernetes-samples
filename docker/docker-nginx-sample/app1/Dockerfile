FROM golang:latest
ENV GOBIN $GOPATH/bin
WORKDIR /app
ADD . /app
RUN go get ./...
RUN go build -o main .
CMD ["/app/main"]