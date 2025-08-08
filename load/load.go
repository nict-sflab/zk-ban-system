package load

import (
	"embed"
	"encoding/json"
	"fmt"

	"github.com/akakou/zk-ban-system/core/core"
	"github.com/akakou/zk-ban-system/dump"
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

func LoadKeyWithRL[T any](format string, fs embed.FS) (map[core.RevocationListSize]T, error) {
	files, err := fs.ReadDir(".")
	if err != nil {
		return nil, err
	}

	keys := new(map[core.RevocationListSize]T)
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

		var key *T
		err = json.Unmarshal(buf, key)

		if err != nil {
			return nil, err
		}

		(*keys)[index] = *key
	}

	return *keys, nil

}

func LoadUserUpdateKey() (core.ProvingKeys, error) {
	return LoadKeyWithRL[core.SnarkProver](dump.UpdateProverKeyFileNameFormat, UpdateProverKey)
}

func LoadGroupManagerUpdateKey() (core.VerifyingKeys, error) {
	return LoadKeyWithRL[core.VerifyingKey](dump.UpdateVerifierKeyFileNameFormat, UpdateVerifierKey)
}
