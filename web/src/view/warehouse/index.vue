<template>
  <div>
    <h1>店小秘</h1>
    <el-table :data="tableData" height="250" style="width: 100%">
      <el-table-column label="更新日期" width="180">
        <template #default="scope">
          <div>{{ scope.row.UpdatedAt.slice(0, 10) }}</div>
        </template>
      </el-table-column>
      <el-table-column label="单品/加工SKU文件" >
        <template #default="scope">
          <el-button link type="primary" size="small" @click="handleClick(scope.row.single_url)">
            下载
          </el-button>
        </template>
      </el-table-column>
      <el-table-column label="组合SKU文件">
        <template #default="scope">
          <el-button link type="primary" size="small" @click="handleClick(scope.row.combined_url)">
            下载
          </el-button>
        </template>
      </el-table-column>
      <el-table-column prop="status" label="任务状态" />
    </el-table>
  </div>
</template>

<script>
export default {
  name: "index.vue"
}
</script>

<style scoped>

</style>

<script setup>
import {ref} from "vue";
import {getWarehouseExportTaskList} from "@/api/warehouseExportTask";
import {toSQLLine} from "@/utils/stringFun";

const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

const onReset = () => {
  searchInfo.value = {}
}
// 搜索

const onSubmit = () => {
  page.value = 1
  pageSize.value = 10
  getTableData()
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 排序
const sortChange = ({ prop, order }) => {
  if (prop) {
    if (prop === 'ID') {
      prop = 'id'
    }
    searchInfo.value.orderKey = toSQLLine(prop)
    searchInfo.value.desc = order === 'descending'
  }
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getWarehouseExportTaskList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

const timeFormat = timeStr => timeStr.slice(0, 10)

getTableData();
const handleClick = (url) => {
  const tokens = url.split('/');
  console.log(tokens)

  const ele = document.createElement('a');
  ele.setAttribute('href', url);
  ele.setAttribute('download', tokens[tokens.length-1]);
  ele.style.display = 'none';
  document.body.appendChild(ele);
  ele.click();
  document.body.removeChild(ele);
}
getTableData()
</script>