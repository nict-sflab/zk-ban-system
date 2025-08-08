package load

import (
	"embed"
	"encoding/json"
	"fmt"

	"github.com/akakou/snark-utils/encode"
	"github.com/akakou/zk-ban-system/core/core"
	"github.com/akakou/zk-ban-system/dump"
	"github.com/akakou/zk-ban/snark"
	"github.com/consensys/gnark/backend/groth16"
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

func LoadKeyWithRL[T any](format string, fs embed.FS, fun func([]byte) (T, error)) (map[core.RevocationListSize]T, error) {
	files, err := fs.ReadDir(".")
	if err != nil {
		return nil, err
	}

	keys := make(map[core.RevocationListSize]T)
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

		key, err := fun(buf)
		keys[index] = key
	}

	return keys, nil

}

func DocodeProver(buf []byte) (*snark.SnarkProver, error) {
	prover := encode.HighLevelSnarkProver{}

	err := json.Unmarshal(buf, &prover)
	if err != nil {
		return nil, err
	}

	prover2, err := prover.ToSnarkProver()
	if err != nil {
		return nil, err
	}

	res := snark.SnarkProver{
		ConstraintSystem: prover2.ConstraintSystem,
		ProveKey:         prover2.ProveKey,
	}

	return &res, nil

}

func LoadUserUpdateKey() (map[core.RevocationListSize]*snark.SnarkProver, error) {
	return LoadKeyWithRL(dump.UpdateProverKeyFileNameFormat, UpdateProverKey, DocodeProver)
}

func LoadGroupManagerUpdateKey() (map[core.RevocationListSize]groth16.VerifyingKey, error) {
	keys, err := LoadKeyWithRL(dump.UpdateVerifierKeyFileNameFormat, UpdateVerifierKey, encode.DecodeVerifierKey)
	if err != nil {
		return nil, err
	}
	return keys, nil
}
