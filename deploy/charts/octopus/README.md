# Octopus

[Octopus](http://192.168.202.74:6869/) 是一个一站式 AI 模型开发平台，面向 AI 模型生产的生命周期，提供了数据集管理、镜像管理、算法管理、训练、部署等功能，方便用户一站式构建AI算法，另外平台还提出了“工作空间”概念，满足不同用户群体的资源使用与管理述求，方便平台管理者更好的管理资源集群。平台是由【鹏城实验室】独立设计开发，并进行维护，结合了一些在大规模生产环境中表现良好的成熟设计，主要为提升学术研究效率，复现学术研究成果而量身打造。

## 简介

这个chart文件能够通过使用[Helm](https://helm.sh)包管理工具，将Octopus服务部署在[Kubernetes](http://kubernetes.io) 集群上。


## 依赖


- Kubernetes v1.16.3+
- Docker v18.09.7+
- Helm v3.5.4+
- Harbor v2.3.0+


## 命令行部署Chart安装包

用octopus作为安装参数release-name进行安装说明：

### 证书设置

为了保证所有节点均能从Harbor仓库获取镜像，需要在所有的节点做证书的配置，这里有两种方式，选择其中一种就行，如下：

#### 校验证书方式

先获取证书，而后将证书存放到docker配置目录下，如下
```console
sudo su -
mkdir -p /etc/docker/certs.d/192.168.202.110:5000
cp ./credentials/harbor/192.168.202.110:5000/* /etc/docker/certs.d/192.168.202.110:5000/
```
然后重启docker:
```console
systemctl restart docker
```
#### 不校验证书方式

直接修改docker配置,跳过对harbor的证书验证
```console
sudo su -
vim /etc/docker/daemon.json

#添加属性
# {
#	"insecure-registries": [...,"192.168.202.110:5000",...]
# }

systemctl restart docker
```

### 添加Chart仓库

这里主要使用Harbor作为Chart仓库,启动https的仓库需要添加证书文件，可通过这里获取[证书] (http://192.168.202.74/octopus/credentials)，如下：

```console
helm repo add --ca-file /path/to/ca.crt --cert-file /path/to/192.168.202.110.cert --username={username} --password={password} harbor https://192.168.202.110:5000/chartrepo/octopus
```
添加成功后同步仓库信息，如下：
```console
helm repo update
```

### 安装Octopus Chart

通过Helm命令安装，｀--debug --dry-run｀为调试模式，只会输出yaml不传给k8s,真正安装需要去掉．如下：
```console
helm install --ca-file /path/to/ca.crt --cert-file /path/to/192.168.202.110.cert --username={username} --password={password} octopus harbor/octopus --version {chart version} --debug --dry-run --values /path/to/custom/values.yaml
```

### 更新Octopus

通过Helm命令更新．如下：
```console
helm upgrade --ca-file /path/to/ca.crt --cert-file /path/to/192.168.202.110.cert --username={username} --password={password} octopus harbor/octopus --version {chart version} --values /path/to/custom/values.yaml
```

## 卸载chart安装包

命令行下执行以下命令：
```console
$ helm delete octopus
```

## 参数

以下表格列出Octopus在chart安装包中所有的可配置参数以及它们的默认值：


### Global 参数

| Parameter                      | Description                                                                                              | Default                                                 |
|--------------------------------|----------------------------------------------------------------------------------------------------------|---------------------------------------------------------|
| `global.image.repository.address`         | 全局镜像仓库地址                                                                             | `nil`                                                   |
| `global.image.repository.pathname`      | 镜像仓库目录名                                                          | `nil`  |
| `global.image.pullPolicy`          | 镜像拉取策略                                                            | `IfNotPresent`                                                   |
| `global.nodeSelector`  | 部署节点标签 | `octopus.openi.pcl.cn/node: "server"`                                                    |

### Common 参数

| Parameter           | Description                                                          | Default                        |
|---------------------|----------------------------------------------------------------------|--------------------------------|
| `Common.resourceTagKey`      | 资源标签关键字                   | `octopus.pcl.ac.cn/resource`                          |
| `Common.resourceTagValuePrefix`  | 资源标签键值前缀                       | `service`                          |


### ingress 参数

| Parameter           | Description                                                          | Default                        |
|---------------------|----------------------------------------------------------------------|--------------------------------|
| `ingress.enabled`      | 是否开启ingress                   | `true`                          |
| `ingress.adminserverPath`  | adminserver服务路径                       | `/adminserver`                          |
| `ingress.openaiserverPath`      | openaiserver服务路径                                | `/openaiserver`                           |
| `ingress.adminportalPath` | adminportal服务路径                           | `/adminportal`                           |
| `ingress.openaiportalPath`     | openaiportal服务路径                                     | `/`                |
| `ingress.loggerHttpdPath`       | 日志服务路径                    | `/log` |
| `ingress.minioPath.web`       | minio服务网页端入口路径 | `/minio`                          |
| `ingress.minioPath.api`       | minio服务接口路径 | `/oss`                          |

### PersistentVolumn 参数

| Parameter           | Description                                                          | Default                        |
|---------------------|----------------------------------------------------------------------|--------------------------------|
| `pv.minio.requests`      | minio服务请求存储空间                   | `100Gi`                          |
| `pv.mysql.requests`  | mysql服务请求存储空间                       | `100Gi`                          |
| `pv.redis.requests`      | redis服务请求存储空间                                | `100Gi`                           |
| `pv.logger.requests` | logger服务请求存储空间                           | `100Gi`                           |
| `pv.logstash.requests`     | logstash服务请求存储空间                                     | `100Gi`                |
| `pv.minio.storageType`      | minio服务存储类型                   | `nil`                          |
| `pv.mysql.storageType`  | mysql服务存储类型                       | `nil`                          |
| `pv.redis.storageType`      | redis服务存储类型                                | `nil`                           |
| `pv.logger.storageType` | logger服务存储类型                           | `nil`                           |
| `pv.logstash.storageType`     | logstash服务存储类型                                     | `nil`                |

```
注意：storageType属性取值可参见PersistentVolumeSource结构体中Spec属性的取值。
若以本地存储，则可配置为：
storageType:
    hostPath:
        path: /test
若以nfs存储，则可配置为:
storageType:
    nfs:
        server:  192.168.203.72
        path:  /test
```

### PersistentVolumnClaim 参数

| Parameter           | Description                                                          | Default                        |
|---------------------|----------------------------------------------------------------------|--------------------------------|
| `pvc.minio.requests`      | minio服务请求存储空间                   | `100Gi`                          |
| `pvc.mysql.requests`  | mysql服务请求存储空间                       | `100Gi`                          |
| `pvc.redis.requests`      | redis服务请求存储空间                                | `100Gi`                           |
| `pvc.logger.requests` | logger服务请求存储空间                           | `100Gi`                           |
| `pvc.logstash.requests`     | logstash服务请求存储空间                                     | `100Gi`                |


### taskset 参数

| Parameter           | Description                                                          | Default                        |
|---------------------|----------------------------------------------------------------------|--------------------------------|
| `taskset.replicaCount`      | taskset服务实例数                   | `1`                          |
| `taskset.adminToken`  | 访问taskset服务的管理员token                       | `KLtmMug9BDvvRjlg`                          |
| `taskset.image.pullPolicy`      | taskset服务镜像拉取策略                                | `nil`                           |
| `taskset.image.address` | taskset服务镜像地址                           | `nil`                           |
| `taskset.image.pathname`     | taskset服务镜像目录名                                     | `nil`                |
| `taskset.image.name`     | taskset服务镜像名称                                     | `pipeline`                |

### base-server 参数

| Parameter           | Description                                                          | Default                        |
|---------------------|----------------------------------------------------------------------|--------------------------------|
| `baseserver.image.pullPolicy`      | baseserver服务镜像拉取策略                                | `nil`                           |
| `baseserver.image.address` | baseserver服务镜像地址                           | `nil`                           |
| `baseserver.image.pathname`     | baseserver服务镜像目录名                                     | `nil`                |
| `baseserver.image.name`     | baseserver服务镜像名称                                     | `base-server`                |
| `baseserver.data.minio.base.accessKeyID`      | minio访问accessKeyID                                | `minioadmin`                           |
| `baseserver.data.minio.base.secretAccessKey` | minio访问secretAccessKey                           | `minioadmin`                           |
| `baseserver.data.minio.base.useSSL`     | 是否使用SSL访问minio                                     | `false`                |
| `baseserver.data.minio.base.mountPath`     | baseserver服务容器存储挂载路径                                     | `/data`                |
| `baseserver.data.harbor.host`      | harbor服务地址                                | `192.168.202.74:5000`                           |
| `baseserver.data.harbor.username` | harbor用户名                           | `openi`                           |
| `baseserver.data.harbor.password`     | harbor密码                                     | `OpenI2018`                |
| `baseserver.data.harbor.useSSL`     | 是否使用SSL访问harbor                                     | `false`                |
| `baseserver.data.redis.password`     | redis密码                                     | `abcde`                |
| `baseserver.data.harbor.useSSL`     | 是否使用SSL访问harbor                                     | `false`                |
| `baseserver.administrator.username`     | 管理员用户名                                     | `admin`                |
| `baseserver.administrator.password`     | 管理员密码                                     | `123456`                |


### openai-server 参数

| Parameter           | Description                                                          | Default                        |
|---------------------|----------------------------------------------------------------------|--------------------------------|
| `openaiserver.image.pullPolicy`      | openaiserver服务镜像拉取策略                                | `nil`                           |
| `openaiserver.image.address` | openaiserver服务镜像地址                           | `nil`                           |
| `openaiserver.image.pathname`     | openaiserver服务镜像目录名                                     | `nil`                |
| `openaiserver.image.name`     | openaiserver服务镜像名称                                     | `openaiserver`                |
| `openaiserver.data.redis.password`     | redis密码                                     | `abcde`                |


### admin-server 参数

| Parameter           | Description                                                          | Default                        |
|---------------------|----------------------------------------------------------------------|--------------------------------|
| `adminserver.image.pullPolicy`      | adminserver服务镜像拉取策略                                | `nil`                           |
| `adminserver.image.address` | adminserver服务镜像地址                           | `nil`                           |
| `adminserver.image.pathname`     | adminserver服务镜像目录名                                     | `nil`                |
| `adminserver.image.name`     | adminserver服务镜像名称                                     | `admin-server`                |
| `adminserver.data.redis.password`     | redis密码                                     | `abcde`                |


### openai-portal 参数

| Parameter           | Description                                                          | Default                        |
|---------------------|----------------------------------------------------------------------|--------------------------------|
| `openaiportal.image.pullPolicy`      | openaiportal服务镜像拉取策略                                | `nil`                           |
| `openaiportal.image.address` | openaiportal服务镜像地址                           | `nil`                           |
| `openaiportal.image.pathname`     | openaiportal服务镜像目录名                                     | `nil`                |
| `openaiportal.image.name`     | openaiportal服务镜像名称                                     | `openai-portal`                |


### admin-portal 参数

| Parameter           | Description                                                          | Default                        |
|---------------------|----------------------------------------------------------------------|--------------------------------|
| `adminportal.image.pullPolicy`      | adminportal服务镜像拉取策略                                | `nil`                           |
| `adminportal.image.address` | adminportal服务镜像地址                           | `nil`                           |
| `adminportal.image.pathname`     | adminportal服务镜像目录名                                     | `nil`                |
| `adminportal.image.name`     | adminportal服务镜像名称                                     | `admin-portal`                |


### scheduler 参数

| Parameter           | Description                                                          | Default                        |
|---------------------|----------------------------------------------------------------------|--------------------------------|
| `scheduler.image.name`     | scheduler服务镜像名称                                     | `scheduler`                |


### controller 参数

| Parameter           | Description                                                          | Default                        |
|---------------------|----------------------------------------------------------------------|--------------------------------|
| `controller.image.name`     |controller服务镜像名称                                     | `controller`                |


### logger 参数

| Parameter           | Description                                                          | Default                        |
|---------------------|----------------------------------------------------------------------|--------------------------------|
| `logger.filebeat.resources.limits.memory`      | filebeat使用内存限制                                | `200Mi`                           |
| `logger.filebeat.resources.requests.cpu` | filebeat使用CPU限制                            | `100m`                           |
| `logger.filebeat.resources.requests.memory`     | filebeat请求内存大小                                     | `100Mi`                |
| `logger.httpd.ingress.path`     | 日志服务nginx路径                                     | `/log`                |
| `logger.httpd.image.pullPolicy`     | 日志服务nginx镜像拉取策略                                     | `Always`                |


### minio 参数

| Parameter           | Description                                                          | Default                        |
|---------------------|----------------------------------------------------------------------|--------------------------------|
| `minio.gateway.enabled`      | 开启网关                                | `true`                           |
| `minio.gateway.type` | 网关类型                            | `nas`                           |
| `minio.gateway.auth.nas.accessKey`     | nas网关登入账号                                     | `minioadmin`                |
| `minio.gateway.auth.nas.secretKey`     | nas网关登入密码                                     | `minioadmin`                |
| `minio.accessKey.password`     | 账号                                     | `minioadmin`                |
| `minio.secretKey.password`     | 密码                                     | `minioadmin`                |
| `minio.resources.requests.memory`     | 请求内存大小                                     | `1Gi`                |
| `minio.persistence.size`     | 请求持久存储大小                                     | `100Gi`                |
| `minio.persistence.existingClaim`     | 请求使用的PVC名称                                     | `octopus-minio-pvc`                |
| `minio.service.type`     | service类型                                     | `NodePort`                |
| `minio.service.nodePort`     | nodePort                                     | `31311`                |
| `minio.service.port`     | service port                                     | `9000`                |
| `minio.volumePermissions.enabled`     | 是否对存储卷有管理员权限                                     | `true`                |


### mysql 参数

| Parameter           | Description                                                          | Default                        |
|---------------------|----------------------------------------------------------------------|--------------------------------|
| `mysql.auth.rootPassword`      | root账号密码                                | `root`                           |
| `mysql.volumePermissions.enabled` | 是否对存储卷有管理员权限                            | `true`                           |
| `mysql.primary.service.type`     | servie类型                                     | `NodePort`                |
| `mysql.primary.service.port`     | service port                                     | `3306`                |
| `mysql.primary.service.nodePort`     | nodePort                                     | `30336`                |
| `mysql.primary.persistence.size`     | 请求持久存储大小                                     | `100Gi`                |
| `mysql.primary.persistence.existingClaim`     | 请求使用的PVC名称                                     | `octopus-mysql-pvc`                |
| `mysql.primary.extraVolumeMounts.name`     | 附加挂载存储卷名称                                    | `mysql-initdb`                |
| `mysql.primary.extraVolumeMounts.mountPath`     | 附加挂载存储卷路径                                     | `/docker-entrypoint-initdb.d`                |
| `mysql.primary.extraVolumes.name`     | 附加存储卷名称                                     | `mysql-initdb`                |
| `mysql.primary.extraVolumes.configMap.name`     | 附加存储卷读取的configMap名称                                     | `mysql-initdb-config`                |



### redis 参数

| Parameter           | Description                                                          | Default                        |
|---------------------|----------------------------------------------------------------------|--------------------------------|
| `redis.master.service.type`      | master servie类型                                | `NodePort`                           |
| `redis.master.service.port` | master service port                            | `6379`                           |
| `redis.master.persistence.size`     | master请求持久存储大小                                     | `50Gi`                |
| `redis.master.persistence.existingClaim`     | master请求使用的PVC名                                     | `octopus-redis-pvc`                |
| `redis.auth.enabled`     | 是否需要登入认证                                    | `true`                |
| `redis.auth.password`     | 登入密码                                     | `abcde`                |
| `mysql.volumePermissions.enabled`     | 是否对存储卷有管理员权限                                     | `true`                |


### nginx-ingress-controller 参数

| Parameter           | Description                                                          | Default                        |
|---------------------|----------------------------------------------------------------------|--------------------------------|
| `nginx-ingress-controller.nodeSelector`      | nginx服务节点选择器                                | `nginx-ingress: "yes"`                           |


### grafana 参数

| Parameter           | Description                                                          | Default                        |
|---------------------|----------------------------------------------------------------------|--------------------------------|
| `grafana.image.tag`      | grafana镜像版本号                                | `nginx-ingress: "6.2.0"`                           |
| `grafana.env.GF_AUTH_BASIC_ENABLED`      | 环境变量GF_AUTH_BASIC_ENABLED                                | `true`                           |
| `grafana.env.GF_AUTH_ANONYMOUS_ENABLED`      | 环境变量GF_AUTH_ANONYMOUS_ENABLED                                | `Viewer`                           |
| `grafana.env.GF_AUTH_ANONYMOUS_ORG_ROLE`      | 环境变量GF_AUTH_ANONYMOUS_ORG_ROLE                                | `admin`                           |
| `grafana.env.GF_SECURITY_ADMIN_USER`      | 环境变量GF_SECURITY_ADMIN_USER                                | `Pa22word`                           |
| `grafana.env.GF_SECURITY_ADMIN_PASSWORD`      | 环境变量GF_SECURITY_ADMIN_PASSWORD                                | `true`                           |
| `grafana.env.GF_SECURITY_ALLOW_EMBEDDING`      | 环境变量GF_SECURITY_ALLOW_EMBEDDING                                | `/grafana`                           |
| `grafana.prometheus.protocol`      | 访问prometheus协议                                | `http`                           |
| `grafana.prometheus.port`      | 访问prometheus端口                                | `9090`                           |
| `grafana.service.type`      | service类型                                | `ClusterIP`                           |
| `grafana.service.port`      | service port                                | `3000`                           |
| `grafana.service.targetPort`      | service targetPort                                | `3000`                           |
| `grafana.ingress.enabled`      | 开启ingress                                | `true`                           |
| `grafana.ingress.path`      | nginx访问路径                                | `/grafana`                           |
| `grafana.resources.limits.cpu`      | CPU资源限制                                | `100m`                           |
| `grafana.resources.limits.memory`      | 内存资源限制                                | `/100Mi`                           |
| `grafana.resources.requests.cpu`      | CPU资源请求                                | `100m`                           |
| `grafana.resources.requests.memory`      | 内存资源请求                                | `/100Mi`                           |


### prometheus 参数

| Parameter           | Description                                                          | Default                        |
|---------------------|----------------------------------------------------------------------|--------------------------------|
| `prometheus.retentionDuration`      | 数据保留时间                                | `365d`                           |
| `prometheus.volumes.dataPath`      | 数据存储路径                                | `/applications/prometheus`                           |
| `prometheus.image.tag`      | 镜像版本号                                | `v2.0.0`                           |
| `prometheus.service.type`      | service类型                                | `NodePort`                           |
| `prometheus.service.hostPort`      | service hostPort                                | `30003`                           |
| `prometheus.service.port`      | service port | `9090`                           |
| `prometheus.service.targetPort`      | service targetPort | `9090`                           |



