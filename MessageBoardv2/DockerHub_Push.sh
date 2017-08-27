#!/bin/sh

export DOCKER_ID_USER=krkartick
docker login
docker images
docker tag go_mb_server_v2 $DOCKER_ID_USER/mbserverv2.project
docker push $DOCKER_ID_USER/mbserverv2.project
