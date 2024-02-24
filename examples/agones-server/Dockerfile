FROM golang:1.22.0 as builder
WORKDIR /go/src
    COPY . github.com/seanhagen/endless_stream/examples/agones-server

    WORKDIR /go/src/github.com/seanhagen/endless_stream/examples/agones-server
    RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server .

    FROM gcr.io/distroless/static-debian12:nonroot
    WORKDIR /

    COPY --from=builder /go/src/github.com/seanhagen/endless_stream/examples/agones-server/server /server

    USER nonroot:nonroot
    ENTRYPOINT ["/server"]
