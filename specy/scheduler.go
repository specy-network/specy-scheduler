package specy

import (
	"fmt"
	"github.com/cosmos/relayer/v2/specy/executor"
	"sync"
	"time"
)

var (
	everyBlockTasks            = make(map[string]*Task)
	timeIntervalTaskGoroutines = make(map[string]chan struct{})
)

type Task struct {
	TaskHash     string
	TaskName     string
	Creator      string
	ConnectionId string
	Msgs         string
	RuleFile     string
	TaskType     string

	Condition Condition
}

type Condition struct {
	IntervalType string
	Interval     int

	StartTime time.Time
}

func NewTask(taskHash string, taskName string, creator string, connectionId string, msgs string, ruleFile string, taskType string, intervalType string, interval int, startTime time.Time) *Task {

	return &Task{
		TaskHash:     taskHash,
		TaskName:     taskName,
		Creator:      creator,
		ConnectionId: connectionId,
		Msgs:         msgs,
		RuleFile:     ruleFile,
		TaskType:     taskType,

		Condition: Condition{
			intervalType,
			interval,
			startTime,
		},
	}
}

func RegisterTask(task *Task) {
	switch task.Condition.IntervalType {
	case "time_interval":
		// 直接触发 task (goroutine)
		triggerTimeIntervalTask(task)

	case "every_block":
		// 将 task 注册到任务列表中 待爬区块的时候遍历触发
		everyBlockTasks[task.TaskHash] = task
	default:
		fmt.Println("unsupportted task type")
	}
}

func triggerTimeIntervalTask(task *Task) {
	stopCh := make(chan struct{})

	// 存储 goroutine 对象
	timeIntervalTaskGoroutines[task.TaskHash] = stopCh

	// 计算距离下一个时间节点的延迟时间
	delay := task.Condition.StartTime.Sub(time.Now())

	// 创建一个定时器，在延迟时间到达后开始触发定时任务
	timer := time.NewTimer(delay)

	go func() {
		// goroutine 中定时执行
		for ; true; <-timer.C {

			// 在定时任务中执行具体的操作
			fmt.Println("定时任务触发了！")

			executor.ExecuteTask(task)

			// 重新设置定时器，按照时间间隔触发下一次定时任务
			timer.Reset(time.Duration(task.Condition.Interval) * time.Second)

			select {
			case <-stopCh:
				fmt.Println(task.TaskHash, "stopped")
				return
			default:
				fmt.Println(task.TaskHash, "running")
			}
		}
	}()
}

func TriggerEveryBlockTasks() {

	// 将task放到数组中 方便后续操作
	tasks := make([]*Task, 0, len(everyBlockTasks))
	for _, task := range everyBlockTasks {
		tasks = append(tasks, task)
	}

	// 每个子列表的大小
	batchSize := 10

	// 计算需要划分的子列表数量
	numBatches := (len(everyBlockTasks) + batchSize - 1) / batchSize

	// 分批处理列表
	for i := 0; i < numBatches; i++ {
		start := i * batchSize
		end := (i + 1) * batchSize

		if end > len(everyBlockTasks) {
			end = len(everyBlockTasks)
		}

		processTriggerEveryBlockTask(tasks[start:end])
	}
}

func processTriggerEveryBlockTask(tasks []*Task) {
	var wg sync.WaitGroup
	defer wg.Done()

	var itemWg sync.WaitGroup

	for _, task := range tasks {
		itemWg.Add(1)

		go func(task *Task) {
			defer itemWg.Done()

			fmt.Println("Processing task:", task)
			executor.ExecuteTask(task)

		}(task)
	}

	itemWg.Wait()
}

func UnregisterTask(taskHash string) {
	if everyBlockTasks[taskHash] != nil {
		removeEveryBlockTask(taskHash)
	} else {
		stopTimeIntervalTaskGoroutine(taskHash)
	}
}

func removeEveryBlockTask(hash string) {
	// 从 map 中移除 task
	delete(everyBlockTasks, hash)
}

func stopTimeIntervalTaskGoroutine(taskHash string) {
	stopCh, ok := timeIntervalTaskGoroutines[taskHash]
	if !ok {
		return
	}

	// 关闭 stop channel 以停止 goroutine
	close(stopCh)

	// 从 map 中移除 goroutine
	delete(timeIntervalTaskGoroutines, taskHash)
}
