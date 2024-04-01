<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="数据源:" prop="source">
           <el-select v-model="formData.source" placeholder="请选择数据源" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in stringOptions" :key="key" :label="item.label" :value="item.value" />
           </el-select>
       </el-form-item>
        <el-form-item label="任务信息:" prop="taskInfo">
          // 此字段为json结构，可以前端自行控制展示和数据绑定模式 需绑定json的key为 formData.taskInfo 后端会按照json的类型进行存取
          {{ formData.taskInfo }}
       </el-form-item>
        <el-form-item label="导出结果:" prop="result">
          // 此字段为json结构，可以前端自行控制展示和数据绑定模式 需绑定json的key为 formData.result 后端会按照json的类型进行存取
          {{ formData.result }}
       </el-form-item>
        <el-form-item label="任务状态:" prop="status">
          <el-input v-model="formData.status" :clearable="true"  placeholder="请输入任务状态" />
       </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createWarehouseExportTask,
  updateWarehouseExportTask,
  findWarehouseExportTask
} from '@/api/warehouseExportTask'

defineOptions({
    name: 'WarehouseExportTaskForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
const stringOptions = ref([])
const formData = ref({
            source: '',
            taskInfo: {},
            result: {},
            status: '',
        })
// 验证规则
const rule = reactive({
               source : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               taskInfo : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               status : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findWarehouseExportTask({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.rewarehouse
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    stringOptions.value = await getDictFunc('string')
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createWarehouseExportTask(formData.value)
               break
             case 'update':
               res = await updateWarehouseExportTask(formData.value)
               break
             default:
               res = await createWarehouseExportTask(formData.value)
               break
           }
           if (res.code === 0) {
             ElMessage({
               type: 'success',
               message: '创建/更改成功'
             })
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
