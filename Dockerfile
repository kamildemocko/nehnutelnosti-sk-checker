FROM golang:1.24-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN apk add --no-cache build-base gcc musl-dev sqlite-dev

RUN go mod download

COPY . .

RUN CGO_ENABLED=1 go build -o /app/nehnutelnosti-checker ./src/cmd/app

FROM alpine:latest

RUN apk add --no-cache libc6-compat sqlite-libs

WORKDIR /app

COPY --from=builder /app/nehnutelnosti-checker .

RUN mkdir -p /app/data

COPY .env .

ENTRYPOINT ["/app/nehnutelnosti-checker"]