!/bin/bash

DOCKER_HUB_HOST=$1
DOCKER_HUB_USERNAME=$2
DOCKER_HUB_PASSWD=$3

# docker:20.10.6-make
(echo $DOCKER_HUB_PASSWD | docker login $DOCKER_HUB_HOST -u $DOCKER_HUB_USERNAME --password-stdin) 1>/dev/null 2>&1

docker build --no-cache -t docker:20.10.6-make -f ./build/basic_environment/docker-make/dockerfile .

docker tag docker:20.10.6-make $DOCKER_HUB_HOST/common_mirror/docker:20.10.6-make
docker push ${DOCKER_HUB_HOST}/common_mirror/docker:20.10.6-make
