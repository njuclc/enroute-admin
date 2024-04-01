import service from '@/utils/request'

// @Tags WarehouseExportTask
// @Summary 创建库存导出
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WarehouseExportTask true "创建库存导出"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /warehouse/createWarehouseExportTask [post]
export const createWarehouseExportTask = (data) => {
  return service({
    url: '/warehouse/createWarehouseExportTask',
    method: 'post',
    data
  })
}

// @Tags WarehouseExportTask
// @Summary 删除库存导出
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WarehouseExportTask true "删除库存导出"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /warehouse/deleteWarehouseExportTask [delete]
export const deleteWarehouseExportTask = (params) => {
  return service({
    url: '/warehouse/deleteWarehouseExportTask',
    method: 'delete',
    params
  })
}

// @Tags WarehouseExportTask
// @Summary 批量删除库存导出
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除库存导出"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /warehouse/deleteWarehouseExportTask [delete]
export const deleteWarehouseExportTaskByIds = (params) => {
  return service({
    url: '/warehouse/deleteWarehouseExportTaskByIds',
    method: 'delete',
    params
  })
}

// @Tags WarehouseExportTask
// @Summary 更新库存导出
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.WarehouseExportTask true "更新库存导出"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /warehouse/updateWarehouseExportTask [put]
export const updateWarehouseExportTask = (data) => {
  return service({
    url: '/warehouse/updateWarehouseExportTask',
    method: 'put',
    data
  })
}

// @Tags WarehouseExportTask
// @Summary 用id查询库存导出
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.WarehouseExportTask true "用id查询库存导出"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /warehouse/findWarehouseExportTask [get]
export const findWarehouseExportTask = (params) => {
  return service({
    url: '/warehouse/findWarehouseExportTask',
    method: 'get',
    params
  })
}

// @Tags WarehouseExportTask
// @Summary 分页获取库存导出列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取库存导出列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /warehouse/getWarehouseExportTaskList [get]
export const getWarehouseExportTaskList = (params) => {
  return service({
    url: '/warehouse/getWarehouseExportTaskList',
    method: 'get',
    params
  })
}
