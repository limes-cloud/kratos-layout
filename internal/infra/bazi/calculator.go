package bazi

import (
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/6tail/tyme4go/tyme"
)

// BaziResult 八字计算结果
type BaziResult struct {
	Name     string
	Gender   string
	Zodiac   string
	LunarDate string
	SolarTerm string

	YearPillar  PillarResult
	MonthPillar PillarResult
	DayPillar   PillarResult
	HourPillar  PillarResult

	FetalOrigin string
	OwnSign     string
	BodySign    string
	Strength    string

	ChildLimit    string
	DecadeFortunes []DecadeFortuneResult
	AnnualFortunes []AnnualFortuneResult
	Wuxing        WuxingScore
}

// WuxingScore 五行占比
type WuxingScore struct {
	Wood  float64 // 木
	Fire  float64 // 火
	Earth float64 // 土
	Metal float64 // 金
	Water float64 // 水
}

// PillarResult 柱信息
type PillarResult struct {
	Stem        string
	Branch      string
	Sound       string
	TenStar     string
	Terrain     string // 地势：日主对本柱地支
	SelfSeat    string // 自坐：本柱天干对本柱地支
	HiddenStems string
	Xukong      string
	Shensha     string
}

// DecadeFortuneResult 大运信息
type DecadeFortuneResult struct {
	SixtyCycle string
	StartAge   int32
	EndAge     int32
	StartYear  string
}

// AnnualFortuneResult 流年信息
type AnnualFortuneResult struct {
	CalYear    int32
	SixtyCycle string
	Age        int32
	TenStar    string
}

// elementStrengthTable 旺相休囚死：[月令地支index][日干五行index] = 状态
// 地支: 子1丑2寅3卯4辰5巳6午7未8申9酉10戌11亥0（按地支索引 0=子...11=亥）
// 五行: 木0木1火2火3土4土5金6金7水8水9（天干索引对应五行：甲乙=木，丙丁=火，戊己=土，庚辛=金，壬癸=水）
// 旺=4 相=3 休=1 囚=2 死=0
var strengthNames = []string{"死", "囚", "休", "相", "旺"}

// 月令地支(0-11: 子丑寅卯辰巳午未申酉戌亥) -> 当令五行index(0=木1=火2=土3=金4=水)
var monthBranchElement = []int{4, 2, 0, 0, 2, 1, 1, 2, 3, 3, 2, 4}

// 生克关系确定旺相休囚死
func getDayStrength(dayHSIdx int, monthBranchIdx int) string {
	// 天干五行: 甲乙=木(0), 丙丁=火(1), 戊己=土(2), 庚辛=金(3), 壬癸=水(4)
	dayElem := dayHSIdx / 2
	// 月令当令五行
	monthElem := monthBranchElement[monthBranchIdx]
	// 五行生克：木(0)->火(1)->土(2)->金(3)->水(4)->木(0)
	// 同 = 旺；生我者 = 相；我生者 = 休；克我者 = 囚；我克者 = 死
	var strength int
	switch {
	case monthElem == dayElem: // 同行旺
		strength = 4
	case (monthElem+1)%5 == dayElem: // 月生日相
		strength = 3
	case (dayElem+1)%5 == monthElem: // 日生月休
		strength = 2
	case (monthElem+3)%5 == dayElem: // 月克日囚
		strength = 1
	default: // 日克月死
		strength = 0
	}
	return strengthNames[strength]
}

