FROM golang as builder
RUN go get github.com/flaccid/go-hello-world/cmd/hello && \
    cd /go/src/github.com/flaccid/go-hello-world/cmd/hello && \
    CGO_ENABLED=0 GOOS=linux go build -o /hello

FROM scratch
COPY --from=builder /hello /hello
ENTRYPOINT ["/hello"]
