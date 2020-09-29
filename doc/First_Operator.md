# First Operator

## 项目描述

编写一个 Operator，实现对于指定命名空间内的资源监控，具体要求如下：

1. 基于 Operator-SDK 来实现 operator
2. 每隔 1 分钟执行一次刷新操作
3. 每次刷新的时候，在 stdout 打印下列资源的信息：
	* deployment
	* statefulset
	* secret
	* configmap
	* pod
	* service
	* endpoint
   每种信息需要打印 name, namespace, kind, status 以及其他你认为适合的信息
4. 编写 chart 来部署这个 operator
5. 设置流水线来支持这个项目的 CI

## 项目目标

1. 熟悉使用 Operator-SDK 来实现 operator 的基本过程
2. 熟悉如何使用 chart 来部署 operator
3. 熟悉如何设置流水线构建部署项目

## 实现过程

### 下载安装 docker、kubectl、kind

* kubectl 可以后续确定好 k8s 版本后再安装

### 搭建本地开发环境，使用 kind 部署本地 k8s 集群

```bash
# 使用 kind 创建 k8s 集群
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
```

### 使用 kind 安装 k8s 集群的本质就是在 docker 运行所有的 k8s 组件

* 启动 kind-control-plane 容器，在容器中运行有所有的 k8s 组件
* 进去 kind-control-plane 容器，里面有类似 docker 命令是 crictl
* 在容器中列出所有镜像可以看到 k8s 的版本和相应镜像的版本，这个可以作为参考信息
* 根据 k8s 版本安装 kubectl

```bash
huzhi@localhost ~ % docker ps -a
CONTAINER ID        IMAGE                  COMMAND                  CREATED             STATUS              PORTS                       NAMES
3bf8e47a72d4        kindest/node:v1.19.1   "/usr/local/bin/entr…"   7 days ago          Up 18 seconds       127.0.0.1:55487->6443/tcp   kind-control-plane
huzhi@localhost ~ %

# 进去容器
huzhi@localhost ~ % docker exec -ti kind-control-plane bash
root@kind-control-plane:/#

# 在容器中列出所有镜像
root@kind-control-plane:/# crictl images
IMAGE                                      TAG                  IMAGE ID            SIZE
docker.io/kindest/kindnetd                 v20200725-4d6bea59   b77790820d015       119MB
docker.io/library/busybox                  latest               6858809bf669c       769kB
docker.io/rancher/local-path-provisioner   v0.0.14              e422121c9c5f9       42MB
k8s.gcr.io/build-image/debian-base         v2.1.0               c7c6c86897b63       53.9MB
k8s.gcr.io/coredns                         1.7.0                bfe3a36ebd252       45.4MB
k8s.gcr.io/etcd                            3.4.13-0             0369cf4303ffd       255MB
k8s.gcr.io/kube-apiserver                  v1.19.1              8cba89a89aaa8       95MB
k8s.gcr.io/kube-controller-manager         v1.19.1              7dafbafe72c90       84.1MB
k8s.gcr.io/kube-proxy                      v1.19.1              47e289e332426       136MB
k8s.gcr.io/kube-scheduler                  v1.19.1              4d648fc900179       65.1MB
k8s.gcr.io/pause                           3.3                  0184c1613d929       686kB
root@kind-control-plane:/#

```

### 初始化 operator 的相关目录

* 使用 operator-sdk 创建一个新的项目，然后使用 git 跟踪每次文件的变化
* 创建 API，也就是自定义资源 crd
* 创建自定义资源 crd 相应的控制器
* 这时就可以编译 go 代码生成二进制文件，使用 dockerfile 构建镜像
	* docker 镜像中运行就是 go 的二进制文件
	* 编译 go 代码时有可能需要下载相关的第三方依赖，所以需要配置 GOPROXY
	* operator-sdk build 命令的本质就是编译 go 代码，构建镜像
	* 每次运行 operator-sdk build 之前需要运行 operator-sdk generate k8s
	* 修改 `* _types.go` 文件内容也需要运行 operator-sdk generate k8s
	* operator-sdk generate crds 命令生成 crd 的 yaml 文件
	* 修改 `* _types.go` 文件内容也需要运行 operator-sdk generate crds

