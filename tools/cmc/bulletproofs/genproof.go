package bulletproofs

import (
	"chainmaker.org/chainmaker/common/crypto/bulletproofs"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
)

func proveCMD() *cobra.Command {
	proveCmd := &cobra.Command{
		Use:   "prove",
		Short: "Bulletproofs prove command",
		Long:  "Bulletproofs prove command",
		RunE: func(_ *cobra.Command, _ []string) error {
			return prove()
		},
	}

	flags := proveCmd.Flags()
	flags.StringVarP(&openingStr, "opening", "", "", "opening")
	flags.Int64VarP(&valueX, "value", "", -1, "value")

	return proveCmd
}

func prove() error {
	if valueX == -1 {
		return errors.New("invalid input, please check it")
	}
	commitmentStr := ""
	proofStr := ""
	if openingStr == "" {
		proof, commitment, opening, err := bulletproofs.Helper().NewBulletproofs().ProveRandomOpening(uint64(valueX))
		if err != nil {
			return err
		}
		proofStr = base64.StdEncoding.EncodeToString(proof)
		commitmentStr = base64.StdEncoding.EncodeToString(commitment)
		openingStr = base64.StdEncoding.EncodeToString(opening)
	} else {
		opening, err := base64.StdEncoding.DecodeString(openingStr)
		if err != nil {
			return err
		}
		proof, commitment, err := bulletproofs.Helper().NewBulletproofs().ProveSpecificOpening(uint64(valueX), opening)
		if err != nil {
			return err
		}
		proofStr = base64.StdEncoding.EncodeToString(proof)
		commitmentStr = base64.StdEncoding.EncodeToString(commitment)
	}

	fmt.Printf("value: [%d]\n", uint64(valueX))
	fmt.Printf("proof: [%s]\n", proofStr)
	fmt.Printf("commitment: [%s]\n", commitmentStr)
	fmt.Printf("opening: [%s]\n", openingStr)

	return nil
}