package signer

import (
	neturl "net/url"

	"github.com/akakou/zk-ban-system/client/utils"
	coresigner "github.com/akakou/zk-ban-system/core/signer"
)

func Sign(message []byte, count int64, signer, gpk []byte, url string) ([]byte, error) {
	signature, err := coresigner.Sign(message, count, signer, gpk)
	if err != nil {
		return nil, err
	}

	signatureStr := string(signature)

	parsed, err := neturl.Parse(url)
	if err != nil {
		return nil, err
	}

	query := parsed.Query()
	query.Add("signature", signatureStr)
	parsed.RawQuery = query.Encode()

	return utils.FetchBinaryWithGET(parsed.String())
}
