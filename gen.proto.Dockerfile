# compile proto files

FROM golang:1.24-alpine

RUN apk add --no-cache protobuf protobuf-dev

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

ENV PATH="/go/bin:${PATH}"

WORKDIR /workspace

ENTRYPOINT ["protoc"]