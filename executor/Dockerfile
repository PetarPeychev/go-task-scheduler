FROM golang:1.22 as builder

WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY executor/ .

RUN CGO_ENABLED=0 GOOS=linux go build -o executor

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/executor .
CMD ["./executor"]
