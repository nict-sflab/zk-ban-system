package gm

import (
	"encoding/json"

	corecore "github.com/akakou/zk-ban-system/core/core"
	"github.com/akakou/zk-ban-system/utils"
)

func SelectVerifierKeyFromSize(size *corecore.RevocationListSize) *PreparableSnarkVerifierKey {
	verifierKey := PreparableSnarkVerifierKeys[0]
	distance := verifierKey.Size.Distance(corecore.RevocationListSizeWeightSetting)

	for _, vk := range PreparableSnarkVerifierKeys[1:] {
		isFitSize1 := vk.Size.NymsNumberPerSession <= size.SessionNumber
		isFitSize2 := vk.Size.SessionNumber <= size.SessionNumber

		if isFitSize1 || isFitSize2 {
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

func (gm *GroupManager[T]) RevocationList(before int64) (string, error) {
	after := utils.Today()

	rl, key, err := gm.QueryRLAndKey(before, after)
	if err != nil {
		return "", err
	}

	rlWithSize := corecore.RevocationList{
		List: rl,
		Size: key.Size,
	}

	res, err := json.Marshal(rlWithSize)

	return string(res), err
}
