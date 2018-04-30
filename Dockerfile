FROM golang:latest

COPY . $GOPATH/src/github.com/lccezinha/twodo
WORKDIR $GOPATH/src/github.com/lccezinha/twodo

# RUN go build -o `go env GOPATH`/bin/twodo .
RUN go install

CMD twodo

EXPOSE 8888
