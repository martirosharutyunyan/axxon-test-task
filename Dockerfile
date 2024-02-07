FROM golang:1.22rc1-alpine3.19 as go-builder

USER root

ARG GOPRIVATE
ARG GITHUB_TOKEN
ARG GITHUB_USERNAME

RUN apk add --no-cache git ca-certificates
RUN go env -w GOPRIVATE=$GOPRIVATE
RUN git config --global url."https://$GITHUB_USERNAME:$GITHUB_TOKEN@github.com/".insteadOf "https://github.com/"

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download -x

COPY . .
ENV GOCACHE="/root/.cache/go-build"
RUN --mount=type=cache,target="/root/.cache/go-build/" go build ./cmd/main-service

FROM alpine:latest

WORKDIR /app

COPY --from=go-builder /app/main-service /usr/local/bin

ENTRYPOINT ["main-service"]
