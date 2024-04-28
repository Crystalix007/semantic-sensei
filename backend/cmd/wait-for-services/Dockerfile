FROM --platform=$BUILDPLATFORM docker.io/library/golang:1.22 as builder

ARG TARGETOS TARGETARCH TARGETVARIANT

WORKDIR /app
WORKDIR /src
RUN \
	--mount=type=cache,target=/go/pkg/mod \
	--mount=type=cache,target=/root/.cache/go-build \
	--mount=type=bind,source=go.mod,target=go.mod \
	--mount=type=bind,source=go.sum,target=go.sum \
	go mod download -x

COPY . .

RUN \
	--mount=type=cache,target=/go/pkg/mod \
	--mount=type=cache,target=/root/.cache/go-build \
	CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} GOARM=${TARGETVARIANT##*v} \
	go build -o /app/wait-for-services ./cmd/wait-for-services

FROM gcr.io/distroless/static-debian12:nonroot

WORKDIR /app
COPY --from=builder /app/wait-for-services ./wait-for-services

ENTRYPOINT ["/app/wait-for-services", "--database"]
