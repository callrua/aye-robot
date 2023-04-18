FROM golang:alpine as builder

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o bin/aye-robot

FROM scratch

EXPOSE 9090

WORKDIR /

COPY --from=builder /app/bin/aye-robot /aye-robot
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENTRYPOINT ["/aye-robot"]
