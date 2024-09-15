FROM golang:alpine AS builder

WORKDIR /app

ADD go.mod .

COPY . .

RUN go build -o main ./cmd


FROM alpine

WORKDIR /build

COPY --from=builder /app/main /build/main
COPY --from=builder /app/config/*.yml /build/config/
COPY --from=builder /app/config/*.sql /build/config/

EXPOSE 5104

CMD ["/build/main"]