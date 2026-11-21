package bazi

import (
	"fmt"
	"strings"
)

// BuildMarkdownReport 将 BaziResult 格式化为 Markdown 报告
func BuildMarkdownReport(r *BaziResult) string {
	var b strings.Builder

	b.WriteString(fmt.Sprintf("# %s 八字命盘\n\n", r.Name))
	b.WriteString(fmt.Sprintf("**性别：** %s　**生肖：** %s\n\n",
		map[string]string{"M": "男", "F": "女"}[r.Gender], r.Zodiac))
	b.WriteString(fmt.Sprintf("**农历：** %s\n\n", r.LunarDate))
	b.WriteString(fmt.Sprintf("**节气：** %s\n\n", r.SolarTerm))

	b.WriteString("---\n\n")

	// 四柱
	b.WriteString("## 四柱\n\n")
	b.WriteString("| | 年柱 | 月柱 | 日柱 | 时柱 |\n")
	b.WriteString("|---|---|---|---|---|\n")
	writePillarRow := func(label string, f func(p *PillarResult) string) {
		b.WriteString(fmt.Sprintf("| **%s** | %s | %s | %s | %s |\n",
			label, f(&r.YearPillar), f(&r.MonthPillar), f(&r.DayPillar), f(&r.HourPillar)))
	}
	writePillarRow("天干", func(p *PillarResult) string { return p.Stem })
	writePillarRow("地支", func(p *PillarResult) string { return p.Branch })
	writePillarRow("十神", func(p *PillarResult) string { return p.TenStar })
	writePillarRow("纳音", func(p *PillarResult) string { return p.Sound })
	writePillarRow("地势", func(p *PillarResult) string { return p.Terrain })
	writePillarRow("自坐", func(p *PillarResult) string { return p.SelfSeat })
	writePillarRow("藏干", func(p *PillarResult) string { return p.HiddenStems })
	writePillarRow("空亡", func(p *PillarResult) string { return p.Xukong })
	writePillarRow("神煞", func(p *PillarResult) string { return p.Shensha })
	b.WriteString("\n")

	// 特殊信息
	b.WriteString("## 命理信息\n\n")
	b.WriteString(fmt.Sprintf("- **胎元：** %s\n", r.FetalOrigin))
	b.WriteString(fmt.Sprintf("- **命宫：** %s\n", r.OwnSign))
	b.WriteString(fmt.Sprintf("- **身宫：** %s\n", r.BodySign))
	b.WriteString(fmt.Sprintf("- **旺相休囚死：** %s\n", r.Strength))
	b.WriteString(fmt.Sprintf("- **童限：** %s\n\n", r.ChildLimit))

	// 五行衰旺
	b.WriteString("## 五行衰旺\n\n")
	b.WriteString("| 五行 | 占比 |\n")
	b.WriteString("|---|---|\n")
	b.WriteString(fmt.Sprintf("| 木 | %.2f%% |\n", r.Wuxing.Wood*100))
	b.WriteString(fmt.Sprintf("| 火 | %.2f%% |\n", r.Wuxing.Fire*100))
	b.WriteString(fmt.Sprintf("| 土 | %.2f%% |\n", r.Wuxing.Earth*100))
	b.WriteString(fmt.Sprintf("| 金 | %.2f%% |\n", r.Wuxing.Metal*100))
	b.WriteString(fmt.Sprintf("| 水 | %.2f%% |\n\n", r.Wuxing.Water*100))

	// 大运
	b.WriteString("## 大运\n\n")
	b.WriteString("| 大运 | 年龄 | 起运年 |\n")
	b.WriteString("|---|---|---|\n")
	for _, df := range r.DecadeFortunes {
		b.WriteString(fmt.Sprintf("| %s | %d～%d岁 | %s年 |\n",
			df.SixtyCycle, df.StartAge, df.EndAge, df.StartYear))
	}
	b.WriteString("\n")

	// 流年（全部）
	b.WriteString("## 流年\n\n")
	b.WriteString("| 公历年 | 干支 | 年龄 | 十神 |\n")
	b.WriteString("|---|---|---|---|\n")
	for _, a := range r.AnnualFortunes {
		b.WriteString(fmt.Sprintf("| %d | %s | %d岁 | %s |\n",
			a.CalYear, a.SixtyCycle, a.Age, a.TenStar))
	}

	return b.String()
}
