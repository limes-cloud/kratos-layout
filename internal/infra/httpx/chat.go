package httpx

import (
	"errors"
	"github.com/baidubce/app-builder/go/appbuilder"
	"io"
	"partyaffairs/internal/core"
)

type BotInfra struct {
}

const (
	defaultURL = "https://qianfan.baidubce.com"
)

func NewChatBot() *BotInfra {
	return &BotInfra{}
}

func (b *BotInfra) client(ctx core.Context) (*appbuilder.AppBuilderClient, error) {
	conf := ctx.Config().AiBot
	config, err := appbuilder.NewSDKConfig(defaultURL, conf.Secret)
	if err != nil {
		return nil, err
	}

	return appbuilder.NewAppBuilderClient(conf.Appid, config)
}

func (b *BotInfra) Session(ctx core.Context) (string, error) {
	client, err := b.client(ctx)
	if err != nil {
		return "", err
	}
	return client.CreateConversation()
}

func (b *BotInfra) Send(ctx core.Context, sid string, msg string) (string, error) {
	client, err := b.client(ctx)
	if err != nil {
		return "", err
	}

	ci, err := client.Run(sid, msg, []string{}, false)
	if err != nil {
		return "", err
	}

	answer, err := ci.Next()
	if err != nil && !errors.Is(err, io.EOF) {
		return "", err
	}
	return answer.Answer, nil
}

func (b *BotInfra) Chat(ctx core.Context, sid string, msg string, reply func(string), event func(tp string, builder any), done func()) error {
	client, err := b.client(ctx)
	if err != nil {
		return err
	}

	ci, err := client.Run(sid, msg, []string{}, true)
	if err != nil {
		return err
	}

	var answer *appbuilder.AppBuilderClientAnswer
	for answer, err = ci.Next(); err == nil; answer, err = ci.Next() {
		if reply != nil {
			reply(answer.Answer)
		}
		if event != nil {
			for _, ev := range answer.Events {
				event(ev.ContentType, ev.Detail)
			}
		}
	}

	if errors.Is(err, io.EOF) {
		done()
		return nil
	}
	return err
}
