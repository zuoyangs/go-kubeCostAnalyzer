
```go
.
├── apps
│   └── kubernetes
│       ├── api
│       │   └── http
│       │       ├── deployment.go
│       │       └── statefulset.go
│       ├── deployment
│       │   ├── impl.go
│       │   └── interface.go
│       ├── interface.go
│       ├── statefulset
│       │   ├── impl.go
│       │   └── interface.go
│       └── util
│           ├── config.go
│           ├── kubeclient.go
│           └── resource_info.go
├── go.mod  
├── main.go
└── ...
```

**详情：**

- `apps/kubernetes`: 主要的Kubernetes相关代码目录。

  - `api/http`: 定义与HTTP API相关的代码，如路由、控制器等。这里只列出`deployment.go`作为示例，实际应用可能需要根据业务需求添加更多文件。

    - `deployment.go`: 实现与Deployment相关的HTTP API逻辑。

  - `deployment`: 存放与Kubernetes Deployment相关的代码。

    - `interface.go`: 定义与Kubernetes Deployment操作相关的接口，可以包括创建、更新、删除、查询等方法。

    - `impl.go`: 提供上述接口的具体实现，利用Kubernetes客户端库与Kubernetes集群交互。

  - `statefulset`: 与`deployment`类似，存放与Kubernetes StatefulSet相关的代码。

    - `interface.go`: 定义与Kubernetes StatefulSet操作相关的接口。

    - `impl.go`: 提供上述接口的具体实现。

  - `interface.go`: 在顶层定义通用的Kubernetes资源操作接口，如`ResourceGetter`，以便于各个具体资源类型（如Deployment、StatefulSet）共享。
  - `model.go`: 存放`getResourceInfo`函数以及`ResourceInfo`、`ClusterResources`、`EnvResources`、`NamespaceResources`等结构体定义。


- `utils`：存放辅助工具类和函数。

    - `config.go`: 封装与配置文件读取相关的逻辑，如`config.Init()`和`config.GetKey()`等。

    - `kubeclient.go`: 封装创建和管理Kubernetes客户端的逻辑，如`clientcmd.BuildConfigFromFlags`和`kubernetes.NewForConfigOrDie`等。


- `main.go`: 应用程序的入口点，可以包含`Kubernete_Querier`函数的调用以及其他初始化逻辑。

- `go.mod`: Go模块定义文件，记录项目依赖。
