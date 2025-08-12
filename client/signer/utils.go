package signer

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/akakou/zk-ban-system/client/utils"
	zkbanw "github.com/akakou/zk-ban/witness"
)

func FetchGroupPublicKey(url string) ([]byte, error) {
	return utils.FetchBase64WithGET(url)
}

func FetchRevocationList(signer []byte, u string) ([]byte, error) {
	signerObj := zkbanw.Signer{}
	err := json.Unmarshal(signer, &signerObj)
	if err != nil {
		return nil, err
	}

	uu, err := url.Parse(u)
	if err != nil {
		return nil, err
	}

	query := uu.Query()
	query.Add("before", fmt.Sprintf("%d", signerObj.Period))
	uu.RawQuery = query.Encode()

	return utils.FetchBinaryWithGET(uu.String())
}