// Calculate 根据出生信息计算八字命盘
// isLunar=true 时 year/month/day 为农历，lunarLeap 表示是否闰月
func Calculate(name, gender string, year, month, day, hour, minute int, location string, isLunar, lunarLeap bool) (*BaziResult, error) {
	// 农历转公历
	solarYear, solarMonth, solarDay := year, month, day
	if isLunar {
		lunarMonthNum := month
		if lunarLeap {
			lunarMonthNum = -month // tyme4go 用负数表示闰月
		}
		ld, err := tyme.LunarDay{}.FromYmd(year, lunarMonthNum, day)
		if err != nil {
			return nil, fmt.Errorf("农历日期无效: %w", err)
		}
		sd := ld.GetSolarDay()
		solarYear, solarMonth, solarDay = sd.GetYear(), sd.GetMonth(), sd.GetDay()
	}

	// 处理真太阳时修正
	birthTime := time.Date(solarYear, time.Month(solarMonth), solarDay, hour, minute, 0, 0, time.Local)
	calcYear, calcMonth, calcDay, calcHour, calcMinute := solarYear, solarMonth, solarDay, hour, minute

	if location != "" {
		loc, err := LookupCity(location)
		if err != nil {
			return nil, err
		}
		corrected := ToTrueSolarTime(birthTime, loc.Lon, loc.TZ)
		calcYear = corrected.Year()
		calcMonth = int(corrected.Month())
		calcDay = corrected.Day()
		calcHour = corrected.Hour()
		calcMinute = corrected.Minute()
	}

	st, err := tyme.SolarTime{}.FromYmdHms(calcYear, calcMonth, calcDay, calcHour, calcMinute, 0)
	if err != nil {
		return nil, fmt.Errorf("invalid birth time: %w", err)
	}

	sctHour := st.GetSixtyCycleHour()
	ec := sctHour.GetEightChar()

	yPillar := ec.GetYear()
	mPillar := ec.GetMonth()
	dPillar := ec.GetDay()
	hPillar := ec.GetHour()

	// 日主天干
	dayStem := dPillar.GetHeavenStem()

	// 神煞：按命理规则分柱计算
	shenshaByPillar := calcMingliShensha(ec)

	// 性别
	var gender_ tyme.Gender
	if gender == "M" {
		gender_ = tyme.MAN
	} else {
		gender_ = tyme.WOMAN
	}

	// 童限
	cl := tyme.ChildLimit{}.FromSolarTime(*st, gender_)
	childLimitStr := buildChildLimitStr(cl)

	// 大运：从起运大运（index=0）开始，取10步
	decades := make([]DecadeFortuneResult, 0, 10)
	df := cl.GetStartDecadeFortune()
	for i := 0; i < 10; i++ {
		decades = append(decades, DecadeFortuneResult{
			SixtyCycle: df.GetSixtyCycle().GetName(),
			StartAge:   int32(df.GetStartAge()),
			EndAge:     int32(df.GetEndAge()),
			StartYear:  df.GetStartSixtyCycleYear().GetSixtyCycle().GetName(),
		})
		df = df.Next(1)
	}

	// 流年：从出生年起70年
	annuals := make([]AnnualFortuneResult, 0, 70)
	birthYear := year
	for i := 0; i < 70; i++ {
		y := birthYear + i
		cy, err2 := tyme.SixtyCycleYear{}.FromYear(y)
		if err2 != nil {
			continue
		}
		tenStar := dayStem.GetTenStar(cy.GetSixtyCycle().GetHeavenStem()).GetName()
		annuals = append(annuals, AnnualFortuneResult{
			CalYear:    int32(y),
			SixtyCycle: cy.GetSixtyCycle().GetName(),
			Age:        int32(y - birthYear + 1),
			TenStar:    tenStar,
		})
	}

	// 农历日期 & 生肖
	ld := st.GetSolarDay().GetLunarDay()
	lunarDateStr := fmt.Sprintf("%d年%s%s", year, ld.GetLunarMonth().String(), ld.GetName())
	zodiac := yPillar.GetEarthBranch().GetZodiac().GetName()

	// 节气信息
	solarTermStr := buildSolarTermStr(*st)

	// 胎元命宫身宫
	fetalOrigin := ec.GetFetalOrigin().GetName()
	ownSign := ec.GetOwnSign().GetName()
	bodySign := ec.GetBodySign().GetName()

	// 旺相休囚死
	strength := getDayStrength(dayStem.GetIndex(), mPillar.GetEarthBranch().GetIndex())

	result := &BaziResult{
		Name:           name,
		Gender:         gender,
		Zodiac:         zodiac,
		LunarDate:      lunarDateStr,
		SolarTerm:      solarTermStr,
		YearPillar:     buildPillar(yPillar, dayStem, shenshaByPillar[0]),
		MonthPillar:    buildPillar(mPillar, dayStem, shenshaByPillar[1]),
		DayPillar:      buildDayPillar(dPillar, dayStem, shenshaByPillar[2]),
		HourPillar:     buildPillar(hPillar, dayStem, shenshaByPillar[3]),
		FetalOrigin:    fetalOrigin,
		OwnSign:        ownSign,
		BodySign:       bodySign,
		Strength:       strength,
		ChildLimit:     childLimitStr,
		DecadeFortunes: decades,
		AnnualFortunes: annuals,
		Wuxing:         calcWuxing(ec),
	}
	return result, nil
}

