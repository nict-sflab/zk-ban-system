package gm

import (
	"math"

	corecore "github.com/akakou/zk-ban-system/core/core"
)

type PreparedSnarkVerifier struct {
	VerifierKey *corecore.SnarkKey
	Prepared    []byte
}

func QueryProperRLWitSize(verifierKeys corecore.SnarkKeys, size *corecore.RevocationListSize) *corecore.RevocationListSize {
	distance := math.MaxFloat64
	var result *corecore.RevocationListSize = nil

	for i := range verifierKeys {
		isNotFitSize1 := i.NymsNumberPerSession < size.NymsNumberPerSession
		isNotFitSize2 := i.SessionNumber < size.SessionNumber

		if isNotFitSize1 || isNotFitSize2 {
			continue
		}

		if verifierKeys == nil {
			result = &i
			continue
		}

		d := i.Distance(corecore.RevocationListSizeWeightSetting)

		if distance > d {
			distance = d
			result = &i
		}
	}

	return result
}
