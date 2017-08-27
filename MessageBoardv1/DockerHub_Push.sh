#!/bin/sh

export DOCKER_ID_USER=krkartick
docker login
docker images
docker tag go_mb_server_v1 $DOCKER_ID_USER/mbserverv1.project
docker push $DOCKER_ID_USER/mbserverv1.project
