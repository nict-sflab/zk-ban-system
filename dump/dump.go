package dump

import (
	"fmt"
	"os"
)

func DumpBasicKeys() {
	joinProver, joinVerify, err := JoinRequestCircuit()
	if err != nil {
		fmt.Printf("failed to generate join keys")
	}

	err = os.WriteFile("../../load/join_prover.key.json", joinProver, 0644)
	if err != nil {
		fmt.Printf("failed to generate join keys")
	}

	err = os.WriteFile("../../load/join_verifier.key.json", joinVerify, 0644)
	if err != nil {
		fmt.Printf("failed to generate join keys")
	}

	signProver, signVerifyKey, err := SignCircuit()
	if err != nil {
		fmt.Printf("failed to generate join keys")
	}

	err = os.WriteFile("../../load/sign_prover.key.json", signProver, 0644)
	if err != nil {
		fmt.Printf("failed to generate join keys")
	}

	err = os.WriteFile("../../load/sign_verifier.key.json", signVerifyKey, 0644)
	if err != nil {
		fmt.Printf("failed to generate join keys")
	}

}

func DumpUpdateKeys(nymsNumberPerSession, sessionNumber int) {
	updateProver, updateVerify, err := UpdateCircuit(nymsNumberPerSession, sessionNumber)
	if err != nil {
		fmt.Printf("failed to generate join keys")
	}

	proverFileName := fmt.Sprintf(UpdateProverKeyFileNameFormat, nymsNumberPerSession, sessionNumber)
	verifierFileName := fmt.Sprintf(UpdateVerifierKeyFileNameFormat, nymsNumberPerSession, sessionNumber)
	err = os.WriteFile(proverFileName, updateProver, 0644)
	if err != nil {
		fmt.Printf("failed to generate join keys")
	}

	err = os.WriteFile(verifierFileName, updateVerify, 0644)
	if err != nil {
		fmt.Printf("failed to generate join keys")
	}
}
