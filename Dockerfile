# File to build the production image
FROM golang:alpine AS builder

WORKDIR /project

# Copy go mod and sum files
COPY src/go.* ./
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY src/ .
RUN go build -o /project/build/main .

# Now copy it into our base image.
FROM scratch
WORKDIR /app
COPY --from=builder /project/build/main /app/build/main

COPY scripts/launch.sh .
RUN chmod +x launch.sh

EXPOSE 8080
CMD [ "/app/launch.sh" ]