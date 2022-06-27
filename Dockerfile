FROM golang:alpine as builder
RUN apk update && apk add --no-cache git
RUN adduser -D -g '' appuser

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o /go/bin/app ./src

FROM alpine

WORKDIR /app

COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /go/bin/app ./app

COPY migrations migrations ./
USER appuser

EXPOSE 8081
ENTRYPOINT ["/app/app"]