// calcWuxing 计算八字五行占比（天干+地支+藏干共15个字）
// 五行索引：木0，火1，土2，金3，水4（天干：甲乙=木，丙丁=火，戊己=土，庚辛=金，壬癸=水）
// 地支五行：子=水，丑=土，寅=木，卯=木，辰=土，巳=火，午=火，未=土，申=金，酉=金，戌=土，亥=水
func calcWuxing(ec tyme.EightChar) WuxingScore {
	// 天干五行：甲(0)乙(1)=木，丙(2)丁(3)=火，戊(4)己(5)=土，庚(6)辛(7)=金，壬(8)癸(9)=水
	stemElement := []int{0, 0, 1, 1, 2, 2, 3, 3, 4, 4}
	// 地支五行：子(0)=水，丑(1)=土，寅(2)=木，卯(3)=木，辰(4)=土，巳(5)=火，午(6)=火，未(7)=土，申(8)=金，酉(9)=金，戌(10)=土，亥(11)=水
	branchElement := []int{4, 2, 0, 0, 2, 1, 1, 2, 3, 3, 2, 4}

	counts := [5]int{}
	pillars := []tyme.SixtyCycle{ec.GetYear(), ec.GetMonth(), ec.GetDay(), ec.GetHour()}
	for _, p := range pillars {
		counts[stemElement[p.GetHeavenStem().GetIndex()]]++
		counts[branchElement[p.GetEarthBranch().GetIndex()]]++
		for _, h := range p.GetEarthBranch().GetHideHeavenStems() {
			counts[stemElement[h.GetHeavenStem().GetIndex()]]++
		}
	}
	total := 0
	for _, c := range counts {
		total += c
	}
	if total == 0 {
		return WuxingScore{}
	}
	t := float64(total)
	return WuxingScore{
		Wood:  math.Round(float64(counts[0])/t*100) / 100,
		Fire:  math.Round(float64(counts[1])/t*100) / 100,
		Earth: math.Round(float64(counts[2])/t*100) / 100,
		Metal: math.Round(float64(counts[3])/t*100) / 100,
		Water: math.Round(float64(counts[4])/t*100) / 100,
	}
}

func buildPillar(sc tyme.SixtyCycle, dayStem tyme.HeavenStem, shensha []string) PillarResult {
	stem := sc.GetHeavenStem()
	branch := sc.GetEarthBranch()
	xukong := buildXukong(sc)
	return PillarResult{
		Stem:        stem.GetName(),
		Branch:      branch.GetName(),
		Sound:       sc.GetSound().GetName(),
		TenStar:     dayStem.GetTenStar(stem).GetName(),
		Terrain:     dayStem.GetTerrain(branch).GetName(),
		SelfSeat:    stem.GetTerrain(branch).GetName(),
		HiddenStems: buildHiddenStems(branch, dayStem),
		Xukong:      xukong,
		Shensha:     strings.Join(shensha, "，"),
	}
}

