package gm

import (
	"encoding/json"

	"github.com/akakou/zk-ban-system/utils"
)

func (gm *GroupManager[T]) RevocationList(before int64) ([]byte, error) {
	after := utils.Today()

	rl, err := gm.QueryRL(before, after)
	if err != nil {
		return nil, err
	}

	res, err := json.Marshal(rl)

	return res, err
}
