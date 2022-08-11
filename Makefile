#外部参数
SERVER_BINARY_DIR=$(binary_dir)
ifeq (${SERVER_BINARY_DIR}, )
	SERVER_BINARY_DIR=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))/server/bin
endif

RELEASE_VER=$(tag)
ifeq (${RELEASE_VER}, )
	RELEASE_VER=latest
endif

DOCKER_HUB_HOST=$(docker_hub_host)
DOCKER_HUB_USERNAME=$(docker_hub_userame)
DOCKER_HUB_PASSWD=$(docker_hub_passwd)
DOCKER_HUB_PROJECT=$(docker_hub_project)

HARBOR_HUB_HOST=$(harbor_hub_host)
HARBOR_HUB_USERNAME=$(harbor_hub_userame)
HARBOR_HUB_PASSWD=$(harbor_hub_passwd)
HARBOR_HUB_PROJECT=$(harbor_hub_project)
HARBOR_HUB_CA_FILE=$(harbor_hub_ca_file)
HARBOR_HUB_CERT_FILE=$(harbor_hub_cert_file)

NEED_LATEST=$(need_latest)


CHARTS_GIT_DIR=./tmp/gitcharts
CHARTS_GIT_CLONE=$(charts_git_clone)
CHARTS_GIT_RAW=$(charts_git_raw)
CHARTS_GIT_USER_NAME=$(charts_git_user_name)
CHARTS_GIT_USER_EMAIL=$(charts_git_user_email)

DRONE_REPO=$(drone_repo)


# 静态变量
Date=`date "+%Y-%m-%d %H:%M:%S"`
LD_FLAGS=" \
    -X 'main.Built=${Date}'   \
    -X 'main.Version=${RELEASE_VER}'"

# 编译
all_build: server_build

server_build: base-server_build admin-server_build openai-server_build volcano_build

init:
	mkdir -p ${SERVER_BINARY_DIR}

base-server_build: init
	cd ./server && go generate

	cd ./server/base-server && go build -ldflags ${LD_FLAGS} -o ${SERVER_BINARY_DIR} ./...

admin-server_build: init
	cd ./server && go generate

	cd ./server/admin-server && go build -ldflags ${LD_FLAGS} -o ${SERVER_BINARY_DIR} ./...

openai-server_build: init
	cd ./server && go generate

	cd ./server/openai-server && go build -ldflags ${LD_FLAGS} -o ${SERVER_BINARY_DIR} ./...

volcano_build: vc-controller_build scheduler_build

vc-controller_build: init
	cd ./server/volcano && go build -ldflags ${LD_FLAGS} -o ${SERVER_BINARY_DIR} ./cmd/controller-manager

scheduler_build: init
	cd ./server/volcano && go build -ldflags ${LD_FLAGS} -o ${SERVER_BINARY_DIR} ./cmd/scheduler

api-doc_build: init
	cd ./server && go generate
# 运行
all_run: server_run

server_run: base-server_run admin-server_run openai-server_run volcano_run

base-server_run:
	cd server && ./bin/base-server -conf base-server/configs &

admin-server_run:
	cd server && ./bin/admin-server -conf admin-server/configs &

openai-server_run:
	cd server && ./bin/openai-server -conf openai-server/configs &

volcano_run: vc-controller_run scheduler_run

vc-controller_run:
	cd server && ./bin/vc-controller &

scheduler_run:
	cd server && ./bin/scheduler &

# 停止
all_stop: server_stop

server_stop: base-server_stop admin-server_stop openai-server_stop volcano_stop

base-server_stop:
	kill -9 `ps -ef|grep "base-server" |grep -v grep  |awk '{print $2}'`

admin-server_stop:
	kill -9 `ps -ef|grep "admin-server" |grep -v grep |awk '{print $2}'`

openai-server_stop:
	kill -9 `ps -ef|grep "openai-server" |grep -v grep |awk '{print $2}'`

