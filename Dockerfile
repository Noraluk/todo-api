FROM golang:1.18.2-alpine3.16

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download
RUN go mod verify

COPY . .

RUN go install github.com/cosmtrek/air@latest

CMD ["air"]