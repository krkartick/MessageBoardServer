#!/bin/bash

# Prepare the MessageBoard Server
go build -o MBServer -v main.go routes.go mysql.go 

# Build with docker
CGO_ENABLED=0 GOOS=linux go build -o mbserverd -a -installsuffix cgo -v main.go routes.go mysql.go 

# Delete any existing docker image
echo "Delete go_mb_server_v2 docker image"
docker rmi -f go_mb_server_v2
docker images

docker ps -a | grep MessageBoardServer | awk -F" " '{print "docker rm -f "$1}' > mypid
sh -x mypid

# Create MB Server docker image
docker build -t go_mb_server_v2 -f Dockerfile .

docker images

docker ps -a
