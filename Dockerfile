FROM golang:1.16-alpine AS builder

RUN go env -w GO111MODULE=on

COPY . /go/src/pmc_server

WORKDIR /go/src/pmc_server

RUN go install ./...

FROM alpine:3.15

COPY ./config.yaml /
COPY --from=builder /go/bin/pmc_server /bin/pmc_server

EXPOSE 8081

ENTRYPOINT ["/bin/pmc_server"]


