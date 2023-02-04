FROM golang:alpine3.16 AS builder

WORKDIR /timely

COPY ./go.mod ./
COPY ./go.sum ./
RUN go mod download

COPY . ./

RUN apk add build-base

RUN go build ./cmd/app

FROM alpine:3.16.0 AS runner

COPY --from=builder /timely/app /app

ENTRYPOINT ["/app"]