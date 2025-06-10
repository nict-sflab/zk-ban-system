package gm

type GroupManager[T any] struct {
	GroupSecretKey  []byte
	GroupPublicKey  []byte
	JoinVerifyKey   []byte
	UpdateVerifyKey []byte
	DB              *DB
}
