#FROM golang:1.14.15 AS builder
FROM golang:alpine3.16 AS builder
RUN mkdir -p /build
WORKDIR /build
COPY . .
RUN go mod download
RUN go build -o agent-server

FROM alpine
WORKDIR /root/
COPY --from=builder /build/agent-server ./
COPY --from=builder /build/config ./

CMD ["./agent-server"]
