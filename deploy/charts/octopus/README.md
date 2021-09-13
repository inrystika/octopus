# Octopus

**Octopus**是一款面向多计算场景的一站式融合计算平台。平台主要针对AI、HPC等场景的计算与资源管理的需求来设计，向算力使用用户提供了对数据、算法、镜像、模型与算力等资源的管理与使用功能，方便用户一站式构建计算环境，实现计算。同时，向集群管理人员提供了集群资源管理与监控，计算任务管理与监控等功能，方便集群管理人员对整体系统进行操作与分析。

**Octopus**平台底层基于容器编排平台[Kubernetes](https://kubernetes.io/zh/docs/concepts/overview/what-is-kubernetes) ，充分利用容器敏捷、轻量、隔离等特点来实现计算场景多样性的需求。

## 文档

详细文档请参考[这里](https://octopus.openi.org.cn/docs/introduction/intro)。

## 特点与场景

Octopus具有如下特点：

- **一站式开发**，为用户提供一站式AI、HPC计算场景的开发功能，通过数据管理、模型开发和模型训练，打通计算全链路；
- **方便管理**，为平台管理者提供一站式的资源管理平台，通过资源配置、监控、权限管控等可视化工具，大大降低平台管理者的管理成本；
- **易于部署**，Octopus 支持[Helm](https://helm.sh)方式的快速部署，简化复杂的部署流程；
- **性能优越**，提供高性能的分布式计算体验，通过多方面优化来保证各个环境的流畅运行，同时通过资源调度优化与分布式计算优化，进一步提高模型训练效率；
- **兼容性好**，平台支持异构硬件，如 GPU、NPU、FPGA 等，满足各种不同的硬件集群部署需求，通过支持多种深度学习框架，如 TensorFlow、Pytorch、PaddlePaddle 等，并可以通过自定义镜像方式支持新增框架。

Octopus适合在如下场景中使用：

- 构建大规模 AI 计算平台；
- 希望共享计算资源；
- 希望在统一的环境下完成模型训练；
- 希望使用集成的插件辅助模型训练，提升效率。

## 使用说明

Octopus详细使用说明请参考[这里](https://octopus.openi.org.cn/docs/deployment/deploy/quick_deploy).

## 配置说明

章鱼服务通过`Helm Charts`包格式进行版本包管理，同样基于该格式的操作方式进行相关安装更新等操作，有关`Helm Charts`具体资料可参考[这里](https://helm.sh/docs/topics/charts/).

### 参数

以下表格列出Octopus在chart安装包中`values.yaml`文件的所有可配置参数以及它们的默认值：


#### 全局参数

| Parameter                      | Description                                                                                              | Default                                                 |
|--------------------------------|----------------------------------------------------------------------------------------------------------|---------------------------------------------------------|
| `global.image.repository.address`         | 全局镜像仓库地址                                                                             | `nil`                                                   |
| `global.image.repository.pathname`      | 镜像仓库目录名                                                          | `nil`  |
| `global.image.pullPolicy`          | 镜像拉取策略                                                            | `IfNotPresent`                                                   |
| `global.nodeSelector`  | 部署节点标签 | `octopus.openi.pcl.cn/node: "server"`                                                    |

#### 通用参数

| Parameter           | Description                                                          | Default                        |
|---------------------|----------------------------------------------------------------------|--------------------------------|
| `Common.resourceTagKey`      | 资源标签关键字                   | `octopus.pcl.ac.cn/resource`                          |
| `Common.resourceTagValuePrefix`  | 资源标签键值前缀                       | `service`                          |


#### 入口参数

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

#### 数据持久化存储参数

| Parameter           | Description                                                          | Default                        |
|---------------------|----------------------------------------------------------------------|--------------------------------|
| `pv.minio.requests`      | minio服务请求存储空间                   | `100Gi`                          |
| `pv.mysql.requests`  | mysql服务请求存储空间                       | `100Gi`                          |
| `pv.redis.requests`      | redis服务请求存储空间                                | `100Gi`                           |
| `pv.logger.requests` | logger服务请求存储空间                           | `100Gi`                           |
| `pv.minio.storageType`      | minio服务存储类型                   | `nil`                          |
| `pv.mysql.storageType`  | mysql服务存储类型                       | `nil`                          |
| `pv.redis.storageType`      | redis服务存储类型                                | `nil`                           |
| `pv.logger.storageType` | logger服务存储类型                           | `nil`                           |

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

#### 数据持久化声明参数

| Parameter           | Description                                                          | Default                        |
|---------------------|----------------------------------------------------------------------|--------------------------------|
| `pvc.minio.requests`      | minio服务请求存储空间                   | `100Gi`                          |
| `pvc.mysql.requests`  | mysql服务请求存储空间                       | `100Gi`                          |
| `pvc.redis.requests`      | redis服务请求存储空间                                | `100Gi`                           |
| `pvc.logger.requests` | logger服务请求存储空间                           | `100Gi`                           |
| `pvc.logstash.requests`     | logstash服务请求存储空间                                     | `100Gi`                |


#### 服务taskset参数

| Parameter           | Description                                                          | Default                        |
|---------------------|----------------------------------------------------------------------|--------------------------------|
| `taskset.replicaCount`      | taskset服务实例数                   | `1`                          |
| `taskset.adminToken`  | 访问taskset服务的管理员token                       | `KLtmMug9BDvvRjlg`                          |
| `taskset.image.pullPolicy`      | taskset服务镜像拉取策略                                | `nil`                           |
| `taskset.image.address` | taskset服务镜像地址                           | `nil`                           |
| `taskset.image.pathname`     | taskset服务镜像目录名                                     | `nil`                |
| `taskset.image.name`     | taskset服务镜像名称                                     | `pipeline`                |

#### 服务base-server参数

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


#### 服务openai-server参数

| Parameter           | Description                                                          | Default                        |
|---------------------|----------------------------------------------------------------------|--------------------------------|
| `openaiserver.image.pullPolicy`      | openaiserver服务镜像拉取策略                                | `nil`                           |
| `openaiserver.image.address` | openaiserver服务镜像地址                           | `nil`                           |
| `openaiserver.image.pathname`     | openaiserver服务镜像目录名                                     | `nil`                |
| `openaiserver.image.name`     | openaiserver服务镜像名称                                     | `openaiserver`                |
| `openaiserver.data.redis.password`     | redis密码                                     | `abcde`                |


#### 服务admin-server参数

| Parameter           | Description                                                          | Default                        |
|---------------------|----------------------------------------------------------------------|--------------------------------|
| `adminserver.image.pullPolicy`      | adminserver服务镜像拉取策略                                | `nil`                           |
| `adminserver.image.address` | adminserver服务镜像地址                           | `nil`                           |
| `adminserver.image.pathname`     | adminserver服务镜像目录名                                     | `nil`                |
| `adminserver.image.name`     | adminserver服务镜像名称                                     | `admin-server`                |
| `adminserver.data.redis.password`     | redis密码                                     | `abcde`                |


#### 服务openai-portal参数

| Parameter           | Description                                                          | Default                        |
|---------------------|----------------------------------------------------------------------|--------------------------------|
| `openaiportal.image.pullPolicy`      | openaiportal服务镜像拉取策略                                | `nil`                           |
| `openaiportal.image.address` | openaiportal服务镜像地址                           | `nil`                           |
| `openaiportal.image.pathname`     | openaiportal服务镜像目录名                                     | `nil`                |
| `openaiportal.image.name`     | openaiportal服务镜像名称                                     | `openai-portal`                |


#### 服务admin-portal参数

| Parameter           | Description                                                          | Default                        |
|---------------------|----------------------------------------------------------------------|--------------------------------|
| `adminportal.image.pullPolicy`      | adminportal服务镜像拉取策略                                | `nil`                           |
| `adminportal.image.address` | adminportal服务镜像地址                           | `nil`                           |
| `adminportal.image.pathname`     | adminportal服务镜像目录名                                     | `nil`                |
| `adminportal.image.name`     | adminportal服务镜像名称                                     | `admin-portal`                |


#### 服务scheduler参数

| Parameter           | Description                                                          | Default                        |
|---------------------|----------------------------------------------------------------------|--------------------------------|
| `scheduler.image.name`     | scheduler服务镜像名称                                     | `scheduler`                |


#### 服务controller参数

| Parameter           | Description                                                          | Default                        |
|---------------------|----------------------------------------------------------------------|--------------------------------|
| `controller.image.name`     |controller服务镜像名称                                     | `controller`                |


#### 服务logger参数

| Parameter           | Description                                                          | Default                        |
|---------------------|----------------------------------------------------------------------|--------------------------------|
| `logger.filebeat.resources.limits.memory`      | filebeat使用内存限制                                | `200Mi`                           |
| `logger.filebeat.resources.requests.cpu` | filebeat使用CPU限制                            | `100m`                           |
| `logger.filebeat.resources.requests.memory`     | filebeat请求内存大小                                     | `100Mi`                |
| `logger.httpd.ingress.path`     | 日志服务nginx路径                                     | `/log`                |
| `logger.httpd.image.pullPolicy`     | 日志服务nginx镜像拉取策略                                     | `Always`                |


#### 服务minio参数

此部分主要基于第三方依赖包的配置，更多详细配置参考[这里](https://artifacthub.io/packages/helm/bitnami/minio).

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


#### 服务mysql参数

此部分主要基于第三方依赖包的配置，更多详细配置参考[这里](https://artifacthub.io/packages/helm/bitnami/mysql).

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



#### 服务redis参数

此部分主要基于第三方依赖包的配置，更多详细配置参考[这里](https://artifacthub.io/packages/helm/bitnami/redis).

| Parameter           | Description                                                          | Default                        |
|---------------------|----------------------------------------------------------------------|--------------------------------|
| `redis.master.service.type`      | master servie类型                                | `NodePort`                           |
| `redis.master.service.port` | master service port                            | `6379`                           |
| `redis.master.persistence.size`     | master请求持久存储大小                                     | `50Gi`                |
| `redis.master.persistence.existingClaim`     | master请求使用的PVC名                                     | `octopus-redis-pvc`                |
| `redis.auth.enabled`     | 是否需要登入认证                                    | `true`                |
| `redis.auth.password`     | 登入密码                                     | `abcde`                |
| `redis.volumePermissions.enabled`     | 是否对存储卷有管理员权限                                     | `true`                |


#### 服务nginx-ingress-controller参数

此部分主要基于第三方依赖包的配置，更多详细配置参考[这里](https://artifacthub.io/packages/helm/bitnami/nginx-ingress-controller).

| Parameter           | Description                                                          | Default                        |
|---------------------|----------------------------------------------------------------------|--------------------------------|
| `nginx-ingress-controller.nodeSelector`      | nginx服务节点选择器                                | `nginx-ingress: "yes"`                           |


#### 服务grafana参数

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


#### 服务prometheus参数

| Parameter           | Description                                                          | Default                        |
|---------------------|----------------------------------------------------------------------|--------------------------------|
| `prometheus.retentionDuration`      | 数据保留时间                                | `365d`                           |
| `prometheus.volumes.dataPath`      | 数据存储路径                                | `/applications/prometheus`                           |
| `prometheus.image.tag`      | 镜像版本号                                | `v2.0.0`                           |
| `prometheus.service.type`      | service类型                                | `NodePort`                           |
| `prometheus.service.hostPort`      | service hostPort                                | `30003`                           |
| `prometheus.service.port`      | service port | `9090`                           |
| `prometheus.service.targetPort`      | service targetPort | `9090`                           |