volcano_stop: vc-controller_stop scheduler_stop


vc-controller_stop:
	kill -9 `ps -ef|grep "vc-controller" |grep -v grep |awk '{print $2}'`

scheduler_stop:
	kill -9 `ps -ef|grep "scheduler" |grep -v grep |awk '{print $2}'`

# 重启
all_stop: server_restart

server_restart: base-server_restart admin-server_restart openai-server_restart volcano_restart

base-server_restart: base-server_stop server_run

admin-server_restart: admin-server_stop admin-server_run

openai-server_restart: openai-server_stop openai-server_run

volcano_restart: vc-controller_restart scheduler_restart

vc-controller_restart: vc-controller_stop vc-controller_run

scheduler_restart: scheduler_stop scheduler_run

# 代码检查
lint_init:
	golangci-lint version

lint: lint_init
	cd ./server && golangci-lint run ./...

common_lint: lint_init
	cd ./server/common && golangci-lint run ./...

base-server_lint: lint_init
	cd ./server/base-server && golangci-lint run ./...

admin-server_lint: lint_init
	cd ./server/admin-server && golangci-lint run ./...

openai-server_lint: lint_init
	cd ./server/openai-server && golangci-lint run ./...

volcano_lint: lint_init
	cd ./server/volcano && golangci-lint run ./...

# 构建镜像
images: base-server_image admin-server_image openai-server_image volcano_image admin-portal_image openai-portal_image api-doc_image node-agent_image

base-server_image:
	docker build --no-cache -t base-server:${RELEASE_VER} -f ./build/application/base-server/dockerfile .

admin-server_image:
	docker build --no-cache -t admin-server:${RELEASE_VER} -f ./build/application/admin-server/dockerfile .

openai-server_image:
	docker build --no-cache -t openai-server:${RELEASE_VER} -f ./build/application/openai-server/dockerfile .

volcano_image: vc-controller_image scheduler_image

vc-controller_image:
	docker build --no-cache -t vc-controller:${RELEASE_VER} -f ./build/application/volcano/vc-controller/dockerfile .

scheduler_image:
	docker build --no-cache -t scheduler:${RELEASE_VER} -f ./build/application/volcano/scheduler/dockerfile .

admin-portal_image:
	docker build --no-cache -t admin-portal:${RELEASE_VER} -f ./build/application/admin-portal/dockerfile .

openai-portal_image:
	docker build --no-cache -t openai-portal:${RELEASE_VER} -f ./build/application/openai-portal/dockerfile .

api-doc_image:
	docker build --no-cache -t api-doc:${RELEASE_VER} -f ./build/application/api-doc/dockerfile .

node-agent_image:
	docker build --no-cache -t node-agent:${RELEASE_VER} -f ./build/application/nodeagent/dockerfile ./controller/nodeagent

# 镜像推送
images_push: base-server_image_push admin-server_image_push openai-server_image_push volcano_image_push admin-portal_image_push openai-portal_image_push api-doc_image_push node-agent_image_push

image_push_init:
	(echo ${DOCKER_HUB_PASSWD} | docker login ${DOCKER_HUB_HOST} -u ${DOCKER_HUB_USERNAME} --password-stdin) 1>/dev/null 2>&1

base-server_image_push: image_push_init
	docker tag base-server:${RELEASE_VER} ${DOCKER_HUB_HOST}/${DOCKER_HUB_PROJECT}/base-server:${RELEASE_VER}
	docker push ${DOCKER_HUB_HOST}/${DOCKER_HUB_PROJECT}/base-server:${RELEASE_VER}

ifneq (${RELEASE_VER}, latest)
ifeq (${NEED_LATEST}, TRUE)
	docker tag base-server:${RELEASE_VER} ${DOCKER_HUB_HOST}/${DOCKER_HUB_PROJECT}/base-server:latest
	docker push ${DOCKER_HUB_HOST}/${DOCKER_HUB_PROJECT}/base-server:latest
