package snark

// type ConstraintSystem struct{ constraint.ConstraintSystem }
// type ProveKey struct{ groth16.ProvingKey }

// type SnarkProver struct {
// 	ConstraintSystem ConstraintSystem
// 	ProveKey         ProveKey
// }

// func (prover *SnarkProver) Prover() *zkbansnark.SnarkProver {
// 	return &zkbansnark.SnarkProver{
// 		ConstraintSystem: prover.ConstraintSystem,
// 		ProveKey:         prover.ProveKey,
// 	}
// }

// func (ccs ConstraintSystem) MarshalJSON() ([]byte, error) {
// 	var buffer bytes.Buffer
// 	_, err := ccs.WriteTo(&buffer)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return buffer.Bytes(), nil
// }

// func (ccs ConstraintSystem) UnmarshalJSON(buf []byte) error {
// 	if string(buf) == "null" {
// 		return nil
// 	}

// 	var z ConstraintSystem
// 	reader := bytes.NewReader(buf)
// 	_, err := z.ReadFrom(reader)

// 	if err != nil {
// 		ccs.ConstraintSystem = z
// 	}

// 	return err
// }

// func (pk ProveKey) MarshalJSON() ([]byte, error) {
// 	var buffer bytes.Buffer
// 	_, err := pk.WriteTo(&buffer)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return buffer.Bytes(), nil
// }

// func (pk ProveKey) UnmarshalJSON(buf []byte) error {
// 	if string(buf) == "null" {
// 		return nil
// 	}

// 	var z ProveKey
// 	reader := bytes.NewReader(buf)
// 	_, err := z.ReadFrom(reader)

// 	if err != nil {
// 		pk.ProvingKey = z
// 	}

// 	return err
// }
