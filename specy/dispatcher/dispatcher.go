package dispatcher

import (
	"context"
	"log"
)

func ExecuteTask(ctx context.Context, taskHash string, calldata string, taskCreator string) {

	// invoke specy engine
	taskResponse, err := InvokeSpecyEngineWithTask(ctx, taskHash)
	if err != nil {
		log.Fatal(err)
		return
	}

	// send task response to chain
	SendTaskResponseToChain(taskResponse, calldata, taskCreator)
	if err != nil {
		log.Fatal(err)
		return
	}
}
