# Infra 层规则

> 完全参照 limes-cloud/manager 的 dbs 风格，使用 kratosx 封装的 GORM（`ctx.DB()` 返回 `*gorm.DB`）。

## dbs/ — 数据仓储实现

### 单例模式

```go
type Conversation struct{}

var (
    conversationIns  *Conversation
    conversationOnce sync.Once
)

func NewConversation() *Conversation {
    conversationOnce.Do(func() {
        conversationIns = &Conversation{}
    })
    return conversationIns
}
```

---

## 方法命名规范

方法名必须与 `domain/repository/` 接口定义**完全一致**，不得自行重命名。本项目接口的主要命名模式：

| 操作 | 本项目命名示例 |
|---|---|
| 创建单条（含 UUID 生成） | `Create` / `Insert` |
| 根据 ID 查单条 | `GetByID` |
| 根据其他条件查单条 | `GetByName` / `GetByContentHash` |
| 无分页列表 | `ListBySession` / `ListTopN` / `ListAll` / `ListActiveByUser` |
| 软删除（设置 valid_to） | `Invalidate` |
| 批量硬删除 | `DeleteByIDs` |
| 写入或更新 | `Upsert` |

**注意：** 接口定义是权威依据，infra 层的签名必须完全匹配，包括参数类型和返回值类型。

---

## CRUD 操作规范

### UUID 生成

本项目主键为 string UUID，**在 dbs 层 Create/Insert 方法内生成**，不依赖数据库自增：

```go
func (c *Conversation) Create(ctx core.Context, conv *entity.Conversation) error {
    if conv.ID == "" {
        conv.ID = uuid.New().String()
    }
    return ctx.DB().Create(conv).Error
}
```

`CreatedAt time.Time` 无需手动赋值，GORM 按约定自动填充。

---

### Get — 查询单条

```go
// GetByID 根据 ID 获取实体
func (kg *KGEntity) GetByID(ctx core.Context, id string) (*entity.KGEntity, error) {
    var e entity.KGEntity
    err := ctx.DB().First(&e, "id = ?", id).Error
    if err != nil {
        return nil, err
    }
    return &e, nil
}

// GetByName 根据用户和名称精确查询
func (kg *KGEntity) GetByName(ctx core.Context, userID, name string) (*entity.KGEntity, error) {
    var e entity.KGEntity
    err := ctx.DB().Where("user_id = ? AND name = ?", userID, name).First(&e).Error
    if err != nil {
        return nil, err
    }
    return &e, nil
}
```

---

### List — 查询列表

```go
// ListBySession 获取会话对话列表（按时间正序）
func (c *Conversation) ListBySession(ctx core.Context, userID, sessionID string, limit int) ([]*entity.Conversation, error) {
    var list []*entity.Conversation
    err := ctx.DB().Where("user_id = ? AND session_id = ?", userID, sessionID).
        Order("created_at ASC").Limit(limit).Find(&list).Error
    return list, err
}

// ListTopN 获取用户的 top-N 实体
func (kg *KGEntity) ListTopN(ctx core.Context, userID string, limit int) ([]*entity.KGEntity, error) {
    var list []*entity.KGEntity
    err := ctx.DB().Where("user_id = ?", userID).
        Order("created_at DESC").Limit(limit).Find(&list).Error
    return list, err
}
```

---

### Update — 更新与软删除

```go
// Invalidate 软删除（设置世界时间 valid_to）
func (m *Memory) Invalidate(ctx core.Context, ids []string, worldTime time.Time) error {
    if len(ids) == 0 {
        return nil
    }
    return ctx.DB().Model(&entity.Memory{}).Where("id IN ?", ids).
        Update("valid_to", worldTime).Error
}

// IncrSupportCount 累加 support_count（非零值字段用 UpdateColumn + Expr）
func (m *Memory) IncrSupportCount(ctx core.Context, ids []string) error {
    if len(ids) == 0 {
        return nil
    }
    return ctx.DB().Model(&entity.Memory{}).Where("id IN ?", ids).
        UpdateColumn("support_count", gorm.Expr("support_count + 1")).Error
}

// MarkTrackDone 动态列名用 switch + map
func (d *Task) MarkTrackDone(ctx core.Context, id string, track string) error {
    var updates map[string]interface{}
    switch track {
    case "kg":
        updates = map[string]interface{}{"kg_done": true}
    case "kv":
        updates = map[string]interface{}{"kv_done": true}
    case "memory":
        updates = map[string]interface{}{"memory_done": true}
    default:
        return nil
    }
    return ctx.DB().Model(&entity.Task{}).Where("id = ?", id).Updates(updates).Error
}
```

