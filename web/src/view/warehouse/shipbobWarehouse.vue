<template>
  <div>
    <el-table style="width: 100%" :data="formatTableData()"
    >
      <el-table-column width="55" />
      <el-table-column align="left" label="导出时间">
        <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
      </el-table-column>

      <el-table-column label="Product url">
        <template #default="scope">
          <el-button v-if="scope.row.link" plain text type="primary" size="small"
                     @click="handleClick(scope.row.link)">
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
  </div>

</template>

<script>
export default {
  name: "ShipbobWarehouse"
}
</script>

<style scoped>

</style>

<script setup>
import { getDictFunc, formatDate, formatBoolean, filterDict, ReturnArrImg, onDownloadFile } from '@/utils/format'

const props = defineProps(['tableData'])
const formatTableData = () => {
  return props.tableData?.shipbob?.map((item) => {
    item.link = item?.result?.downloadLink
    return item
  })
}

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
</script>