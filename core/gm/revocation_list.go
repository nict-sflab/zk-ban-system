package gm

import (
	"math"

	"github.com/akakou/zk-ban/witness"
)

type RevocationListSize struct {
	NymsNumberPerSession int
	SessionNumber        int
}

type RevocationListSizeWeight struct {
	NymsNumberPerSession float64
	SessionNumber        float64
}

var RevocationListSizeWeightSetting = &RevocationListSizeWeight{
	NymsNumberPerSession: 3,
	SessionNumber:        1,
}

func (size *RevocationListSize) Distance(weight *RevocationListSizeWeight) float64 {
	weightedNymNumber := float64(size.NymsNumberPerSession) * weight.NymsNumberPerSession
	weightedSessionNumber := float64(size.SessionNumber) * weight.SessionNumber

	return math.Pow(weightedNymNumber, 2.0) + math.Pow(weightedSessionNumber, 2.0)
}

type RevocationList struct {
	List witness.RevocationList
	Size RevocationListSize
}

func SelectVerifierKeyFromSize(size *RevocationListSize) *PreparableSnarkVerifierKey {
	verifierKey := PreparableSnarkVerifierKeys[0]
	distance := verifierKey.Size.Distance(RevocationListSizeWeightSetting)

	for _, vk := range PreparableSnarkVerifierKeys[1:] {
		isFitSize1 := vk.Size.NymsNumberPerSession <= size.SessionNumber
		isFitSize2 := vk.Size.SessionNumber <= size.SessionNumber

		if isFitSize1 || isFitSize2 {
			continue
		}

		d := vk.Size.Distance(RevocationListSizeWeightSetting)

		if distance > d {
			distance = d
			verifierKey = vk
		}
	}

	return verifierKey
}