```bash
# 进入工作目录
huzhi@localhost imooc % pwd
/Users/huzhi/work/code/go_code/imooc
huzhi@localhost imooc %

# 使用 operator-sdk 创建一个新的项目
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
│   ├── Dockerfile
│   └── bin
│       ├── entrypoint
│       └── user_setup
├── cmd
│   └── manager
│       └── main.go
├── deploy
│   ├── operator.yaml
│   ├── role.yaml
│   ├── role_binding.yaml
│   └── service_account.yaml
├── go.mod
├── pkg
│   ├── apis
│   │   └── apis.go
│   └── controller
│       └── controller.go
├── tools.go
└── version
    └── version.go

9 directories, 14 files
huzhi@localhost imooc %

# 进入项目目录
huzhi@localhost imooc-operator % pwd
/Users/huzhi/work/code/go_code/imooc/imooc-operator
huzhi@localhost imooc-operator %

# 使用 git 跟踪每次文件的变化
# git init .
# git add .
# git commit -m ''

# 创建 API，也就是自定义资源 crd
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
	modified:   go.mod

Untracked files:
  (use "git add <file>..." to include in what will be committed)

	deploy/crds/
	go.sum
	pkg/apis/addtoscheme_k8s_v1alpha1.go
	pkg/apis/k8s/

no changes added to commit (use "git add" and/or "git commit -a")
huzhi@localhost imooc-operator %

huzhi@localhost imooc-operator % tree -a deploy/crds/
deploy/crds/
├── k8s.imooc.com_imoocpods_crd.yaml
└── k8s.imooc.com_v1alpha1_imoocpod_cr.yaml

0 directories, 2 files
huzhi@localhost imooc-operator %

huzhi@localhost imooc-operator % tree -a pkg/apis/k8s/
pkg/apis/k8s/
├── group.go
└── v1alpha1
    ├── doc.go
    ├── imoocpod_types.go
    ├── register.go
    └── zz_generated.deepcopy.go

1 directory, 5 files
huzhi@localhost imooc-operator %

# 创建自定义资源 crd 相应的控制器
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

huzhi@localhost imooc-operator % tree -a pkg/controller/imoocpod/
pkg/controller/imoocpod/
└── imoocpod_controller.go

0 directories, 1 file
huzhi@localhost imooc-operator %


# 这时就可以编译 go 代码生成二进制文件，使用 dockerfile 构建镜像
# docker 镜像中运行就是 go 的二进制文件
# 编译 go 代码时有可能需要下载相关的第三方依赖，所以需要配置 GOPROXY
# operator-sdk build 命令的本质就是编译 go 代码，构建镜像
# 每次运行 operator-sdk build 之前需要运行 operator-sdk generate k8s
# 修改 * _types.go 文件内容也需要运行 operator-sdk generate k8s
# operator-sdk generate crds 命令生成 crd 的 yaml 文件
# 修改 * _types.go 文件内容也需要运行 operator-sdk generate crds

# 最好使用 go env -w GOPROXY=https://goproxy.io 命令
huzhi@localhost imooc-operator % export GOPROXY=https://goproxy.io
huzhi@localhost imooc-operator % go env GOPROXY
https://goproxy.io

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

huzhi@localhost imooc-operator % operator-sdk build rubinus/imoocpod-operator
# 根据需要下载 go 依赖的第三方包
# 编译 go 生成二进制文件
# 使用二进制文件构建镜像

# 在本地开发环境运行
1、将生成的镜像导入 kind 运行的容器中，导入之后的名称发生了变化，需要注意
2、修改 operator.yaml 文件，修改使用的镜像名称，按需修改 imagePullPolicy 配置，在 kind 中运行需要修改为 imagePullPolicy: Never
3、在 k8s 中部署相关资源

huzhi@localhost imooc-operator % kubectl apply -f deploy/role.yaml

huzhi@localhost imooc-operator % kubectl apply -f deploy/role_binding.yaml

huzhi@localhost imooc-operator % kubectl apply -f deploy/service_account.yaml

huzhi@localhost imooc-operator % kubectl apply -f deploy/crds/k8s.imooc.com_imoocpods_crd.yaml

huzhi@localhost imooc-operator % kubectl apply -f deploy/operator.yaml

huzhi@localhost imooc-operator % kubectl apply -f deploy/crds/k8s.imooc.com_v1alpha1_imoocpod_cr.yaml

```

