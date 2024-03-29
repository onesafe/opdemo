## 安装

#### 安装go
https://golang.org/dl/


#### 安装operator-sdk
https://github.com/operator-framework/operator-sdk/blob/master/doc/user/install-operator-sdk.md



安装完成后，查看版本信息
```sh
$ operator-sdk version
operator-sdk version: v0.7.0
$ go version
go version go1.12.7 darwin/amd64
```
<br>

## 使用


### 使用operator-sdk创建新项目
```sh
$ cd $GOPATH/src/gitlab.4pd.io/wangyiping
$ operator-sdk new opdemo
```


### 添加API
```sh
$ operator-sdk add api --api-version=app.4paradigm.com/v1 --kind=AppService
```

### 添加Controller
```sh
$ operator-sdk add controller --api-version=app.4paradigm.com/v1 --kind=AppService
```

### 书写自己的业务逻辑
打开pkg/apis/app/v1/appservice_types.go，需要我们根据我们的需求去自定义结构体 AppServiceSpec

定义API SPEC完成后，可以通过命令生成对应的code
```sh
$ operator-sdk generate k8s
```

具体的业务逻辑在pkg/controller里面去实现

<br>

## 调试
首先，在k8s集群里面安装crd对象
```sh
$ kubectl create -f deploy/crds/app_v1_appservice_crd.yaml
```

在本地项目中启动 Operator 的调试
```sh
$ operator-sdk up local 
```

创建这个资源对象
```sh
$ kubectl create -f deploy/crds/app_v1_appservice_cr.yaml
```

观察Operator的调试窗口出现的信息

<br>

## 部署
执行下面的命令构建 Operator 应用打包成 Docker 镜像
```sh
$ operator-sdk build docker02:35000/opdemo:wyp
$ docker push docker02:35000/opdemo:wyp
```

修改deploy/operator.yaml里面，替换镜像名

创建Operator RBAC
```sh
# Setup Service Account
$ kubectl create -f deploy/service_account.yaml
# Setup RBAC
$ kubectl create -f deploy/role.yaml
$ kubectl create -f deploy/role_binding.yaml
```

安装 CRD 和 Operator
```sh
$ kubectl create -f deploy/operator.yaml
```

此时，集群里面已经有了CRD和对应的Operator了

当我们创建CRD对象时，会有对应的Operator来处理相应的业务逻辑


