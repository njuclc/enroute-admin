package biz

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/biz"
    bizReq "github.com/flipped-aurora/gin-vue-admin/server/model/biz/request"
)

type WarehouseExportTaskService struct {
}

// CreateWarehouseExportTask 创建库存导出记录
// Author [piexlmax](https://github.com/piexlmax)
func (warehouseService *WarehouseExportTaskService) CreateWarehouseExportTask(warehouse *biz.WarehouseExportTask) (err error) {
	err = global.GVA_DB.Create(warehouse).Error
	return err
}

// DeleteWarehouseExportTask 删除库存导出记录
// Author [piexlmax](https://github.com/piexlmax)
func (warehouseService *WarehouseExportTaskService)DeleteWarehouseExportTask(ID string) (err error) {
	err = global.GVA_DB.Delete(&biz.WarehouseExportTask{},"id = ?",ID).Error
	return err
}

// DeleteWarehouseExportTaskByIds 批量删除库存导出记录
// Author [piexlmax](https://github.com/piexlmax)
func (warehouseService *WarehouseExportTaskService)DeleteWarehouseExportTaskByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]biz.WarehouseExportTask{},"id in ?",IDs).Error
	return err
}

// UpdateWarehouseExportTask 更新库存导出记录
// Author [piexlmax](https://github.com/piexlmax)
func (warehouseService *WarehouseExportTaskService)UpdateWarehouseExportTask(warehouse biz.WarehouseExportTask) (err error) {
	err = global.GVA_DB.Save(&warehouse).Error
	return err
}

// GetWarehouseExportTask 根据ID获取库存导出记录
// Author [piexlmax](https://github.com/piexlmax)
func (warehouseService *WarehouseExportTaskService)GetWarehouseExportTask(ID string) (warehouse biz.WarehouseExportTask, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&warehouse).Error
	return
}

// GetWarehouseExportTaskInfoList 分页获取库存导出记录
// Author [piexlmax](https://github.com/piexlmax)
func (warehouseService *WarehouseExportTaskService)GetWarehouseExportTaskInfoList(info bizReq.WarehouseExportTaskSearch) (list []biz.WarehouseExportTask, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&biz.WarehouseExportTask{})
    var warehouses []biz.WarehouseExportTask
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Source != "" {
        db = db.Where("source = ?",info.Source)
    }
    if info.Status != "" {
        db = db.Where("status = ?",info.Status)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Order("id desc").Find(&warehouses).Error
	return  warehouses, total, err
}
