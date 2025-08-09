package gm_test

import (
	"testing"

	"github.com/akakou/zk-ban-system/core/gm"
	"github.com/akakou/zk-ban-system/utils"
	"github.com/akakou/zk-ban/witness"
	"github.com/stretchr/testify/assert"
)

var period = 1

func today() int64 {
	return int64(period)
}

func passDay() {
	period += 1
}

func TestQuery(t *testing.T) {
	gsk, _, _ := witness.RandomGroupKeyPair()
	g, err := gm.Default[string](gsk.Bytes(), &gm.DBConfig{
		Type:   "sqlite3",
		Config: "file::memory:?cache=shared&_fk=1",
	})
	assert.NoError(t, err)

	utils.Period = today

	for range 100 {
		g.DB.Client.Revocation.Create().
			SetCount(0).
			SetNym([]byte("aaa")).
			SetRevokedPeriod(3).
			SetSignedPeriod(0).
			SaveX(*g.DB.Ctx)
	}

	for range 100 {
		g.DB.Client.Revocation.Create().
			SetCount(0).
			SetNym([]byte("gggg")).
			SetRevokedPeriod(3).
			SetSignedPeriod(1).
			SaveX(*g.DB.Ctx)
	}
	passDay()

	rl, err := g.QueryRL(2)
	assert.NoError(t, err)
	assert.Equal(t, 100, rl.Size.NymsNumberPerSession)
	assert.Equal(t, 10, rl.Size.SessionNumber)

	t.Fatalf("%v", rl.List)
}