**注意：**
- `Updates(struct)` 会忽略零值字段（如 `false`、`0`）；更新布尔/数字字段时用 `Updates(map[string]interface{}{...})`
- 单列原子更新用 `UpdateColumn`，避免触发 GORM 钩子

---

### Delete — 删除记录

```go
// DeleteByIDs 批量硬删除
func (c *Conversation) DeleteByIDs(ctx core.Context, ids []string) error {
    if len(ids) == 0 {
        return nil
    }
    return ctx.DB().Where("id IN ?", ids).Delete(&entity.Conversation{}).Error
}
```

---

### Upsert — 写入或更新

使用 `clause.OnConflict`，`Columns` 需与数据库唯一约束列完全对应：

```go
// Upsert 写入或更新 KV 记忆
func (d *KVMemory) Upsert(ctx core.Context, kv *entity.KVMemory) error {
    if kv.ID == "" {
        kv.ID = uuid.New().String()
    }
    return ctx.DB().Clauses(clause.OnConflict{
        Columns:   []clause.Column{{Name: "user_id"}, {Name: "key"}, {Name: "time_key"}},
        DoUpdates: clause.AssignmentColumns([]string{"value", "unit", "valid_from", "valid_to"}),
    }).Create(kv).Error
}
```

---

## 事务操作

统一使用 `ctx.Transaction(func(ctx core.Context) error { ... })` 闭包，不使用原生 `Begin/Commit`：

```go
err = ctx.Transaction(func(txCtx core.Context) error {
    if err := t.convRepo.DeleteByIDs(txCtx, ids); err != nil {
        return err
    }
    return t.taskRepo.Complete(txCtx, taskID)
})
```

---

## Redis 缓存

### 手动 Redis 操作（窗口缓存模式）

对于简单的 JSON 序列化缓存，直接操作 `ctx.Redis()`：

```go
const (
    windowCacheTTL    = 10 * time.Minute
    windowCachePrefix = "conv:window:"
)

// GetWindow 获取窗口对话（Redis 缓存优先，miss 则查 PG 并回填）
func (c *Conversation) GetWindow(ctx core.Context, userID, sessionID string, limit int) ([]*entity.Conversation, error) {
    key := windowCachePrefix + sessionID

    // 尝试读缓存
    val, err := ctx.Redis().Get(ctx.Ctx(), key).Result()
    if err == nil && val != "" {
        var cached []*entity.Conversation
        if err := json.Unmarshal([]byte(val), &cached); err == nil {
            return cached, nil
        }
    }

    // cache miss：查 PG
    list, err := c.ListBySession(ctx, userID, sessionID, limit)
    if err != nil {
        return nil, err
    }

    // 回填缓存
    if data, err := json.Marshal(list); err == nil {
        ctx.Redis().Set(ctx.Ctx(), key, data, windowCacheTTL)
    }
    return list, nil
}

// InvalidateWindowCache 使缓存失效
func (c *Conversation) InvalidateWindowCache(ctx core.Context, sessionID string) error {
    return ctx.Redis().Del(ctx.Ctx(), windowCachePrefix+sessionID).Err()
}
```

### Redis key 命名规则

```
conv:window:{sessionID}          // 对话窗口缓存
```

---

## 高级查询

### 原则：GORM 链式 API 优先

**所有查询统一使用 GORM 链式 API**，包括向量检索和 BM25 全文检索。`ctx.DB().Raw().Scan()` 仅保留给 `ExecRawQuery`（Text-to-SQL 专用）。

