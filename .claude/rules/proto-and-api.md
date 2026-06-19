# Proto 与 API 规则

## api/ 独立 Go 模块

`api/` 目录使用独立的 `go.mod`，可被其他服务引用：

```
api/
├── conversation/
│   ├── proto/
│   │   ├── conversation.proto          # 消息定义
│   │   └── conversation_service.proto  # 服务定义
│   ├── conversation.pb.go
│   ├── conversation.pb.validate.go
│   ├── conversation_service.pb.go
│   ├── conversation_service_grpc.pb.go
│   ├── conversation_service_http.pb.go
│   └── openapi.yaml
├── search/
│   ├── proto/
│   │   ├── search.proto
│   │   └── search_service.proto
│   └── ...
├── errors/
│   ├── proto/
│   │   └── errors.proto               # 错误码定义
│   ├── errors.pb.go
│   └── errors_errors.pb.go            # 生成的错误函数
├── go.mod                              # 独立模块
└── go.sum
```

## Proto 文件规范

### package 命名规则

package 使用 `memory.api.<service>` 格式，以项目名为前缀：

| 文件路径 | package |
|---|---|
| `api/conversation/proto/xxx.proto` | `memory.api.conversation` |
| `api/search/proto/xxx.proto` | `memory.api.search` |
| `api/errors/proto/xxx.proto` | `memory.api.errors` |

规则：`memory.api.<服务名>`，不包含 `proto` 子目录层级（已禁用 `PACKAGE_DIRECTORY_MATCH` lint）。

### go_package 格式

使用项目内相对路径，**不含完整 module 前缀**（如 `github.com/xxx/`）：

```protobuf
option go_package = "memory/api/conversation;conversation";
```

格式：`memory/api/<service>;<package_alias>`

### import 路径

import 使用相对于项目根目录的完整路径：

```protobuf
// 导入同服务的消息定义
import "api/conversation/proto/conversation.proto";

// 导入第三方 proto（通过 third_party/ 软链接解析）
import "google/api/annotations.proto";

// 导入项目内错误码扩展
import "errorsx/errors.proto";
```

### 消息定义（xxx.proto）

```protobuf
syntax = "proto3";

package memory.api.conversation;

option go_package = "memory/api/conversation;conversation";

import "validate/validate.proto";

message CreateConversationRequest {
  string user_id = 1 [(validate.rules).string = {min_len: 1, max_len: 64}];
  string session_id = 2 [(validate.rules).string = {min_len: 1, max_len: 64}];
  string role = 3 [(validate.rules).string = {in: ["user", "assistant"]}];
  string content = 4 [(validate.rules).string = {min_len: 1, max_len: 20000}];
}

message CreateConversationResponse {
  string id = 1;
  int64 created_at = 2;
}
```

### 服务定义（xxx_service.proto）

```protobuf
syntax = "proto3";

package memory.api.conversation;

option go_package = "memory/api/conversation;conversation";

import "api/conversation/proto/conversation.proto";
import "google/api/annotations.proto";

service ConversationService {
  // CreateConversation 写入对话（同步）
  rpc CreateConversation(CreateConversationRequest) returns (CreateConversationResponse) {
    option (google.api.http) = {
      post: "/api/v1/conversation"
      body: "*"
    };
  }
}
```

### 编写规则

1. 每个服务拆分为 `xxx.proto`（消息）+ `xxx_service.proto`（服务）
2. `xxx_service.proto` 必须 import 对应的 `xxx.proto`（使用完整相对路径）
3. `xxx_service.proto` 必须 import `google/api/annotations.proto` 以定义 HTTP 路由
4. HTTP 路由通过 `google.api.http` 注解定义
5. 字段命名使用 snake_case（proto 规范），生成 Go 代码自动转为 CamelCase
6. 时间戳使用 `int64`（Unix 毫秒），不使用 `google.protobuf.Timestamp`
7. 同一 `proto/` 目录下的所有文件必须使用相同的 `package` 和 `go_package`
8. RPC 请求/响应消息命名必须使用 `<MethodName>Request` 和 `<MethodName>Response`；不要使用 `<MethodName>Reply`
9. 所有 RPC Request 消息必须使用 `validate/validate.proto` 限制请求参数；字符串至少设置非空规则，外部输入字段必须设置合理长度上限或枚举范围
10. Response 消息不需要添加 validate 规则；返回内容由服务端构造，边界控制应在业务代码中完成

## 错误码定义

### errors.proto

```protobuf
syntax = "proto3";

package memory.api.errors;

import "errorsx/errors.proto";
option go_package = "memory/api/errors;errors";

enum ErrorReason {
  option (errors.default_code) = 500;

  ParamsError     = 0 [(errors.message)="参数错误"];
  SystemError     = 1 [(errors.message)="系统错误"];
  DatabaseError   = 2 [(errors.message)="数据库错误"];
  TransformError  = 3 [(errors.message)="数据转换失败"];
  GetError        = 4 [(errors.message)="获取数据失败"];
  ListError       = 5 [(errors.message)="获取列表数据失败"];
  CreateError     = 6 [(errors.message)="创建数据失败"];
  UpdateError     = 7 [(errors.message)="更新数据失败"];
  DeleteError     = 8 [(errors.message)="删除数据失败"];

  // 业务错误
  WindowTriggerError  = 20 [(errors.message)="窗口触发失败"];
  ExtractionError     = 21 [(errors.message)="记忆提取失败"];
  EmbeddingError      = 22 [(errors.message)="向量嵌入失败"];
  LLMCallError        = 23 [(errors.message)="LLM调用失败"];
  TextToSQLError      = 24 [(errors.message)="SQL生成失败"];
  SQLValidationError  = 25 [(errors.message)="SQL校验失败"];
}
```

### 使用生成的错误函数

```go
import "memory/api/errors"

// 无参数 — 返回默认消息
return nil, errors.GetError()

// 带参数 — 追加详细信息
return nil, errors.CreateError("窗口大小不足")

// 判断错误类型
if errors.IsParamsError(err) {
    // 处理参数错误
}
```

## 代码生成

使用 kratosx proto 插件生成代码：

```bash
# 生成消息 + 服务 + gRPC + HTTP + OpenAPI
kratosx proto client api/conversation/proto/conversation_service.proto

# 生成错误码
kratosx proto client api/errors/proto/errors.proto

# 全量生成
make proto
```

生成文件包括：
- `*.pb.go`：消息结构体
- `*_grpc.pb.go`：gRPC 服务端/客户端
- `*_http.pb.go`：HTTP 网关
- `openapi.yaml`：OpenAPI 文档
- `*_errors.pb.go`：错误函数

**注意：**
- 生成代码不手动修改，修改 proto 后重新生成
- `kratosx proto client` 接收**相对路径**参数（相对于项目根目录）
- VSCode 保存 `.proto` 文件时由 RunOnSave 自动触发生成（见 vscode.md）

## buf.yaml 配置

项目根目录的 `buf.yaml` 定义 proto 模块结构：

```yaml
version: v2
modules:
  - path: .
    excludes:
      - third_party
  - path: third_party
```

- 模块 `.`：项目自有 proto（排除 third_party 避免重复扫描）
- 模块 `third_party`：第三方 proto 依赖（google/api、validate、errorsx）

buf 扩展依据此文件解析 import 路径和执行 lint 检查。
