package service

import (
	"github.com/limes-cloud/kratosx"
	"github.com/limes-cloud/usercenter/api/usercenter/auth"
	"regexp"
	"strings"

	"github.com/limes-cloud/kratosx/library/pool"
	"poverty/api/poverty/errors"
	"poverty/internal/conf"
	"poverty/internal/domain/entity"
	"poverty/internal/domain/repository"
	"poverty/internal/types"
)

type ChatService struct {
	conf *conf.Config
	repo repository.ChatRepository
	bot  repository.BotRepository
	user repository.UserRepository
}

func NewChatService(
	conf *conf.Config,
	repo repository.ChatRepository,
	bot repository.BotRepository,
	user repository.UserRepository,
) *ChatService {
	return &ChatService{conf: conf, repo: repo, bot: bot, user: user}
}

func (srv *ChatService) GenSessionId(ctx kratosx.Context) (string, error) {
	session, err := srv.bot.Session()
	if err != nil {
		ctx.Logger().Warnw("msg", "ai chat session error", "err", err.Error())
		return "", errors.AIChatServerError()
	}
	return session, err
}

// SendChatMessage 发送会话信息
func (srv *ChatService) SendChatMessage(ctx kratosx.Context, req *types.SendChatMessageRequest) error {
	var text string
	info, err := auth.Get(ctx)
	if err != nil {
		return errors.SystemError()
	}

	re := regexp.MustCompile(`\^\[(?:[0-9]+)\](?:\[(?:[0-9]+)\])*\^`)

	reply := func(msg string) {
		msg = strings.ReplaceAll(msg, "  ", "")
		msg = re.ReplaceAllString(msg, "")
		if msg == "" {
			return
		}
		text += msg
		req.Reply(msg)
	}

	done := func() {
		req.Done()
		// 写入回复数据
		_ = ctx.Go(pool.AddRunner(func() error {
			_, err = srv.repo.CreateChatRecord(ctx.Clone(), &entity.ChatRecord{
				UserId:    info.UserId,
				SessionId: req.SessionId,
				Message:   text,
				Type:      "reply",
			})
			return err
		}))
	}

	_ = ctx.Go(pool.AddRunner(func() error {
		// 写入发送数据
		_, err = srv.repo.CreateChatRecord(ctx.Clone(), &entity.ChatRecord{
			UserId:    info.UserId,
			SessionId: req.SessionId,
			Message:   req.Message,
			Type:      "send",
		})
		return err
	}))

	if err := srv.bot.Chat(req.SessionId, req.Message, reply, req.Event, done); err != nil {
		ctx.Logger().Warnw("msg", "ai chat send error", "err", err.Error())
		return errors.AIChatServerError()
	}
	return nil
}

//// SendChatMessage 发送会话信息
//func (srv *ChatService) SendChatMessage(ctx kratosx.Context, req *types.SendChatMessageRequest) (*types.SendChatMessageReply, error) {
//	if req.SessionId == nil {
//		session, err := srv.bot.Session()
//		if err != nil {
//			ctx.Logger().Warnw("msg", "ai chat session error", "err", err.Error())
//			return nil, errors.AIChatServerError()
//		}
//		req.SessionId = &session
//	}
//	err := srv.bot.Chat(*req.SessionId, req.Message, req.Reply, req.Event)
//	if err != nil {
//		ctx.Logger().Warnw("msg", "ai chat send error", "err", err.Error())
//		return nil, errors.AIChatServerError()
//	}
//	return &types.SendChatMessageReply{
//		SessionId: *req.SessionId,
//		Message:   msg,
//	}, nil
//}

// ListChatRecord 获取会话记录信息列表
func (srv *ChatService) ListChatRecord(ctx kratosx.Context, req *types.ListChatRecordRequest) ([]*entity.ChatRecord, uint32, error) {
	if req.UserId == nil && req.UserName != nil {
		users, err := srv.user.ListUserByName(ctx, *req.UserName)
		if err != nil {
			return nil, 0, errors.ListError(err.Error())
		}
		req.UserIds = make([]uint32, 0)
		for _, item := range users {
			req.UserIds = append(req.UserIds, item.Id)
		}
	}

	list, total, err := srv.repo.ListChatRecord(ctx, req)
	if err != nil {
		return nil, 0, errors.ListError(err.Error())
	}

	for ind, item := range list {
		user, err := srv.user.GetUser(ctx, item.UserId)
		if err != nil {
			continue
		}
		if user.RealName != nil {
			list[ind].UserName = *user.RealName
		}
		if user.AvatarUrl != nil {
			list[ind].UserAvatar = *user.AvatarUrl
		}
	}
	return list, total, nil
}

// DeleteChatRecord 删除会话记录信息
func (srv *ChatService) DeleteChatRecord(ctx kratosx.Context, id uint32) error {
	if err := srv.repo.DeleteChatRecord(ctx, id); err != nil {
		return errors.DeleteError(err.Error())
	}
	return nil
}
