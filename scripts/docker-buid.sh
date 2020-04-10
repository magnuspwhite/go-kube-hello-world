#! /bin/bash

IMAGE="magnuspwhite/validity"
IMAGE_WITH_TAG=${IMAGE}:latest

echo ${DOCKERHUB_PASS} | docker login --username "$DOCKERHUB_USER" --password-stdin
docker build . -t $IMAGE_WITH_TAG
docker push $IMAGE_WITH_TAG