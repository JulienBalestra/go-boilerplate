FROM golang:1.10 as builder

ENV PROJECT go-boilerplate

RUN git clone --depth=1 https://github.com/JulienBalestra/$PROJECT.git /go/src/github.com/JulienBalestra/$PROJECT && \
    make -C /go/src/github.com/JulienBalestra/$PROJECT

FROM busybox:latest

COPY --from=builder /go/src/github.com/JulienBalestra/$PROJECT/$PROJECT /usr/local/bin/$PROJECT
