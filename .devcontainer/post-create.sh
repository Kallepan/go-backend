#!/bin/bash

# install go modules
cd src
go mod download
go mod tidy

# install go tools
go install github.com/google/wire/cmd/wire@latest