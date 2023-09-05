#!/bin/bash
# This script is used to launch the 'staging' environment
# Launch: bash hack/launch.staging.sh

# Load environment variables
export $(grep -v '^#' .staging.env | xargs)

# launch the backend and database
docker-compose -f docker-compose.staging.yaml down -v
docker-compose -f docker-compose.staging.yaml up --build