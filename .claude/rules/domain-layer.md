# Domain 层规则

## entity/ — 领域实体

### 结构体规范

- 纯结构体，只包含字段和简单方法，不含业务逻辑
- DB tag 使用 `gorm`（`gorm:"column:xxx"`，主键用 `gorm:"primaryKey"`）
- JSON tag 使用小驼峰
- 公共字段（id、created_at、updated_at）通过嵌入 `model.BaseModel` 提供，放在结构体最后一行

```go
package entity

import "github.com/limes-cloud/kratosx/model"

type Conversation struct {
    UserID    string `json:"userId"    gorm:"column:user_id"`
    SessionID string `json:"sessionId" gorm:"column:session_id"`
    Role      string `json:"role"      gorm:"column:role"`
    Content   string `json:"content"   gorm:"column:content"`
    model.BaseModel  // 提供 Id、CreatedAt、UpdatedAt，放最后
}
```

### nullable 字段

可空数据库字段使用指针：

```go
type Memory struct {
    RawText    string   `json:"rawText"    gorm:"column:raw_text"`
    Confidence *float64 `json:"confidence" gorm:"column:confidence"` // 可空
    ValidTo    *int64   `json:"validTo"    gorm:"column:valid_to"`   // 可空，NULL = 当前有效
    model.BaseModel
}
```

### 关联字段

关联关系在 repository 层手动查询，不定义外键标签：

```go
type Relation struct {
    FromID    string `json:"fromId"    gorm:"column:from_id"`
    ToID      string `json:"toId"      gorm:"column:to_id"`
    Predicate string `json:"predicate" gorm:"column:predicate"`
    FactText  string `json:"factText"  gorm:"column:fact_text"`
    model.BaseModel
}
```

### 向量字段

pgvector 的 embedding 字段使用 `[]float32`：

```go
type Memory struct {
    Embedding pgvector.Vector `json:"-" gorm:"column:embedding;type:vector(1536)"` // 向量字段，JSON 不序列化
    model.BaseModel
}
```

## repository/ — 仓储接口

### 接口定义规范

- 第一个参数必须是 `ctx core.Context`
- 方法命名：`GetXxx` / `GetXxxByYyy` / `ListXxx` / `CreateXxx` / `UpdateXxx` / `DeleteXxx`（动词 + 实体名，与 manager 完全一致）
- 查询单条返回 `(*entity.Xxx, error)`
- 查询列表（无分页）返回 `([]*entity.Xxx, error)`
- 查询列表（分页）返回 `([]*entity.Xxx, uint32, error)`（第二个值为总数 uint32）
- 创建方法返回 `(string, error)`（本项目主键为 string UUID）

```go
package repository

import (
    "github.com/xxx/memory/internal/core"
    "github.com/xxx/memory/internal/domain/entity"
    "github.com/xxx/memory/internal/types"
)

// Conversation 对话仓储接口
type Conversation interface {
    // GetConversation 根据 ID 查询对话
    GetConversation(ctx core.Context, id string) (*entity.Conversation, error)

    // ListConversationBySession 查询会话对话列表
    ListConversationBySession(ctx core.Context, req *types.ListConversationRequest) ([]*entity.Conversation, error)

    // CreateConversation 创建对话记录
    CreateConversation(ctx core.Context, ent *entity.Conversation) (string, error)

    // DeleteConversationByIDs 批量删除对话
    DeleteConversationByIDs(ctx core.Context, ids []string) error
}
```

### 特殊查询方法

对于有复杂检索需求（向量、BM25），在接口中定义专用方法：

```go
// Memory 记忆仓储接口
type Memory interface {
    // ListMemory 分页查询记忆
    ListMemory(ctx core.Context, req *types.ListMemoryRequest) ([]*entity.Memory, uint32, error)

    // VectorSearchMemory 语义向量检索
    VectorSearchMemory(ctx core.Context, userID string, vec []float32, threshold float64, limit int) ([]*entity.MemorySearchResult, error)

    // BM25SearchMemory BM25 关键词检索
    BM25SearchMemory(ctx core.Context, userID string, queryTokens string, limit int) ([]*entity.MemorySearchResult, error)

    // InvalidateMemory 软删除（设置 valid_to）
    InvalidateMemory(ctx core.Context, ids []string) error
}
```

### 事务方法

事务不在 repository 接口中定义，而是通过 `ctx.Transaction()` 在 service 层编排：

```go
// service 层使用事务
err := ctx.Transaction(func(ctx core.Context) error {
    if err := u.repo.DeleteRelation(ctx, id); err != nil {
        return err
    }
    return u.repo.InvalidateMemory(ctx, ids)
})
```

## service/ — 领域服务

### 构造函数

通过 `NewXxx()` 注入 repository 接口：

```go
type Conversation struct {
    repo   repository.Conversation
    window repository.Window
}

func NewConversation(
    repo repository.Conversation,
    window repository.Window,
) *Conversation {
    return &Conversation{
        repo:   repo,
        window: window,
    }
}
```

### 业务方法

- 第一个参数 `ctx core.Context`
- 错误日志 + 返回 proto 错误函数
- 复杂业务逻辑可拆分为私有方法

```go
// Create 写入对话并检查窗口
func (s *Conversation) Create(ctx core.Context, req *types.CreateConversationRequest) (string, error) {
    conv := &entity.Conversation{
        UserID:    req.UserID,
        SessionID: req.SessionID,
        Role:      req.Role,
        Content:   req.Content,
    }
    id, err := s.repo.Create(ctx, conv)
    if err != nil {
        ctx.Logger().Warnw("msg", "create conversation error", "err", err.Error())
        return "", errors.CreateError()
    }

    // 检查窗口是否触发提取
    if err := s.window.CheckAndTrigger(ctx, req.SessionID); err != nil {
        ctx.Logger().Warnw("msg", "check window trigger error", "err", err.Error())
        // 提取触发失败不影响对话写入，仅记录日志
    }

    return id, nil
}
```

### 禁止事项

1. **禁止在 service 中直接 import DB / Redis 客户端**：数据库操作必须通过 repository 接口
2. **禁止在 service 中直接 import infra 包**：依赖注入在 app 层完成
3. **禁止在 service 中返回原生 error**：必须使用 proto 生成的错误函数
4. **禁止在 service 中处理 HTTP/gRPC 协议细节**：proto 转换在 app 层完成