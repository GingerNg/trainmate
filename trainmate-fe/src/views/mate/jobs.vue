<template>
  <div class="app-container">
    <el-table v-loading="listLoading" :data="listObjs" border fit highlight-current-row style="width: 100%">
      <el-table-column align="center" label="exp_id" width="200">
        <template slot-scope="{row}">
          <span>{{ row.exp_id }}</span>
        </template>
      </el-table-column>

      <el-table-column align="center" label="job_id" width="250px">
        <template slot-scope="{row}">
          <span>
            {{ row.id }}
          </span>
        </template>
      </el-table-column>

      <el-table-column width="100px" align="center" label="name">
        <template slot-scope="{row}">
          <span>{{ row.name }}</span>
        </template>
      </el-table-column>

      <el-table-column width="100px" align="center" label="metrics">
        <template slot-scope="{row}">
          <span>{{ row.metrics }}</span>
        </template>
      </el-table-column>

      <el-table-column width="100px" align="center" label="model_url">
        <template slot-scope="{row}">
          <span>{{ row.model_url }}</span>
        </template>
      </el-table-column>

      <el-table-column width="100px" align="center" label="config">
        <template slot-scope="{row}">
          <span>{{ row.config }}</span>
        </template>
      </el-table-column>

      <el-table-column width="100px" align="center" label="创建时间">
        <template slot-scope="{row}">
          <span>{{ row.create_time }}</span>
        </template>
      </el-table-column>

      <el-table-column width="100px" align="center" label="更新时间">
        <template slot-scope="{row}">
          <span>{{ row.update_time }}</span>
        </template>
      </el-table-column>

      <el-table-column width="100px" align="center" label="操作">
        <template slot-scope="{row}">
          <span class="link-type" @click="deleteOne(row.id)">
            删除
          </span>
        </template>
      </el-table-column>

    </el-table>
  </div>
</template>
<script>
import { jobQueryList, jobQueryOne, jobInsert, jobDelete } from '@/api/mate'
// import { fetchList } from '@/api/article'

export default {
  name: 'InlineEditTable',
  filters: {
    statusFilter(status) {
      const statusMap = {
        published: 'success',
        draft: 'info',
        deleted: 'danger'
      }
      return statusMap[status]
    }
  },
  data() {
    return {
      listObjs: null,
      listLoading: true,
      expId: ''
    }
  },
  created() {
    // 如果是跳转来的，则接受初始化参数
    if (this.$route.query) {
      this.expId = this.$route.query.exp_id
      // this.listQuery.date_time = this.$route.query.date_time
    }
    this.getList()
    // this.getOne()
    // this.newOne()
  },
  methods: {
    getList() {
      var query = { 'exp_id': this.expId }
      this.listLoading = true
      jobQueryList(query).then(response => {
        this.listObjs = response.data
        this.listLoading = false
      })
    },
    getOne() {
      var query = { 'job_id': '' }
      this.listLoading = true
      jobQueryOne(query).then(response => {
        this.listLoading = false
      })
    },

    newOne() {
      var query = { 'name': '121212112121', 'exp_id': '1212121' }
      this.listLoading = true
      jobInsert(query).then(response => {
        this.listLoading = false
      })
    },

    deleteOne(id) {
      var query = { 'job_id': id }
      this.listLoading = true
      jobDelete(query).then(response => {
        this.listLoading = false
      })
      this.getList()
    },

    refresh() {
      this.getList(),
      this.$message({
        message: '刷新成功',
        type: 'success'
      })
    }
  }
}
</script>

<style scoped>
.edit-input {
  padding-right: 100px;
}
.cancel-btn {
  position: absolute;
  right: 15px;
  top: 10px;
}
</style>
