package core

import (
	zkbanw "github.com/akakou/zk-ban/witness"
)

type RevocationList struct {
	List  *zkbanw.RevocationList
	Index int
	// Size *RevocationListSize
}

type RevocationListSize struct {
	NymsNumberPerSession []int
}

func (size *RevocationListSize) Distance(weight *RevocationListSizeWeight) float64 {
	var deltaL = 0.0
	for _, nymNum := range size.NymsNumberPerSession {
		deltaL += float64(nymNum)
	}

	tDash := float64(len(size.NymsNumberPerSession))

	return deltaL*weight.NymsNumberPerSession + tDash*weight.SessionNumber
}

type RevocationListSizeWeight struct {
	NymsNumberPerSession float64
	SessionNumber        float64
}

var RevocationListSizeWeightSetting = &RevocationListSizeWeight{
	NymsNumberPerSession: 3,
	SessionNumber:        1,
}
