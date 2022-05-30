FROM golang:1.18.2-alpine3.16

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /todo

EXPOSE 8080

RUN pwd

CMD [ "/todo" ]