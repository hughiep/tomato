# Install dependencies
FROM golang:1.23.1-alpine3.20 AS base

WORKDIR /

RUN set -x \ 
  && apk add --no-cache git

# Build the binary
FROM base AS builder

WORKDIR /payment-api

COPY go.mod go.sum ./

RUN go mod download 

COPY . ./

RUN mkdir -p /api \
  && go install \
  && go build -o /api .

# Copy the binary to the runtime image
FROM alpine:3.14 AS runtime

ARG WORK_DIR="/home/api"

WORKDIR ${WORK_DIR}

# Add the user
USER root

RUN addgroup -S api && adduser -S api -G api

COPY --from=builder --chown=api:api \
  /api \
  ${WORK_DIR}/bin/payment-api

USER api

RUN ls -la ${WORK_DIR}/bin/payment-api

ENTRYPOINT [ "./bin/payment-api/tomato" ]
