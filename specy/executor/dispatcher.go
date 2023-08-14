package executor

import (
	"github.com/cosmos/relayer/v2/specy"
	"log"
)

func ExecuteTask(task *specy.Task) {

	// invoke specy engine
	taskResponse, err := InvokeEngineWithTask(task.TaskHash)
	if err != nil {
		log.Fatal(err)
		return
	}

	// send task response to chain
	SendTaskResponseToChain(taskResponse, task)
	if err != nil {
		log.Fatal(err)
		return
	}
}
