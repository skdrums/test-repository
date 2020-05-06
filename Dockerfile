FROM golang:1.13-alpine

WORKDIR /go/src/app
COPY . .
ENV GO111MODULE=on

RUN go get github.com/pilu/fresh
CMD ["fresh"]
