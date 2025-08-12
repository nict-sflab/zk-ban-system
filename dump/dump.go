package dump

import (
	"fmt"
	"os"
)

const BASE_PATH = "../../load/"

func DumpBasicKeys() {
	joinProver, joinVerify, err := JoinRequestCircuit()
	if err != nil {
		fmt.Printf("failed to dump join keys")
	}

	err = os.WriteFile(BASE_PATH+"join_prover.key.json", joinProver, 0644)
	if err != nil {
		fmt.Printf("failed to dump join keys")
	}

	err = os.WriteFile(BASE_PATH+"join_verifier.key.json", joinVerify, 0644)
	if err != nil {
		fmt.Printf("failed to dump join keys")
	}

	signProver, signVerifyKey, err := SignCircuit()
	if err != nil {
		fmt.Printf("failed to dump sign keys")
	}

	err = os.WriteFile(BASE_PATH+"sign_prover.key.json", signProver, 0644)
	if err != nil {
		fmt.Printf("failed to dump sign keys")
	}

	err = os.WriteFile(BASE_PATH+"sign_verifier.key.json", signVerifyKey, 0644)
	if err != nil {
		fmt.Printf("failed to dump sign keys")
	}

}

func DumpUpdateKeys(nymsNumberPerSession, sessionNumber int) {
	updateProver, updateVerify, err := UpdateCircuit(nymsNumberPerSession, sessionNumber)
	if err != nil {
		fmt.Printf("failed to dump update keys")
	}

	proverFileName := fmt.Sprintf(UpdateProverKeyFileNameFormat, nymsNumberPerSession, sessionNumber)
	verifierFileName := fmt.Sprintf(UpdateVerifierKeyFileNameFormat, nymsNumberPerSession, sessionNumber)

	err = os.WriteFile(BASE_PATH+proverFileName, updateProver, 0644)
	if err != nil {
		fmt.Printf("failed to dump update keys")
	}

	err = os.WriteFile(BASE_PATH+verifierFileName, updateVerify, 0644)
	if err != nil {
		fmt.Printf("failed to dump update keys")
	}
}
