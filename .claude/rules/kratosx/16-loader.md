# 文件加载器 (library/loader)

> 包路径：`github.com/limes-cloud/kratosx/library/loader`

## 概述

文件内容加载器，在应用启动时将文件内容加载到内存，运行时通过 key 快速获取。

## 配置

```yaml
loader:
  prompt_kg: "prompts/kg_extract.txt"
  prompt_kv: "prompts/kv_extract.txt"
  prompt_memory: "prompts/memory_extract.txt"
  email_template: "templates/email.html"
```

## 使用方式

```go
// 通过 context 获取文件内容
content := ctx.Loader("prompt_kg")
// content 为 []byte 类型

// 转为字符串使用
promptText := string(ctx.Loader("prompt_kg"))
```

## 适用场景

- Prompt 模板文件
- 邮件 HTML 模板
- SQL 模板文件
- 其他需要启动时加载的静态文件
