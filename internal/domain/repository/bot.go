package repository

import "partyaffairs/internal/core"

type Bot interface {
	Session(ctx core.Context) (string, error)
	Send(ctx core.Context, sid string, msg string) (string, error)
	Chat(ctx core.Context, sid string, msg string, reply func(string), event func(tp string, builder any), done func()) error
}
