package keys

import (
	"embed"

	"github.com/akakou/zk-ban/load"
	"github.com/akakou/zk-ban/snark"
)

//go:embed join_prover.key.json
var JoinProverKey []byte

//go:embed sign_prover.key.json
var SignProverKey []byte

//go:embed all:update_prover-*.key.json
var UpdateProverKey embed.FS

//go:embed join_verifier.key.json
var JoinVerifierKey []byte

//go:embed sign_verifier.key.json
var SignVerifierKey []byte

//go:embed all:update_verifier-*.key.json
var UpdateVerifierKey embed.FS

func LoadGroupManagerUpdateKey(fs embed.FS) ([]*snark.SizedSnarkVerifier, error) {
	return load.LoadGroupManagerUpdateKey(fs)
}

func LoadUserUpdateKey(fs embed.FS) ([]*snark.SnarkProver, error) {
	return load.LoadUserUpdateKey(fs)
}
