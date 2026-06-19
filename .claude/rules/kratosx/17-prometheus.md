# Prometheus 监控 (library/prometheus)

> 包路径：`github.com/limes-cloud/kratosx/library/prometheus`

## 概述

Prometheus 指标管理，支持 Counter、Gauge、Histogram、Summary 四种指标类型，配置驱动。

## 配置

```yaml
metrics: true    # 启用 /metrics 端点

prometheus:
  - name: extraction_total
    type: counter
    namespace: memory
    subsystem: pipeline
    help: "提取流水线执行总数"
    labels: ["track", "status"]

  - name: recall_latency
    type: histogram
    namespace: memory
    subsystem: recall
    help: "召回延迟"
    labels: ["channel"]
    buckets: [0.01, 0.05, 0.1, 0.25, 0.5, 1.0, 2.5, 5.0]

  - name: active_sessions
    type: gauge
    namespace: memory
    subsystem: session
    help: "活跃会话数"
    labels: ["user_type"]

  - name: embedding_duration
    type: summary
    namespace: memory
    subsystem: infra
    help: "Embedding 调用耗时"
    labels: ["provider"]
    objectives:
      0.5: 0.05
      0.9: 0.01
      0.99: 0.001
```

## Prometheus 接口

```go
type Prometheus interface {
    CounterVec(name string) *prometheus.CounterVec
    GaugeVec(name string) *prometheus.GaugeVec
    HistogramVec(name string) *prometheus.HistogramVec
    SummaryVec(name string) *prometheus.SummaryVec
}
```

## 使用方式

```go
prom := prometheus.Instance()

// Counter — 计数
prom.CounterVec("extraction_total").
    WithLabelValues("kg", "success").Inc()

// Histogram — 延迟
prom.HistogramVec("recall_latency").
    WithLabelValues("vector").Observe(0.123)

// Gauge — 当前值
prom.GaugeVec("active_sessions").
    WithLabelValues("vip").Set(42)

// Summary — 分位数
prom.SummaryVec("embedding_duration").
    WithLabelValues("openai").Observe(0.5)
```

## 辅助类型

框架提供了更简洁的封装类型：

```go
// Counter
c := prometheus.NewCounter(prom.CounterVec("extraction_total"))
c.With("kg", "success").Inc()
c.With("kg", "failed").Add(3)

// Gauge
g := prometheus.NewGauge(prom.GaugeVec("active_sessions"))
g.With("vip").Set(42)
g.With("vip").Add(1)
g.With("vip").Sub(1)

// Histogram / Summary (统一 Observer 接口)
h := prometheus.NewHistogram(prom.HistogramVec("recall_latency"))
h.With("vector").Observe(0.123)
```

## /metrics 端点

启用 `metrics: true` 后，HTTP Server 自动注册 `/metrics` 路由，暴露 Prometheus 指标。
