FROM golang:latest

COPY . $GOPATH/src/github.com/lccezinha/twodo
WORKDIR $GOPATH/src/github.com/lccezinha/twodo

RUN apt-get update \
  && apt-get install -y --allow-unauthenticated postgresql-client --no-install-recommends \
  && rm -rf /var/lib/apt/lists/*

RUN go get github.com/julienschmidt/httprouter
RUN go get github.com/lib/pq
RUN go build -o `go env GOPATH`/bin/twodo ./cmd/web

CMD twodo

EXPOSE 8888
