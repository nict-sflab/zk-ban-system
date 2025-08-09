package core

import (
	"math"

	zkbanw "github.com/akakou/zk-ban/witness"
)

type RevocationList struct {
	List *zkbanw.RevocationList
	Size *RevocationListSize
}

type RevocationListSize struct {
	NymsNumberPerSession int
	SessionNumber        int
}

func (size *RevocationListSize) Distance(weight *RevocationListSizeWeight) float64 {
	weightedNymNumber := float64(size.NymsNumberPerSession) * weight.NymsNumberPerSession
	weightedSessionNumber := float64(size.SessionNumber) * weight.SessionNumber

	return math.Pow(weightedNymNumber, 2.0) + math.Pow(weightedSessionNumber, 2.0)
}

type RevocationListSizeWeight struct {
	NymsNumberPerSession float64
	SessionNumber        float64
}

var RevocationListSizeWeightSetting = &RevocationListSizeWeight{
	NymsNumberPerSession: 3,
	SessionNumber:        1,
}
