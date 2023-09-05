#!/bin/bash

# This script is used to launch the application in development mode.
export $(grep -v '^#' .dev.env | xargs)

go run main.go