### 生成 CSV 

```bash
# 生成 CSV 
huzhi@localhost imooc-operator % operator-sdk generate csv
INFO[0000] Generating CSV manifest version
INFO[0006] Fill in the following required fields in file imooc-operator.clusterserviceversion.yaml:
	spec.description
	spec.keywords
	spec.provider
INFO[0006] CSV manifest generated successfully
huzhi@localhost imooc-operator %
huzhi@localhost imooc-operator % git status
On branch master
Your branch is ahead of 'origin/master' by 23 commits.
  (use "git push" to publish your local commits)

Untracked files:
  (use "git add <file>..." to include in what will be committed)

	deploy/olm-catalog/

nothing added to commit but untracked files present (use "git add" to track)
huzhi@localhost imooc-operator %
huzhi@localhost imooc-operator % tree -a deploy/olm-catalog/
deploy/olm-catalog/
└── imooc-operator
    └── manifests
        ├── imooc-operator.clusterserviceversion.yaml  # operator 的描述文件
        └── k8s.imooc.com_imoocpods_crd.yaml           # 和之前生成的 crd 文件一模一样

2 directories, 2 files
huzhi@localhost imooc-operator %

```

### 构建 bundle image file

* 执行 operator-sdk bundle 命令之前需要先执行 operator-sdk generate csv

