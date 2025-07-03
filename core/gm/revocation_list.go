package gm

import (
	"encoding/json"

	corecore "github.com/akakou/zk-ban-system/core/core"
	"github.com/akakou/zk-ban-system/utils"
)

func (gm *GroupManager[T]) RevocationList(before int64) ([]byte, error) {
	after := utils.Today()

	rl, key, err := gm.QueryRLAndKey(before, after)
	if err != nil {
		return nil, err
	}

	rlWithSize := corecore.RevocationList{
		List: rl,
		Size: key.Size,
	}

	res, err := json.Marshal(rlWithSize)

	return res, err
}
