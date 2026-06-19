# Types DTO 规则

> 参考 limes-cloud/manager/internal/types/ 的 DTO 模式。

## 请求 DTO

### 基本规则

- 仅包含 JSON tag，不含 DB tag
- 查询条件中的可选参数使用指针
- 必填字段使用值类型

```go
package types

// CreateConversationRequest 创建对话请求
type CreateConversationRequest struct {
    UserID    string `json:"userId"`     // 必填
    SessionID string `json:"sessionId"`  // 必填
    Role      string `json:"role"`       // 必填
    Content   string `json:"content"`    // 必填
}

// GetMemoryRequest 获取记忆请求
type GetMemoryRequest struct {
    ID   *string `json:"id"`   // 可选，ID 或 Type 二选一
    Type *string `json:"type"` // 可选
}
```

### 列表查询 DTO

列表查询嵌入 `page.Search`，分页/排序由 `page.SearchScopes` 统一处理：

```go
import "github.com/limes-cloud/kratosx/model/page"

// ListMemoryRequest 记忆列表查询请求
type ListMemoryRequest struct {
    page.Search        // 含 Page、PageSize、OrderBy、Order

    // 过滤条件
    UserID     string   `json:"userId"`     // 必填
    MemoryType *string  `json:"memoryType"` // 可选，按类型过滤
    InIDs      []string `json:"inIds"`      // 可选，ID 包含过滤
}
```

### 中转参数

有些参数不在 proto 中定义，而是在 service 层逻辑中推导后设置，称为"中转参数"：

```go
type ListUserRequest struct {
    // ... 查询条件

    // 中转参数（由 service 层根据权限推导设置）
    InDeptIds []uint32 `json:"inDeptIds"`  // 由 DeptId 推导
    InJobIds  []uint32 `json:"inJobIds"`   // 由 JobId 推导
}
```

中转参数不在 proto 请求中暴露，由 service 层设置后传给 repository 用于数据库查询。

## 响应 DTO

通常不单独定义响应 DTO，而是直接使用 entity 结构体 + `value.Transform` 转换为 proto 响应。

对于列表查询，repository 返回 `([]*entity.Xxx, uint32, error)`，第二个值为总数：

```go
// repository 接口
type Memory interface {
    ListMemory(ctx core.Context, req *types.ListMemoryRequest) ([]*entity.Memory, uint32, error)
}

// service 层
func (s *Memory) ListMemory(ctx core.Context, req *types.ListMemoryRequest) ([]*entity.Memory, uint32, error) {
    list, total, err := s.repo.ListMemory(ctx, req)
    if err != nil {
        ctx.Logger().Warnw("msg", "list memory error", "err", err.Error())
        return nil, 0, errors.ListError()
    }
    return list, total, nil
}

// app 层
func (h *Memory) ListMemory(c context.Context, req *memory.ListMemoryRequest) (*memory.ListMemoryReply, error) {
    var (
        ctx = core.MustContext(c)
        in  types.ListMemoryRequest
    )
    if err := value.Transform(req, &in); err != nil {
        return nil, errors.TransformError()
    }
    list, total, err := h.srv.ListMemory(ctx, &in)
    if err != nil {
        return nil, err
    }
    reply := &memory.ListMemoryReply{Total: total}
    if err := value.Transform(list, &reply.List); err != nil {
        return nil, errors.TransformError()
    }
    return reply, nil
}
```

## 搜索结果 DTO

对于向量检索和 BM25 检索，返回带 score 的结果：

```go
// MemorySearchResult 记忆检索结果（带分数），定义在 entity 包内
type MemorySearchResult struct {
    ID              string  `json:"id"         gorm:"column:id"`
    RawText         string  `json:"rawText"    gorm:"column:raw_text"`
    MemoryType      string  `json:"memoryType" gorm:"column:memory_type"`
    Score           float64 `json:"score"      gorm:"column:score"`
    RetrievalSource string  `json:"retrievalSource"` // vector / bm25，应用层赋值，无 gorm tag
}
```
