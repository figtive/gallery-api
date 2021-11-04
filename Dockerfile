FROM golang:1.17 AS builder
ARG CGO_ENABLED=0
WORKDIR /app
COPY . .
RUN ["go", "mod", "tidy"]
RUN ["go", "build"]

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/gallery-api /usr/local/bin/
CMD ["/usr/local/bin/gallery-api"]
