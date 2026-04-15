package codec

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"

	"github.com/akakou/zk-ban/snark"
	"github.com/consensys/gnark/backend/groth16"
)

func init() {
	gob.Register(groth16.NewProof(snark.EcCurve))
}

func Marshal(v any) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(v); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func Unmarshal(data []byte, v any) error {
	dec := gob.NewDecoder(bytes.NewReader(data))
	return dec.Decode(v)
}

func MarshalBase64(v any) (string, error) {
	buf, err := Marshal(v)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(buf), nil
}

func UnmarshalBase64(data string, v any) error {
	buf, err := base64.URLEncoding.DecodeString(data)
	if err != nil {
		return err
	}
	return Unmarshal(buf, v)
}
