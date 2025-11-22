FROM golang:1.20 AS builder
WORKDIR /src
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /out/capital-gains ./cmd/capital-gains

FROM alpine:3.18
RUN apk add --no-cache ca-certificates
COPY --from=builder /out/capital-gains /usr/local/bin/capital-gains
USER nobody
ENTRYPOINT ["/usr/local/bin/capital-gains"]
