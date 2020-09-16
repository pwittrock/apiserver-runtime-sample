module github.com/pwittrock/apiserver-runtime-sample

go 1.15

require (
	github.com/evanphx/json-patch v4.5.0+incompatible // indirect
	github.com/go-openapi/spec v0.19.5
	github.com/google/go-cmp v0.5.0 // indirect
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/imdario/mergo v0.3.9 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	github.com/onsi/ginkgo v1.13.0 // indirect
	github.com/pwittrock/apiserver-runtime v0.0.0-20200908143328-237f7f6cdf62
	github.com/stretchr/testify v1.6.1 // indirect
	gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f // indirect
	k8s.io/apimachinery v0.19.0
	k8s.io/client-go v0.19.0
	k8s.io/component-base v0.19.0
	k8s.io/klog v1.0.0
	k8s.io/kube-openapi v0.0.0-20200811211545-daf3cbb84823
)

replace (
	k8s.io/api => k8s.io/api v0.0.0-20200726131424-9540e4cac147
	k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20200726131235-945d4ebf362b
	k8s.io/apiserver => k8s.io/apiserver v0.0.0-20200726132555-3b0c6c609142
	k8s.io/client-go => k8s.io/client-go v0.0.0-20200726131703-36233866f1c7
	k8s.io/code-generator => k8s.io/code-generator v0.0.0-20200726131043-26c52896b75b
	k8s.io/component-base => k8s.io/component-base v0.0.0-20200726132252-a5fb6b31bf34
)
