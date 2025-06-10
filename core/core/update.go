package core

type UpdateRequest struct {
	Before        int64
	After         int64
	UserPublicKey []byte
	Proof         []byte
}
