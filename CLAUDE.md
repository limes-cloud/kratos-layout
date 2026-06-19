# Kratos Layout 基础开发模板

## 项目简介

这是基于 `limes-cloud/manager` 开发规范整理的 kratosx 服务模板，用于快速创建符合 DDD 分层、proto-first、独立 API 模块、自注册 handler 模式的 Go 微服务。

模板内置基础示例：

- `api/errors`：统一错误码 proto 与生成函数
- `api/dictionary`：字典 CRUD 示例
- `api/entity`：实体、字段、规则注册与 CRUD 示例
- `internal/app`：gRPC/HTTP handler 自注册
- `internal/domain`：entity、repository 接口、service 业务逻辑
- `internal/infra`：dbs 仓储实现
- `internal/types`：内部 DTO

## 技术栈

| 维度 | 选型 |
|---|---|
| 语言 | Go |
| 微服务框架 | go-kratos/kratos（gRPC + HTTP 网关） |
| 基础框架 | github.com/limes-cloud/kratosx |
| 数据库 | GORM（通过 `ctx.DB()` 使用） |
| 缓存 | Redis（通过 kratosx 管理，按业务需要启用） |
| Proto | protoc + kratosx proto 插件 |

## DDD 分层架构

```
app/       → 应用层（gRPC/HTTP handler，组装 service，自注册）
domain/    → 领域层（entity + repository 接口 + service 业务逻辑）
infra/     → 基础设施层（dbs 仓储实现 + 外部客户端实现）
types/     → 请求/响应 DTO
api/       → 独立 Go 模块，存放 proto 与生成代码
```

**依赖方向：** app → domain ← infra，domain 层不依赖任何具体实现。

## 核心编码原则

1. **proto-first**：API 先定义 proto，再生成代码，错误码也由 proto 生成。
2. **API 独立模块**：`api/` 拥有独立 `go.mod`，主模块通过 `replace layout/api => ./api` 引用。
3. **接口在 domain 层定义，实现在 infra 层**：repository 接口在 `domain/repository/`，实现类在 `infra/dbs/`。
4. **自注册模式**：app 层通过 `init() + register()` 自动注册 gRPC/HTTP handler。
5. **手动依赖注入**：`NewXxx()` 构造函数注入 repository 接口，不使用 DI 容器。
6. **错误使用 proto 函数**：service/app 层返回 `errors.XxxError()`，不暴露内部 error。
7. **nullable 字段用指针**：`*string` / `*bool` 等，通过 `value.Pointer()` / `value.Value()` 转换。
8. **使用 kratosx.Context**：handler 入口使用 `kratosx.MustContext(ctx)` 转换。
9. **使用 GORM**：数据库操作通过 kratosx 封装的 GORM 进行，`ctx.DB()` 返回 `*gorm.DB`。

## 详细规则索引

| 规则文件 | 内容 |
|---|---|
| [architecture.md](.claude/rules/architecture.md) | 分层职责与依赖约束 |
| [code-style.md](.claude/rules/code-style.md) | 命名规范、错误处理、构造函数、日志、注释 |
| [app-layer.md](.claude/rules/app-layer.md) | App 层 handler 与自注册、三步转换模式 |
| [domain-layer.md](.claude/rules/domain-layer.md) | Domain 层 entity/repository/service |
| [infra-layer.md](.claude/rules/infra-layer.md) | Infra 层仓储实现与客户端封装 |
| [proto-and-api.md](.claude/rules/proto-and-api.md) | Proto 定义、错误码、API 独立模块 |
| [main-and-config.md](.claude/rules/main-and-config.md) | main.go 入口、配置加载与热更新 |
| [types-dto.md](.claude/rules/types-dto.md) | 请求/响应 DTO、分页、中转参数 |
| [build-and-deploy.md](.claude/rules/build-and-deploy.md) | Docker 多阶段构建、Makefile、配置管理 |
| [vscode.md](.claude/rules/vscode.md) | VSCode 扩展、保存自动化、Proto 代码生成 |

## kratosx 框架文档索引

| 文档 | 内容 |
|---|---|
| [kratosx/README.md](.claude/rules/kratosx/README.md) | 总览与快速上手 |
| [kratosx/01-core.md](.claude/rules/kratosx/01-core.md) | App 初始化、Context 接口、Option |
| [kratosx/02-config.md](.claude/rules/kratosx/02-config.md) | 配置加载、热更新、类型定义 |
| [kratosx/03-middleware.md](.claude/rules/kratosx/03-middleware.md) | 中间件链及使用方式 |
| [kratosx/04-database.md](.claude/rules/kratosx/04-database.md) | GORM 封装 |
| [kratosx/05-redis.md](.claude/rules/kratosx/05-redis.md) | Redis 多实例管理 |
| [kratosx/06-logger.md](.claude/rules/kratosx/06-logger.md) | zap 结构化日志 |
| [kratosx/07-jwt.md](.claude/rules/kratosx/07-jwt.md) | JWT 认证 |
| [kratosx/08-pool.md](.claude/rules/kratosx/08-pool.md) | 协程池与 WaitRunner |
| [kratosx/09-request.md](.claude/rules/kratosx/09-request.md) | HTTP 请求客户端 |
| [kratosx/10-client.md](.claude/rules/kratosx/10-client.md) | gRPC 客户端连接 |
| [kratosx/11-email.md](.claude/rules/kratosx/11-email.md) | 邮件服务 |
| [kratosx/12-captcha.md](.claude/rules/kratosx/12-captcha.md) | 验证码 |
| [kratosx/13-signature.md](.claude/rules/kratosx/13-signature.md) | API 签名验证 |
| [kratosx/14-lock.md](.claude/rules/kratosx/14-lock.md) | 分布式锁 |
| [kratosx/15-env.md](.claude/rules/kratosx/15-env.md) | 环境管理 |
| [kratosx/16-loader.md](.claude/rules/kratosx/16-loader.md) | 文件加载器 |
| [kratosx/17-prometheus.md](.claude/rules/kratosx/17-prometheus.md) | Prometheus 监控 |
| [kratosx/18-tasker.md](.claude/rules/kratosx/18-tasker.md) | 生命周期管理 |
| [kratosx/19-registry.md](.claude/rules/kratosx/19-registry.md) | 服务注册与发现 |
| [kratosx/20-pkg.md](.claude/rules/kratosx/20-pkg.md) | 工具包（value/Transform、crypto、tree） |
| [kratosx/21-model.md](.claude/rules/kratosx/21-model.md) | 模型与分页 |
| [kratosx/22-server.md](.claude/rules/kratosx/22-server.md) | CORS、响应格式化、静态文件 |
| [kratosx/23-cmd.md](.claude/rules/kratosx/23-cmd.md) | CLI 工具与 Proto 插件 |
| [kratosx/24-pprof.md](.claude/rules/kratosx/24-pprof.md) | 性能分析 |
| [kratosx/25-ip.md](.claude/rules/kratosx/25-ip.md) | 客户端 IP 获取 |
