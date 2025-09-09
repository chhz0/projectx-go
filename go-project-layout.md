# Go-Project-Layout

本项目的布局参考了[golang-standards/project-layout](https://github.com/golang-standards/project-layout)

```plaintext
.
├── api
├── assets
├── build
├── cmd
├── configs
├── deploys
├── docs
├── example
├── githooks
├── init
├── internal
├── pkg
├── scripts
├── test
├── third_party
├── tools
├── vendor
├── web
├── website
├── LICENSE
├── go.mod
└── README.md
```

## 核心目录

- `cmd` 项目的入口，所有可执行的`main.go`均存储在该目录
- `internal` 存放私有程序和库，任何不希望被其它项目引用的代码都应该放在这里
- `pkg` 存放公共库，外部项目可以调用 (/pkg在项目是很小时或者允许使用的Go组件很少时可以考虑不使用)
- `vendor` 应用程序依赖项，使用 **go mod vendor** 可以创建

## 通用目录

- `configs` 配置文件模板或者默认配置文件
- `init` System init(systemd, upstart, sysv) 和 process manager/supervisor 配置
- `scripts` 执行各种构建、安装、分析等操作的脚本
- `build` 打包与持续集成
- `deploy` IaaS、PaaS、系统和容器编排部署配置和模板(docker-compose、k8s/helm)
- `test` 额外的外部测试应用程序和测试数据

## 其它目录

- `api` OpenAPI/Swagger 规范，JSON模式文件，协议定义文本
- `web` Web应用的组件：静态Web资源、服务器端模板和SPAs
- `docs` 设计和用户文档
- `tools` 该项目支持的工具(例如代码生成工具)
- `example` 应用程序/公共库的示例
- `third_party` 外部辅助工具，分叉代码和其它第三方工具(例如 Swagger UI)
- `githooks` git hooks
- `assets` 和存储库一起使用的其它资源(图片、图标等)
- `website` 如果不使用 Github pages，则存放项目的网站数据

