package httpx

import (
	"errors"
	"github.com/baidubce/app-builder/go/appbuilder"
	"io"
)

type BotInfra struct {
	client *appbuilder.AppBuilderClient
}

const (
	defaultURL = "https://qianfan.baidubce.com"
)

func NewChatBotInfra(secret string, appid string) (*BotInfra, error) {
	config, err := appbuilder.NewSDKConfig(defaultURL, secret)
	if err != nil {
		return nil, err
	}

	client, err := appbuilder.NewAppBuilderClient(appid, config)
	if err != nil {
		return nil, err
	}

	return &BotInfra{
		client: client,
	}, nil
}

func (b *BotInfra) Session() (string, error) {
	return b.client.CreateConversation()
}

func (b *BotInfra) Send(sid string, msg string) (string, error) {
	ci, err := b.client.Run(sid, msg, []string{}, false)
	if err != nil {
		return "", err
	}

	answer, err := ci.Next()
	if err != nil && !errors.Is(err, io.EOF) {
		return "", err
	}
	return answer.Answer, nil
}

func (b *BotInfra) Chat(sid string, msg string, reply func(string), event func(tp string, builder any), done func()) error {
	ci, err := b.client.Run(sid, msg, []string{}, true)
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