- 向量距离 `<=>` 通过 `gorm.Expr(...)` 内嵌到 `Where` / `Order` 子句
- BM25 `to_tsvector`/`@@` 同样通过 `Where` / `Order(gorm.Expr(...))` 表达
- 复杂对称条件（`OR`）直接写入 `Where` 字符串
- 读取完整 struct 后在 Go 层计算布尔结果（如 IsAllDone）

---

### 向量检索（pgvector）

向量参数必须用 `pgvector.NewVector(vec)` 包装后传入，不能直接传 `[]float32`。

**返回 DTO（Scan）：**

```go
// VectorSearch 语义向量检索（返回 MemorySearchResult DTO）
func (m *Memory) VectorSearch(ctx core.Context, userID string, vec []float32, threshold float64, limit int) ([]*entity.MemorySearchResult, error) {
    var results []*entity.MemorySearchResult
    v := pgvector.NewVector(vec)
    err := ctx.DB().Model(&entity.Memory{}).
        Select("id, raw_text, memory_type, (1 - (embedding <=> ?)) AS score", v).
        Where("user_id = ?", userID).
        Where("valid_to IS NULL").
        Where("1 - (embedding <=> ?) > ?", v, threshold).
        Order(gorm.Expr("embedding <=> ?", v)).
        Limit(limit).
        Scan(&results).Error
    return results, err
}
```

**返回模型（Find）：**

```go
// VectorSearch 向量检索实体（返回 entity 模型）
func (kg *KGEntity) VectorSearch(ctx core.Context, userID string, vec []float32, threshold float64, limit int) ([]*entity.KGEntity, error) {
    var list []*entity.KGEntity
    v := pgvector.NewVector(vec)
    err := ctx.DB().Model(&entity.KGEntity{}).
        Where("user_id = ?", userID).
        Where("1 - (embedding <=> ?) > ?", v, threshold).
        Order(gorm.Expr("embedding <=> ?", v)).
        Limit(limit).Find(&list).Error
    return list, err
}
```

**entity 向量字段定义：**

```go
type Memory struct {
    ID        string          `json:"id"        gorm:"primaryKey;column:id"`
    Embedding pgvector.Vector `json:"-"         gorm:"column:embedding;type:vector(1536)"`
    ValidTo   *time.Time      `json:"validTo"   gorm:"column:valid_to"`
    // ...
}
```

**检索结果 DTO 定义在 entity 包内：**

```go
type MemorySearchResult struct {
    ID              string  `json:"id"              gorm:"column:id"`
    RawText         string  `json:"rawText"         gorm:"column:raw_text"`
    MemoryType      string  `json:"memoryType"      gorm:"column:memory_type"`
    Score           float64 `json:"score"           gorm:"column:score"`
    RetrievalSource string  `json:"retrievalSource"` // 应用层赋值，无 gorm tag
}
```

---

### BM25 全文检索

```go
// BM25Search BM25 关键词检索（Memory）
func (m *Memory) BM25Search(ctx core.Context, userID string, queryTokens string, limit int) ([]*entity.MemorySearchResult, error) {
    var results []*entity.MemorySearchResult
    err := ctx.DB().Model(&entity.Memory{}).
        Select("id, raw_text, memory_type, ts_rank(to_tsvector('simple', keywords_text), plainto_tsquery('simple', ?), 2) AS score", queryTokens).
        Where("user_id = ?", userID).
        Where("valid_to IS NULL").
        Where("to_tsvector('simple', keywords_text) @@ plainto_tsquery('simple', ?)", queryTokens).
        Order("score DESC").Limit(limit).Scan(&results).Error
    return results, err
}

// BM25Search BM25 关键词检索（KVKey）
func (d *KVKey) BM25Search(ctx core.Context, queryTokens string, limit int) ([]*entity.KVKey, error) {
    var list []*entity.KVKey
    err := ctx.DB().Model(&entity.KVKey{}).
        Where("to_tsvector('simple', keywords_aliases) @@ plainto_tsquery('simple', ?)", queryTokens).
        Order(gorm.Expr("ts_rank(to_tsvector('simple', keywords_aliases), plainto_tsquery('simple', ?), 2) DESC", queryTokens)).
        Limit(limit).Find(&list).Error
    return list, err
}
```

---

### 复杂 OR 条件与 Go 层计算

