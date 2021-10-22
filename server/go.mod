module server

go 1.16

require (
	github.com/antihax/optional v1.0.0
	github.com/bsm/redislock v0.7.1
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/envoyproxy/protoc-gen-validate v0.1.0
	github.com/evanphx/json-patch v4.9.0+incompatible
	github.com/fsnotify/fsnotify v1.4.9
	github.com/fsouza/go-dockerclient v1.7.2
	github.com/go-kratos/kratos/v2 v2.0.0-beta3
	github.com/go-oauth2/oauth2/v4 v4.4.2
	github.com/go-oauth2/redis/v4 v4.1.1
	github.com/go-redis/redis/v8 v8.10.0
	github.com/golang/groupcache v0.0.0-20200121045136-8c9f03a8e57e // indirect
	github.com/golang/protobuf v1.4.3
	github.com/google/uuid v1.1.2
	github.com/gorilla/mux v1.8.0
	github.com/imdario/mergo v0.3.11
	github.com/jinzhu/copier v0.2.5
	github.com/json-iterator/go v1.1.10
	github.com/kr/text v0.2.0 // indirect
	github.com/minio/minio-go/v7 v7.0.11
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	github.com/sony/sonyflake v1.0.0
	golang.org/x/crypto v0.0.0-20200709230013-948cd5f35899
	golang.org/x/net v0.0.0-20210410081132-afb366fc7cd1 // indirect
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d
	golang.org/x/text v0.3.6
	gonum.org/v1/gonum v0.6.2
	google.golang.org/genproto v0.0.0-20210212180131-e7f2df4ecc2d
	google.golang.org/grpc v1.36.0
	google.golang.org/protobuf v1.25.0
	gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f // indirect
	gopkg.in/errgo.v2 v2.1.0
	gopkg.in/resty.v1 v1.12.0
	gopkg.in/yaml.v2 v2.4.0
	gorm.io/driver/mysql v1.0.4
	gorm.io/gorm v1.21.7
	gorm.io/plugin/soft_delete v1.0.0
	gotest.tools v2.2.0+incompatible
	k8s.io/api v0.18.18
	k8s.io/apimachinery v0.18.18
	k8s.io/client-go v0.18.18
	volcano.sh/volcano v0.0.0-00010101000000-000000000000
)

replace (
	k8s.io/api => k8s.io/api v0.18.18
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.18.18
	k8s.io/apimachinery => k8s.io/apimachinery v0.18.18
	k8s.io/apiserver => k8s.io/apiserver v0.18.18
	k8s.io/cli-runtime => k8s.io/cli-runtime v0.18.18
	k8s.io/client-go => k8s.io/client-go v0.18.18
	k8s.io/cloud-provider => k8s.io/cloud-provider v0.18.18
	k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.18.18
	k8s.io/code-generator => k8s.io/code-generator v0.18.18
	k8s.io/component-base => k8s.io/component-base v0.18.18
	k8s.io/cri-api => k8s.io/cri-api v0.18.18
	k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.18.18
	k8s.io/klog => k8s.io/klog v1.0.0
	k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.18.18
	k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.18.18
	k8s.io/kube-proxy => k8s.io/kube-proxy v0.18.18
	k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.18.18
	k8s.io/kubectl => k8s.io/kubectl v0.18.18
	k8s.io/kubelet => k8s.io/kubelet v0.18.18
	k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.18.18
	k8s.io/metrics => k8s.io/metrics v0.18.18
	k8s.io/node-api => k8s.io/node-api v0.18.18
	k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.18.18
	k8s.io/sample-cli-plugin => k8s.io/sample-cli-plugin v0.18.18
	k8s.io/sample-controller => k8s.io/sample-controller v0.18.18
	volcano.sh/volcano => ./taskset/pkg/volcano
)