endif
endif

admin-server_image_push: image_push_init
	docker tag admin-server:${RELEASE_VER} ${DOCKER_HUB_HOST}/${DOCKER_HUB_PROJECT}/admin-server:${RELEASE_VER}
	docker push ${DOCKER_HUB_HOST}/${DOCKER_HUB_PROJECT}/admin-server:${RELEASE_VER}

ifneq (${RELEASE_VER}, latest)
ifeq (${NEED_LATEST}, TRUE)
	docker tag admin-server:${RELEASE_VER} ${DOCKER_HUB_HOST}/${DOCKER_HUB_PROJECT}/admin-server:latest
	docker push ${DOCKER_HUB_HOST}/${DOCKER_HUB_PROJECT}/admin-server:latest
endif
endif

openai-server_image_push: image_push_init
	docker tag openai-server:${RELEASE_VER} ${DOCKER_HUB_HOST}/${DOCKER_HUB_PROJECT}/openai-server:${RELEASE_VER}
	docker push ${DOCKER_HUB_HOST}/${DOCKER_HUB_PROJECT}/openai-server:${RELEASE_VER}

ifneq (${RELEASE_VER}, latest)
ifeq (${NEED_LATEST}, TRUE)
	docker tag openai-server:${RELEASE_VER} ${DOCKER_HUB_HOST}/${DOCKER_HUB_PROJECT}/openai-server:latest
	docker push ${DOCKER_HUB_HOST}/${DOCKER_HUB_PROJECT}/openai-server:latest
endif
endif

volcano_image_push: vc-controller_image_push scheduler_image_push

vc-controller_image_push: image_push_init
	docker tag vc-controller:${RELEASE_VER} ${DOCKER_HUB_HOST}/${DOCKER_HUB_PROJECT}/vc-controller:${RELEASE_VER}
	docker push ${DOCKER_HUB_HOST}/${DOCKER_HUB_PROJECT}/vc-controller:${RELEASE_VER}

ifneq (${RELEASE_VER}, latest)
ifeq (${NEED_LATEST}, TRUE)
	docker tag vc-controller:${RELEASE_VER} ${DOCKER_HUB_HOST}/${DOCKER_HUB_PROJECT}/vc-controller:latest
	docker push ${DOCKER_HUB_HOST}/${DOCKER_HUB_PROJECT}/vc-controller:latest
endif
endif

scheduler_image_push: image_push_init
	docker tag scheduler:${RELEASE_VER} ${DOCKER_HUB_HOST}/${DOCKER_HUB_PROJECT}/scheduler:${RELEASE_VER}
	docker push ${DOCKER_HUB_HOST}/${DOCKER_HUB_PROJECT}/scheduler:${RELEASE_VER}

ifneq (${RELEASE_VER}, latest)
ifeq (${NEED_LATEST}, TRUE)
	docker tag scheduler:${RELEASE_VER} ${DOCKER_HUB_HOST}/${DOCKER_HUB_PROJECT}/scheduler:latest
	docker push ${DOCKER_HUB_HOST}/${DOCKER_HUB_PROJECT}/scheduler:latest
endif
endif

admin-portal_image_push: image_push_init
	docker tag admin-portal:${RELEASE_VER} ${DOCKER_HUB_HOST}/${DOCKER_HUB_PROJECT}/admin-portal:${RELEASE_VER}
	docker push ${DOCKER_HUB_HOST}/${DOCKER_HUB_PROJECT}/admin-portal:${RELEASE_VER}

ifneq (${RELEASE_VER}, latest)
ifeq (${NEED_LATEST}, TRUE)
	docker tag admin-portal:${RELEASE_VER} ${DOCKER_HUB_HOST}/${DOCKER_HUB_PROJECT}/admin-portal:latest
	docker push ${DOCKER_HUB_HOST}/${DOCKER_HUB_PROJECT}/admin-portal:latest
endif
endif

