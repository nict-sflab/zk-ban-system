package gm

import (
	"encoding/base64"

	"github.com/akakou/zk-ban-system/core"
	"github.com/akakou/zk-ban-system/core/gm/ent/updateticket"
	"github.com/cockroachdb/errors"
)

var ErrAlreadyIssueCredential = errors.New("already issue credential")
var ErrCredentialNotFound = errors.New("ent: credential not found")
var ErrTicketAlreadyUsed = errors.New("ticket already has benn used")

func (gm *GroupManager[T]) UpdateCredential(req *core.UpdateRequest, after int64) (string, error) {
	registeredCred, err := gm.queryAlreadyRegisteredCredential(req.UpdateRequest.PublicKey.Number.Bytes())
	if err != nil {
		return "", err
	}

	if registeredCred != "" {
		return registeredCred, nil
	}

	if after == req.Before {
		return "", errors.New("no update")
	}

	ticketExist := gm.DB.Client.UpdateTicket.Query().
		Where(updateticket.Ticket(req.UpdateRequest.PublicKey.Number.Bytes())).
		ExistX(*gm.DB.Ctx)

	if ticketExist {
		return "", ErrTicketAlreadyUsed
	}

	index := core.KeyIndex{
		First:  after,
		Second: req.Before,
	}

	verifier, hasVerifier := gm.PreparedSnarkVerifiers[index]

	if !hasVerifier {
		verifier, err = gm.precomputesVerifyUpdateRequest(req.Before, after)
		if err != nil {
			return "", err
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

	gm.DB.Client.Credential.Create().
		SetCredential(cred.Signature).
		SetPublicKey(req.UpdateRequest.PublicKey.Number.Bytes()).
		SaveX(*gm.DB.Ctx)

	resp := base64.URLEncoding.EncodeToString(cred.Signature)
	return resp, nil
}
