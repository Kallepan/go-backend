# Generic GoLang Backend (REST API)

## Description

This is a generic backend for a REST API written in GoLang. It uses the [Gin](https://github.com/gin-gonic/gin) framework. It is meant to be used as a starting point for a new project. It includes basic modules for logging, configuration, database, and authentication. It also includes a Dockerfile for building a Docker image. The Docker image is based on the [Scratch](https://hub.docker.com/_/scratch/).

## Usage

- Use VSCode to open the project folder
- Copy the `.env.example` file to `.dev.env` and update the values
- Create the devcontainer. 
- Now launch `hack/launch.dev.sh` to start the development container