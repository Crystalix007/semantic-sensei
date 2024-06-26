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
	go build -o /app/semantic-sensei ./cmd/semantic-sensei

FROM gcr.io/distroless/static-debian12:nonroot

WORKDIR /app
COPY --from=builder /app/semantic-sensei ./semantic-sensei

ENTRYPOINT ["/app/semantic-sensei", "serve", "--address=:8080"]
EXPOSE 8080/tcp
