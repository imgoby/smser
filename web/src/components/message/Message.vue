<template>
  <div>
    <el-form :inline="true" class="demo-form-inline">
      <el-form-item>
        <el-input placeholder="审批人" size="small"></el-input>
      </el-form-item>
      <el-form-item>
        <el-button size="small" type="primary">查询</el-button>
      </el-form-item>
    </el-form>
    <el-table
        :data="tableData"
        style="width: 100%">
      <el-table-column
          prop="date"
          label="日期"
          width="200"
      >
      </el-table-column>
      <el-table-column
          prop="name"
          label="姓名"
          width="200"
      >
      </el-table-column>
      <el-table-column
          prop="province"
          label="省份"
          width="300"
      >
      </el-table-column>
      <el-table-column
          prop="city"
          label="市区"
          width="200"
      >
      </el-table-column>
      <el-table-column
          prop="address"
          label="地址"
          width="300"
      >
      </el-table-column>
      <el-table-column
          prop="zip"
          label="邮编"
          width="200"
      >
      </el-table-column>
      <el-table-column
          fixed="right"
          label="操作"
          width="200">
        <template slot-scope="scope">
          <el-button
              @click.native.prevent="deleteRow(scope.$index, tableData)"
              type="text"
              size="small">
            移除
          </el-button>
          <el-button
              @click.native.prevent="deleteRow(scope.$index, tableData)"
              type="text"
              size="small">
            添加
          </el-button>
        </template>
      </el-table-column>
    </el-table>
    <el-pagination
        background
        layout="prev, pager, next"
        :total="1000">
    </el-pagination>
  </div>
</template>

<script>
export default {
  name: "Message"
}
</script>

<style scoped>

</style>
