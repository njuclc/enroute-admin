<template>
  <div>
    <el-table style="width: 100%" :data="formatTableData()"
              >
      <el-table-column width="55" />
      <el-table-column align="left" label="导出时间" width="180">
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
  </div>
</template>

<script>
export default {
  name: "DxmWarehouse"
}
</script>

<style scoped>
</style>

<script setup>
import { getDictFunc, formatDate, formatBoolean, filterDict, ReturnArrImg, onDownloadFile } from '@/utils/format'

const props = defineProps(['tableData'])

const formatTableData = () => {
   return props.tableData?.dianxiaomi?.map((item) => {
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