FROM golang:alpine AS builder

RUN apk add --no-cache git

WORKDIR /app

ADD go.mod .
RUN go mod download

COPY . .

RUN go build -o main ./cmd


FROM alpine

WORKDIR /build

COPY --from=builder /app/main /build/main
COPY --from=builder /app/config/*.yml /build/config/
COPY --from=builder /app/config/*.sql /build/config/

EXPOSE 5103

CMD ["/build/main"]