# IDRM Go Common

[English](README.md)

## 概述

`idrm-go-common` 是 kweaver-ai 开发的**智能数据权限管理（IDRM）**系统的综合 Go 通用库。它为构建数据治理和管理应用程序提供共享组件、API 定义、工具和服务集成。

## 功能特性

- **审计日志**：支持分布式日志的综合审计跟踪系统
- **REST 客户端**：用于微服务集成的预构建 HTTP 客户端
- **中间件**：用于身份验证、授权和日志记录的 HTTP 中间件
- **工作流引擎**：支持消息队列的事件驱动工作流处理
- **会话管理**：基于 Redis 的会话存储和生命周期管理
- **数据库回调**：变更数据捕获和数据血缘跟踪
- **错误处理**：跨服务的集中式错误代码管理

## 要求

- Go 1.24 或更高版本

## 安装

```bash
go get github.com/kweaver-ai/idrm-go-common
```

## 项目结构

```
idrm-go-common/
├── api/                  # Protocol Buffer 和 API 定义
│   ├── auth-service/     # 认证和授权 API
│   ├── audit/            # 审计事件定义
│   ├── configuration-center/  # 配置管理 API
│   ├── data-view/        # 数据视图 API
│   └── task_center/      # 任务和工作流 API
├── audit/                # 审计日志实现
├── callback/             # gRPC 回调客户端
├── d_session/            # 分布式会话管理
├── database_callback/    # 数据库变更跟踪
├── errorcode/            # 错误代码定义
├── interception/         # 令牌和用户拦截
├── middleware/           # HTTP 中间件
├── reconcile/            # 对账工具
├── rest/                 # REST 客户端实现
├── trace/                # OpenTelemetry 链路追踪
├── util/                 # 通用工具
└── workflow/             # 工作流引擎
```

## 使用示例

### 审计日志

```go
import "github.com/kweaver-ai/idrm-go-common/audit"

// 创建审计上下文
ctx := audit.NewContext().
    WithOperator("user-id").
    WithOperation("create").
    WithResource("data-source")

// 记录审计事件
audit.Log(ctx, "数据源已创建")
```

### REST 客户端

```go
import "github.com/kweaver-ai/idrm-go-common/rest/auth-service/v1"

// 创建认证服务客户端
client := authservice.NewAuthClient(&authservice.Config{
    BaseURL: "http://localhost:8080",
})

// 检查授权
resp, err := client.CheckPermission(ctx, &authservice.CheckPermissionRequest{
    Subject: "user:123",
    Action:  "read",
    Object:  "resource:456",
})
```

### 中间件

```go
import "github.com/kweaver-ai/idrm-go-common/middleware/v1"

// 添加认证中间件
router.Use(middleware.TokenInterception())

// 添加审计日志中间件
router.Use(middleware.AuditLogger())

// 添加访问控制中间件
router.Use(middleware.AccessControl())
```

## 模块文档

### API 定义 (`api/`)

用于服务集成的 Protocol Buffer 定义和 Go 类型：

- **auth-service**：认证策略和授权请求
- **audit**：事件类型、操作者和资源
- **configuration-center**：用户、角色和权限管理
- **data-view**：数据视图和子视图定义
- **task_center**：工单和任务管理

### REST 客户端 (`rest/`)

用于服务集成的 HTTP 客户端实现：

- **auth-service**：授权和策略执行
- **anyrobot**：数据模型和统一查询客户端
- **configuration-center**：配置和用户管理
- **task_center**：任务和工作流操作
- **data_catalog**：数据目录集成

### 审计 (`audit/`)

综合审计日志系统：

- 代理检测（Web、移动端）
- 带源跟踪的事件日志记录
- 用于分布式日志的 Kafka 集成
- 结构化审计事件

### 中间件 (`middleware/`)

用于请求处理的 HTTP 中间件：

- 令牌拦截和验证
- 用户上下文管理
- 权限检查
- 审计日志记录
- 错误处理

### 工作流 (`workflow/`)

事件驱动的工作流处理：

- 消息队列抽象（Kafka、NSQ）
- 工作流定义和执行
- 事件消费和处理

## 依赖项

主要依赖包括：

- `github.com/gin-gonic/gin` - HTTP 框架
- `github.com/google/uuid` - UUID 生成
- `github.com/redis/go-redis/v9` - Redis 客户端
- `github.com/IBM/sarama` - Kafka 客户端
- `github.com/nsqio/go-nsq` - NSQ 客户端
- `go.opentelemetry.io/otel` - OpenTelemetry 链路追踪
- `go.uber.org/zap` - 结构化日志
- `gorm.io/gorm` - ORM 库
- `google.golang.org/grpc` - gRPC 框架

## 开发

### 运行测试

```bash
go test ./...
```

### 运行测试并生成覆盖率报告

```bash
go test -cover ./...
```

### 生成 Protocol Buffers

```bash
# 从 .proto 文件生成 Go 代码
make proto
```

## 许可证

Copyright © kweaver-ai

## 贡献

欢迎贡献！请随时提交 Pull Request。
