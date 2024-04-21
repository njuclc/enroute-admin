<template>
  <div>
    <el-tabs v-model="activeName" tab-position="left">
      <el-tab-pane label="店小秘库存" name="first">
        <DxmWarehouse :tableData="sourceMap" />
      </el-tab-pane>
      <el-tab-pane label="Shopify Raise" name="second">
        dd
      </el-tab-pane>
      <el-tab-pane label="Shopify销量" name="third">
        ee
      </el-tab-pane>
      <el-tab-pane label="Shipbob" name="fourth">
        <ShipbobWarehouse :tableData="sourceMap" />
      </el-tab-pane>
    </el-tabs>
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
import {
  getWarehouseExportTaskList
} from '@/api/warehouseExportTask'
import _ from 'lodash'
import DxmWarehouse from "@/view/warehouse/dxmWarehouse.vue";
import ShipbobWarehouse from "@/view/warehouse/shipbobWarehouse.vue";
import { ref, reactive } from 'vue'
const activeName = ref("first")
let sourceMap = ref({})

// 查询
const getTableData = async () => {
  const table = await getWarehouseExportTaskList({ page: 1, pageSize: 40 })
  if (table.code === 0) {
    sourceMap.value = _.groupBy(table.data.list,  (item) =>item.source )
  }
  console.log(sourceMap.value)
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
getTableData()
</script>