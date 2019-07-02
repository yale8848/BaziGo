package SiZhu

// 干支

import (
	. "github.com/yale8848/BaziGo/Common"
)

// 获得八字年的干支，0-59 对应 甲子到癸亥
func GetGanZhiFromYear(nYear int) int {
	if nYear > 0 {
		return (nYear - 4) % 60
	} else { // 需要独立判断公元前的原因是没有公元 0 年
		return (nYear - 3) % 60
	}
}

// 获得八字年的干支，0-59 对应 甲子到癸亥, 加字符串
func GetGanZhiFromYear2(pGanZhi *TGanZhi, nYear int) TGanZhi {
	pGanZhi.Value = GetGanZhiFromYear(nYear)
	return *pGanZhi
}

// 获得某公/农历年的天干，0-9 对应 甲到癸
func GetGanFromYear(nYear int) int {
	if nYear > 0 {
		return (nYear - 4) % 10
	} else { // 需要独立判断公元前的原因是没有公元 0 年
		return (nYear - 3) % 10
	}
}

// 获得某公/农历年的地支，0-11 对应 子到亥
func GetZhiFromYear(nYear int) int {
	if nYear > 0 {
		return (nYear - 4) % 12
	} else { // 需要独立判断公元前的原因是没有公元 0 年
		return (nYear - 3) % 12
	}
}

// 将干支拆分成天干地支，0-59 转换成 0-9 0-11
func ExtractGanZhi(nGanZhi int) (int, int) {
	nGanZhi = nGanZhi % 60
	if nGanZhi < 0 {
		nGanZhi += 60
	}
	return (nGanZhi % 10), (nGanZhi % 12)
}

// 将干支拆分成天干地支，0-59 转换成 0-9 0-11, 带字符串
func ExtractGanZhi2(pGanZhi *TGanZhi, pGan *TGan, pZhi *TZhi) (TGan, TZhi) {
	pGan.Value, pZhi.Value = ExtractGanZhi(pGanZhi.Value)
	return *pGan, *pZhi
}

// 将天干地支组合成干支，0-9 0-11 转换成 0-59
func CombineGanZhi(nGan, nZhi int) int {
	if (nGan <= 9) && (nZhi <= 11) {
		for i := 0; i <= 6; i++ {
			if (i*10+nGan)%12 == nZhi {
				return i*10 + nGan
			}
		}
	}
	return -1
}

// 将天干地支组合成干支，0-9 0-11 转换成 0-59 带字符串
func CombineGanZhi2(pGanZhi *TGanZhi, pGan *TGan, pZhi *TZhi) TGanZhi {
	pGanZhi.Value = CombineGanZhi(pGan.Value, pZhi.Value)
	return *pGanZhi
}

// 获得某干的五行，0-4 对应 金木水火土
// 甲乙为木，丙丁为火，戊己为土，庚辛为金，壬癸为水，
func Get5XingFromGan(nGan int) int {
	switch nGan {
	case 0, 1:
		return 1 // 甲 阳木 乙 阴木
	case 2, 3:
		return 3 // 丙 阳火 丁 阴火
	case 4, 5:
		return 4 // 戊 阳土 己 阴土
	case 6, 7:
		return 0 // 庚 阳金 辛 阴金
	case 8, 9:
		return 2 // 壬 阳水 癸 阴水
	}
	return -1
}

// 获得某干的五行，0-4 对应 金木水火土, 加字符串
// 甲乙为木，丙丁为火，戊己为土，庚辛为金，壬癸为水，
func Get5XingFromGan2(pWuXing *TWuXing, nGan int) TWuXing {
	pWuXing.Value = Get5XingFromGan(nGan)
	return *pWuXing
}

// 获得某支的五行，0-4 对应 金木水火土
func Get5XingFromZhi(nZhi int) int {
	switch nZhi {
	case 8, 9:
		return 0
	case 2, 3:
		return 1
	case 0, 11:
		return 2
	case 5, 6:
		return 3
	case 1, 4, 7, 10:
		return 4
	}
	return -1
}

// 获得某支的五行，0-4 对应 金木水火土, 加字符串
func Get5XingFromZhi2(pWuXing *TWuXing, nZhi int) TWuXing {
	pWuXing.Value = Get5XingFromZhi(nZhi)
	return *pWuXing
}
