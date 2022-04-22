<template>
  <div class="app-container">
    <el-form :inline="true" class="filter-form" :show-message="false">
      <el-form-item label="类型：">
        <el-select v-model="type" placeholder="请选择类型" style="max-width:140px">
          <el-option v-for="item in typelist" :key="item.ID" :value="item.ID" :label="item.Name" />
        </el-select>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" icon="el-icon-plus" @click="openAddFrom()">新增数据</el-button>
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
      <el-table-column label="分类ID">
        <template slot-scope="scope">
          {{ scope.row.TypeID }}
        </template>
      </el-table-column>
      <el-table-column label="Title">
        <template slot-scope="scope">
          {{ scope.row.Name }}
        </template>
      </el-table-column>
      <el-table-column label="Url">
        <template slot-scope="scope">
          {{ scope.row.URL }}
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

    <!--  添加导航网址  -->
    <el-dialog title="添加导航url" :visible.sync="formAddVisible" width="450px" center>
      <el-form ref="addform" :inline="true">
        <el-form-item label="选择分类">
          <el-select v-model="add_type" placeholder="请选择类型" style="max-width:140px">
            <el-option v-for="item in typelist" :key="item.ID" :value="item.ID" :label="item.Name" />
          </el-select>
        </el-form-item>
        <el-form-item label="输入名称">
          <el-input v-model="addForm.name" placeholder="请输入导航名称" />
        </el-form-item>
        <el-form-item label="输入域名">
          <el-input v-model="addForm.url" placeholder="请输入导航url" />
        </el-form-item>
      </el-form>
      <el-button native-type="primary" @click="onSubmitAdd">确认添加</el-button>
    </el-dialog>

    <!--  修改导航url  -->
    <el-dialog title="编辑导航url" :visible.sync="formEditVisible" width="450px" center>
      <el-form ref="addform" :inline="true">
        <el-form-item label="选择分类">
          <el-select v-model="edit_type" placeholder="请选择类型" style="max-width:140px">
            <el-option v-for="item in typelist" :key="item.ID" :value="item.ID" :label="item.Name" />
          </el-select>
        </el-form-item>
        <el-form-item label="输入名称">
          <el-input v-model="editForm.name" placeholder="请输入导航名称" />
        </el-form-item>
        <el-form-item label="输入域名">
          <el-input v-model="editForm.url" placeholder="请输入导航url" />
        </el-form-item>
      </el-form>
      <el-button native-type="primary" @click="onSubmitEdit">确认修改</el-button>
    </el-dialog>
  </div>
</template>

<script>
import { getTypeList, getUrlList, editUrl, addUrl, delUrl } from '@/api/url'
export default {

  data() {
    return {
      add_type: '全部',
      edit_type: '',
      typelist: [],
      list: [],
      type: 0,
      formAddVisible: false,
      formEditVisible: false,
      addForm: {
        type_id: 0,
        name: '',
        url: ''
      },
      editForm: {
        id: 0,
        type_id: 0,
        name: '',
        url: ''
      }
    }
  },
  watch: {
    type(val) {
      this.fetchData()
    },
    add_type(val) {
      this.addForm.type_id = val
    }
  },
  created() {
    this.fetchData()
  },
  methods: {
    fetchData() {
      this.listLoading = true
      if (this.typelist.length === 0) {
        getTypeList().then(response => {
          response.data.unshift({
            ID: 0,
            Name: '全部'
          })
          this.typelist = response.data
        })
      }

      getUrlList({
        id: this.type
      }).then(response => {
        this.list = response.data.Value
        this.listLoading = false
      })
    },
    openAddFrom() {
      this.formAddVisible = true
      this.add_type = this.type
    },
    bindEdit(data) {
      this.editForm.id = data.ID
      this.editForm.name = data.Name
      this.editForm.url = data.URL
      this.editForm.type_id = data.TypeID
      this.edit_type = data.TypeID
      this.formEditVisible = true
    },
    deteleType(data) {
      this.$confirm('是否删除导航url?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'danger'
      }).then(() => {
        delUrl({
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
      addUrl({
        type_id: this.addForm.type_id,
        name: this.addForm.name,
        url: this.addForm.url
      }).then(res => {
        this.$message({
          type: 'success',
          message: '添加成功'
        })
        this.fetchData()
        this.formAddVisible = false
      })
    },
    onSubmitEdit() {
      editUrl({
        id: this.editForm.id,
        type_id: this.edit_type,
        name: this.editForm.name,
        url: this.editForm.url
      }).then(res => {
        this.$message({
          type: 'success',
          message: '修改成功'
        })
        this.fetchData()
        this.formEditVisible = false
      })
    }
  }
}
</script>

