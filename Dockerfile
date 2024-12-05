FROM oven/bun:1.0.3 AS css-build
WORKDIR /app
COPY package.json bun.lockb ./
RUN bun install
COPY . .
RUN bun run build:css

FROM golang:1.23-alpine AS go-fetch-stage

WORKDIR /app
COPY go.mod go.sum ./
COPY ./static ./static
COPY ./content/ ./content/
RUN go mod download 


RUN --mount=type=cache,target=/go/pkg/mod/ \
  --mount=type=bind,source=go.sum,target=go.sum \
  --mount=type=bind,source=go.mod,target=go.mod \
  go mod download -x


# COPY . .


FROM ghcr.io/a-h/templ:latest AS generate-stage
COPY --chown=65532:65532 . /app
WORKDIR /app

RUN ["templ", "generate"]


FROM golang:1.23-alpine AS go-build-stage
COPY --from=generate-stage /app /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/app

# RUN go build -o ./bin/web main.go



FROM ubuntu:22.04

WORKDIR /app

COPY --from=go-build-stage /app/app .
COPY --from=go-fetch-stage /app/static ./static
COPY --from=go-fetch-stage /app/content ./content
COPY --from=css-build /app/static/css/output.css ./static/css/output.css

EXPOSE 8080

ENTRYPOINT ["/app"]
