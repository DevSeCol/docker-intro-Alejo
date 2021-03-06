FROM golang:1.14
WORKDIR /go/src
COPY ./main.go .
RUN go build main.go 
EXPOSE 8080
CMD ["./main"]
