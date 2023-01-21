FROM golang:alpine AS builder

RUN apk update && apk add bash

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

COPY app/go.mod .
COPY app/go.sum .
RUN go mod download

COPY app/ .

RUN go build -o main .

WORKDIR /dist

RUN cp /build/main .

FROM scratch

COPY --from=builder /dist/main /


EXPOSE 3002

ENTRYPOINT ["/main"]