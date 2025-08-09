package core

import zkban "github.com/akakou/zk-ban"

type UpdateRequest struct {
	Before        int64
	UpdateRequest *zkban.UpdateRequest
}
