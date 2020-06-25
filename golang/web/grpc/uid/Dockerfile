FROM golang:1.14 AS builder

RUN mkdir /uidapp
ADD . /uidapp
WORKDIR /uidapp/server
RUN CGO_ENABLED=0 GOOS=linux go build -o uidserver ./...

FROM alpine:latest AS execution
COPY --from=builder /uidapp/server .
CMD ["./uidserver"]