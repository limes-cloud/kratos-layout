package service

import (
	"github.com/limes-cloud/kratosx"

	"tianji/api/errors"
	"tianji/internal/infra/bazi"
)

type Bazi struct{}

func NewBazi() *Bazi { return &Bazi{} }

type BaziRequest struct {
	Name, Gender, Location string
	Year, Month, Day, Hour, Minute int
	IsLunar   bool
	LunarLeap bool
}

func (s *Bazi) Calculate(ctx kratosx.Context, req *BaziRequest) (*bazi.BaziResult, error) {
	result, err := bazi.Calculate(req.Name, req.Gender, req.Year, req.Month, req.Day, req.Hour, req.Minute, req.Location, req.IsLunar, req.LunarLeap)
	if err != nil {
		ctx.Logger().Warnw("msg", "bazi calculate error", "err", err.Error())
		return nil, errors.GetError(err.Error())
	}
	return result, nil
}
