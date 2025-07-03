package gm

import (
	"math"

	corecore "github.com/akakou/zk-ban-system/core/core"
)

type PreparedSnarkVerifier struct {
	VerifierKey *corecore.SnarkKey
	Prepared    []byte
	KeyIndex    corecore.KeyIndex
}

func QueryProperVerifier(verifierKeys []*corecore.SnarkKey, size *corecore.RevocationListSize) *corecore.SnarkKey {
	distance := math.MaxFloat64
	verifierKey := verifierKeys[0]

	for _, vk := range verifierKeys[1:] {
		isNotFitSize1 := vk.Size.NymsNumberPerSession <= size.SessionNumber
		isNotFitSize2 := vk.Size.SessionNumber <= size.SessionNumber

		if isNotFitSize1 || isNotFitSize2 {
			continue
		}

		d := vk.Size.Distance(corecore.RevocationListSizeWeightSetting)

		if distance > d {
			distance = d
			verifierKey = vk
		}
	}

	return verifierKey
}
