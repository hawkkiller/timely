FROM golang:alpine3.17 AS builder

COPY . /timely

WORKDIR /timely

RUN apk add build-base

RUN go mod download

RUN go build ./cmd/app

FROM alpine:3.17.0 AS runner

COPY --from=builder /timely/app /app

ENTRYPOINT ["/app"]