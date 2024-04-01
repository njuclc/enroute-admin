// 自动生成模板WarehouseExportTask
package biz

import (
    "github.com/flipped-aurora/gin-vue-admin/server/global"
    "gorm.io/datatypes"
)

// 库存导出 结构体  WarehouseExportTask
type WarehouseExportTask struct {
 global.GVA_MODEL 
      Source  string `json:"source" form:"source" gorm:"column:source;comment:;size:128;" binding:"required"`  //数据源 
      TaskInfo  datatypes.JSON `json:"taskInfo" form:"taskInfo" gorm:"column:task_info;comment:;type:text;" binding:"required"`  //任务信息 
      Result  datatypes.JSON `json:"result" form:"result" gorm:"column:result;comment:;type:text;"`  //导出结果 
      Status  string `json:"status" form:"status" gorm:"column:status;comment:;size:128;" binding:"required"`  //任务状态 
}


// TableName 库存导出 WarehouseExportTask自定义表名 warehouse_export_task
func (WarehouseExportTask) TableName() string {
  return "biz_warehouse_export_task"
}

type ExternalAuth struct {
    global.GVA_MODEL
    Source  string `json:"source" form:"source" gorm:"column:source;comment:;size:128;" binding:"required"`  //数据源
    Auth  datatypes.JSON `json:"auth" form:"auth" gorm:"column:auth;comment:;type:text;"`  //任务信息
}

// TableName 库存导出 WarehouseExportTask自定义表名 warehouse_export_task
func (ExternalAuth) TableName() string {
    return "biz_external_auth"
}

