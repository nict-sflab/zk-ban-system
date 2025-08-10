package gm

import (
	"encoding/base64"
	"errors"
	"fmt"

	corecore "github.com/akakou/zk-ban-system/core"
	"github.com/akakou/zk-ban-system/core/gm/ent/updateticket"
	"github.com/akakou/zk-ban/precomputes"
)

var ErrAlreadyIssueCredential = errors.New("already issue credential")

func (gm *GroupManager[T]) UpdateCredential(req *corecore.UpdateRequest, after int64) (string, error) {
	fmt.Printf("period: %v\n", after)
	if after == req.Before {
		return "", errors.New("no update")
	}

	exist, err := gm.DB.Client.UpdateTicket.
		Query().
		Where(updateticket.Ticket(req.UpdateRequest.UpdateTicket.Number.Bytes())).
		Exist(*gm.DB.Ctx)

	if err != nil {
		return "", err
	}

	if exist {
		return "", ErrAlreadyIssueCredential
	}

	index := corecore.KeyIndex{
		First:  after,
		Second: req.Before,
	}

	verifier, hasVerifier := gm.PreparedSnarkVerifiers[index]

	if !hasVerifier {
		rl, err := gm.QueryRL(req.Before)
		if err != nil {
			return "", err
		}

		gk := gm.VerifierKeys[*rl.Size]

		vk, err := precomputes.NewUpdateVerificationKeyBLS12381(gk.VerifyingKey)
		if err != nil {
			return "", err
		}

		prepared, err := vk.PrecomputeVerify(*rl.List, &gm.GroupPublicKey)
		if err != nil {
			return "", err
		}

		verifier = &PreparedSnarkVerifier{
			VerifierKey: vk,
			Prepared:    *prepared,
		}

		gm.PreparedSnarkVerifiers[index] = verifier
	}

	err = verifier.VerifierKey.VerifyPrepared(verifier.Prepared, req.UpdateRequest, after, req.Before)
	if err != nil {
		return "", err
	}

	cred, err := gm.GroupSecretKey.IssueCredential(req.UpdateRequest.PublicKey)
	if err != nil {
		return "", err
	}

	gm.DB.Client.UpdateTicket.Create().
		SetTicket(req.UpdateRequest.UpdateTicket.Number.Bytes()).
		SaveX(*gm.DB.Ctx)

	resp := base64.URLEncoding.EncodeToString(cred.Signature)
	return resp, nil
}
