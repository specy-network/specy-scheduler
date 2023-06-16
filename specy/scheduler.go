package specy

import (
	"context"
	"fmt"
	"github.com/cosmos/relayer/v2/specy/dispatcher"
	"time"
)

var (
	goroutines = make(map[string]chan struct{})
)

//type Task struct {
//	id              string
//	contractAddress string
//	method          string
//	schema          string
//	single          bool
//
//	startTime time.Time
//	interval  time.Duration
//}
//
//func NewTask(id string, contractAddress string, method string, schema string, single bool) *Task {
//	return &Task{
//		id:              id,
//		contractAddress: contractAddress,
//		method:          method,
//		schema:          schema,
//		single:          single,
//	}
//}

func StartScheduler(ctx context.Context, taskHash string, calldata string, taskCreator string, startTime time.Time, interval time.Duration) {
	startGoroutine(ctx, goroutines, taskHash, calldata, taskCreator, startTime, interval)
}

func startGoroutine(ctx context.Context, goroutines map[string]chan struct{}, taskHash string, calldata string, taskCreator string, startTime time.Time, interval time.Duration) {
	stopCh := make(chan struct{})

	// 存储 goroutine 对象
	goroutines[taskHash] = stopCh

	// 计算距离下一个时间节点的延迟时间
	delay := startTime.Sub(time.Now())

	// 创建一个定时器，在延迟时间到达后开始触发定时任务
	timer := time.NewTimer(delay)

	go func() {
		// 执行 goroutine 的逻辑
		for {
			// 等待定时器触发
			<-timer.C

			// 在定时任务中执行具体的操作
			fmt.Println("定时任务触发了！")

			dispatcher.ExecuteTask(ctx, taskHash, calldata, taskCreator)

			// 重新设置定时器，按照时间间隔触发下一次定时任务
			timer.Reset(interval)

			select {
			case <-stopCh:
				fmt.Println(taskHash, "stopped")
				return
			default:
				fmt.Println(taskHash, "running")
				// 执行其他操作
			}
		}
	}()
}

func StopScheduler(taskHash string) {
	stopGoroutine(goroutines, taskHash)
}

func stopGoroutine(goroutines map[string]chan struct{}, taskHash string) {
	stopCh, ok := goroutines[taskHash]
	if !ok {
		return
	}

	// 关闭 stop channel 以停止 goroutine
	close(stopCh)

	// 从 map 中移除 goroutine
	delete(goroutines, taskHash)
}
