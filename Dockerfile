FROM golang:1.19 as builder

ENV GOPATH /c4c-ir
ENV GOBIN /c4c-ir/app
WORKDIR ${GOBIN}
COPY go.mod go.sum ${GOBIN}/
RUN go mod download
ADD . ${GOBIN}
RUN go build -o ./app

FROM cgr.dev/chainguard/glibc-dynamic:latest
COPY --from=builder /c4c-ir/app/app /app
ENTRYPOINT ["/app"]
