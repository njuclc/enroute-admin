package initialize

import (
	"fmt"
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/task"

	"github.com/robfig/cron/v3"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

func Timer() {
	go func() {
		var option []cron.Option
		option = append(option, cron.WithSeconds())
		// 清理DB定时任务
		_, err := global.GVA_Timer.AddTaskByFunc("ClearDB", "@daily", func() {
			err := task.ClearTable(global.GVA_DB) // 定时任务方法定在task文件包中
			if err != nil {
				fmt.Println("timer error:", err)
			}
		}, "定时清理数据库【日志，黑名单】内容", option...)
		if err != nil {
			fmt.Println("add timer error:", err)
		}

		// 其他定时任务定在这里 参考上方使用方法

		var warehouseApi = v1.ApiGroupApp.BizApiGroup.WarehouseExportTaskApi
		_, err = global.GVA_Timer.AddTaskByFunc("触发导出", "0 0 1 * * *", func() {
			warehouseApi.TriggerDianxiaomiExport(nil)
		}, "触发导出", option...)
		if err != nil {
			fmt.Println("add timer error:", err)
		}

		_, err = global.GVA_Timer.AddTaskByFunc("监听导出任务", "0 0/1 * * * *", func() {
			warehouseApi.MonitorExport(nil)
		}, "监听任务是否完成", option...)
		if err != nil {
			fmt.Println("add timer error:", err)
		}

	}()
}
