<template>
  <div class="app-container">
    <el-table v-loading="listLoading" :data="listObjs" border fit highlight-current-row style="width: 100%">
      <el-table-column align="center" label="task_id" width="200">
        <template slot-scope="{row}">
          <!-- <span>{{ row.id }}</span> -->
          <span class="link-type" @click="datasetRouter(row.id)">
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
import { taskQueryList, taskDelete } from '@/api/mate'
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
      listLoading: true
    }
  },
  created() {
    this.getList()
  },
  methods: {
    datasetRouter(id) {
      this.$router.push({
        path: '/mate/dataset',
        query: {
          task_id: id
          // date_time: 1606899965
        }
      }) // 带参跳转
    },
    getList() {
      var query = { }
      this.listLoading = true
      taskQueryList(query).then(response => {
        this.listObjs = response.data
        this.listLoading = false
      })
    },

    deleteOne(id) {
      var query = { 'task_id': id }
      this.listLoading = true
      taskDelete(query).then(response => {
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
