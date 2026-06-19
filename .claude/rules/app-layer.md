# App 层规则

> 参考 limes-cloud/manager/internal/app/ 的 handler + register 模式。

## 目录结构

```
internal/app/
├── register.go           # 自注册框架（全局变量 + register 函数）
├── conversation.go       # 对话写入 handler
└── recall.go             # 记忆召回 handler
```

## Handler 结构体

每个 handler 必须遵循以下模式：

```go
type Conversation struct {
    conversation.UnimplementedConversationServiceServer  // 嵌入 proto 生成的基类
    srv *service.Conversation                             // 注入 domain service
}
```

## 构造函数与依赖注入

在 `NewXxx()` 中完成 service → repo 的注入链：

```go
func NewConversation() *Conversation {
    return &Conversation{
        srv: service.NewConversation(
            dbs.NewConversation(),       // infra/dbs/ 仓储实现（含 PG + Redis 缓存）
        ),
    }
}
```

## 自注册模式

每个 handler 通过 `init() + register()` 自动注册 gRPC 和 HTTP 服务：

```go
func init() {
    register(func(hs *http.Server, gs *grpc.Server) {
        srv := NewConversation()
        conversation.RegisterConversationServiceHTTPServer(hs, srv)
        conversation.RegisterConversationServiceServer(gs, srv)
    })
}
```

**register.go 框架：**

```go
package app

import (
    "github.com/go-kratos/kratos/v2/transport/grpc"
    "github.com/go-kratos/kratos/v2/transport/http"
)

type registerFunc func(hs *http.Server, gs *grpc.Server)

var registerList []registerFunc

func register(fn registerFunc) {
    registerList = append(registerList, fn)
}

func Register(hs *http.Server, gs *grpc.Server) {
    for _, registry := range registerList {
        registry(hs, gs)
    }
}
```

## Handler 方法实现

每个方法遵循三步模式：

### 1. 上下文转换

```go
ctx := core.MustContext(c)   // c 是 context.Context，转为 core.Context
```

### 2. 请求转换（proto → types/entity）

```go
var in types.GetRoleRequest
if err := value.Transform(req, &in); err != nil {
    ctx.Logger().Errorw("msg", "req transform error", "err", err)
    return nil, errors.TransformError()
}
```

或直接转换为 entity：

```go
var in entity.Role
if err := value.Transform(req, &in); err != nil {
    ctx.Logger().Errorw("msg", "req transform error", "err", err)
    return nil, errors.TransformError()
}
```

### 3. 调用 service + 响应转换

```go
// 调用 service
result, err := h.srv.GetRole(ctx, &in)
if err != nil {
    return nil, err
}

// 响应转换（entity → proto）
reply := role.GetRoleReply{}
if err := value.Transform(result, &reply); err != nil {
    ctx.Logger().Errorw("msg", "resp transform error", "err", err)
    return nil, errors.TransformError()
}
return &reply, nil
```

### 简单场景

对于只传 ID 的简单请求，可以跳过 Transform：

```go
func (h *Conversation) CreateConversation(c context.Context, req *conversation.CreateConversationRequest) (*conversation.CreateConversationReply, error) {
    ctx := core.MustContext(c)
    id, err := h.srv.Create(ctx, &types.CreateConversationRequest{
        UserID:    req.GetUserId(),
        SessionID: req.GetSessionId(),
        Role:      req.GetRole(),
        Content:   req.GetContent(),
    })
    if err != nil {
        return nil, err
    }
    return &conversation.CreateConversationReply{Id: id}, nil
}
```

## 新增 Handler 的步骤

1. 在 `api/xxx/proto/` 定义 proto 文件
2. 运行 protoc 生成代码
3. 在 `internal/app/` 新建 `xxx.go` 文件
4. 定义 handler 结构体，嵌入 `UnimplementedXxxServer`
5. 实现 `NewXxx()` 构造函数，注入 service 和 dbs
6. 编写 `init()` + `register()` 自注册
7. 实现各 RPC 方法（三步模式：上下文转换 → 请求转换 → 调用 + 响应转换）
