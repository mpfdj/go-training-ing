FROM golang:1.20.4-alpine

WORKDIR /app

COPY . .

RUN go build -o ./go-docker-demo

EXPOSE 8080

CMD [ "./go-docker-demo" ]