package core

import (
	"context"

	"github.com/limes-cloud/kratosx"
)

// Context 封装 kratosx.Context，提供业务扩展方法
type Context struct {
	kratosx.Context
}

// Transaction 事务，闭包内自动使用事务连接，ctx 类型为 core.Context
func (c Context) Transaction(fn func(ctx Context) error, name ...string) error {
	return c.Context.Transaction(func(ctx kratosx.Context) error {
		return fn(Context{Context: ctx})
	}, name...)
}

// Config 获取业务配置单例
func (Context) Config() *Conf {
	return conf
}

// Clone 克隆上下文，用于异步任务（脱离原始 context 的 cancel/timeout）
func (c Context) Clone() Context {
	return MustContext(context.WithoutCancel(c.Context))
}

// MustContext 将标准 context 转为 core.Context
func MustContext(ctx context.Context, opts ...kratosx.ContextOptionFunc) Context {
	return Context{Context: kratosx.MustContext(ctx, opts...)}
}
