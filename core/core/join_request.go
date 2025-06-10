package core

type JoinRequest[T any] struct {
	Period        int64
	UserPublicKey []byte
	Proof         []byte
	Option        T
}
