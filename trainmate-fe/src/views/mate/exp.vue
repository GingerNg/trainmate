<template>
  <div class="app-container">
    <el-table v-loading="listLoading" :data="listObjs" border fit highlight-current-row style="width: 100%">
      <el-table-column align="center" label="dataset_id" width="250px">
        <template slot-scope="{row}">
          <span>{{ row.dataset_id }}</span>
        </template>
      </el-table-column>

      <el-table-column align="center" label="exp_id" width="200">
        <template slot-scope="{row}">
          <!-- <span>{{ row.exp_id }}</span> -->
          <!-- <el-button class="filter-item" type="primary" @click="$router.push('/mate/jobs')">{{ row.exp_id }}</el-button> -->
          <span class="link-type" @click="jobRouter(row.id)">
            {{ row.id }}
          </span>
        </template>
      </el-table-column>

      <el-table-column width="100px" align="center" label="name">
        <template slot-scope="{row}">
          <span>{{ row.name }}</span>
        </template>
      </el-table-column>

      <el-table-column width="100px" align="center" label="描述">
        <template slot-scope="{row}">
          <span>{{ row.desp }}</span>
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
import { expQueryList, expDelete } from '@/api/mate'

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
      listLoading: true
    }
  },
  created() {
    // 如果是跳转来的，则接受初始化参数
    if (this.$route.query) {
      this.dataset_id = this.$route.query.dataset_id
      // this.listQuery.date_time = this.$route.query.date_time
    }
    this.getList()
  },
  methods: {
    jobRouter(id) {
      this.$router.push({
        path: '/mate/jobs',
        query: {
          exp_id: id
          // date_time: 1606899965
        }
      }) // 带参跳转
    },

    getList() {
      var query = { }
      this.listLoading = true
      expQueryList(query).then(response => {
        this.listObjs = response.data
        this.listLoading = false
      })
    },

    deleteOne(id) {
      var query = { 'exp_id': id }
      this.listLoading = true
      expDelete(query).then(response => {
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
