package core

import (
	"math"

	"github.com/akakou/zk-ban/highlevel"
)

type RevocationList struct {
	List *highlevel.HighLevelRevocationList
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
