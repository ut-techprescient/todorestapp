FROM golang:1.21-bullseye

WORKDIR /app

COPY . .

RUN go mod download

RUN go build ./cmd/web

EXPOSE 4000

CMD ["./web"]