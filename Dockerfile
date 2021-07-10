FROM golang:1.16-alpine AS builder

WORKDIR /app
COPY . /app
RUN   apk add build-base && \
      go test -v -cover ./... && \
      go build -o d2iq

FROM alpine:3

COPY --from=builder /app/d2iq /app/d2iq
EXPOSE 8080/tcp
CMD ["/app/d2iq"]
