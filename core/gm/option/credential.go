package option

// func UpdateCredential[T any](req *core.UpdateRequest, after int64, gm gm.GroupManager[T]) (string, error) {
// 	cred, err := QueryCredential(req.UpdateRequest.UpdateTicket.Number.Bytes(), gm.DB)

// 	if err == nil {
// 		return cred, nil
// 	} else {
// 		fmt.Printf("db query error: %v\n", err)
// 		return gm.UpdateCredential(req, after)
// 	}
// }

// func QueryCredential(ticket []byte, db *gm.DB) (string, error) {

// 	if err != nil {
// 		return "", err
// 	}

// 	result := base64.URLEncoding.EncodeToString(cred.Credential)

// 	return result, nil
// }
