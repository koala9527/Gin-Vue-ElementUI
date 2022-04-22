<template>
  <div id="root">
    <div style="display: flex;minHeight:98vh">

    <!-- 根节点必须有，不然报错 -->
    <el-aside id="part_left">
      <!-- 左边导航栏 -->
      <div id="ico">
        <!-- 网站图标 -->
        <img alt="Vue logo" src="./assets/logo.png" />
      </div>
      <div id="jump">
        <!-- 跳转的网站 -->
        <ul>
          <li v-for="item in jump_items" :key="item.name">
            <!-- v:for把数据进行循环渲染 -->
            <a :href="item.url">{{ item.name }}</a>
          </li>
        </ul>
      </div>
      <div id="only_menu">
        <!-- 坐侧大菜单按钮，用于右侧菜单滑动到指定位置 -->
        <ul id="menu_list">
          <li v-for="item in url_lists" :key="item.ID" >
            <a :href="'#'+ item.Name "><div >{{ item.Name }}</div></a>
          </li>
        </ul>
      </div>
    </el-aside>
    <el-main id="part_right">
      <!-- 右侧网站导航栏 -->
      <ul id="url_list">
        <li v-for="item in url_lists" :key="item.ID">
          <!-- v:for嵌套循环，循环里面就不能用id定位元素了，要用class -->
          <div class="one_part">
            <div class="menu_target" :id='item.Name'>{{ item.Name }}</div>
            <div class="url_list_data card">
              <li
                v-for="url_data in item.UrlLists"
                :key="url_data.ID"
                class="data-container"
              >
                <a :href="url_data.URL" class="btn">{{ url_data.Name }}</a>
              </li>
            </div>
          </div>
        </li>
      </ul>
    </el-main>
    </div>
    <el-footer>
      <div style="text-align: center;">
        <a style="color:#000" href="http://beian.miit.gov.cn/" target="_blank">鄂ICP备2021017303号</a>
      </div>
   
    </el-footer>

  </div>
</template>

<script>
export default {
  name: "App",
  components: {},
  data() {
    //渲染数据后期采用api动态请求的方式
    return {
      // 假的数据列表
      jump_items: [
        { name: "情报处1", url: "http://www.baidu.com" },
        { name: "情报处2", url: "http://www.baidu.com" },
        { name: "情报处3", url: "http://www.baidu.com" },
        { name: "情报处4", url: "http://www.baidu.com" },
        { name: "情报处5", url: "http://www.baidu.com" },
        { name: "情报处6", url: "http://www.baidu.com" },
      ],
      //这里后期通过一些
      url_lists: [],
    };
  },
  created() {
    this.listUlr();
  },

  methods: {
    listUlr() {
      this.$axios.get("/api").then((res) => {
        this.url_lists = res.data.data;
      });
    },
  },
};
</script>

<style>
/*元素样式，这个我最不在行，先做一个简单的出来，后面慢慢打磨*/
ul {
  overflow: hidden;
} /* 这个是浏览器的原因默认会有边距和填充 */
/* *{ margin:0; padding:0; } */

#part_left {
  display: flex;
  flex-direction: column;
  align-items: stretch;
  width: 350px;
  flex-shrink: 0;
  background-color: #e4e8ec;
}

#ico {
  display: flex;
  margin-top: 20px;
  justify-content: center;
}
img {
  width: 120px;
}
#jump {
  height: 180px;
  display: flex;
  flex-direction: row;
  justify-content: space-between;
}
#jump ul {
  display: flex;
  width: 100%;
  flex-wrap: wrap; /*项目允许换行*/
  justify-content: space-around;
  margin: 0;
  padding: 30px;
  text-align: center;
}
#jump ul li {
  list-style-type: none;
  margin-top: 12px;
  width: 80px;
  font-weight: 600;
  font-size: 22px;
}

#only_menu {
  display: flex;
}
#menu_list {
  display: flex;
  width: 100%;
  flex-wrap: wrap;
  margin: 0;
  justify-content: space-around;
  padding: 30px;
  text-align: center;
}

#menu_list li {
  list-style-type: none;
  margin-top: 12px;
  width: 80px;
  font-weight: 500;
  font-size: 25px;
}

#part_right {
  display: flex;
  width: 100%;
  top: 50px;
  right: 50px;
  background-color: #e4e8ec;
}
#url_list li {
  list-style-type: none;
}
.menu_target {
  list-style-type: none;
  font-weight: 800;
  font-size: 28px;
}
.url_list_data {
  display: flex;
  flex-wrap: wrap;
  width: 100%;
}

.url_list_data li {
  margin-top: 12px;
  margin-bottom: 18px;
  width: 200px;
  height: 40px;
  background: #e4e8ec;
  display: flex;
  justify-content: center;
  align-items: center;
}
.url_list_data li a {
  font-weight: 500;
  font-size: 25px;
}

.btn {
  position: relative;
  padding: 10px 20px;
  border-top-right-radius: 10px;
  border-bottom-left-radius: 10px;
  transition: all 0.5s;
}

.btn::after,
.btn::before {
  content: " ";
  width: 10px;
  height: 10px;
  position: absolute;
  border: 0px solid #e4e8ec;
  transition: all 0.3s;
}

.btn::after {
  top: -1px;
  left: -1px;
  border-top: 5px solid #41B883;
  border-left: 5px solid #41B883;
}

.btn::before {
  bottom: -1px;
  right: -1px;
  border-bottom: 5px solid #41B883;
  border-right: 5px solid #41B883;
}

.btn:hover {
  border-top-right-radius: 10px;
  border-bottom-left-radius: 10px;
}
.btn:hover::before,
.btn:hover::after {
  width: 100%;
  height: 100%;
}
</style>


