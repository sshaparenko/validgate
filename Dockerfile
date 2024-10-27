FROM golang:alpine AS builder

LABEL maintainer="Mykola Shaparenko"

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download 

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/main

# Start a new stage from scratch
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage.
COPY --from=builder /app/main .
COPY --from=builder /app/internal ./internal

# Expose port to the outside world
EXPOSE 8080

#Command to run the executable
CMD [ "./main" ]