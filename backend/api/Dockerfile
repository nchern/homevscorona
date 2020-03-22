FROM golang:1.13 as builder

WORKDIR /go/src/github.com/nchern/homevscorona/backend/api
COPY . .

# Build the project inside an intermediate container
RUN make install-deps install

FROM golang:1.13

WORKDIR /

COPY --from=builder /go/bin/api /api

# Main service port
EXPOSE 8080

# (optional) metrics exposing port
EXPOSE 8081

ENTRYPOINT /api
