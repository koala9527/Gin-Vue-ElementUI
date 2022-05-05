# Gin-Vue-ElementUI
Gin + Element UI + Vue 组合成一套 简陋套件，教程异常详细，非常适合想用新技术又没有多少技术的新手
# 特点
没有花里胡哨的的新奇的技术  
前后端分离  
粗糙，前端实际功能就是单页面导航数据请求渲染，CSS,JS,HTML代码纯手敲没有复制  
Golang的Gin 框架，引入了非常简单的基础包，主要是做mysql增删改查，MD5混淆  
ElementUI框架也实现了简单的登录注销，接口token鉴权，数据增删改查，注册管理员账号请手动添加mysql数据  
有Docker CI相关的高级功能

# 查看效果
目前还是可以查看效果,运行成本几乎没有，不能查看效果可以直接联系我（邮件&iss）  
前端访问：http://www-test.tuwei.space/  
管理后台访问：http://admin-test.tuwei.space/  账号：admin 密码：111111  

# 简单启动介绍
**这里只介绍本地运行，不介绍容器打包，CI的高级功能**  

## mysql 数据导入
* 首选需要一个数据库，创建一个my_web数据库
* 运行my_web.sql文件，导入数据
* ![image](https://user-images.githubusercontent.com/36888009/164601097-2d18d5ea-ab6d-48ef-905a-0c9c82869ed9.png)

## gin-api启动

* 修改gin-api框架中的config.ini文件中的[database]数据库配置  
* go mod tidy 拉取依赖  
* go run main.go 直接运行（数据库连不上启动会失败）  
* ![image](https://user-images.githubusercontent.com/36888009/164600171-385c0d9f-f53e-4b68-ad75-3fe2960b1444.png)

## vue-admin启动
* 查看.env文件配置`VUE_APP_BASE_API`是不是http://localhost:8080， 应该不用改哦
* npm install 安装依赖
* npm run dev 启动脚手架
* 接下来根据提示访问本地服务地址使用账号密码进去测试一下就行了
* ![image](https://user-images.githubusercontent.com/36888009/164602983-8f4a0337-7b40-41f4-b016-cefc10cc3a99.png)


## vue-web启动
* 同上修改.env文件
* npm install 安装依赖
* npm run serve 启动脚手架
* 直接访问查看效果了 
* ![image](https://user-images.githubusercontent.com/36888009/164603248-e6fdb2ff-2dff-41b2-a0c8-ba1c6bcdbca0.png)
# 新手向详细教程
教程所有经历的文章总结都在我的掘金账号：   
[土圭垚墝](https://juejin.cn/user/2893570335056494)    
2021年11月整个月的系列28篇文章：  
《从零开始摸索VUE，配合Golang搭建导航网站》系列   
还有些项目相关的运维相关的可以参考一下树莓派搭建内网穿透，GitLab搭建的教程 

[从零开始摸索VUE，配合Golang搭建导航网站（一.项目初始化）](https://juejin.cn/post/7025414336608731173)   
[从零开始摸索VUE，配合Golang搭建导航网站（二.了解项目结构）](https://juejin.cn/post/7025785301410775071)   
[从零开始摸索VUE，配合Golang搭建导航网站（三.做一个简单的单页面）](https://juejin.cn/post/7026151054836236325)   
[从零开始摸索VUE，配合Golang搭建导航网站（四.项目运行环境搭建和CI脚本编写）](https://juejin.cn/post/7026553910496067620)   
[从零开始摸索VUE，配合Golang搭建导航网站（五.使用doker部署启动）](https://juejin.cn/post/7026900855731257381)   
[从零开始摸索VUE，配合Golang搭建导航网站（六.CSS容器布局学习总结）](https://juejin.cn/post/7027268843441487885)   
[从零开始摸索VUE，配合Golang搭建导航网站（七.CSS Flex容器布局实战总结）](https://juejin.cn/post/7027643852626329613)   
[从零开始摸索VUE，配合Golang搭建导航网站（八.基于Golang的Gin框架的介绍）](https://juejin.cn/post/7028007794871631885)   
[从零开始摸索VUE，配合Golang搭建导航网站（九.Gin框架中GORM使用）](https://juejin.cn/post/7028378570812571656)   
[从零开始摸索VUE，配合Golang搭建导航网站（十.Gin框架优化，DockerFile编写）](https://juejin.cn/post/7028748827481866254)   
[从零开始摸索VUE，配合Golang搭建导航网站（十一.Gin容器化部署上线，CI脚本编写）](https://juejin.cn/post/7029164244554203167)   
[从零开始摸索VUE，配合Golang搭建导航网站（十二.使用Docker 新建Mysql应用，持久化数据保存，修改CI流程）](https://juejin.cn/post/7029490852858691621)   
[从零开始摸索VUE，配合Golang搭建导航网站（十三.Vue cli axios 简单使用）](https://juejin.cn/post/7029886862663614478)   
[从零开始摸索VUE，配合Golang搭建导航网站（十四.Vue cli env环境变量 ，后端跨域设置）](https://juejin.cn/post/7030287215108292645)   
[从零开始摸索VUE，配合Golang搭建导航网站（十五.添加数据后CSS样式优化）](https://juejin.cn/post/7030606687459344392)   
[从零开始摸索VUE，配合Golang搭建导航网站（十六.CSS动画初探）](https://juejin.cn/post/7030983896950898719)   
[从零开始摸索VUE，配合Golang搭建导航网站（十七.VUE锚点跳转，基础模板语法总结）](https://juejin.cn/post/7031332298846896165)   
[从零开始摸索VUE，配合Golang搭建导航网站（十八.Gin框架分层优化）](https://juejin.cn/post/7031719508251262990)   
[从零开始摸索VUE，配合Golang搭建导航网站（十九.GORM数据增删改查和Gin验证器）](https://juejin.cn/post/7032074651191017486)   
[从零开始摸索VUE，配合Golang搭建导航网站（二十.vue-element-admin 快速上手，认识基本架构）](https://juejin.cn/post/7032449009293738014)   
[从零开始摸索VUE，配合Golang搭建导航网站（二十一.vue-admin-template模拟账号登录）](https://juejin.cn/post/7032880603229847566)   
[从零开始摸索VUE，配合Golang搭建导航网站（二十二.vue-admin-template接入后端接口）](https://juejin.cn/post/7033204767195136036)   
[从零开始摸索VUE，配合Golang搭建导航网站（二十三.vue-admin-template接入后端增删改接口）](https://juejin.cn/post/7033572476164505636)   
[从零开始摸索VUE，配合Golang搭建导航网站（二十四.vue-admin-template带筛选的列表展示，初识双向绑定）](https://juejin.cn/post/7033940928465207332)   
[从零开始摸索VUE，配合Golang搭建导航网站（二十五.vue-admin-template分类详情数据删改查）](https://juejin.cn/post/7034305270289268773)   
[从零开始摸索VUE，配合Golang搭建导航网站（二十六.vue-admin-template完善用户登录后台接口）](https://juejin.cn/post/7034883872290504741)   
[从零开始摸索VUE，配合Golang搭建导航网站（二十七.vue-admin-template完善前端登陆逻辑）](https://juejin.cn/post/7035158665489285133)   
[从零开始摸索VUE，配合Golang搭建导航网站（二十八.vue-admin-template完善CI脚本，整体完成上线，汇总开源！！！）](https://juejin.cn/post/7035427424544227358)   

# 如有帮助可以请我喝杯咖啡

<img src="https://user-images.githubusercontent.com/36888009/164605674-da62f7b8-c4c2-4be3-a57b-91081e0f796b.jpg" width="300" height="300" alt="喝杯咖啡"/><br/>
