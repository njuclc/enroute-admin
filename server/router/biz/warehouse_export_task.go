package biz

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type WarehouseExportTaskRouter struct {
}

// InitWarehouseExportTaskRouter 初始化 库存导出 路由信息
func (s *WarehouseExportTaskRouter) InitWarehouseExportTaskRouter(Router *gin.RouterGroup) {
	warehouseRouter := Router.Group("warehouse").Use(middleware.OperationRecord())
	warehouseRouterWithoutRecord := Router.Group("warehouse")
	var warehouseApi = v1.ApiGroupApp.BizApiGroup.WarehouseExportTaskApi
	{
		warehouseRouter.POST("createWarehouseExportTask", warehouseApi.CreateWarehouseExportTask)   // 新建库存导出
		warehouseRouter.DELETE("deleteWarehouseExportTask", warehouseApi.DeleteWarehouseExportTask) // 删除库存导出
		warehouseRouter.DELETE("deleteWarehouseExportTaskByIds", warehouseApi.DeleteWarehouseExportTaskByIds) // 批量删除库存导出
		warehouseRouter.PUT("updateWarehouseExportTask", warehouseApi.UpdateWarehouseExportTask)    // 更新库存导出
	}
	{
		warehouseRouterWithoutRecord.GET("manualMonitorExport", warehouseApi.MonitorExport)
		warehouseRouterWithoutRecord.GET("manualTriggerDianxiaomiExport", warehouseApi.TriggerDianxiaomiExport)
		warehouseRouterWithoutRecord.GET("findWarehouseExportTask", warehouseApi.FindWarehouseExportTask)        // 根据ID获取库存导出
		warehouseRouterWithoutRecord.GET("getWarehouseExportTaskList", warehouseApi.GetWarehouseExportTaskList)  // 获取库存导出列表
	}
}
