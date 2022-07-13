module server

go 1.16

require (
	github.com/antihax/optional v1.0.0
	github.com/bsm/redislock v0.7.1
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/envoyproxy/protoc-gen-validate v0.1.0
	github.com/evanphx/json-patch v4.12.0+incompatible
	github.com/fsnotify/fsnotify v1.5.1
	github.com/fsouza/go-dockerclient v1.7.2
	github.com/go-kratos/kratos/v2 v2.0.0-beta3
	github.com/go-redis/redis/v8 v8.10.0
	github.com/golang/protobuf v1.5.2
	github.com/google/uuid v1.2.0
	github.com/gorilla/mux v1.8.0
	github.com/imdario/mergo v0.3.12
	github.com/influxdata/influxdb v1.9.4
	github.com/jinzhu/copier v0.2.5
	github.com/minio/minio-go/v7 v7.0.11
	github.com/prometheus/client_golang v1.12.0
	github.com/seldonio/seldon-core/operator v1.11.2
	github.com/sony/sonyflake v1.0.0
	github.com/spf13/pflag v1.0.5
	go.uber.org/automaxprocs v1.4.0
	golang.org/x/crypto v0.0.0-20210817164053-32db794688a5
	golang.org/x/oauth2 v0.0.0-20210819190943-2bc19b11175f
	golang.org/x/text v0.3.7
	golang.org/x/time v0.0.0-20210723032227-1f47c861a9ac
	gonum.org/v1/gonum v0.8.2
	google.golang.org/genproto v0.0.0-20210831024726-fe130286e0e2
	google.golang.org/grpc v1.40.0
	google.golang.org/protobuf v1.27.1
	gopkg.in/errgo.v2 v2.1.0
	gopkg.in/yaml.v2 v2.4.0
	gorm.io/driver/mysql v1.0.4
	gorm.io/gorm v1.21.7
	gorm.io/plugin/soft_delete v1.0.0
	gotest.tools v2.2.0+incompatible
	k8s.io/api v0.23.0
	k8s.io/apimachinery v0.23.0
	k8s.io/apiserver v0.23.0
	k8s.io/client-go v12.0.0+incompatible
	k8s.io/component-base v0.23.0
	k8s.io/klog v1.0.0
	k8s.io/kubernetes v1.23.0
	k8s.io/utils v0.0.0-20210930125809-cb0fa318a74b
	nodeagent v0.0.0-00010101000000-000000000000
	volcano.sh/apis v1.6.0
	volcano.sh/volcano v1.6.0
)

replace (
	github.com/opencontainers/runc => github.com/opencontainers/runc v1.0.3
	k8s.io/api => k8s.io/api v0.23.0
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.23.0
	k8s.io/apimachinery => k8s.io/apimachinery v0.23.0
	k8s.io/apiserver => k8s.io/apiserver v0.23.0
	k8s.io/cli-runtime => k8s.io/cli-runtime v0.23.0
	k8s.io/client-go => k8s.io/client-go v0.23.0
	k8s.io/cloud-provider => k8s.io/cloud-provider v0.23.0
	k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.23.0
	k8s.io/code-generator => k8s.io/code-generator v0.23.0
	k8s.io/component-base => k8s.io/component-base v0.23.0
	k8s.io/component-helpers => k8s.io/component-helpers v0.23.0
	k8s.io/controller-manager => k8s.io/controller-manager v0.23.0
	k8s.io/cri-api => k8s.io/cri-api v0.23.0
	k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.23.0
	k8s.io/klog => k8s.io/klog v1.0.0
	k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.23.0
	k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.23.0
	k8s.io/kube-proxy => k8s.io/kube-proxy v0.23.0
	k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.23.0
	k8s.io/kubectl => k8s.io/kubectl v0.23.0
	k8s.io/kubelet => k8s.io/kubelet v0.23.0
	k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.23.0
	k8s.io/metrics => k8s.io/metrics v0.23.0
	k8s.io/mount-utils => k8s.io/mount-utils v0.23.0
	k8s.io/node-api => k8s.io/node-api v0.23.0
	k8s.io/pod-security-admission => k8s.io/pod-security-admission v0.23.0
	k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.23.0
	k8s.io/sample-cli-plugin => k8s.io/sample-cli-plugin v0.23.0
	k8s.io/sample-controller => k8s.io/sample-controller v0.23.0
	nodeagent => ../controller/nodeagent
	sigs.k8s.io/controller-runtime => sigs.k8s.io/controller-runtime v0.11.0
)
