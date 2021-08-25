!/bin/bash

# ./docker_build.sh 3.6.0-rc.1 192.168.202.74:5000 username password

HELM_VERSION=$1
DOCKER_HUB_HOST=$2
DOCKER_HUB_USERNAME=$3
DOCKER_HUB_PASSWD=$4

if [ "$HELM_VERSION" = "" ]; then
  HELM_VERSION="latest"
fi
HELM_ORIGIN_VERSION=$HELM_VERSION

if [ "$HELM_VERSION" != "latest" ]; then
  HELM_VERSION="${HELM_VERSION}-make"
fi

(echo $DOCKER_HUB_PASSWD | docker login $DOCKER_HUB_HOST -u $DOCKER_HUB_USERNAME --password-stdin) 1>/dev/null 2>&1

docker build --no-cache -t helm:${HELM_VERSION} --build-arg helm_version=${HELM_ORIGIN_VERSION} -f ./build/basic_environment/helm-make/dockerfile .

docker tag helm:${HELM_VERSION} ${DOCKER_HUB_HOST}/common_mirror/helm:${HELM_VERSION}
docker push ${DOCKER_HUB_HOST}/common_mirror/helm:${HELM_VERSION}
