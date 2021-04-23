FROM golang:1.16

WORKDIR /go/src/app

# manage dependencies
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o kc_server cmd/server/main.go
CMD ["/go/src/app/kc_server"]