package core

import (
	zkbanw "github.com/akakou/zk-ban/witness"
)

type RevocationList struct {
	List    *zkbanw.RevocationList
	KeyName string
}

type RevocationListSize struct {
	NymsNumberPerPeriod []int
}

func (size *RevocationListSize) Distance(weight *RevocationListSizeWeight) float64 {
	var deltaL = 0.0
	for _, nymNum := range size.NymsNumberPerPeriod {
		deltaL += float64(nymNum)
	}

	tDash := float64(len(size.NymsNumberPerPeriod))

	return deltaL*weight.NymsNumberPerPeriod + tDash*weight.PeriodNumber
}

type RevocationListSizeWeight struct {
	NymsNumberPerPeriod float64
	PeriodNumber        float64
}

var RevocationListSizeWeightSetting = &RevocationListSizeWeight{
	NymsNumberPerPeriod: 3,
	PeriodNumber:        1,
}
