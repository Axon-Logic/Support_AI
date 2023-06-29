FROM alpine:3.14 as production
WORKDIR /app
# Add certificates
RUN apk add --no-cache ca-certificates
# Copy built binary from builder
COPY ./app .
# Exec built binary
CMD ./app
