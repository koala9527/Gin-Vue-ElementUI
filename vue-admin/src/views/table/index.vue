<template>
  <div class="app-container">

    <el-form class="filter-form">
      <el-form-item>
        <el-button type="primary" icon="el-icon-plus" @click="openAddFrom()">新增类别</el-button>
      </el-form-item>
    </el-form>

    <el-table
      v-loading="listLoading"
      :data="list"
      element-loading-text="加载中..."
      border
      fit
      highlight-current-row
    >
      <el-table-column align="center" label="ID" width="95">
        <template slot-scope="scope">
          {{ scope.row.ID }}
        </template>
      </el-table-column>
      <el-table-column label="Title">
        <template slot-scope="scope">
          {{ scope.row.Name }}
        </template>
      </el-table-column>
      <el-table-column align="center" fixed="right" label="操作" width="190px">
        <template slot-scope="scope">
          <div class="flex-btns">
            <el-button :loading="scope.row.loading" type="primary" @click="bindEdit(scope.row)">编辑</el-button>
            <el-button type="info" @click="deteleType(scope.row)">删除</el-button>

          </div>
        </template>
      </el-table-column>
    </el-table>
    <!--  添加分类  -->
    <el-dialog title="添加分类" :visible.sync="formAddVisible" width="450px" center>
      <el-form ref="addform">
        <el-form-item label="分类名称">
          <el-input v-model="addForm.name" placeholder="请输入内容" />
        </el-form-item>
        <el-form-item>
          <el-button native-type="primary" @click="onSubmitAdd">确认添加</el-button>
        </el-form-item>
      </el-form>
    </el-dialog>

    <!--  修改分类  -->
    <el-dialog title="编辑分类" :visible.sync="formEditVisible" width="450px" center>
      <el-form ref="addform">
        <el-form-item label="分类名称">
          <el-input v-model="editForm.name" placeholder="请输入分类" />
        </el-form-item>
        <el-form-item>
          <el-button native-type="primary" @click="onSubmitEdit">确认修改</el-button>
        </el-form-item>
      </el-form>
    </el-dialog>

  </div>
</template>

<script>
import { getTypeList, delType, addType, editType } from '@/api/url'

export default {
  filters: {
    statusFilter(status) {
      const statusMap = {
        published: 'success',
        draft: 'gray',
        deleted: 'danger'
      }
      return statusMap[status]
    }
  },
  data() {
    return {
      list: null,
      listLoading: true,
      formAddVisible: false,
      formEditVisible: false,
      addForm: {
        name: ''
      },
      editForm: {
        id: 0,
        name: ''
      }
    }
  },
  created() {
    this.fetchData()
  },
  methods: {
    fetchData() {
      this.listLoading = true
      getTypeList().then(response => {
        this.list = response.data
        this.listLoading = false
      })
    },
    openAddFrom() {
      this.formAddVisible = true
    },
    bindEdit(data) {
      this.editForm.id = data.ID
      this.editForm.name = data.Name
      this.formEditVisible = true
    },
    deteleType(data) {
      this.$confirm('是否删除分类?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'danger'
      }).then(() => {
        delType({
          id: data.ID
        }).then(res => {
          this.$message({
            type: 'success',
            message: '删除成功'
          })
          this.fetchData()
        })
      }).catch(() => {
        this.$message({
          type: 'info',
          message: '取消删除'
        })
      })
    },

    onSubmitAdd() {
      addType({
        name: this.addForm.name
      }).then(res => {
        this.$message({
          type: 'success',
          message: '添加成功'
        })
        this.fetchData()
      })
      this.formAddVisible = false
    },

    onSubmitEdit() {
      editType({
        name: this.editForm.name,
        id: this.editForm.id
      }).then(res => {
        this.$message({
          type: 'success',
          message: '修改成功'
        })
        this.fetchData()
      })
      this.formEditVisible = false
    }
  }
}
</script>
