# Multi-stage build
# Best to use Debian instead of alpine
FROM golang:1.20.4-alpine AS builder

WORKDIR /app

COPY . .

RUN go build -o ./go-docker-demo .


FROM gcr.io/distroless/static-debian11

COPY --from=builder /app/go-docker-demo /go-docker-demo



EXPOSE 8080

CMD [ "./go-docker-demo" ]