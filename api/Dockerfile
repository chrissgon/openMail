FROM golang:1.17-alpine

WORKDIR /api

COPY . .

RUN go build

EXPOSE 3000

CMD [ "go", "run", "main.go" ]