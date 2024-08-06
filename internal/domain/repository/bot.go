package repository

type BotRepository interface {
	Session() (string, error)
	Send(sid string, msg string) (string, error)
	Chat(sid string, msg string, reply func(string), event func(tp string, builder any), done func()) error
}
