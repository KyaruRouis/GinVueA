# 后端项目搭建

## 登录功能实现

### User数据结构

#### models/sys_user.go，设计用户数据结构

#### 自动建表：修改init.go中的NewGromDB函数

### 登录参数结构

#### service/types.go，设计LoginPasswordRequest数据结构

### 登录业务实现

#### service/login.go，新增LoginPassword函数

#### models/sys_user.go，新增GetUserByUsernamePassword函数

### jwt-go整合

#### define/define.go

##### 新增token配置

##### 新增用户声明数据结构

#### helper/helper.go

##### 新增GenerateToken函数

#### service/login.go

##### LoginPassword函数添加生成token、刷新token功能

#### service/types.go

##### 设计LoginPasswordReply数据结构

### 登录路由和配置跨域

#### router/app.go

##### App函数添加登录路由

##### 添加跨域中间件

#### middleware/cors.go

##### 新增Cors跨域函数

