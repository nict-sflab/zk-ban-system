package verifier

import (
	coreverifier "github.com/akakou/zk-ban-system/core/verifier"
)

type VerifierServer struct {
	Verifier *coreverifier.Verifier
}
