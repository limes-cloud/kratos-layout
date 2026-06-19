# 代码风格规则

> 参考 limes-cloud/manager 和 kratosx 的代码风格。

## 命名规范

| 类型 | 规范 | 示例 |
|---|---|---|
| 文件名 | 全小写，下划线分隔 | `kv_memory.go`, `entity_resolver.go` |
| 包名 | 全小写，无下划线，与目录名一致 | `entity`, `repository`, `dbs` |
| 结构体 | 大驼峰，与业务领域对应 | `KVMemory`, `EntityResolver`, `Role` |
| 接口 | 大驼峰，与领域名词对应 | `Conversation`, `Memory`, `KGEntity` |
| 方法 | 大驼峰，动词开头 | `GetRole()`, `CreateConversation()`, `VectorSearch()` |
| 私有方法 | 小驼峰 | `appendRoleChildren()`, `buildEntityHints()` |
| 常量 | 大驼峰或全大写 | `EmbedModeFact`, `defaultBatchSize` |
| 错误函数 | 大驼峰 + Error 后缀 | `ParamsError()`, `GetError()`, `CreateError()` |
| 错误判断函数 | Is 前缀 | `IsParamsError()`, `IsNotFoundError()` |

## Nullable 字段

DB 中可空字段在 Go 中使用指针表示：

```go
// 正确
type Entity struct {
    Status      *bool       `db:"status"`
    Description *string     `db:"description"`
    Age         *int32      `db:"age"`
}

// 错误 — 不可空字段不应使用指针
type Entity struct {
    ID   *uint32  `db:"id"`    // ID 不可空，不要用指针
    Name *string  `db:"name"`  // Name 不可空，不要用指针
}
```

指针操作使用 `value.Pointer()` 和 `value.Value()`：

```go
// 创建指针
status := value.Pointer(true)
name := value.Pointer("test")

// 解引用（nil 安全）
isEnabled := value.Value(entity.Status)  // false if nil
desc := value.Value(entity.Description)  // "" if nil
```

条件查询中使用指针判断：

```go
if req.Status != nil {
    db = db.Where("status = ?", *req.Status)
}
```

## 错误处理

### 原则

1. **service 层**：捕获底层 error，日志记录，返回 proto 生成的错误函数
2. **app 层**：直接透传 service 返回的错误
3. **infra 层**：返回原生 error，不做包装

### 正确示例

```go
// service 层
func (u *Role) GetRole(ctx core.Context, id uint32) (*entity.Role, error) {
    res, err := u.repo.GetRole(ctx, id)
    if err != nil {
        ctx.Logger().Warnw("msg", "get role error", "err", err.Error())
        return nil, errors.GetError()
    }
    return res, nil
}

// app 层
func (h *Role) GetRole(c context.Context, req *role.GetRoleRequest) (*role.GetRoleReply, error) {
    ctx := core.MustContext(c)
    ent, err := h.srv.GetRole(ctx, req.Id)
    if err != nil {
        return nil, err
    }
    // ... 转换响应
}
```

### 错误示例

```go
// 错误 — 直接暴露底层错误
return nil, fmt.Errorf("database query failed: %w", err)

// 错误 — 在 app 层重新包装错误
return nil, errors.New(500, "GET_ERROR", err.Error())
```

## 构造函数模式

使用 `NewXxx()` 手动依赖注入，不使用 DI 容器：

```go
// service 层
type Role struct {
    repo  repository.Role
    rm    repository.RoleMenu
}

func NewRole(
    repo repository.Role,
    rm repository.RoleMenu,
) *Role {
    return &Role{
        repo: repo,
        rm:   rm,
    }
}

// infra 层 — 单例模式
type Role struct{}

var (
    roleIns  *Role
    roleOnce sync.Once
)

func NewRole() *Role {
    roleOnce.Do(func() {
        roleIns = &Role{}
    })
    return roleIns
}
```

## 日志规范

使用结构化日志，通过 `ctx.Logger()` 获取：

```go
// 警告级别 — 可预期的业务错误
ctx.Logger().Warnw("msg", "get role error", "err", err.Error())

// 错误级别 — 不可预期的系统错误
ctx.Logger().Errorw("msg", "database connection failed", "err", err.Error())

// 键名规范
// "msg" — 操作描述
// "err" — 错误信息
// "key" — 业务键名
// "id"  — 资源ID
```

## 注释规范

### 公开方法注释

- 每个公开方法必须有中文注释，描述其功能
- 注释放在方法上方，以方法名开头

```go
// GetRole 获取指定的角色信息
func (u *Role) GetRole(ctx core.Context, id uint32) (*entity.Role, error) {

// CreateConversation 写入对话记录
func (h *Conversation) CreateConversation(c context.Context, req *conversation.CreateConversationRequest) (*conversation.CreateConversationReply, error) {
```

### 私有方法注释

- 私有方法（小写开头）同样需要注释，描述其职责

