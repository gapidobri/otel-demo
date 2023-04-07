FROM golang AS builder

WORKDIR /go/src/app

COPY . .

RUN go mod download

ENTRYPOINT [ "go", "run", "cmd/app/main.go" ]