FROM alpine:3.14 as production
# Add certificates
RUN apk add --no-cache ca-certificates
# Copy built binary from builder
COPY ./app .
# Expose port for GIN and GRPC
EXPOSE 7776
EXPOSE 7777
# Exec built binary
CMD ./app