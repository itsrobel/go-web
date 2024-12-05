FROM oven/bun:1.0.3 AS css-build


WORKDIR /app


COPY package.json bun.lockb ./
RUN bun install

COPY . .

RUN bun run build:css

FROM golang:1.23-alpine AS go-build

WORKDIR /app
COPY go.mod go.sum ./
COPY ./static ./static
COPY ./content/ ./content/
RUN go mod download 


RUN --mount=type=cache,target=/go/pkg/mod/ \
  --mount=type=bind,source=go.sum,target=go.sum \
  --mount=type=bind,source=go.mod,target=go.mod \
  go mod download -x


COPY . .

RUN go build -o ./bin/web main.go

FROM ubuntu:22.04


WORKDIR /app


COPY --from=go-build /app/bin/web .
COPY --from=go-build /app/static ./static
COPY --from=go-build /app/content ./content
COPY --from=css-build /app/static/css/output.css ./static/css/output.css

EXPOSE 8080

CMD ["./web"]
