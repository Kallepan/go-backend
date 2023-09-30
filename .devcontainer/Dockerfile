FROM mcr.microsoft.com/devcontainers/go:1.21-bookworm

LABEL name="GoLang Dev Container" \
    description="GoLang Dev Container" \
    version="1.0.0"

# Install Go tools
RUN go install github.com/google/wire/cmd/wire@latest

EXPOSE ${PORT}