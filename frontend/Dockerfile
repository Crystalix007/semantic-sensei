FROM --platform=$BUILDPLATFORM docker.io/library/node:20-alpine AS build

WORKDIR /app

RUN npm install -g pnpm
COPY package.json pnpm-lock.yaml ./
RUN \
	--mount=type=cache,target=/app/node_modules \
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