func buildDayPillar(sc tyme.SixtyCycle, dayStem tyme.HeavenStem, shensha []string) PillarResult {
	stem := sc.GetHeavenStem()
	branch := sc.GetEarthBranch()
	xukong := buildXukong(sc)
	return PillarResult{
		Stem:        stem.GetName(),
		Branch:      branch.GetName(),
		Sound:       sc.GetSound().GetName(),
		TenStar:     "日元",
		Terrain:     dayStem.GetTerrain(branch).GetName(),
		SelfSeat:    stem.GetTerrain(branch).GetName(),
		HiddenStems: buildHiddenStems(branch, dayStem),
		Xukong:      xukong,
		Shensha:     strings.Join(shensha, "，"),
	}
}

func buildXukong(sc tyme.SixtyCycle) string {
	extras := sc.GetExtraEarthBranches()
	names := make([]string, len(extras))
	for i, eb := range extras {
		names[i] = eb.GetName()
	}
	return strings.Join(names, "")
}

func buildHiddenStems(branch tyme.EarthBranch, dayStem tyme.HeavenStem) string {
	hides := branch.GetHideHeavenStems()
	stems := make([]string, len(hides))
	stars := make([]string, len(hides))
	for i, h := range hides {
		stems[i] = h.GetHeavenStem().GetName()
		stars[i] = dayStem.GetTenStar(h.GetHeavenStem()).GetName()
	}
	return strings.Join(stems, "") + " " + strings.Join(stars, "")
}

