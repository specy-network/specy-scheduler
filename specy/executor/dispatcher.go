package executor

import (
	"encoding/json"
	"fmt"
	"log"

	specytypes "github.com/cosmos/relayer/v2/specy/types"
)

func ExecuteTask(task *specytypes.Task) {

	// invoke specy engine
	//engineOutput, err := InvokeEngineWithTask(task.RuleFile, task.CheckData)
	engineOutput, err := mockEngine(task)
	if err != nil {
		log.Fatal(err)
		return
	}

	executeMsg, err := assembleExecuteMsgWithEngineOutput(task.Msg, engineOutput)
	if err != nil {
		fmt.Errorf("failed assemble execute msg: %s", err)
		return
	}

	// send task response to chain
	err = SendTaskResponseToChain(executeMsg, task)
	if err != nil {
		fmt.Errorf("SendTaskResponseToChain error: %s", err)
		return
	}
}

func mockEngine(task *specytypes.Task) (string, error) {
	return task.Msg, nil
}

func assembleExecuteMsgWithEngineOutput(taskMsg string, engineOutput string) (string, error) {
	var taskMsgData map[string]interface{}
	var engineOutputData map[string]interface{}

	err := json.Unmarshal([]byte(taskMsg), &taskMsgData)
	if err != nil {
		return "", fmt.Errorf("failed parsing task msg: %s", err)
	}

	err = json.Unmarshal([]byte(engineOutput), &engineOutputData)
	if err != nil {
		return "", fmt.Errorf("failed parsing engine output: %s", err)
	}

	assembleExecuteMsg(taskMsgData, engineOutputData)

	executeMsgData, err := json.Marshal(taskMsgData)
	if err != nil {
		return "", fmt.Errorf("failed marshaling execute msg data: %s", err)
	}
	executeMsg := string(executeMsgData)

	fmt.Println("execute msg:", executeMsg)
	return executeMsg, nil
}

func assembleExecuteMsg(taskMsgValueMap, engineOutputValueMap map[string]interface{}) {
	for key, value := range engineOutputValueMap {
		if taskMsgValue, exists := taskMsgValueMap[key]; exists {
			if engineOutputInnerMap, isTMap := value.(map[string]interface{}); isTMap {
				if taskMsgInnerMap, isEMap := taskMsgValue.(map[string]interface{}); isEMap {
					assembleExecuteMsg(taskMsgInnerMap, engineOutputInnerMap)
				} else {
					taskMsgValueMap[key] = value
				}
			} else {
				taskMsgValueMap[key] = value
			}
		}
	}
}
