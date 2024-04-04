<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" :rules="searchRule"
        @keyup.enter="onSubmit">
        <el-form-item label="创建日期" prop="createdAt">
          <template #label>
            <span>
              创建日期
              <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
                <el-icon>
                  <QuestionFilled />
                </el-icon>
              </el-tooltip>
            </span>
          </template>
          <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="开始日期"
            :disabled-date="time => searchInfo.endCreatedAt ? time.getTime() > searchInfo.endCreatedAt.getTime() : false"></el-date-picker>
          —
          <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="结束日期"
            :disabled-date="time => searchInfo.startCreatedAt ? time.getTime() < searchInfo.startCreatedAt.getTime() : false"></el-date-picker>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <el-table ref="multipleTable" style="width: 100%" tooltip-effect="dark" :data="tableData" row-key="ID"
      @selection-change="handleSelectionChange">
      <el-table-column type="selection" width="55" />

      <el-table-column align="left" label="日期" width="180">
        <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
      </el-table-column>

      <el-table-column label="单品/加工SKU文件">
        <template #default="scope">
          <el-button v-if="scope.row.single_url" plain text type="primary" size="small"
            @click="handleClick(scope.row.single_url)">
            下载
          </el-button>
          <div v-else>请等待导出任务完成</div>
        </template>
      </el-table-column>
      <el-table-column label="组合SKU文件">
        <template #default="scope">
          <el-button v-if="scope.row.single_url" plain text type="primary" size="small"
            @click="handleClick(scope.row.combined_url)">
            下载
          </el-button>
          <div v-else>请等待导出任务完成</div>
        </template>
      </el-table-column>
      <el-table-column label="任务状态">
        <template #default="scope">
          <el-tag v-if="scope.row.status === 'success'" type="success"> {{ scope.row.status }} </el-tag>
          <el-tag v-else-if="scope.row.status === 'pending'" type="primary"> {{ scope.row.status }} </el-tag>
          <el-tag v-else type="danger"> {{ scope.row.status }} </el-tag>
        </template>
      </el-table-column>
    </el-table>
    <el-pagination :current-page="page" :page-size="pageSize" :page-sizes="[10, 30, 50, 100]" :total="total"
      layout="total, prev, pager, next, sizes" @current-change="handleCurrentChange"
      @size-change="handleSizeChange" />
  </div>
</template>

<script setup>
import {
  getWarehouseExportTaskList
} from '@/api/warehouseExportTask'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict, ReturnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

defineOptions({
  name: 'WarehouseExportTask'
})

// 自动化生成的字典（可能为空）以及字段
const stringOptions = ref([])
const formData = ref({
  source: '',
  taskInfo: {},
  result: {},
  status: '',
})


// 验证规则
const rule = reactive({
  source: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  },
  {
    whitespace: true,
    message: '不能只输入空格',
    trigger: ['input', 'blur'],
  }
  ],
  taskInfo: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  },
  ],
  status: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  },
  {
    whitespace: true,
    message: '不能只输入空格',
    trigger: ['input', 'blur'],
  }
  ],
})

const searchRule = reactive({
  createdAt: [
    {
      validator: (rule, value, callback) => {
        if (searchInfo.value.startCreatedAt && !searchInfo.value.endCreatedAt) {
          callback(new Error('请填写结束日期'))
        } else if (!searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt) {
          callback(new Error('请填写开始日期'))
        } else if (searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt && (searchInfo.value.startCreatedAt.getTime() === searchInfo.value.endCreatedAt.getTime() || searchInfo.value.startCreatedAt.getTime() > searchInfo.value.endCreatedAt.getTime())) {
          callback(new Error('开始日期应当早于结束日期'))
        } else {
          callback()
        }
      }, trigger: 'change'
    }
  ],
})

const elSearchFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 搜索
const onSubmit = () => {
  elSearchFormRef.value?.validate(async (valid) => {
    if (!valid) return
    page.value = 1
    pageSize.value = 10
    getTableData()
  })
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async () => {
  const table = await getWarehouseExportTaskList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list.map((item) => {
      if (item?.result?.processMsg?.msg) {
        let msg = JSON.parse(item?.result?.processMsg?.msg)
        if (msg.downUrlList?.[0]) {
          item.single_url = msg.downUrlList?.[0]
        }
        if (msg.downUrlListZu?.[0]) {
          item.combined_url = msg.downUrlListZu?.[0]
        }
      }
      return item
    })
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

</script>

<style></style>
