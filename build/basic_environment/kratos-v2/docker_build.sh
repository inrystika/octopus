!/bin/bash

DOCKER_HUB_HOST=$1
DOCKER_HUB_USERNAME=$2
DOCKER_HUB_PASSWD=$3
THIRD_PARTY=$4

# 构建kratos-ci:V2
(echo $DOCKER_HUB_PASSWD | docker login $DOCKER_HUB_HOST -u $DOCKER_HUB_USERNAME --password-stdin) 1>/dev/null 2>&1

docker build --no-cache -t kratos:v2 -f ./build/basic_environment/kratos-v2/dockerfile $THIRD_PARTY

docker tag kratos:v2 $DOCKER_HUB_HOST/common_mirror/kratos:v2
docker push ${DOCKER_HUB_HOST}/common_mirror/kratos:v2