```go
// QueryBetweenEntities 查询两实体间双向关系（OR 条件保留在一个 Where，其他拆开）
func (d *Relation) QueryBetweenEntities(ctx core.Context, entityA, entityB, userID string) ([]*entity.Relation, error) {
    var list []*entity.Relation
    err := ctx.DB().Where(
        "((from_id = ? AND to_id = ?) OR (from_id = ? AND to_id = ?))",
        entityA, entityB, entityB, entityA).
        Where("user_id = ?", userID).
        Where("valid_to IS NULL").Find(&list).Error
    return list, err
}

// IsAllDone 读取完整 struct，在 Go 层计算布尔结果（无需 Raw SQL）
func (d *Task) IsAllDone(ctx core.Context, id string) (bool, error) {
    var task entity.Task
    if err := ctx.DB().First(&task, "id = ?", id).Error; err != nil {
        return false, err
    }
    return task.KGDone && task.KVDone && task.MemoryDone, nil
}
```

---

## Text-to-SQL 执行

LLM 生成的 SQL 在执行前必须通过安全校验，校验失败直接返回 error（不使用 fallback）：

```go
// ExecRawQuery 执行 Text-to-SQL 生成的查询（含安全校验）
func (d *KVMemory) ExecRawQuery(ctx core.Context, sql string, args ...interface{}) ([]*entity.KVMemoryRow, error) {
    if err := validateKVRawSQL(sql); err != nil {
        return nil, err
    }
    var results []*entity.KVMemoryRow
    return results, ctx.DB().Raw(sql, args...).Scan(&results).Error
}

// validateKVRawSQL 校验规则：只读、只访问 kv_memory、必须含 user_id
func validateKVRawSQL(sql string) error {
    lower := strings.ToLower(strings.TrimSpace(sql))
    if !strings.HasPrefix(lower, "select") {
        return fmt.Errorf("only SELECT is allowed")
    }
    if !strings.Contains(lower, "kv_memory") {
        return fmt.Errorf("only kv_memory table is allowed")
    }
    if !strings.Contains(lower, "user_id") {
        return fmt.Errorf("WHERE must include user_id")
    }
    return nil
}
```

---

## llm/ — LLM 客户端

实现 `domain/repository.LLM` 接口，接口定义在 `domain/repository/llm.go`。

- 通过配置切换后端（OpenAI / 百度 / 自部署），infra 层只做实现
- 所有 System Prompt 集中在 `infra/llm/prompts.go` 中维护，按轨道分组：

```go
const (
    PromptKGExtract     = "你是一个知识图谱提取助手..."
    PromptKVExtract     = "你是一个键值对提取助手..."
    PromptMemoryExtract = "你是一个记忆总结助手..."
    PromptCutoffDecide  = "你是一个对话截断点判断助手..."
    PromptTextToSQL     = "你是一个 SQL 生成器..."
)
```

---

## embedding/ — Embedding 客户端

实现 `domain/repository.Embedder` 接口，接口定义在 `domain/repository/embedder.go`。

```go
// EmbedBatch 批量向量化（非对称嵌入：写入与查询用不同指令前缀）
func (e *OpenAIEmbedder) EmbedBatch(ctx context.Context, texts []string, mode EmbedMode) ([][]float32, error) {
    prefix := modePrefix[mode]
    prefixed := make([]string, len(texts))
    for i, t := range texts {
        prefixed[i] = prefix + t
    }

    vecs64, err := e.client.EmbedStrings(ctx, prefixed)
    if err != nil {
        return nil, fmt.Errorf("embed strings: %w", err)
    }

    // float64 → float32（pgvector 使用 float32 存储）
    result := make([][]float32, len(vecs64))
    for i, v := range vecs64 {
        result[i] = make([]float32, len(v))
        for j, f := range v {
            result[i][j] = float32(f)
        }
    }
    return result, nil
}
```

---

## tokenizer/ — jieba 分词

```go
type Tokenizer interface {
    // TokenizeForFTS jieba 分词 + 过滤标点 + 空格拼接（写入 keywords_text / BM25 查询共用）
    TokenizeForFTS(text string) string
    // Cut 原始分词（返回词列表）
    Cut(text string) []string
}
```
