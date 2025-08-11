package gm

import (
	"encoding/base64"

	corecore "github.com/akakou/zk-ban-system/core"
	"github.com/akakou/zk-ban-system/core/gm/ent/credential"
	"github.com/akakou/zk-ban-system/core/gm/ent/updateticket"
	"github.com/akakou/zk-ban/precomputes"
	"github.com/cockroachdb/errors"
)

var ErrAlreadyIssueCredential = errors.New("already issue credential")
var ErrCredentialNotFound = errors.New("ent: credential not found")

func (gm *GroupManager[T]) UpdateCredential(req *corecore.UpdateRequest, after int64) (string, error) {
	existCred, err := gm.DB.Client.Credential.Query().Where(
		credential.Or(
			credential.HasUpdateTicketWith(
				updateticket.TicketEQ(req.UpdateRequest.PublicKey.Number.Bytes()),
			),
		)).Only(*gm.DB.Ctx)

	if err == nil {
		resp := base64.URLEncoding.EncodeToString(existCred.Credential)
		return resp, nil
	} else if err.Error() != ErrCredentialNotFound.Error() {
		return "", err
	}

	if after == req.Before {
		return "", errors.New("no update")
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

	updateTicket := gm.DB.Client.UpdateTicket.Create().
		SetTicket(req.UpdateRequest.UpdateTicket.Number.Bytes()).
		SaveX(*gm.DB.Ctx)

	gm.DB.Client.Credential.Create().
		SetCredential(cred.Signature).
		SetUpdateTicket(updateTicket).
		SaveX(*gm.DB.Ctx)

	resp := base64.URLEncoding.EncodeToString(cred.Signature)
	return resp, nil
}
