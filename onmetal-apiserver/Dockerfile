# Build the manager binary
FROM --platform=$BUILDPLATFORM golang:1.19 AS builder

WORKDIR /workspace
# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum

# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download

# Copy the go source
COPY cmd/ cmd/
COPY internal/ internal/

ARG TARGETOS
ARG TARGETARCH

RUN mkdir bin

RUN --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/go/pkg \
    CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH GO111MODULE=on go build -ldflags="-s -w" -a -o bin/onmetal-apiserver ./cmd/apiserver

FROM gcr.io/distroless/static:nonroot AS onmetal-apiserver
WORKDIR /
COPY --from=builder /workspace/bin/onmetal-apiserver .
USER 65532:65532

ENTRYPOINT ["/onmetal-apiserver"]
