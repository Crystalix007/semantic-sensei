FROM --platform=$BUILDPLATFORM docker.io/library/node:20-alpine AS build

WORKDIR /app

RUN npm install -g pnpm
RUN \
	--mount=type=cache,target=/app/node_modules \
	--mount=type=bind,source=package.json,target=package.json \
	--mount=type=bind,source=pnpm-lock.yaml,target=pnpm-lock.yaml \
	pnpm install

COPY . .
RUN \
	--mount=type=cache,target=/app/node_modules \
	pnpm run build

FROM docker.io/library/node:20-alpine AS release

WORKDIR /app

COPY --from=build /app/build .

ENV HOST=0.0.0.0
ENV PORT=3000
ENV NODE_OPTIONS=--experimental-default-type=module

ENTRYPOINT [ "node", "index.js" ]
EXPOSE 3000/tcp
