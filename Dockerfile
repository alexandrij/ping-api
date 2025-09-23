FROM golang:1.24.6-alpine AS builder
WORKDIR /build
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /main main.go
FROM alpine:latest
COPY --from=builder main /bin/main
ENTRYPOINT ["/bin/main"]