openai-portal_image_push: image_push_init
	docker tag openai-portal:${RELEASE_VER} ${DOCKER_HUB_HOST}/${DOCKER_HUB_PROJECT}/openai-portal:${RELEASE_VER}
	docker push ${DOCKER_HUB_HOST}/${DOCKER_HUB_PROJECT}/openai-portal:${RELEASE_VER}

ifneq (${RELEASE_VER}, latest)
ifeq (${NEED_LATEST}, TRUE)
	docker tag openai-portal:${RELEASE_VER} ${DOCKER_HUB_HOST}/${DOCKER_HUB_PROJECT}/openai-portal:latest
	docker push ${DOCKER_HUB_HOST}/${DOCKER_HUB_PROJECT}/openai-portal:latest
endif
endif

api-doc_image_push: image_push_init
	docker tag api-doc:${RELEASE_VER} ${DOCKER_HUB_HOST}/${DOCKER_HUB_PROJECT}/api-doc:${RELEASE_VER}
	docker push ${DOCKER_HUB_HOST}/${DOCKER_HUB_PROJECT}/api-doc:${RELEASE_VER}

ifneq (${RELEASE_VER}, latest)
ifeq (${NEED_LATEST}, TRUE)
	docker tag api-doc:${RELEASE_VER} ${DOCKER_HUB_HOST}/${DOCKER_HUB_PROJECT}/api-doc:latest
	docker push ${DOCKER_HUB_HOST}/${DOCKER_HUB_PROJECT}/api-doc:latest
endif
endif

node-agent_image_push: image_push_init
	docker tag node-agent:${RELEASE_VER} ${DOCKER_HUB_HOST}/${DOCKER_HUB_PROJECT}/node-agent:${RELEASE_VER}
	docker push ${DOCKER_HUB_HOST}/${DOCKER_HUB_PROJECT}/node-agent:${RELEASE_VER}

ifneq (${RELEASE_VER}, latest)
ifeq (${NEED_LATEST}, TRUE)
	docker tag node-agent:${RELEASE_VER} ${DOCKER_HUB_HOST}/${DOCKER_HUB_PROJECT}/node-agent:latest
	docker push ${DOCKER_HUB_HOST}/${DOCKER_HUB_PROJECT}/node-agent:latest
endif
endif

# helm chart
charts: charts_build charts_push

charts_build:
	-mkdir -p ./tmp/charts
	helm package ./deploy/charts/octopus --version ${RELEASE_VER} --app-version ${RELEASE_VER} -d ./tmp/charts

charts_push:
	-helm repo add --ca-file=${HARBOR_HUB_CA_FILE} --cert-file=${HARBOR_HUB_CERT_FILE} --username=${HARBOR_HUB_USERNAME} --password=${HARBOR_HUB_PASSWD} chartrepo ${HARBOR_HUB_HOST}/chartrepo/${HARBOR_HUB_PROJECT}
	helm push --ca-file=${HARBOR_HUB_CA_FILE} --cert-file=${HARBOR_HUB_CERT_FILE} --username=${HARBOR_HUB_USERNAME} --password=${HARBOR_HUB_PASSWD} ./tmp/charts/octopus-${RELEASE_VER}.tgz chartrepo
ifeq (${DRONE_REPO}, OpenI/octopus)
	git clone ${CHARTS_GIT_CLONE} ${CHARTS_GIT_DIR}
	cp ./tmp/charts/octopus-${RELEASE_VER}.tgz ${CHARTS_GIT_DIR}
	helm repo index ${CHARTS_GIT_DIR} --url ${CHARTS_GIT_RAW}
	cd ${CHARTS_GIT_DIR} && git config --global user.email ${CHARTS_GIT_USER_EMAIL} && git config --global user.name ${CHARTS_GIT_USER_NAME} && git add index.yaml octopus-${RELEASE_VER}.tgz && git commit -m "${RELEASE_VER}" && git push
endif