```go
// selectModel 根据名称选择模型实例，默认返回 extraction
func (l *OpenAILLM) selectModel(name string) model.BaseChatModel {

// isPunctuation 判断字符串是否全由标点/符号组成
func isPunctuation(s string) bool {
```

### 字段注释

#### Entity 字段（强制）

**Entity 结构体的每个字段都必须有行尾注释**，注释内容根据字段类型有不同要求：

| 字段特征 | 必须说明的内容 |
|---|---|
| 枚举/有限取值 | 列出所有合法值（如 `user \| assistant`）|
| 时间字段（`int64`） | 说明时间单位：`Unix 秒` 或 `Unix 毫秒` |
| 时间字段（`time.Time`）| 直接说明用途即可 |
| 可空指针字段（`*T`）| 说明 NULL 的语义（如 "NULL = 当前有效"）|
| 向量字段（`pgvector.Vector`）| 说明嵌入对象 + 用途；如 JSON 不序列化需注明 |
| 数组字段（`pq.StringArray`、`pq.Int64Array`）| 说明数组元素含义及典型值 |
| 嵌入的 `model.CreateModel` / `model.BaseModel` | 说明提供的字段：`Id(uint32 自增主键) + CreatedAt(int64 Unix 秒)` |
| 普通字段（即使名称自解释）| 简短说明用途，加上业务约束或取值范围 |

```go
// 正确：entity 字段全部有注释
type Conversation struct {
    UserID    string `json:"userId"    gorm:"column:user_id"`    // 所属用户（记忆隔离）
    SessionID string `json:"sessionId" gorm:"column:session_id"` // 来源会话，同一会话共享滑动窗口
    Role      string `json:"role"      gorm:"column:role"`       // 发言角色：user | assistant
    Content   string `json:"content"   gorm:"column:content"`    // 消息正文
    model.CreateModel                                             // Id(uint32 自增主键) + CreatedAt(int64 Unix 秒)
}

type Memory struct {
    Confidence   float64         `json:"confidence"   gorm:"column:confidence"`  // 记忆置信度（0~1），由 LLM 提取时赋值
    ValidFrom    *int64          `json:"validFrom"    gorm:"column:valid_from"`  // 世界时间：事实开始成立（Unix 毫秒）；NULL 表示历史起点未知
    ValidTo      *int64          `json:"validTo"      gorm:"column:valid_to"`    // 世界时间：事实失效（Unix 毫秒）；NULL = 仍成立
    Embedding    pgvector.Vector `json:"-"            gorm:"column:embedding;type:vector(1536)"` // RawText 的语义向量，用于相似度检索；JSON 不序列化
    model.CreateModel                                                              // Id(uint32 自增主键) + CreatedAt(int64 Unix 秒)
}

// 错误：缺少注释或注释不完整
type Memory struct {
    ValidTo   *int64          // 缺少 NULL 语义说明
    Embedding pgvector.Vector // 缺少"用于什么"和"JSON 不序列化"说明
}
```

#### 配置类字段（强制）

配置结构体字段必须注释，标明含义和典型值：

```go
type RecallConf struct {
    VectorTopK int     `json:"vectorTopK"` // 向量检索返回条数
    RRFK       int     `json:"rrfK"`       // RRF 融合常数 k（通常 60）
}
```

#### 服务/仓储结构体字段（推荐）

服务和仓储的私有字段，当注入的接口用途不明显时加注释：

```go
type Recall struct {
    memoryRepo   repository.Memory   // 记忆仓储（向量+BM25 检索）
    kgEntityRepo repository.KGEntity // KG 实体仓储（实体查询+向量检索）
    relationRepo repository.Relation // KG 关系仓储（关系语义检索+BFS 遍历）
}

### 函数体内逻辑注释

- 函数内部的关键逻辑步骤需要注释，描述"做什么"和"为什么这样做"
- 注释放在代码块上方，不放在行尾（除非极短的行内说明）
- 重点关注：分支逻辑、数据转换、异步启动、错误处理策略、非直觉的算法步骤

```go
func (e *OpenAIEmbedder) EmbedBatch(ctx context.Context, texts []string, mode EmbedMode) ([][]float32, error) {
    // 根据 mode 为每段文本添加指令前缀（非对称嵌入：写入和查询使用不同前缀以提升检索效果）
    prefix := modePrefix[mode]
    prefixed := make([]string, len(texts))
    for i, t := range texts {
        prefixed[i] = prefix + t
    }

    // 调用 eino embedder，返回 float64 向量
    vecs64, err := e.client.EmbedStrings(ctx, prefixed)
    if err != nil {
        return nil, fmt.Errorf("embed strings: %w", err)
    }

    // 将 float64 转为 float32（pgvector 使用 float32 存储，节省空间）
    result := make([][]float32, len(vecs64))
    ...
}
```

- 以下情况**不需要**逻辑注释：
  - 一眼能看懂的简单赋值或返回
  - 变量命名已完全自解释的代码
  - TODO 桩方法（尚未实现的空函数）
