package dispatcher

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	specyconfig "github.com/cosmos/relayer/v2/specy/config"
	specytypes "github.com/cosmos/relayer/v2/specy/types"
	"os/exec"
	"time"
)

func SendTaskResponseToChain(specyResp specytypes.TaskResponse, calldata string) error {
	taskResult := string(specyResp.Result.TaskResult)
	//taskResult := "FM2vKqiPHN0XCQ=="
	//taskResult, _ = decodeTaskResult(taskResult)
	completeCalldata, err := assembleCalldata(calldata, taskResult)
	if err != nil {
		return err
	}

	cmd := exec.Command(specyconfig.Config.ChainBinaryLocation, "tx", "specy", "execute-task", string(specyResp.Taskhash), completeCalldata, string(specyResp.RuleFileHash), string(specyResp.Signature))
	// 执行命令并获取输出
	output, err := cmd.Output()
	if err != nil {
		return err
	}

	// 输出结果
	fmt.Println(string(output))
	return nil
}

func assembleCalldata(calldata string, taskResult string) (string, error) {
	// 解析 JSON 字符串
	var data Data
	err := json.Unmarshal([]byte(calldata), &data)
	if err != nil {
		fmt.Println("解析 JSON 失败:", err)
		return "", err
	}

	// 根据 index 赋值
	if data.Index >= 0 && data.Index < len(data.Params) {
		data.Params[data.Index].Value = taskResult
	}

	// 提取 value 值
	values := make([]string, len(data.Params))
	for i, param := range data.Params {
		values[i] = param.Value
	}
	currentTime := time.Now()
	truncatedTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 0, 0, 0, 0, currentTime.Location())
	timestamp := truncatedTime.Unix()
	values[0] = string(timestamp)

	var newData ExecuteData
	newData.Params = values
	newData.Index = data.Index
	fmt.Printf("ExecuteData: %+v\n", newData)

	jsonStr, err := json.Marshal(newData)
	fmt.Printf("jsonStr: %+v\n", jsonStr)

	return string(jsonStr), err
}

type Param struct {
	Value string `json:"value"`
}

type Data struct {
	Params []Param `json:"params"`
	Index  int     `json:"index"`
}

type ExecuteData struct {
	Params []string `json:"params"`
	Index  int      `json:"index"`
}

func decodeTaskResult(taskResult string) (string, error) {
	// 解码 Base64 字符串
	decoded, err := base64.StdEncoding.DecodeString(taskResult)
	if err != nil {
		fmt.Println("Base64 解码失败:", err)
		return "", err
	}

	// 将解码后的字节切片转换为十六进制字符串
	hexString := hex.EncodeToString(decoded)
	fmt.Println("转换后的十六进制字符串:", hexString)
	return hexString, err
}

func SendProofResponseToChain(txSpecResp specytypes.ProofResponse) error {
	jsonData, err := json.Marshal(txSpecResp.Proofs)
	if err != nil {
		fmt.Println("JSON encoding error:", err)
		return err
	}
	cmd := exec.Command(specyconfig.Config.ChainBinaryLocation, "tx", "regulatory", "submit-spec-value", string(txSpecResp.TxHash), string(jsonData), string(txSpecResp.ProofsHash), string(txSpecResp.TeeSignature))
	// 执行命令并获取输出
	output, err := cmd.Output()
	if err != nil {
		return err
	}

	// 输出结果
	fmt.Println(string(output))
	return nil
}
