# First Operator

## 项目描述

编写一个 Operator，实现对于指定环境变量内的资源监控，具体要求如下：

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

## 项目运行方式

运行环境需求：

1. 运行该项目需要有一个 k8s 的基础环境，建议使用 [kind](ttps://kind.sigs.k8s.io/) 在本地部署一个 k8s 集群

2. 安装 Operator SDK，具体参考 [Operator SDK 安装文档](https://sdk.operatorframework.io/docs/installation/install-operator-sdk/)

### 直接在 k8s 创建相关资源

1. 克隆项目

```bash
$ git clone https://gitlab-ce.alauda.cn/zhihu/first-operator.git
```

2. 生成 docker 镜像

```bash
# 镜像名称和 tag 自定，这里以 first_operator_image:v1 为例
$ operator-sdk build first_operator_image:v1
```

3. 修改相关 yaml 文件中的镜像配置

```bash
# 修改 deploy/operator.yaml 文件中的 image 配置
$ sed "s#image: REPLACE_IMAGE#image: first_operator_image:v1#g" ./deploy/operator.yaml
```

4. 在 k8s 上部署相关的资源

```bash
# 部署 ServiceAccount
$ kubectl apply -f deploy/service_account.yaml

# 部署 Role
$ kubectl apply -f deploy/role.yaml

# 部署 RoleBinding
$ kubectl apply -f deploy/role_binding.yaml

# 部署 CustomResourceDefinition
$ kubectl apply -f deploy/crds/k8s.imooc.com_imoocpods_crd.yaml

# 部署 operator Deployment
$ kubectl apply -f deploy/operator.yaml

# 根据 crd 中定义的资源部署相关实例 cr 
$ kubectl apply -f deploy/crds/k8s.imooc.com_v1alpha1_imoocpod_cr.yaml

```

### 使用 chart 方式安装

note：chart 方式安装需要预安装 helm，参考 [helm 安装文档](https://helm.sh/docs/intro/install/)

1. 在相关 Harbor 仓库中下载 chart 包到本地

```bash
# 替换 url 为 chart 存储地址 
$ wget url -O mychart-0.1.0.tgz
```

2. 解压安装包并安装

```bash
$ tar -zxvf mychart-0.1.0.tgz
$ helm install ./mychart/
```

## 验证输出结果

打印相关 pod 的 stdout 输出查看结果

```bash
$ kubectl get pods | grep imooc-operator

$ kubectl logs -f imooc-operator-64bdf69895-rl4q
...
定义返回值
watchedNamespace default
[imooc-operator-64bdf69895-rl4qj redisoperator-d8898fb46-wcv6m example-imoocpod-pod]
[imooc-operator redisoperator]
service: kubernetes
service: imooc-operator-metrics
...

```

## 参考

[项目具体实现过程](http://confluence.alauda.cn/pages/viewpage.action?pageId=75083324)