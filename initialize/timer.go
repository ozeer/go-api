package initialize

import (
	"fmt"

	"github.com/robfig/cron/v3"

	"github.com/ozeer/go-api/config"
	"github.com/ozeer/go-api/global"
	"github.com/ozeer/go-api/utils"
)

func Timer() {
	if global.CONFIG.Timer.Start {
		for i := range global.CONFIG.Timer.Detail {
			go func(detail config.Detail) {
				var option []cron.Option
				if global.CONFIG.Timer.WithSeconds {
					option = append(option, cron.WithSeconds())
				}
				_, err := global.Timer.AddTaskByFunc("ClearDB", global.CONFIG.Timer.Spec, func() {
					err := utils.ClearTable(global.DB, detail.TableName, detail.CompareField, detail.Interval)
					if err != nil {
						fmt.Println("timer error:", err)
					}
				}, option...)
				if err != nil {
					fmt.Println("add timer error:", err)
				}
			}(global.CONFIG.Timer.Detail[i])
		}
	}
}
