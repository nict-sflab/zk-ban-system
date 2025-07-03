package dump

import (
	"embed"
	"fmt"

	"github.com/akakou/zk-ban-system/core/core"
)

//go:embed join_prover.key.json
var JoinProverKey []byte

//go:embed sign_prover.key.json
var SignProverKey []byte

//go:embed all:update_prover-*-*.key.json
var UpdateProverKey embed.FS

//go:embed join_verifier.key.json
var JoinVerifierKey []byte

//go:embed sign_verifier.key.json
var SignVerifierKey []byte

//go:embed all:update_verifier-*-*.key.json
var UpdateVerifierKey embed.FS

func LoadKeyWithRL(format string, fs embed.FS) (core.SnarkKeys, error) {
	files, err := fs.ReadDir(".")
	if err != nil {
		return nil, err
	}

	keys := make(core.SnarkKeys)
	for _, f := range files {
		name := f.Name()

		index := core.RevocationListSize{
			NymsNumberPerSession: 0,
			SessionNumber:        0,
		}

		_, err = fmt.Sscanf(name, format, &index.NymsNumberPerSession, &index.SessionNumber)
		if err != nil {
			return nil, err
		}

		buf, err := fs.ReadFile(name)
		if err != nil {
			return nil, err
		}

		k := core.SnarkKey(buf)
		keys[index] = &k
	}

	return keys, nil

}

func LoadUserUpdateKey() (core.SnarkKeys, error) {
	return LoadKeyWithRL(UpdateProverKeyFileNameFormat, UpdateProverKey)
}

func LoadVerifierKey() []byte {
	return SignVerifierKey
}

func LoadGroupManagerUpdateKey() (core.SnarkKeys, error) {
	return LoadKeyWithRL(UpdateVerifierKeyFileNameFormat, UpdateVerifierKey)
}
