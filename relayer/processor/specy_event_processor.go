package processor

import (
	"context"
	"fmt"
	"github.com/cosmos/relayer/v2/specy"
	"regexp"
	"strconv"
	"time"

	abci "github.com/cometbft/cometbft/abci/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/relayer/v2/utils"

	"go.uber.org/zap"
)

func HandleEventWithSpecy(
	ctx context.Context,
	log *zap.Logger,
	events []abci.Event,
	chainID string,
	height uint64,
	base64Encoded bool,
) {

	for _, event := range events {
		var evt sdk.StringEvent
		if base64Encoded {
			evt = utils.ParseBase64Event(event)
		} else {
			evt = sdk.StringifyEvent(event)
		}

		// listen regulatory events and init or update info in ctx map
		switch evt.Type {
		case "register":
			// TODO init register info map
		case "relation":
			// TODO init relation contract name map

		case "create_task":
			var taskHash string
			var startTime time.Time
			var intervalDuration time.Duration
			var single bool
			for _, attr := range event.Attributes {
				switch attr.Key {
				case "task_hash":
					taskHash = attr.Value
				case "task_rule_file":
					reg := regexp.MustCompile(`after (\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}\+\d{2}:\d{2})`)
					match := reg.FindStringSubmatch(attr.Value)

					if len(match) > 1 {
						dateTimeStr := match[1]
						fmt.Println(dateTimeStr)

						// 将字符串转换为 time.Time 类型
						dateTime, err := time.Parse("2006-01-02T15:04:05-07:00", dateTimeStr)
						if err != nil {
							fmt.Println("日期时间解析失败:", err)
						} else {
							fmt.Println(dateTime)
						}

						now := time.Now()
						nextTime := time.Date(now.Year(), now.Month(), now.Day(), dateTime.Hour(), dateTime.Minute(), dateTime.Second(), 0, now.Location())
						if nextTime.Before(now) || nextTime.Equal(now) {
							nextTime = nextTime.AddDate(0, 0, 1)
						}

						startTime = nextTime
					} else {
						fmt.Println("未找到匹配的结果")
					}

					intervalDuration = 24 * time.Hour
				case "task_single":
					single, _ = strconv.ParseBool(attr.Value)
				default:
					continue
				}
			}
			if !single {
				// 周期任务
				specy.StartScheduler(ctx, taskHash, startTime, intervalDuration)
			}

		case "cancle_task":
			var taskHash string
			for _, attr := range event.Attributes {
				switch attr.Key {
				case "task_hash":
					taskHash = attr.Value
				}
			}
			specy.StopScheduler(taskHash)
		}
	}
}
