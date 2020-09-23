# First Operator

```bash
huzhi@localhost ~ % kind create cluster
Creating cluster "kind" ...
 ✓ Ensuring node image (kindest/node:v1.19.1) 🖼
 ✓ Preparing nodes 📦
 ✓ Writing configuration 📜
 ✓ Starting control-plane 🕹️
 ✓ Installing CNI 🔌
 ✓ Installing StorageClass 💾
Set kubectl context to "kind-kind"
You can now use your cluster with:

kubectl cluster-info --context kind-kind

Have a question, bug, or feature request? Let us know! https://kind.sigs.k8s.io/#community 🙂
huzhi@localhost ~ %

huzhi@localhost imooc % pwd
/Users/huzhi/work/code/go_code/imooc
huzhi@localhost imooc %
huzhi@localhost imooc % operator-sdk new imooc-operator --skip-validation=true --repo=github.com/imooc-com/imooc-operator
INFO[0000] Creating new Go operator 'imooc-operator'.
INFO[0000] Created go.mod
INFO[0000] Created tools.go
INFO[0000] Created cmd/manager/main.go
INFO[0000] Created build/Dockerfile
INFO[0000] Created build/bin/entrypoint
INFO[0000] Created build/bin/user_setup
INFO[0000] Created deploy/service_account.yaml
INFO[0000] Created deploy/role.yaml
INFO[0000] Created deploy/role_binding.yaml
INFO[0000] Created deploy/operator.yaml
INFO[0000] Created pkg/apis/apis.go
INFO[0000] Created pkg/controller/controller.go
INFO[0000] Created version/version.go
INFO[0000] Created .gitignore
INFO[0000] Project creation complete.
huzhi@localhost imooc %
huzhi@localhost imooc % ll
total 0
drwxr-xr-x   3 huzhi  staff   96  9 22 13:48 ./
drwxr-xr-x   3 huzhi  staff   96  9 22 13:46 ../
drwxr-x---  10 huzhi  staff  320  9 22 13:48 imooc-operator/
huzhi@localhost imooc %
huzhi@localhost imooc % tree -a imooc-operator
imooc-operator
├── .gitignore
├── build
│   ├── Dockerfile
│   └── bin
│       ├── entrypoint
│       └── user_setup
├── cmd
│   └── manager
│       └── main.go
├── deploy
│   ├── operator.yaml
│   ├── role.yaml
│   ├── role_binding.yaml
│   └── service_account.yaml
├── go.mod
├── pkg
│   ├── apis
│   │   └── apis.go
│   └── controller
│       └── controller.go
├── tools.go
└── version
    └── version.go

9 directories, 14 files
huzhi@localhost imooc %

huzhi@localhost imooc-operator % go env
GO111MODULE=""
GOARCH="amd64"
GOBIN=""
GOCACHE="/Users/huzhi/Library/Caches/go-build"
GOENV="/Users/huzhi/Library/Application Support/go/env"
GOEXE=""
GOFLAGS=""
GOHOSTARCH="amd64"
GOHOSTOS="darwin"
GOINSECURE=""
GOMODCACHE="/Users/huzhi/go/pkg/mod"
GONOPROXY=""
GONOSUMDB=""
GOOS="darwin"
GOPATH="/Users/huzhi/go"
GOPRIVATE=""
GOPROXY="https://proxy.golang.org,direct"
GOROOT="/usr/local/go"
GOSUMDB="sum.golang.org"
GOTMPDIR=""
GOTOOLDIR="/usr/local/go/pkg/tool/darwin_amd64"
GCCGO="gccgo"
AR="ar"
CC="clang"
CXX="clang++"
CGO_ENABLED="1"
GOMOD="/Users/huzhi/work/code/go_code/imooc/imooc-operator/go.mod"
CGO_CFLAGS="-g -O2"
CGO_CPPFLAGS=""
CGO_CXXFLAGS="-g -O2"
CGO_FFLAGS="-g -O2"
CGO_LDFLAGS="-g -O2"
PKG_CONFIG="pkg-config"
GOGCCFLAGS="-fPIC -m64 -pthread -fno-caret-diagnostics -Qunused-arguments -fmessage-length=0 -fdebug-prefix-map=/var/folders/4f/gy92m6hd2xj3c6_yz0dht6880000gn/T/go-build588738101=/tmp/go-build -gno-record-gcc-switches -fno-common"
huzhi@localhost imooc-operator %
huzhi@localhost imooc-operator % ll /Users/huzhi/go
ls: /Users/huzhi/go: No such file or directory
huzhi@localhost imooc-operator %


huzhi@localhost imooc-operator % pwd
/Users/huzhi/work/code/go_code/imooc/imooc-operator
huzhi@localhost imooc-operator %
huzhi@localhost imooc-operator % operator-sdk add api --api-version=k8s.imooc.com/v1alpha1 --kind=ImoocPod
INFO[0000] Generating api version k8s.imooc.com/v1alpha1 for kind ImoocPod.
INFO[0000] Created pkg/apis/k8s/group.go
INFO[0060] Created pkg/apis/k8s/v1alpha1/imoocpod_types.go
INFO[0060] Created pkg/apis/addtoscheme_k8s_v1alpha1.go
INFO[0060] Created pkg/apis/k8s/v1alpha1/register.go
INFO[0060] Created pkg/apis/k8s/v1alpha1/doc.go
INFO[0060] Created deploy/crds/k8s.imooc.com_v1alpha1_imoocpod_cr.yaml
INFO[0060] Running deepcopy code-generation for Custom Resource group versions: [k8s:[v1alpha1], ]
F0922 14:08:47.478086   10952 deepcopy.go:885] Hit an unsupported type invalid type for invalid type, from ./pkg/apis/k8s/v1alpha1.ImoocPod
huzhi@localhost imooc-operator %
huzhi@localhost imooc-operator % git status
On branch master
Changes not staged for commit:
  (use "git add <file>..." to update what will be committed)
  (use "git checkout -- <file>..." to discard changes in working directory)

	modified:   deploy/role.yaml

Untracked files:
  (use "git add <file>..." to include in what will be committed)

	deploy/crds/
	pkg/apis/addtoscheme_k8s_v1alpha1.go
	pkg/apis/k8s/

no changes added to commit (use "git add" and/or "git commit -a")
huzhi@localhost imooc-operator %

huzhi@localhost imooc-operator % operator-sdk add controller --api-version=k8s.imooc.com/v1alpha1 --kind=ImoocPod
INFO[0000] Generating controller version k8s.imooc.com/v1alpha1 for kind ImoocPod.
INFO[0000] Created pkg/controller/imoocpod/imoocpod_controller.go
INFO[0000] Created pkg/controller/add_imoocpod.go
INFO[0000] Controller generation complete.
huzhi@localhost imooc-operator %
huzhi@localhost imooc-operator % git status
On branch master
Untracked files:
  (use "git add <file>..." to include in what will be committed)

	pkg/controller/add_imoocpod.go
	pkg/controller/imoocpod/

nothing added to commit but untracked files present (use "git add" to track)
huzhi@localhost imooc-operator %
huzhi@localhost imooc-operator % export GOPROXY=https://goproxy.io
huzhi@localhost imooc-operator % go env GOPROXY
https://goproxy.io
huzhi@localhost imooc-operator % go mod download github.com/operator-framework/operator-sdk@v0.17.0
huzhi@localhost imooc-operator % ll /Users/huzhi/go
total 0
drwxr-xr-x   3 huzhi  staff    96  9 22 14:33 ./
drwxr-xr-x+ 38 huzhi  staff  1216  9 22 14:33 ../
drwxr-xr-x   4 huzhi  staff   128  9 22 14:33 pkg/
huzhi@localhost imooc-operator % ll /Users/huzhi/go/pkg
total 0
drwxr-xr-x  4 huzhi  staff  128  9 22 14:33 ./
drwxr-xr-x  3 huzhi  staff   96  9 22 14:33 ../
drwxr-xr-x  4 huzhi  staff  128  9 22 14:34 mod/
drwxr-xr-x  3 huzhi  staff   96  9 22 14:33 sumdb/
huzhi@localhost imooc-operator % ll /Users/huzhi/go/pkg/mod
total 0
drwxr-xr-x  4 huzhi  staff  128  9 22 14:34 ./
drwxr-xr-x  4 huzhi  staff  128  9 22 14:33 ../
drwxr-xr-x  4 huzhi  staff  128  9 22 14:34 cache/
drwxr-xr-x  3 huzhi  staff   96  9 22 14:34 github.com/
huzhi@localhost imooc-operator % ll /Users/huzhi/go/pkg/mod/github.com/operator-framework/operator-sdk@v0.17.0
total 600
dr-x------  28 huzhi  staff     896  9 22 14:34 ./
drwxr-xr-x   3 huzhi  staff      96  9 22 14:34 ../
dr-xr-xr-x   5 huzhi  staff     160  9 22 14:34 .github/
-r--r--r--   1 huzhi  staff    1427  9 22 14:34 .gitignore
-r--r--r--   1 huzhi  staff     107  9 22 14:34 .gitmodules
-r--r--r--   1 huzhi  staff    7897  9 22 14:34 .travis.yml
-r--r--r--   1 huzhi  staff   81069  9 22 14:34 CHANGELOG.md
-r--r--r--   1 huzhi  staff    2911  9 22 14:34 CONTRIBUTING.MD
-r--r--r--   1 huzhi  staff   11357  9 22 14:34 LICENSE
-r--r--r--   1 huzhi  staff     657  9 22 14:34 MAINTAINERS
-r--r--r--   1 huzhi  staff    8318  9 22 14:34 Makefile
-r--r--r--   1 huzhi  staff     328  9 22 14:34 OWNERS
-r--r--r--   1 huzhi  staff    9064  9 22 14:34 README.md
-r--r--r--   1 huzhi  staff     348  9 22 14:34 SECURITY.md
dr-xr-xr-x   3 huzhi  staff      96  9 22 14:34 cmd/
-r--r--r--   1 huzhi  staff    2971  9 22 14:34 code-of-conduct.md
dr-xr-xr-x  17 huzhi  staff     544  9 22 14:34 doc/
dr-xr-xr-x   3 huzhi  staff      96  9 22 14:34 example/
-r--r--r--   1 huzhi  staff    2386  9 22 14:34 go.mod
-r--r--r--   1 huzhi  staff  142549  9 22 14:34 go.sum
dr-xr-xr-x  11 huzhi  staff     352  9 22 14:34 hack/
dr-xr-xr-x   4 huzhi  staff     128  9 22 14:34 images/
dr-xr-xr-x   9 huzhi  staff     288  9 22 14:34 internal/
dr-xr-xr-x  17 huzhi  staff     544  9 22 14:34 pkg/
-r--r--r--   1 huzhi  staff    2140  9 22 14:34 release.sh
dr-xr-xr-x   6 huzhi  staff     192  9 22 14:34 test/
dr-xr-xr-x   3 huzhi  staff      96  9 22 14:34 version/
dr-xr-xr-x   6 huzhi  staff     192  9 22 14:34 website/
huzhi@localhost imooc-operator %

huzhi@localhost imooc-operator % operator-sdk build rubinus/imoocpod-operator
go: downloading sigs.k8s.io/controller-runtime v0.5.2
go: downloading k8s.io/client-go v0.17.4
go: downloading k8s.io/apimachinery v0.17.4
go: downloading k8s.io/klog v1.0.0
go: downloading k8s.io/kube-state-metrics v1.7.2
go: downloading github.com/spf13/pflag v1.0.5
go: downloading k8s.io/api v0.17.4
go: downloading github.com/gogo/protobuf v1.3.1
go: downloading go.uber.org/zap v1.14.1
go: downloading github.com/google/gofuzz v1.0.0
go: downloading golang.org/x/oauth2 v0.0.0-20190604053449-0f29369cfe45
go: downloading github.com/go-logr/logr v0.1.0
go: downloading golang.org/x/net v0.0.0-20200226121028-0de0cce0169b
go: downloading github.com/golang/groupcache v0.0.0-20191027212112-611e8accdfc9
go: downloading github.com/imdario/mergo v0.3.7
go: downloading github.com/evanphx/json-patch v4.5.0+incompatible
go: downloading golang.org/x/time v0.0.0-20191024005414-555d28b269f0
go: downloading github.com/go-logr/zapr v0.1.1
go: downloading k8s.io/utils v0.0.0-20191114200735-6ca3b61696b6
go: downloading github.com/gophercloud/gophercloud v0.6.0
go: downloading github.com/pkg/errors v0.9.1
go: downloading gomodules.xyz/jsonpatch/v2 v2.0.1
go: downloading github.com/davecgh/go-spew v1.1.1
go: downloading github.com/coreos/prometheus-operator v0.38.0
go: downloading sigs.k8s.io/yaml v1.1.0
go: downloading github.com/googleapis/gnostic v0.3.1
go: downloading github.com/prometheus/client_golang v1.5.1
go: downloading cloud.google.com/go v0.49.0
go: downloading golang.org/x/text v0.3.2
go: downloading github.com/hashicorp/golang-lru v0.5.3
go: downloading github.com/Azure/go-autorest v13.3.2+incompatible
go: downloading github.com/Azure/go-autorest/autorest v0.9.3-0.20191028180845-3492b2aff503
go: downloading go.uber.org/atomic v1.6.0
go: downloading github.com/prometheus/procfs v0.0.8
go: downloading github.com/golang/protobuf v1.3.2
go: downloading github.com/Azure/go-autorest/logger v0.1.0
go: downloading github.com/prometheus/common v0.9.1
go: downloading github.com/Azure/go-autorest/tracing v0.5.0
go: downloading gopkg.in/yaml.v2 v2.2.8
go: downloading github.com/prometheus/client_model v0.2.0
go: downloading gopkg.in/fsnotify.v1 v1.4.7
go: downloading github.com/beorn7/perks v1.0.1
go: downloading github.com/google/go-cmp v0.4.0
go: downloading k8s.io/kube-openapi v0.0.0-20191107075043-30be4d16710a
go: downloading github.com/modern-go/reflect2 v1.0.1
go: downloading github.com/json-iterator/go v1.1.9
go: downloading github.com/cespare/xxhash v1.1.0
go: downloading gopkg.in/inf.v0 v0.9.1
go: downloading github.com/cespare/xxhash/v2 v2.1.1
go: downloading go.uber.org/multierr v1.5.0
go: downloading golang.org/x/crypto v0.0.0-20200220183623-bac4c82f6975
go: downloading github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd
go: downloading github.com/matttproud/golang_protobuf_extensions v1.0.1
go: downloading github.com/Azure/go-autorest/autorest/adal v0.8.1-0.20191028180845-3492b2aff503
go: downloading golang.org/x/sys v0.0.0-20200122134326-e047566fdf82
go: downloading github.com/google/uuid v1.1.1
go: downloading github.com/dgrijalva/jwt-go v3.2.0+incompatible
go: downloading github.com/Azure/go-autorest/autorest/date v0.2.0
# github.com/imooc-com/imooc-operator/pkg/apis/k8s/v1alpha1
pkg/apis/k8s/v1alpha1/imoocpod_types.go:47:25: cannot use &ImoocPod literal (type *ImoocPod) as type runtime.Object in argument to SchemeBuilder.Register:
	*ImoocPod does not implement runtime.Object (missing DeepCopyObject method)
pkg/apis/k8s/v1alpha1/imoocpod_types.go:47:38: cannot use &ImoocPodList literal (type *ImoocPodList) as type runtime.Object in argument to SchemeBuilder.Register:
	*ImoocPodList does not implement runtime.Object (missing DeepCopyObject method)
FATA[0057] Failed to build operator binary: failed to exec []string{"go", "build", "-o", "/Users/huzhi/work/code/go_code/imooc/imooc-operator/build/_output/bin/imooc-operator", "-gcflags", "all=-trimpath=/Users/huzhi/work/code/go_code/imooc", "-asmflags", "all=-trimpath=/Users/huzhi/work/code/go_code/imooc", "github.com/imooc-com/imooc-operator/cmd/manager"}: exit status 2
huzhi@localhost imooc-operator % git status
On branch master
Changes not staged for commit:
  (use "git add <file>..." to update what will be committed)
  (use "git checkout -- <file>..." to discard changes in working directory)

	modified:   go.mod

Untracked files:
  (use "git add <file>..." to include in what will be committed)

	go.sum

no changes added to commit (use "git add" and/or "git commit -a")
huzhi@localhost imooc-operator % ll /Users/huzhi/go/
total 0
drwxr-xr-x   3 huzhi  staff    96  9 22 14:33 ./
drwxr-xr-x+ 38 huzhi  staff  1216  9 22 14:33 ../
drwxr-xr-x   4 huzhi  staff   128  9 22 14:33 pkg/
huzhi@localhost imooc-operator % ll /Users/huzhi/go/pkg/mod/
total 0
drwxr-xr-x  11 huzhi  staff  352  9 22 14:36 ./
drwxr-xr-x   4 huzhi  staff  128  9 22 14:33 ../
drwxr-xr-x   4 huzhi  staff  128  9 22 14:34 cache/
drwxr-xr-x   3 huzhi  staff   96  9 22 14:36 cloud.google.com/
drwxr-xr-x  24 huzhi  staff  768  9 22 14:36 github.com/
drwxr-xr-x   5 huzhi  staff  160  9 22 14:36 go.uber.org/
drwxr-xr-x   3 huzhi  staff   96  9 22 14:35 golang.org/
drwxr-xr-x   3 huzhi  staff   96  9 22 14:35 gomodules.xyz/
drwxr-xr-x   5 huzhi  staff  160  9 22 14:36 gopkg.in/
drwxr-xr-x   9 huzhi  staff  288  9 22 14:36 k8s.io/
drwxr-xr-x   4 huzhi  staff  128  9 22 14:35 sigs.k8s.io/
huzhi@localhost imooc-operator %


huzhi@localhost imooc-operator % operator-sdk build rubinus/imoocpod-operator
# github.com/imooc-com/imooc-operator/pkg/apis/k8s/v1alpha1
pkg/apis/k8s/v1alpha1/imoocpod_types.go:47:25: cannot use &ImoocPod literal (type *ImoocPod) as type runtime.Object in argument to SchemeBuilder.Register:
	*ImoocPod does not implement runtime.Object (missing DeepCopyObject method)
pkg/apis/k8s/v1alpha1/imoocpod_types.go:47:38: cannot use &ImoocPodList literal (type *ImoocPodList) as type runtime.Object in argument to SchemeBuilder.Register:
	*ImoocPodList does not implement runtime.Object (missing DeepCopyObject method)
FATA[0000] Failed to build operator binary: failed to exec []string{"go", "build", "-o", "/Users/huzhi/work/code/go_code/imooc/imooc-operator/build/_output/bin/imooc-operator", "-gcflags", "all=-trimpath=/Users/huzhi/work/code/go_code/imooc", "-asmflags", "all=-trimpath=/Users/huzhi/work/code/go_code/imooc", "github.com/imooc-com/imooc-operator/cmd/manager"}: exit status 2
huzhi@localhost imooc-operator %


解决方式是执行：operator-sdk generate k8s


huzhi@localhost imooc-operator % operator-sdk generate k8s
INFO[0000] Running deepcopy code-generation for Custom Resource group versions: [k8s:[v1alpha1], ]
INFO[0007] Code-generation complete.
huzhi@localhost imooc-operator %
huzhi@localhost imooc-operator % git status
On branch master
Untracked files:
  (use "git add <file>..." to include in what will be committed)

	pkg/apis/k8s/v1alpha1/zz_generated.deepcopy.go

nothing added to commit but untracked files present (use "git add" to track)
huzhi@localhost imooc-operator %
huzhi@localhost imooc-operator % operator-sdk build rubinus/imoocpod-operator
INFO[0003] Building OCI image rubinus/imoocpod-operator
Sending build context to Docker daemon  43.22MB
Step 1/7 : FROM registry.access.redhat.com/ubi8/ubi-minimal:latest
latest: Pulling from ubi8/ubi-minimal
0fd3b5213a9b: Pull complete
aebb8c556853: Pull complete
Digest: sha256:5cfbaf45ca96806917830c183e9f37df2e913b187aadb32e89fd83fa455ebaa6
Status: Downloaded newer image for registry.access.redhat.com/ubi8/ubi-minimal:latest
 ---> 28095021e526
Step 2/7 : ENV OPERATOR=/usr/local/bin/imooc-operator     USER_UID=1001     USER_NAME=imooc-operator
 ---> Running in 4eb465ec6005
Removing intermediate container 4eb465ec6005
 ---> 76d2e87b0708
Step 3/7 : COPY build/_output/bin/imooc-operator ${OPERATOR}
 ---> fc9d47754af6
Step 4/7 : COPY build/bin /usr/local/bin
 ---> d6a84c55d362
Step 5/7 : RUN  /usr/local/bin/user_setup
 ---> Running in 1d51b996b409
+ echo 'imooc-operator:x:1001:0:imooc-operator user:/root:/sbin/nologin'
+ mkdir -p /root
+ chown 1001:0 /root
+ chmod ug+rwx /root
+ rm /usr/local/bin/user_setup
Removing intermediate container 1d51b996b409
 ---> c6e6f496dc4f
Step 6/7 : ENTRYPOINT ["/usr/local/bin/entrypoint"]
 ---> Running in e8e80609bd7b
Removing intermediate container e8e80609bd7b
 ---> 531269d4bd98
Step 7/7 : USER ${USER_UID}
 ---> Running in a3e22ada93f5
Removing intermediate container a3e22ada93f5
 ---> a12ab398317b
Successfully built a12ab398317b
Successfully tagged rubinus/imoocpod-operator:latest
INFO[0081] Operator build complete.
huzhi@localhost imooc-operator %

huzhi@localhost imooc-operator % operator-sdk generate crds
INFO[0000] Running CRD generator.
INFO[0003] CRD generation complete.
huzhi@localhost imooc-operator % git status
On branch master
Untracked files:
  (use "git add <file>..." to include in what will be committed)

	deploy/crds/k8s.imooc.com_imoocpods_crd.yaml

nothing added to commit but untracked files present (use "git add" to track)
huzhi@localhost imooc-operator %


mkdir -p ~/docker-data/docker_registry
docker run -d -p 5000:5000 --restart always --name registry -v ~/docker-data/docker_registry:/var/lib/registry registry:2

# kind 容器内部的 docker 命令是 crictl
root@kind-control-plane:/# crictl images
IMAGE                                      TAG                  IMAGE ID            SIZE
docker.io/kindest/kindnetd                 v20200725-4d6bea59   b77790820d015       119MB
docker.io/rancher/local-path-provisioner   v0.0.14              e422121c9c5f9       42MB
docker.io/rubinus/imoocpod-operator        latest               a12ab398317ba       188MB
k8s.gcr.io/build-image/debian-base         v2.1.0               c7c6c86897b63       53.9MB
k8s.gcr.io/coredns                         1.7.0                bfe3a36ebd252       45.4MB
k8s.gcr.io/etcd                            3.4.13-0             0369cf4303ffd       255MB
k8s.gcr.io/kube-apiserver                  v1.19.1              8cba89a89aaa8       95MB
k8s.gcr.io/kube-controller-manager         v1.19.1              7dafbafe72c90       84.1MB
k8s.gcr.io/kube-proxy                      v1.19.1              47e289e332426       136MB
k8s.gcr.io/kube-scheduler                  v1.19.1              4d648fc900179       65.1MB
k8s.gcr.io/pause                           3.3                  0184c1613d929       686kB
root@kind-control-plane:/#




# 在 imoocpod_types.go 文件中增加字段
huzhi@localhost imooc-operator % git status
On branch master
Changes not staged for commit:
  (use "git add <file>..." to update what will be committed)
  (use "git checkout -- <file>..." to discard changes in working directory)

	modified:   pkg/apis/k8s/v1alpha1/imoocpod_types.go

no changes added to commit (use "git add" and/or "git commit -a")
huzhi@localhost imooc-operator %
huzhi@localhost imooc-operator % git diff
diff --git a/pkg/apis/k8s/v1alpha1/imoocpod_types.go b/pkg/apis/k8s/v1alpha1/imoocpod_types.go
index 6c843ce..65d996f 100644
--- a/pkg/apis/k8s/v1alpha1/imoocpod_types.go
+++ b/pkg/apis/k8s/v1alpha1/imoocpod_types.go
@@ -12,6 +12,7 @@ type ImoocPodSpec struct {
        // INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
        // Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
        // Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
+       Replicas int `json:"replicas"`  // 期望 pod 的数量
 }

 // ImoocPodStatus defines the observed state of ImoocPod
@@ -19,6 +20,8 @@ type ImoocPodStatus struct {
        // INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
        // Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
        // Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
+       Replicas int `json:"replicas"`  // 实际运行的 pod 的数量
+       PodNames []string `json:"podNames"`  // 实际运行的 pod 的名称
 }

 // +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

huzhi@localhost imooc-operator %
huzhi@localhost imooc-operator % operator-sdk generate k8s
INFO[0000] Running deepcopy code-generation for Custom Resource group versions: [k8s:[v1alpha1], ]
INFO[0006] Code-generation complete.
huzhi@localhost imooc-operator %
huzhi@localhost imooc-operator % git status
On branch master
Changes not staged for commit:
  (use "git add <file>..." to update what will be committed)
  (use "git checkout -- <file>..." to discard changes in working directory)

	modified:   pkg/apis/k8s/v1alpha1/imoocpod_types.go
	modified:   pkg/apis/k8s/v1alpha1/zz_generated.deepcopy.go

no changes added to commit (use "git add" and/or "git commit -a")
huzhi@localhost imooc-operator %
huzhi@localhost imooc-operator %
huzhi@localhost imooc-operator % git diff
diff --git a/pkg/apis/k8s/v1alpha1/imoocpod_types.go b/pkg/apis/k8s/v1alpha1/imoocpod_types.go
index 6c843ce..65d996f 100644
--- a/pkg/apis/k8s/v1alpha1/imoocpod_types.go
+++ b/pkg/apis/k8s/v1alpha1/imoocpod_types.go
@@ -12,6 +12,7 @@ type ImoocPodSpec struct {
        // INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
        // Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
        // Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
+       Replicas int `json:"replicas"`  // 期望 pod 的数量
 }

 // ImoocPodStatus defines the observed state of ImoocPod
@@ -19,6 +20,8 @@ type ImoocPodStatus struct {
        // INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
        // Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
        // Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
+       Replicas int `json:"replicas"`  // 实际运行的 pod 的数量
+       PodNames []string `json:"podNames"`  // 实际运行的 pod 的名称
 }

 // +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
diff --git a/pkg/apis/k8s/v1alpha1/zz_generated.deepcopy.go b/pkg/apis/k8s/v1alpha1/zz_generated.deepcopy.go
index 72ab512..0a23ff0 100644
--- a/pkg/apis/k8s/v1alpha1/zz_generated.deepcopy.go
+++ b/pkg/apis/k8s/v1alpha1/zz_generated.deepcopy.go
@@ -14,7 +14,7 @@ func (in *ImoocPod) DeepCopyInto(out *ImoocPod) {
        out.TypeMeta = in.TypeMeta
        in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
        out.Spec = in.Spec
-       out.Status = in.Status
+       in.Status.DeepCopyInto(&out.Status)
        return
 }

@@ -88,6 +88,11 @@ func (in *ImoocPodSpec) DeepCopy() *ImoocPodSpec {
 // DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
 func (in *ImoocPodStatus) DeepCopyInto(out *ImoocPodStatus) {
        *out = *in
+       if in.PodNames != nil {
+               in, out := &in.PodNames, &out.PodNames
+               *out = make([]string, len(*in))
+               copy(*out, *in)
+       }
        return
 }

huzhi@localhost imooc-operator %

huzhi@localhost imooc-operator %
huzhi@localhost imooc-operator % operator-sdk generate crds
INFO[0000] Running CRD generator.
INFO[0000] CRD generation complete.
huzhi@localhost imooc-operator %
huzhi@localhost imooc-operator % git status
On branch master
Changes not staged for commit:
  (use "git add <file>..." to update what will be committed)
  (use "git checkout -- <file>..." to discard changes in working directory)

	modified:   deploy/crds/k8s.imooc.com_imoocpods_crd.yaml

no changes added to commit (use "git add" and/or "git commit -a")
huzhi@localhost imooc-operator % git diff
diff --git a/deploy/crds/k8s.imooc.com_imoocpods_crd.yaml b/deploy/crds/k8s.imooc.com_imoocpods_crd.yaml
index a9d3b7c..e8b8cc6 100644
--- a/deploy/crds/k8s.imooc.com_imoocpods_crd.yaml
+++ b/deploy/crds/k8s.imooc.com_imoocpods_crd.yaml
@@ -30,9 +30,32 @@ spec:
           type: object
         spec:
           description: ImoocPodSpec defines the desired state of ImoocPod
+          properties:
+            replicas:
+              description: 'INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
+                Important: Run "operator-sdk generate k8s" to regenerate code after
+                modifying this file Add custom validation using kubebuilder tags:
+                https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html'
+              type: integer
+          required:
+          - replicas
           type: object
         status:
           description: ImoocPodStatus defines the observed state of ImoocPod
+          properties:
+            podNames:
+              items:
+                type: string
+              type: array
+            replicas:
+              description: 'INSERT ADDITIONAL STATUS FIELD - define observed state
+                of cluster Important: Run "operator-sdk generate k8s" to regenerate
+                code after modifying this file Add custom validation using kubebuilder
+                tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html'
+              type: integer
+          required:
+          - podNames
+          - replicas
           type: object
       type: object
   version: v1alpha1
huzhi@localhost imooc-operator %

huzhi@localhost imooc-operator % git diff
diff --git a/pkg/controller/imoocpod/imoocpod_controller.go b/pkg/controller/imoocpod/imoocpod_controller.go
index 46f7ead..9721291 100644
--- a/pkg/controller/imoocpod/imoocpod_controller.go
+++ b/pkg/controller/imoocpod/imoocpod_controller.go
@@ -100,6 +100,39 @@ func (r *ReconcileImoocPod) Reconcile(request reconcile.Request) (reconcile.Resu
                return reconcile.Result{}, err
        }

+       reqLogger.Info("自定义日志")
+
+       /*
+       // 通过自定义的资源 ImoocPod 获取已经存在的 pod
+       lbls := labels.Set{
+               "app": instance.Name,
+       }
+       existingPods := &corev1.PodList{}
+       err = r.client.List(
+               context.TODO(),
+               existingPods,
+               &client.ListOptions{
+                       Namespace: request.Namespace,
+                       LabelSelector: labels.SelectorFromSet(lbls),
+               },
+       )
+       if err != nil {
+               reqLogger.Error(err, "获取已经存在的 pod 失败")
+               return reconcile.Result{}, nil
+       }
+
+       // 获取 pod 的名称
+       var existingPodNames []string
+       for _, pod := range existingPods.Items {
+               if pod.GetObjectMeta().GetDeletionTimestamp() != nil {
+                       continue
+               }
+               if pod.Status.Phase == corev1.PodPending || pod.Status.Phase == corev1.PodRunning {
+                       existingPodNames = append(existingPodNames, pod.GetObjectMeta().GetName())
+               }
+       }
+       */
+
        // Define a new Pod object
        pod := newPodForCR(instance)

huzhi@localhost imooc-operator %
huzhi@localhost imooc-operator %
huzhi@localhost imooc-operator %
huzhi@localhost imooc-operator %
huzhi@localhost imooc-operator % git commit -am '更新代码'
[master 9234886] 更新代码
 1 file changed, 33 insertions(+)
huzhi@localhost imooc-operator % git diff
huzhi@localhost imooc-operator % git status
On branch master
nothing to commit, working tree clean
huzhi@localhost imooc-operator % operator-sdk generate k8s
INFO[0000] Running deepcopy code-generation for Custom Resource group versions: [k8s:[v1alpha1], ]
INFO[0007] Code-generation complete.
huzhi@localhost imooc-operator %

huzhi@localhost imooc-operator %
huzhi@localhost imooc-operator % operator-sdk build rubinus/imoocpod-operator:v1
INFO[0003] Building OCI image rubinus/imoocpod-operator:v1
Sending build context to Docker daemon  43.28MB
Step 1/7 : FROM registry.access.redhat.com/ubi8/ubi-minimal:latest
 ---> 28095021e526
Step 2/7 : ENV OPERATOR=/usr/local/bin/imooc-operator     USER_UID=1001     USER_NAME=imooc-operator
 ---> Using cache
 ---> 76d2e87b0708
Step 3/7 : COPY build/_output/bin/imooc-operator ${OPERATOR}
 ---> 70e02489cdc2
Step 4/7 : COPY build/bin /usr/local/bin
 ---> e0c9cec43538
Step 5/7 : RUN  /usr/local/bin/user_setup
 ---> Running in 8b3b6d06bacc
+ echo 'imooc-operator:x:1001:0:imooc-operator user:/root:/sbin/nologin'
+ mkdir -p /root
+ chown 1001:0 /root
+ chmod ug+rwx /root
+ rm /usr/local/bin/user_setup
Removing intermediate container 8b3b6d06bacc
 ---> 0f744ac3bf58
Step 6/7 : ENTRYPOINT ["/usr/local/bin/entrypoint"]
 ---> Running in ce89fe52b485
Removing intermediate container ce89fe52b485
 ---> 0664af241ae1
Step 7/7 : USER ${USER_UID}
 ---> Running in 6aefa2f43a2b
Removing intermediate container 6aefa2f43a2b
 ---> 39353446f42c
Successfully built 39353446f42c
Successfully tagged rubinus/imoocpod-operator:v1
INFO[0006] Operator build complete.
huzhi@localhost imooc-operator %

huzhi@localhost imooc-operator %
huzhi@localhost imooc-operator % kind load docker-image rubinus/imoocpod-operator:v1
Image: "rubinus/imoocpod-operator:v1" with ID "sha256:39353446f42c23e2d688f0ae73484fd8f0b642f78508f47560220b88bb8d94a6" not yet present on node "kind-control-plane", loading...
huzhi@localhost imooc-operator %

huzhi@localhost ~ % docker ps -a
CONTAINER ID        IMAGE                  COMMAND                  CREATED             STATUS              PORTS                       NAMES
3bf8e47a72d4        kindest/node:v1.19.1   "/usr/local/bin/entr…"   8 hours ago         Up 4 hours          127.0.0.1:55487->6443/tcp   kind-control-plane
huzhi@localhost ~ % docker exec -ti kind-control-plane bash
root@kind-control-plane:/# crictl images
IMAGE                                      TAG                  IMAGE ID            SIZE
docker.io/kindest/kindnetd                 v20200725-4d6bea59   b77790820d015       119MB
docker.io/library/busybox                  latest               6858809bf669c       769kB
docker.io/rancher/local-path-provisioner   v0.0.14              e422121c9c5f9       42MB
docker.io/rubinus/imoocpod-operator        latest               a12ab398317ba       188MB
k8s.gcr.io/build-image/debian-base         v2.1.0               c7c6c86897b63       53.9MB
k8s.gcr.io/coredns                         1.7.0                bfe3a36ebd252       45.4MB
k8s.gcr.io/etcd                            3.4.13-0             0369cf4303ffd       255MB
k8s.gcr.io/kube-apiserver                  v1.19.1              8cba89a89aaa8       95MB
k8s.gcr.io/kube-controller-manager         v1.19.1              7dafbafe72c90       84.1MB
k8s.gcr.io/kube-proxy                      v1.19.1              47e289e332426       136MB
k8s.gcr.io/kube-scheduler                  v1.19.1              4d648fc900179       65.1MB
k8s.gcr.io/pause                           3.3                  0184c1613d929       686kB
root@kind-control-plane:/#
root@kind-control-plane:/#
root@kind-control-plane:/# crictl images
IMAGE                                      TAG                  IMAGE ID            SIZE
docker.io/kindest/kindnetd                 v20200725-4d6bea59   b77790820d015       119MB
docker.io/library/busybox                  latest               6858809bf669c       769kB
docker.io/rancher/local-path-provisioner   v0.0.14              e422121c9c5f9       42MB
docker.io/rubinus/imoocpod-operator        latest               a12ab398317ba       188MB
docker.io/rubinus/imoocpod-operator        v1                   39353446f42c2       188MB
k8s.gcr.io/build-image/debian-base         v2.1.0               c7c6c86897b63       53.9MB
k8s.gcr.io/coredns                         1.7.0                bfe3a36ebd252       45.4MB
k8s.gcr.io/etcd                            3.4.13-0             0369cf4303ffd       255MB
k8s.gcr.io/kube-apiserver                  v1.19.1              8cba89a89aaa8       95MB
k8s.gcr.io/kube-controller-manager         v1.19.1              7dafbafe72c90       84.1MB
k8s.gcr.io/kube-proxy                      v1.19.1              47e289e332426       136MB
k8s.gcr.io/kube-scheduler                  v1.19.1              4d648fc900179       65.1MB
k8s.gcr.io/pause                           3.3                  0184c1613d929       686kB
root@kind-control-plane:/#

huzhi@localhost imooc-operator % git diff
diff --git a/deploy/operator.yaml b/deploy/operator.yaml
index 60c3795..6a64e26 100644
--- a/deploy/operator.yaml
+++ b/deploy/operator.yaml
@@ -19,7 +19,7 @@ spec:
           # todo 替换为生成的镜像
           # kind load docker-image rubinus/imoocpod-operator
           # image: REPLACE_IMAGE
-          image: docker.io/rubinus/imoocpod-operator:latest
+          image: docker.io/rubinus/imoocpod-operator:v1
           command:
           - imooc-operator
           # 由于使用 kind 的缘故，imagePullPolicy 需要修改为 Never
huzhi@localhost imooc-operator %

```



































