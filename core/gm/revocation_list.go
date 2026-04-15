package gm

import (
	"github.com/akakou/zk-ban-system/utils/codec"
)

func (gm *GroupManager[T]) RevocationList(before, after int64) ([]byte, error) {
	rl, _, err := gm.QueryRL(before, after)
	if err != nil {
		return nil, err
	}

	res, err := codec.Marshal(rl)

	return res, err
}
