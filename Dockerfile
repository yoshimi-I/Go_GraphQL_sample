# Development stage
FROM golang:1.22.2 AS dev

ENV ROOT=/go/src/app

WORKDIR ${ROOT}

COPY backend ./

RUN go mod tidy

RUN go install github.com/cosmtrek/air@v1.51.0
RUN go install github.com/go-delve/delve/cmd/dlv@latest

CMD ["air", "-c", ".air.toml"]


# Production build stage
FROM golang:1.22.2 AS builder

ENV ROOT=/go/src/app

WORKDIR ${ROOT}

COPY . .

RUN go build -o main ./cmd/graphql

# Production stage
FROM alpine:3.19.1 AS deploy

ENV ROOT=/go/src/app

WORKDIR ${ROOT}

RUN apk add --no-cache tzdata

COPY --from=builder ${ROOT}/main .

CMD ["/go/src/app/main"]