// calcMingliShensha 计算命理神煞（主流寄宫法：神煞标注在起例地支所在柱）
func calcMingliShensha(ec tyme.EightChar) [4][]string {
	y := ec.GetYear()
	m := ec.GetMonth()
	d := ec.GetDay()
	h := ec.GetHour()

	yBr := y.GetEarthBranch().GetIndex()
	dBr := d.GetEarthBranch().GetIndex()
	dSt := d.GetHeavenStem().GetIndex()
	ySt := y.GetHeavenStem().GetIndex()
	mBr := m.GetEarthBranch().GetIndex()

	pillars := [4]tyme.SixtyCycle{y, m, d, h}
	result := [4][]string{{}, {}, {}, {}}

	add := func(idx int, name string) { result[idx] = append(result[idx], name) }

	hitBr := func(target int, name string) {
		for i, p := range pillars {
			if p.GetEarthBranch().GetIndex() == target {
				add(i, name)
			}
		}
	}
	hitSt := func(target int, name string) {
		for i, p := range pillars {
			if p.GetHeavenStem().GetIndex() == target {
				add(i, name)
			}
		}
	}

	// ── 年干见支（startPillar: year） ──
	// 太极贵人：年干/2分组，rule:[[0,6],[3,9],[4,10,1,7],[2,11],[5,8]]
	// 甲乙(0)→子午, 丙丁(1)→卯酉, 戊己庚辛(2)→辰戌丑未, 壬(3)→寅亥, 癸(4)→巳申
	taijiBrs := [][]int{{0, 6}, {3, 9}, {4, 10, 1, 7}, {2, 11}, {5, 8}}
	for _, br := range taijiBrs[ySt>>1] {
		hitBr(br, "太极贵人")
	}

	// ── 年干、日干见支（startPillar: year, day） ──
	for _, st := range []int{ySt, dSt} {
		// 天乙贵人 rule(10项): [[1,7],[0,8],[11,9],[11,9],[1,7],[0,8],[1,7],[2,6],[3,5],[3,5]]
		tianyiRule := [][]int{{1, 7}, {0, 8}, {11, 9}, {11, 9}, {1, 7}, {0, 8}, {1, 7}, {2, 6}, {3, 5}, {3, 5}}
		for _, br := range tianyiRule[st] {
			hitBr(br, "天乙贵人")
		}
		// 文昌贵人 rule: [5,6,8,9,8,9,11,0,2,3]
		hitBr([]int{5, 6, 8, 9, 8, 9, 11, 0, 2, 3}[st], "文昌贵人")
		// 国印贵人 rule: [10,11,1,2,1,2,4,5,7,8]
		hitBr([]int{10, 11, 1, 2, 1, 2, 4, 5, 7, 8}[st], "国印贵人")
		// 天厨贵人 rule: [5,6,5,6,8,9,11,0,2,3]
		hitBr([]int{5, 6, 5, 6, 8, 9, 11, 0, 2, 3}[st], "天厨贵人")
		// 学堂 rule(9项,癸无): [11,11,2,2,8,8,5,8,8]
		if st < 9 {
			hitBr([]int{11, 11, 2, 2, 8, 8, 5, 8, 8}[st], "学堂")
		}
	}

	// ── 日干见支（startPillar: day） ──
	hitBr([]int{3, 2, 6, 5, 6, 5, 9, 8, 0, 11}[dSt], "羊刃")
	hitBr([]int{4, 5, 7, 8, 7, 8, 10, 11, 1, 2}[dSt], "金舆")
	hitBr([]int{2, 3, 5, 6, 5, 6, 8, 9, 11, 0}[dSt], "绿神")
	hitBr([]int{11, 11, 2, 9, 2, 9, 5, 0, 8, 3}[dSt], "垣城")

	// 年官：年干见时支 rule: [7,4,5,9,10,3,11,8,2,8]
	if h.GetEarthBranch().GetIndex() == []int{7, 4, 5, 9, 10, 3, 11, 8, 2, 8}[ySt] {
		add(3, "天官")
	}

	// ── 月支见干支 ──
	// 天德 rule: ['b5','s6','s3','b8','s8','s7','b11','s0','s9','b2','s2','s1']
	tiandeRule := [][2]int{{1, 5}, {0, 6}, {0, 3}, {1, 8}, {0, 8}, {0, 7}, {1, 11}, {0, 0}, {0, 9}, {1, 2}, {0, 2}, {0, 1}}
	// [0=stem/1=branch, value]
	if r := tiandeRule[mBr]; r[0] == 0 {
		hitSt(r[1], "天德贵人")
	} else {
		hitBr(r[1], "天德贵人")
	}
	// 天德合 rule: ['b8','s1','s8','b5','s3','s2','b2','s5','s4','b11','s7','s6']
	tiandehRule := [][2]int{{1, 8}, {0, 1}, {0, 8}, {1, 5}, {0, 3}, {0, 2}, {1, 2}, {0, 5}, {0, 4}, {1, 11}, {0, 7}, {0, 6}}
	if r := tiandehRule[mBr]; r[0] == 0 {
		hitSt(r[1], "天德合")
	} else {
		hitBr(r[1], "天德合")
	}
	// 月德 rule: bV%4→[8,6,2,0]
	hitSt([]int{8, 6, 2, 0}[mBr%4], "月德贵人")
	// 月德合 rule: bV%4→[3,1,7,5]
	hitSt([]int{3, 1, 7, 5}[mBr%4], "月德合")
	// 天医：月支→前一支
	hitBr((mBr+11)%12, "天医")
	// 太白星：月支bV%3→[5,1,9]（午子酉）
	hitBr([]int{5, 1, 9}[mBr%3], "太白星")

	// ── 年支见支 ──
	// 驿马/华盖/将星/劫煞/亡神/桃花/灾煞 lunisolar:年支和日支都起例，bV%4分组
	for _, startBr := range []int{yBr, dBr} {
		g := startBr % 4
		hitBr([]int{2, 11, 8, 5}[g], "驿马")
		hitBr([]int{4, 1, 10, 7}[g], "华盖")
		hitBr([]int{0, 9, 6, 3}[g], "将星")
		hitBr([]int{5, 2, 11, 8}[g], "劫煞")
		hitBr([]int{11, 8, 5, 2}[g], "亡神")
		hitBr([]int{9, 6, 3, 0}[g], "桃花")
		hitBr([]int{6, 3, 0, 9}[g], "灾煞")
	}
	// 孤辰 rule: [2,2,5,5,5,8,8,8,11,11,11,2]
	hitBr([]int{2, 2, 5, 5, 5, 8, 8, 8, 11, 11, 11, 2}[yBr], "孤辰")
	// 寡宿 rule: [10,10,1,1,1,4,4,4,7,7,7,10]
	hitBr([]int{10, 10, 1, 1, 1, 4, 4, 4, 7, 7, 7, 10}[yBr], "寡宿")
	// 丧门：年支+2
	hitBr((yBr+2)%12, "丧门")
	// 吊客：年支+10
	hitBr((yBr+10)%12, "吊客")

	// 天罗地网（年支或日支见支）
	for _, startBr := range []int{yBr, dBr} {
		if startBr == 10 {
			hitBr(11, "天罗")
		} else if startBr == 11 {
			hitBr(10, "天罗")
		}
		if startBr == 4 {
			hitBr(5, "地网")
		} else if startBr == 5 {
			hitBr(4, "地网")
		}
	}

	// ── 日柱特定干支 ──
	dName := d.GetName()
	if dName == "庚辰" || dName == "壬辰" || dName == "庚戌" || dName == "戊戌" {
		add(2, "魁罡")
	}
	if dName == "甲寅" || dName == "丙辰" || dName == "戊辰" || dName == "庚辰" || dName == "壬戌" {
		add(2, "日德")
	}
	if dName == "丁酉" || dName == "癸巳" || dName == "丁亥" || dName == "癸卯" {
		add(2, "日贵")
	}
	badDays := map[string]bool{"甲辰": true, "甲申": true, "甲午": true, "甲寅": true, "乙巳": true,
		"丙申": true, "丁亥": true, "戊戌": true, "己丑": true, "庚辰": true}
	if badDays[dName] {
		add(2, "十恶大败")
	}

	return result
}

