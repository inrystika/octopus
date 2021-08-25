!/bin/bash

DOCKER_HUB_HOST=$1
DOCKER_HUB_USERNAME=$2
DOCKER_HUB_PASSWD=$3
IMAGE_TAG=$4

(echo $DOCKER_HUB_PASSWD | docker login $DOCKER_HUB_HOST -u $DOCKER_HUB_USERNAME --password-stdin) 1>/dev/null 2>&1

docker build --no-cache -t notebook:$IMAGE_TAG -f ./build/notebook_environment/$IMAGE_TAG/dockerfile .

docker tag notebook:$IMAGE_TAG $DOCKER_HUB_HOST/common_mirror/notebook:$IMAGE_TAG
docker push ${DOCKER_HUB_HOST}/common_mirror/notebook:$IMAGE_TAG
