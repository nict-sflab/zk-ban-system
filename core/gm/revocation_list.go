package gm

import (
	"encoding/json"
)

func (gm *GroupManager[T]) RevocationList(before int64) ([]byte, error) {
	rl, err := gm.QueryRL(before)
	if err != nil {
		return nil, err
	}

	res, err := json.Marshal(rl)

	return res, err
}