func buildChildLimitStr(cl tyme.ChildLimit) string {
	years := cl.GetYearCount()
	months := cl.GetMonthCount()
	days := cl.GetDayCount()

	// 童限时长描述
	var duration string
	if years > 0 {
		duration = fmt.Sprintf("%d年%d月%d日", years, months, days)
	} else if months > 0 {
		duration = fmt.Sprintf("%d月%d日", months, days)
	} else {
		duration = fmt.Sprintf("%d日", days)
	}

	// 起运时刻
	endTime := cl.GetEndTime()
	endSCY := cl.GetEndSixtyCycleYear()
	// 干支年名
	sixtyCycleName := endSCY.GetSixtyCycle().GetName()

	return fmt.Sprintf("出生后%s起大运，每逢%s年%d月%d日前后交运",
		duration, sixtyCycleName, endTime.GetMonth(), endTime.GetDay())
}

func buildSolarTermStr(st tyme.SolarTime) string {
	// 当前节气（之前的）
	prevTerm := st.GetTerm()
	prevTermTime := prevTerm.GetJulianDay().GetSolarTime()
	diffSec := st.Subtract(prevTermTime)
	prevDays := diffSec / 86400
	prevHours := (diffSec % 86400) / 3600

	// 下一个节气
	nextTerm := prevTerm.Next(1)
	nextTermTime := nextTerm.GetJulianDay().GetSolarTime()
	diffSec2 := nextTermTime.Subtract(st)
	nextDays := diffSec2 / 86400
	nextHours := (diffSec2 % 86400) / 3600

	return fmt.Sprintf("%s后%d天%d小时，%s前%d天%d小时",
		prevTerm.GetName(), prevDays, prevHours,
		nextTerm.GetName(), nextDays, nextHours)
}
