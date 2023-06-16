package dispatcher

import (
	"context"
	"fmt"
	"log"
	"time"

	specyconfig "github.com/cosmos/relayer/v2/specy/config"
	"github.com/cosmos/relayer/v2/specy/types"
	"google.golang.org/grpc"
)

func InvokeSpecyEngineWithTx(
	ctx context.Context,
	cts []*types.ContractEvent,
	txHash []byte,
	msgSender []byte,
	chainID string,
	height uint64,
) (types.ProofResponse, error) {

	tmd := &types.TxMetaData{
		FromAddress: msgSender,
		ToAddress:   []byte{},
		Value:       0,
	}
	data := &types.Data{
		Meta:   tmd,
		Events: cts,
	}
	//构建监管请求
	pr := &types.ProofRequest{
		ChainType:    "cosmos",
		ChainID:      chainID,
		Data:         data,
		TxHash:       txHash,
		Height:       height,
		TxIndex:      1,
		OriginalData: nil,
	}

	cproof, err := SendProofRequest(ctx, *pr)

	if err != nil {
		fmt.Println("监管出错！！！！！")
		return types.ProofResponse{}, err
	}
	//对cproof进行检查，后期确定好如何检查在完善
	fmt.Println("监管成功，结束调用")
	fmt.Print(cproof)
	return cproof, err
}

func InvokeSpecyEngineWithTask(ctx context.Context, taskHash string) (types.TaskResponse, error) {
	// 构建请求
	request := &types.TaskRequest{
		Taskhash: []byte(taskHash),
	}

	response, err := SendTaskRequest(ctx, *request)
	return response, err
}

func SendTaskRequest(ctx context.Context, request types.TaskRequest) (types.TaskResponse, error) {
	//如果stream失效则应该进行重新构造尝试
	log.Default().Println("开始specy逻辑")
	cs := getCachedEngineStream(ctx)
	if cs == nil {
		specyEngineAddress := specyconfig.Config.EngineNodeAddress
		fmt.Printf("-------------specyEngineAddress: %s \n", specyEngineAddress)

		clientCon, err := grpc.Dial(specyEngineAddress, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(10*time.Second))
		if err != nil {
			log.Fatal(err)
			return types.TaskResponse{}, err
		}

		client := types.NewRegulatorClient(clientCon)
		fmt.Printf("-------------client: %+v \n", client)

		stream, err := client.GetTaskResult(context.Background())
		fmt.Printf("-------------stream: %+v \n", stream)

		if err != nil {
			return types.TaskResponse{}, err
		}
		ctx = cacheEngineStream(ctx, &stream)
		cs = stream
	}
	cs.Send(&request)
	resp, err := cs.Recv()
	fmt.Printf("-------------resp: %+v \n", resp)

	if err != nil {
		return types.TaskResponse{}, err
	}
	return *resp, nil
}

func cacheEngineStream(ctx context.Context, stream *types.Regulator_GetTaskResultClient) context.Context {
	specyInfoMap := *ctx.Value(types.SpecyInfoKey).(*map[string]any)
	specyInfoMap[types.EngineStreamKey] = &stream

	return context.WithValue(ctx, types.SpecyInfoKey, &specyInfoMap)
}

func getCachedEngineStream(ctx context.Context) types.Regulator_GetTaskResultClient {
	specyInfoMap, ok := ctx.Value(types.SpecyInfoKey).(map[string]any)
	if !ok {
		return nil
	}
	cs, ok := specyInfoMap[types.EngineStreamKey].(types.Regulator_GetTaskResultClient)
	if !ok {
		return nil
	}
	return cs
}

func SendProofRequest(ctx context.Context, pr types.ProofRequest) (types.ProofResponse, error) {
	//	//如果stream失效则应该进行重新构造尝试
	//	log.Default().Println("开始监管逻辑")
	//	cs := getCachedComplianceStream(ctx)
	//	if cs == nil {
	//		// TODO query registation endpoint info in global context, if does not exist, query chain and store in global context
	//		//registrationList := k.GetAllRegistration(ctx)
	//		//log.Default().Println(registrationList[0])
	//		//if len(registrationList) == 0 {
	//		//	return types.ProofResponse{}, types.ErrEmptyRegistrationList
	//		//}
	//
	//		//complianceAddress := registrationList[0].Endpoint.IpAddress + ":" + fmt.Sprint(registrationList[0].Endpoint.Port)
	//		complianceAddress := getRegulatoryEndpoint(ctx)
	//		fmt.Printf("-------------complianceAddress: %s \n", complianceAddress)
	//
	//		clientCon, err := grpc.Dial(complianceAddress, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(10*time.Second))
	//		if err != nil {
	//			log.Fatal(err)
	//		}
	//
	//		client := types.NewRegulatorClient(clientCon)
	//		fmt.Printf("-------------client: %+v \n", client)
	//
	//		stream, err := client.GetComplianceProof(context.Background())
	//		fmt.Printf("-------------stream: %+v \n", stream)
	//
	//		if err != nil {
	//			return types.ProofResponse{}, err
	//		}
	//		cacheComplianceStream(ctx, &stream)
	//		cs = stream
	//	}
	//	cs.Send(&pr)
	//	resp, err := cs.Recv()
	//	if err != nil {
	//		return types.ProofResponse{}, err
	//	}
	//	return *resp, nil
	return types.ProofResponse{}, nil
}

//
//func cacheComplianceStream(ctx context.Context, stream *types.Regulator_GetComplianceProofClient) {
//	specyInfoMap := ctx.Value(types.SpecyInfoKey).(map[string]any)
//	specyInfoMap[types.ComplianceStreamKey] = &stream
//
//	ctx = context.WithValue(ctx, types.SpecyInfoKey, &specyInfoMap)
//}
//
//func getCachedComplianceStream(ctx context.Context) types.Regulator_GetComplianceProofClient {
//	specyInfoMap, ok := ctx.Value(types.SpecyInfoKey).(map[string]any)
//	if !ok {
//		return nil
//	}
//	cs, ok := specyInfoMap[types.ComplianceStreamKey].(types.Regulator_GetComplianceProofClient)
//	if !ok {
//		return nil
//	}
//	return cs
//}

func getRegulatoryEndpoint(ctx context.Context) string {
	specyInfoMap := ctx.Value(types.SpecyInfoKey).(map[string]any)
	return specyInfoMap[types.RegistrationEndpointKey].(string)
}
