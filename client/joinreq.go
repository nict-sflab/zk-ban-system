package client

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/akakou/zk-ban-system/core"
	"github.com/akakou/zk-ban-system/utils"
	"github.com/akakou/zk-ban/highlevel"
)

func RequestJoin(idToken []byte, prover []byte, url string) (*highlevel.HighLevelSigner, error) {
	period := utils.Today()

	proverObj := highlevel.HighLevelSnarkProver{}
	err := json.Unmarshal(prover, &proverObj)
	if err != nil {
		return nil, err
	}

	proof, reqObj, err := highlevel.JoinRequest(period, &proverObj)
	if err != nil {
		return nil, err
	}

	req := core.JoinRequest[[]byte]{
		Period:        reqObj.Period,
		UserPublicKey: reqObj.UserPublicKey,
		Proof:         proof,
		Option:        idToken,
	}

	reqBytes, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	res, err := http.Post(url, "application/json", bytes.NewBuffer(reqBytes))
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(res.Body)
	if err != nil {
		return nil, err
	}

	signer := highlevel.HighLevelSigner{
		Credential:    buf.Bytes(),
		Secret:        reqObj.UserSecretKey,
		UserPublicKey: reqObj.UserPublicKey,
		Period:        period,
	}

	return &signer, nil
}
