# Generic GoLang Backend (REST API)

## Description

This is a generic backend for a REST API written in GoLang. It uses the [Gin](https://github.com/gin-gonic/gin) framework. It is meant to be used as a starting point for a new project. It includes basic modules for logging, configuration, database, and authentication. It also includes a Dockerfile for building a Docker image. The Docker image is based on [Scratch](https://hub.docker.com/_/scratch/).

## Usage
- Launch the devcontainer using the preconfigured files
- Now launch `hack/launch.dev.sh` to start the webserver in development mode
- Tests can be run using `hack/test.sh` if the devcontainer is running.
