package app

import (
	"context"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/limes-cloud/kratosx"

	"tianji/api/bazi"
	"tianji/internal/domain/service"
	infrabazi "tianji/internal/infra/bazi"
)

type BaziHandler struct {
	bazi.UnimplementedBaziServer
	srv *service.Bazi
}

func init() {
	register(func(hs *http.Server, gs *grpc.Server) {
		h := &BaziHandler{srv: service.NewBazi()}
		bazi.RegisterBaziHTTPServer(hs, h)
		bazi.RegisterBaziServer(gs, h)
	})
}

// Calculate 计算八字命盘
func (h *BaziHandler) Calculate(c context.Context, req *bazi.BaziRequest) (*bazi.BaziResponse, error) {
	ctx := kratosx.MustContext(c)
	result, err := h.srv.Calculate(ctx, &service.BaziRequest{
		Name:      req.Name,
		Gender:    req.Gender,
		Year:      int(req.Year),
		Month:     int(req.Month),
		Day:       int(req.Day),
		Hour:      int(req.Hour),
		Minute:    int(req.Minute),
		Location:  req.Location,
		IsLunar:   req.IsLunar,
		LunarLeap: req.LunarLeap,
	})
	if err != nil {
		return nil, err
	}
	return toProtoResponse(result), nil
}

// Report 生成八字命盘 Markdown 报告
func (h *BaziHandler) Report(c context.Context, req *bazi.BaziRequest) (*bazi.BaziReportResponse, error) {
	ctx := kratosx.MustContext(c)
	result, err := h.srv.Calculate(ctx, &service.BaziRequest{
		Name:      req.Name,
		Gender:    req.Gender,
		Year:      int(req.Year),
		Month:     int(req.Month),
		Day:       int(req.Day),
		Hour:      int(req.Hour),
		Minute:    int(req.Minute),
		Location:  req.Location,
		IsLunar:   req.IsLunar,
		LunarLeap: req.LunarLeap,
	})
	if err != nil {
		return nil, err
	}
	return &bazi.BaziReportResponse{Md: infrabazi.BuildMarkdownReport(result)}, nil
}

func toProtoResponse(r *infrabazi.BaziResult) *bazi.BaziResponse {
	resp := &bazi.BaziResponse{
		Name:        r.Name,
		Gender:      r.Gender,
		Zodiac:      r.Zodiac,
		LunarDate:   r.LunarDate,
		SolarTerm:   r.SolarTerm,
		FetalOrigin: r.FetalOrigin,
		OwnSign:     r.OwnSign,
		BodySign:    r.BodySign,
		Strength:    r.Strength,
		ChildLimit:  r.ChildLimit,
		YearPillar:  toPillar(&r.YearPillar),
		MonthPillar: toPillar(&r.MonthPillar),
		DayPillar:   toPillar(&r.DayPillar),
		HourPillar:  toPillar(&r.HourPillar),
		Wuxing: &bazi.WuxingScore{
			Wood:  r.Wuxing.Wood,
			Fire:  r.Wuxing.Fire,
			Earth: r.Wuxing.Earth,
			Metal: r.Wuxing.Metal,
			Water: r.Wuxing.Water,
		},
	}
	for _, d := range r.DecadeFortunes {
		resp.DecadeFortunes = append(resp.DecadeFortunes, &bazi.DecadeFortune{
			SixtyCycle: d.SixtyCycle,
			StartAge:   d.StartAge,
			EndAge:     d.EndAge,
			StartYear:  d.StartYear,
		})
	}
	for _, a := range r.AnnualFortunes {
		resp.AnnualFortunes = append(resp.AnnualFortunes, &bazi.AnnualFortune{
			CalYear:    a.CalYear,
			SixtyCycle: a.SixtyCycle,
			Age:        a.Age,
			TenStar:    a.TenStar,
		})
	}
	return resp
}

func toPillar(p *infrabazi.PillarResult) *bazi.Pillar {
	if p == nil {
		return nil
	}
	return &bazi.Pillar{
		Stem:        p.Stem,
		Branch:      p.Branch,
		Sound:       p.Sound,
		TenStar:     p.TenStar,
		Terrain:     p.Terrain,
		SelfSeat:    p.SelfSeat,
		HiddenStems: p.HiddenStems,
		Xukong:      p.Xukong,
		Shensha:     p.Shensha,
	}
}
