FROM golang as builder
COPY . /go/src/github/flaccid/go-hello-world
ENV CGO_ENABLED=0 GOOS=linux
RUN cd /go/src/github/flaccid/go-hello-world/cmd/hello && \
     go mod download && \
     go build -o /hello

FROM scratch
COPY --from=builder /hello /hello
ENTRYPOINT ["/hello"]
