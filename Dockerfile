FROM golang AS builder

WORKDIR /src
# Download dependencies
COPY go.mod go.sum /
RUN go mod download && \
    go install github.com/GeertJohan/go.rice && \
    go install github.com/GeertJohan/go.rice/rice


# Add source code
COPY . .
RUN CGO_ENABLED=0 go generate 'rice embed-go' && \
    go build -o main .

# Multi-Stage production build
FROM alpine AS production
RUN apk --no-cache add ca-certificates

WORKDIR /app
# Retrieve the binary from the previous stage
COPY --from=builder /src/main .
# Expose port
EXPOSE 8080
# Set the binary as the entrypoint of the container
CMD ["./main"]