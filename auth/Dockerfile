FROM --platform=linux/amd64 golang:1.19.2-alpine

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o ./app

EXPOSE 3000

CMD ["./app"]