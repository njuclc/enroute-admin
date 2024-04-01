package biz

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/biz"
	bizReq "github.com/flipped-aurora/gin-vue-admin/server/model/biz/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"github.com/tidwall/gjson"
	"go.uber.org/zap"
	"gorm.io/datatypes"
	"log"
)

type WarehouseExportTaskApi struct {
}

var warehouseService = service.ServiceGroupApp.BizServiceGroup.WarehouseExportTaskService
var externalAuthService = service.ServiceGroupApp.BizServiceGroup.ExternalAuthService

// CreateWarehouseExportTask 创建库存导出
// @Tags WarehouseExportTask
// @Summary 创建库存导出
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body biz.WarehouseExportTask true "创建库存导出"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /warehouse/createWarehouseExportTask [post]
func (warehouseApi *WarehouseExportTaskApi) CreateWarehouseExportTask(c *gin.Context) {
	var warehouse biz.WarehouseExportTask
	err := c.ShouldBindJSON(&warehouse)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := warehouseService.CreateWarehouseExportTask(&warehouse); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteWarehouseExportTask 删除库存导出
// @Tags WarehouseExportTask
// @Summary 删除库存导出
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body biz.WarehouseExportTask true "删除库存导出"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /warehouse/deleteWarehouseExportTask [delete]
func (warehouseApi *WarehouseExportTaskApi) DeleteWarehouseExportTask(c *gin.Context) {
	ID := c.Query("ID")
	if err := warehouseService.DeleteWarehouseExportTask(ID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteWarehouseExportTaskByIds 批量删除库存导出
// @Tags WarehouseExportTask
// @Summary 批量删除库存导出
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /warehouse/deleteWarehouseExportTaskByIds [delete]
func (warehouseApi *WarehouseExportTaskApi) DeleteWarehouseExportTaskByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := warehouseService.DeleteWarehouseExportTaskByIds(IDs); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateWarehouseExportTask 更新库存导出
// @Tags WarehouseExportTask
// @Summary 更新库存导出
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body biz.WarehouseExportTask true "更新库存导出"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /warehouse/updateWarehouseExportTask [put]
func (warehouseApi *WarehouseExportTaskApi) UpdateWarehouseExportTask(c *gin.Context) {
	var warehouse biz.WarehouseExportTask
	err := c.ShouldBindJSON(&warehouse)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := warehouseService.UpdateWarehouseExportTask(warehouse); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindWarehouseExportTask 用id查询库存导出
// @Tags WarehouseExportTask
// @Summary 用id查询库存导出
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query biz.WarehouseExportTask true "用id查询库存导出"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /warehouse/findWarehouseExportTask [get]
func (warehouseApi *WarehouseExportTaskApi) FindWarehouseExportTask(c *gin.Context) {
	ID := c.Query("ID")
	if rewarehouse, err := warehouseService.GetWarehouseExportTask(ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rewarehouse": rewarehouse}, c)
	}
}

// GetWarehouseExportTaskList 分页获取库存导出列表
// @Tags WarehouseExportTask
// @Summary 分页获取库存导出列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bizReq.WarehouseExportTaskSearch true "分页获取库存导出列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /warehouse/getWarehouseExportTaskList [get]
func (warehouseApi *WarehouseExportTaskApi) GetWarehouseExportTaskList(c *gin.Context) {
	var pageInfo bizReq.WarehouseExportTaskSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := warehouseService.GetWarehouseExportTaskInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}

func (warehouseApi *WarehouseExportTaskApi) TriggerDianxiaomiExport(c *gin.Context) {
	cookies, e := getHeader("dianxiaomi")
	if e!=nil {
		return
	}

	client := resty.New()
	fieldStr := "SKU±sku≠仓库±warehose≠货架位±hjw≠库存量±stockNum≠单价±warehosePrice≠库存备注±warehoseComment≠安全库存量±safeAmount≠总价±amount≠采购在途±onPassageNum≠库内在途±purWaitInNum≠调拨在途±onPassageNumMove≠未发±unbilledOrderNum≠预售±lockStockNum≠可用库存±availableStockNum≠平台SKU±otherSku≠识别码±sbmId≠中文名称±name≠图片URL±imgUrl≠商品净重±productWeight≠采购参考价±productPrice≠采购员±agentName≠长±length≠宽±width≠高±height≠来源URL±sourceUrl≠商品备注±productComment≠商品编码±skuCode≠商品状态±productStatus≠包含子SKU±containSku≠中文报关±cnCustom≠英文报关±enCustom≠申报重量±customWeight≠申报金额±customPrice≠危险运输品±customDangerDes≠材质±customMaterial≠用途±customPurpose≠海关编码±customHgbm≠创建时间±createTime≠更新时间±updateTime≠图片±picture"
	resp, err := client.R().SetHeader("cookie", cookies["cookies"]).
		SetFormData(map[string]string{"templateField": fieldStr, "warehouseId": "6285210"}).
		Post("https://www.dianxiaomi.com/warehouseProduct/export.json")
	global.GVA_LOG.Info(fmt.Sprintf("resp: %s", utils.ToString(resp)))
	if err != nil {
		log.Fatal(err)
	}

	uuid := gjson.Get(resp.String(), "uuid").Str
	global.GVA_LOG.Info(fmt.Sprintf("uuid: %s", utils.ToString(uuid)))

	taskInfo := map[string]string{
		"uuid": uuid,
	}

	err = warehouseService.CreateWarehouseExportTask(&biz.WarehouseExportTask{
		Source:   "dianxiaomi",
		TaskInfo: datatypes.JSON(utils.ToString(taskInfo)),
		Result:   nil,
		Status:   "pending",
	})
	if err != nil {
		log.Fatal(err)
	}
}

func (warehouseApi *WarehouseExportTaskApi) MonitorExport(c *gin.Context) {
	reqInfo := bizReq.WarehouseExportTaskSearch{
		Status:         "pending",
	}
	taskList, total, err :=warehouseService.GetWarehouseExportTaskInfoList(reqInfo)
	msg := fmt.Sprintf("total: %d, err: %s, taskLen: %d", total, err, len(taskList))
	global.GVA_LOG.Info(msg)
	client := resty.New()

	for _, task := range taskList {
		fmt.Printf("start task, id: %d",task.ID)
		cookies, e := getHeader(task.Source)
		if e != nil {
			global.GVA_LOG.Error(fmt.Sprintf("get header err: %s", cookies))
			continue
		}

		uuid :=gjson.Get(task.TaskInfo.String(), "uuid").Str
		resp, err := client.R().SetHeader("cookie", cookies["cookies"]).
			SetFormData(map[string]string{"uuid": uuid}).
			Post("https://www.dianxiaomi.com/checkProcess.json")
		if err != nil {
			log.Fatal(err)
		}

		respCode := gjson.Get(resp.String(), "processMsg.code").Int()

		if respCode==1  {
			task.Status = "success"
			//task.Result,_ = respCodeNode.String()
		} else if respCode == -1 {
			task.Status = "expired"
		}

		task.Result = datatypes.JSON(resp.String())
		e = warehouseService.UpdateWarehouseExportTask(task)
		if e!=nil {
			global.GVA_LOG.Error(fmt.Sprintf("UpdateWarehouseExportTask err: %s", cookies))
		}
	}
}

func getHeader(source string) (map[string]string, error) {
	authInfo, err := externalAuthService.GetAuthInfo(source)
	if err != nil {
		return nil, err
	}

	if source == "dianxiaomi" {
		cookies := gjson.Get(authInfo, "header.cookies").Str
		return map[string]string{
			"cookies": cookies,
		},nil
	}

	return nil, nil
}