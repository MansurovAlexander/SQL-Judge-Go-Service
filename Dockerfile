FROM golang:1.21.6

RUN go version

COPY ./ ./

RUN apt-get update
RUN apt-get -y install postgresql-client

RUN chmod +x wait-for-postgres.sh

RUN go mod download
RUN go build -o sql-judge ./cmd/main.go

CMD ["./main"]