```bash
# 如果不执行 operator-sdk generate csv 无法成功执行 operator-sdk bundle
huzhi@localhost imooc-operator % operator-sdk bundle create --generate-only
FATA[0000] error generating bundle image files: stat /Users/huzhi/work/code/go_code/imooc/imooc-operator/deploy/olm-catalog/imooc-operator/manifests: no such file or directory
huzhi@localhost imooc-operator %

huzhi@localhost imooc-operator % mkdir -p deploy/olm-catalog/imooc-operator/manifests

huzhi@localhost imooc-operator % operator-sdk bundle create --generate-only
FATA[0000] error generating bundle image files: The directory /Users/huzhi/work/code/go_code/imooc/imooc-operator/deploy/olm-catalog/imooc-operator/manifests contains no yaml files
huzhi@localhost imooc-operator %

# 执行 operator-sdk bundle 命令之前需要先执行 operator-sdk generate csv

huzhi@localhost imooc-operator % operator-sdk generate csv
INFO[0000] Generating CSV manifest version
INFO[0006] Fill in the following required fields in file imooc-operator.clusterserviceversion.yaml:
	spec.description
	spec.keywords
	spec.provider
INFO[0006] CSV manifest generated successfully
huzhi@localhost imooc-operator %

huzhi@localhost imooc-operator % git status
On branch master
Your branch is ahead of 'origin/master' by 23 commits.
  (use "git push" to publish your local commits)

Untracked files:
  (use "git add <file>..." to include in what will be committed)

	deploy/olm-catalog/

nothing added to commit but untracked files present (use "git add" to track)
huzhi@localhost imooc-operator %
huzhi@localhost imooc-operator % tree -a deploy/olm-catalog
deploy/olm-catalog
└── imooc-operator
    └── manifests
        ├── imooc-operator.clusterserviceversion.yaml
        └── k8s.imooc.com_imoocpods_crd.yaml

2 directories, 2 files
huzhi@localhost imooc-operator %
huzhi@localhost imooc-operator % git add .
huzhi@localhost imooc-operator % git commit -m '生成 csv'
[master 772fed5] 生成 csv
 2 files changed, 239 insertions(+)
 create mode 100644 deploy/olm-catalog/imooc-operator/manifests/imooc-operator.clusterserviceversion.yaml
 create mode 100644 deploy/olm-catalog/imooc-operator/manifests/k8s.imooc.com_imoocpods_crd.yaml
huzhi@localhost imooc-operator %
huzhi@localhost imooc-operator % git status
On branch master
Your branch is ahead of 'origin/master' by 24 commits.
  (use "git push" to publish your local commits)

nothing to commit, working tree clean
huzhi@localhost imooc-operator %
huzhi@localhost imooc-operator %
huzhi@localhost imooc-operator % operator-sdk bundle create --generate-only
INFO[0000] Building annotations.yaml
INFO[0000] Generating output manifests directory
INFO[0000] Building Dockerfile
huzhi@localhost imooc-operator %
huzhi@localhost imooc-operator % git status
On branch master
Your branch is ahead of 'origin/master' by 24 commits.
  (use "git push" to publish your local commits)

Changes not staged for commit:
  (use "git add <file>..." to update what will be committed)
  (use "git checkout -- <file>..." to discard changes in working directory)

	modified:   deploy/olm-catalog/imooc-operator/manifests/imooc-operator.clusterserviceversion.yaml
	modified:   deploy/olm-catalog/imooc-operator/manifests/k8s.imooc.com_imoocpods_crd.yaml

Untracked files:
  (use "git add <file>..." to include in what will be committed)

	bundle.Dockerfile
	deploy/olm-catalog/imooc-operator/metadata/

no changes added to commit (use "git add" and/or "git commit -a")
huzhi@localhost imooc-operator %

# 该命令把之前生成的 csv 文件和 crd 文件清空了，也许需要再次运行 operator-sdk generate csv
huzhi@localhost imooc-operator % cat deploy/olm-catalog/imooc-operator/manifests/imooc-operator.clusterserviceversion.yaml
huzhi@localhost imooc-operator %
huzhi@localhost imooc-operator % cat deploy/olm-catalog/imooc-operator/manifests/k8s.imooc.com_imoocpods_crd.yaml
huzhi@localhost imooc-operator %
huzhi@localhost imooc-operator % git add .
huzhi@localhost imooc-operator % git commit -m 'operator-sdk bundle'
[master f0095db] operator-sdk bundle
 4 files changed, 18 insertions(+), 239 deletions(-)
 create mode 100644 bundle.Dockerfile
 rewrite deploy/olm-catalog/imooc-operator/manifests/imooc-operator.clusterserviceversion.yaml (100%)
 rewrite deploy/olm-catalog/imooc-operator/manifests/k8s.imooc.com_imoocpods_crd.yaml (100%)
 create mode 100644 deploy/olm-catalog/imooc-operator/metadata/annotations.yaml
huzhi@localhost imooc-operator %
huzhi@localhost imooc-operator % git status
On branch master
Your branch is ahead of 'origin/master' by 25 commits.
  (use "git push" to publish your local commits)

nothing to commit, working tree clean
huzhi@localhost imooc-operator %

huzhi@localhost imooc-operator % operator-sdk generate csv
INFO[0000] Generating CSV manifest version
FATA[0000] error generating CSV: no CSV manifest in deploy/olm-catalog/imooc-operator/manifests
huzhi@localhost imooc-operator %
huzhi@localhost imooc-operator %
huzhi@localhost imooc-operator % git status
On branch master
Your branch is ahead of 'origin/master' by 25 commits.
  (use "git push" to publish your local commits)

nothing to commit, working tree clean
huzhi@localhost imooc-operator %

```

## 遗留的问题

* operator-sdk bundle 清空 csv 文件和 crd 文件之后后续再次如何生成？
* 使用 CSV 描述文件之后，operator.yaml 文件去哪里了？

### 使用 CSV 描述文件之后，operator.yaml 文件去哪里

理论上应该是将 operator.yaml 文件中的内容写入 csv 的某个固定字段 
