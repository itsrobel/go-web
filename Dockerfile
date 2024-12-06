# Stage 1: Build CSS using Bun
FROM oven/bun:1.0.3 AS css-build
WORKDIR /app
COPY package.json bun.lockb ./
RUN bun install
COPY . .
RUN bun run build:css

# Stage 2: Fetch Go dependencies
FROM golang:1.23-alpine AS go-fetch-stage
WORKDIR /app
COPY go.mod go.sum ./
COPY ./static ./static
COPY ./content/ ./content/
RUN go mod download

# Stage 3: Generate templates using Templ
FROM ghcr.io/a-h/templ:v0.2.793 AS generate-stage
WORKDIR /app
COPY --chown=65532:65532 . .
RUN ["templ", "generate"]

# Stage 4: Build Go application
FROM golang:1.23-alpine AS go-build-stage
WORKDIR /app
COPY --from=generate-stage /app /app
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/app

# Final Stage: Create runtime image
# FROM ubuntu:22.04 AS runtime
FROM alpine:3.21 AS runtime

WORKDIR /app

# Copy built application and required assets from previous stages
COPY --from=go-build-stage /app/app .
COPY --from=go-fetch-stage /app/static ./static
COPY --from=go-fetch-stage /app/content ./content
COPY --from=css-build /app/static/css/output.css ./static/css/output.css

# Expose application port and set entrypoint
EXPOSE 8080
ENTRYPOINT ["/app/app"]
