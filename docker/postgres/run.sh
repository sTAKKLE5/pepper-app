#!/bin/bash

# Run the Docker container
docker run --name pepper-app -p 5432:5432 -d pepper-app
