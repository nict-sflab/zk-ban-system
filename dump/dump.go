package dump

import (
	"fmt"
	"os"

	"github.com/akakou/zk-ban-system/core/core"
)

func DumpBasicKeys() {
	joinProver, joinVerify, err := core.JoinRequestCircuit()
	if err != nil {
		fmt.Printf("failed to generate join keys")
	}

	err = os.WriteFile("./join_prover.key.json", joinProver, 0644)
	if err != nil {
		fmt.Printf("failed to generate join keys")
	}

	err = os.WriteFile("./join_verifier.key.json", joinVerify, 0644)
	if err != nil {
		fmt.Printf("failed to generate join keys")
	}

	signProver, signVerifyKey, err := core.SignCircuit()
	if err != nil {
		fmt.Printf("failed to generate join keys")
	}

	err = os.WriteFile("./sign_prover.key.json", signProver, 0644)
	if err != nil {
		fmt.Printf("failed to generate join keys")
	}

	err = os.WriteFile("./sign_verifier.key.json", signVerifyKey, 0644)
	if err != nil {
		fmt.Printf("failed to generate join keys")
	}

}

func DumpUpdateKeys(a, b int) {
	updateProver, updateVerify, err := core.UpdateCircuit(a, b)
	if err != nil {
		fmt.Printf("failed to generate join keys")
	}

	err = os.WriteFile("./update_prover.key.json", updateProver, 0644)
	if err != nil {
		fmt.Printf("failed to generate join keys")
	}

	err = os.WriteFile("./update_verifier.key.json", updateVerify, 0644)
	if err != nil {
		fmt.Printf("failed to generate join keys")
	}
}
