FROM oven/bun:1.0.3 AS css-build


WORKDIR /app


COPY package.json bun.lockb ./
RUN bun install

COPY . .

RUN bun run build:css

FROM golang:1.23-alpine AS go-build

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download 

COPY . .

RUN go build -o ./bin/web main.go

FROM ubuntu:22.04


WORKDIR /app


COPY --from=go-build /app/bin/web .
COPY --from=css-build /app/static/css/output.css ./static/css/output.css

CMD ["./web"]
