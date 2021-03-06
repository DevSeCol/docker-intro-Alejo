FROM golang:1.14
WORKDIR /go/src
CMD ["ls -Rl"]
COPY ./main.go .
RUN go build main.go 
EXPOSE 8080
CMD ["./main"]
