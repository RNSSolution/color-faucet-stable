From golang:1.12-alpine
WORKDIR /go/src/project/
COPY . /go/src/project/
RUN go build faucet.go
CMD ["./faucet"]
