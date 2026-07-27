# Enterprise Docker Container for smart-grid-energy-go-fiber-v2026-75
FROM alpine:3.19
RUN apk add --no-cache bash curl ca-certificates
WORKDIR /app
COPY . /app
EXPOSE 8080
CMD ["echo", "Container active for smart-grid-energy-go-fiber-v2026-75 (Go / Fiber Microservice)"]
