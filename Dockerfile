FROM golang:1.17-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
COPY config.json ./

RUN go mod download

COPY *.go ./

RUN go build -o /kana

CMD [ "/kana" ]
