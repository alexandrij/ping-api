FROM golang:1.24.6 AS builder

WORKDIR /build

COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod download

COPY main.go main.go
COPY configs/ configs/
COPY cmd/ cmd/
COPY config/ config/
COPY internal/ internal/
COPY pkg/ pkg/

RUN CGO_ENABLED=0 GOOS=linux go build -a -o bin/ping-api main.go

FROM alpine:latest

WORKDIR /

COPY --from=builder /build/bin/ping-api /ping-api/bin/ping-api
COPY --from=builder /build/configs/app.yml /ping-api/configs/app.yml

ENTRYPOINT ["/bin/ping-api"]