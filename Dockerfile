FROM golang:1.20.5-alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

COPY *.mod *.sum *.go ./

RUN go mod vendor
RUN go build -a -ldflags '-w -s' -o main .

WORKDIR /dist

RUN cp /build/main .

FROM scratch
COPY --from=builder /dist/main .
ENTRYPOINT ["/main"]