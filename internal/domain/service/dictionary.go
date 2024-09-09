package service

import (
	"github.com/limes-cloud/kratosx"

	"layout/api/layout/errors"
	"layout/internal/conf"
	"layout/internal/domain/entity"
	"layout/internal/domain/repository"
	"layout/internal/types"
)

type Dictionary struct {
	conf *conf.Config
	repo repository.Dictionary
}

func NewDictionary(conf *conf.Config, repo repository.Dictionary) *Dictionary {
	return &Dictionary{conf: conf, repo: repo}
}

// GetDictionary 获取指定字典
func (u *Dictionary) GetDictionary(ctx kratosx.Context, req *types.GetDictionaryRequest) (*entity.Dictionary, error) {
	var (
		ent *entity.Dictionary
		err error
	)
	if req.Id != nil {
		ent, err = u.repo.GetDictionary(ctx, *req.Id)
	} else if req.Keyword != nil {
		ent, err = u.repo.GetDictionaryByKeyword(ctx, *req.Keyword)
	} else {
		return nil, errors.ParamsError()
	}
	if err != nil {
		return nil, errors.GetError(err.Error())
	}
	return ent, nil
}

// ListDictionary 获取字典目录列表
func (u *Dictionary) ListDictionary(ctx kratosx.Context, req *types.ListDictionaryRequest) ([]*entity.Dictionary, uint32, error) {
	list, total, err := u.repo.ListDictionary(ctx, req)
	if err != nil {
		return nil, 0, errors.ListError(err.Error())
	}
	return list, total, nil
}

// CreateDictionary 创建字典目录
func (u *Dictionary) CreateDictionary(ctx kratosx.Context, req *entity.Dictionary) (uint32, error) {
	id, err := u.repo.CreateDictionary(ctx, req)
	if err != nil {
		return 0, errors.CreateError(err.Error())
	}
	return id, nil
}

// UpdateDictionary 更新字典目录
func (u *Dictionary) UpdateDictionary(ctx kratosx.Context, req *entity.Dictionary) error {
	if err := u.repo.UpdateDictionary(ctx, req); err != nil {
		return errors.UpdateError(err.Error())
	}
	return nil
}

// DeleteDictionary 删除字典目录
func (u *Dictionary) DeleteDictionary(ctx kratosx.Context, id uint32) error {
	if err := u.repo.DeleteDictionary(ctx, id); err != nil {
		return errors.DeleteError(err.Error())
	}
	return nil
}
