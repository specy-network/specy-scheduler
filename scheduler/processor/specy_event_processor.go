package processor

import (
	"fmt"

	abci "github.com/cometbft/cometbft/abci/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/relayer/v2/specy"

	specytypes "github.com/cosmos/relayer/v2/specy/types"

	"regexp"
	"time"

	"github.com/cosmos/relayer/v2/utils"
)

func HandleEventWithSpecy(
	events []abci.Event,
	base64Encoded bool,
) {

	for _, event := range events {
		var evt sdk.StringEvent
		if base64Encoded {
			evt = utils.ParseBase64Event(event)
		} else {
			evt = sdk.StringifyEvent(event)
		}

		switch evt.Type {
		case "create_task":
			registerTaskOnScheduler(evt)

		case "cancle_task":
			unregisterTaskOnScheduler(evt)
		}

		// listen regulatory events and init or update info
		//case "register":
		//case "relation":
	}
}

func registerTaskOnScheduler(evt sdk.StringEvent) {
	// 入参打印
	fmt.Printf("registerTaskOnScheduler event: %+v \n", evt)

	var creator string
	var taskName string
	var taskHash string
	var connectionId string
	var msg string
	var ruleFile string
	var taskType = "0"
	var intervalType = "0"
	var interval = 1
	var startTime time.Time
	var checkData string

	for _, attr := range evt.Attributes {
		switch attr.Key {
		case "creator":
			creator = attr.Value
		case "task_name":
			taskName = attr.Value
		case "task_hash":
			taskHash = attr.Value
		case "connect_id":
			connectionId = attr.Value
		case "task_msgs":
			msg = attr.Value
		case "task_rule_file":
			ruleFile = attr.Value
			reg := regexp.MustCompile(`after (\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}\+\d{2}:\d{2})`)
			match := reg.FindStringSubmatch(ruleFile)

			if len(match) > 1 {
				// ruleFile 中指定了执行时间的情况
				dateTimeStr := match[1]
				// 解析
				startTime = parseAndCalcStartTime(dateTimeStr)
				// 执行间隔默认是24小时
				if interval == 0 {
					interval = 24 * 60 * 60
				}
			} else {
				// ruleFile 中没有指定执行时间 默认是现在开始(延迟30s)
				startTime = time.Now().Add(30 * time.Second)
			}
		//case "task_type":
		//	taskType = attr.Value
		//case "task_interval_type":
		//	intervalType = attr.Value
		//case "task_interval_number":
		//	interval, _ = strconv.Atoi(attr.Value)
		case "task_check_data":
			checkData = attr.Value
		default:
			continue
		}
	}

	// 注册任务
	task := specytypes.NewTask(taskHash, taskName, creator, connectionId, msg, ruleFile, taskType, intervalType, interval, startTime, checkData)
	specy.RegisterTask(task)
}

func parseAndCalcStartTime(dateTimeStr string) time.Time {
	fmt.Println(dateTimeStr)

	// 将字符串转换为 time.Time 类型
	dateTime, err := time.Parse("2006-01-02T15:04:05-07:00", dateTimeStr)
	if err != nil {
		fmt.Println("日期时间解析失败:", err)
	} else {
		fmt.Println(dateTime)
	}

	now := time.Now()
	var startTime time.Time
	if dateTime.Before(now) || dateTime.Equal(now) {
		startTime = time.Date(now.Year(), now.Month(), now.Day(), dateTime.Hour(), dateTime.Minute(), dateTime.Second(), 0, now.Location())
	} else {
		startTime = dateTime
	}
	if startTime.Before(now) || startTime.Equal(now) {
		startTime = startTime.AddDate(0, 0, 1)
	}

	return startTime
}

func unregisterTaskOnScheduler(evt sdk.StringEvent) {
	var taskHash string
	for _, attr := range evt.Attributes {
		switch attr.Key {
		case "task_hash":
			taskHash = attr.Value
		}
	}

	// 取消注册任务
	specy.UnregisterTask(taskHash)
}
