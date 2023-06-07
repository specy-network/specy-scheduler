package dispatcher

import (
	"encoding/json"
	"fmt"
	specytypes "github.com/cosmos/relayer/v2/specy/types"
	"os/exec"
)

func SendTaskResponseToChain(specyResp specytypes.TaskResponse) error {
	jsonData, err := json.Marshal(specyResp.Result)
	if err != nil {
		fmt.Println("JSON encoding error:", err)
		return err
	}
	cmd := exec.Command("stargazed", "tx", "specy", "execute-task", string(specyResp.Taskhash), string(jsonData))
	// 执行命令并获取输出
	output, err := cmd.Output()
	if err != nil {
		return err
	}

	// 输出结果
	fmt.Println(string(output))
	return nil
}

func SendProofResponseToChain(txSpecResp specytypes.ProofResponse) error {
	jsonData, err := json.Marshal(txSpecResp.Proofs)
	if err != nil {
		fmt.Println("JSON encoding error:", err)
		return err
	}
	cmd := exec.Command("iris", "tx", "regulatory", "submit-spec-value", string(txSpecResp.TxHash), string(jsonData), string(txSpecResp.ProofsHash), string(txSpecResp.TeeSignature))
	// 执行命令并获取输出
	output, err := cmd.Output()
	if err != nil {
		return err
	}

	// 输出结果
	fmt.Println(string(output))
	return nil
}
