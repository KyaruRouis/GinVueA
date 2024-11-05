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

## 管理员功能实现

### 获取管理员列表数据

#### models/sys_user.go

##### 新增GetUserList的函数用于获取管理员信息列表数据

#### service/types.go

##### 新增GetUserListRequest结构体

##### 新增QueryRequest结构体

##### 新增GetUserListReply结构体

#### service/init.go

##### 新增init.go文件

##### 新增NewQueryRequest函数

#### define/define.go

##### 新增默认分页显示条数

#### service/user.go

##### 新增user.go文件

##### 新增GetUserList函数

#### router/app.go

##### 新增获取管理员列表数据的路由

### 添加管理员

#### models/sys_user.go

##### 修改Sysuser结构

#### service/types.go

##### 新增AddUserRequest结构体

#### service/user.go

##### 新增AddUser函数

#### router/app.go

##### 添加管理员信息路由地址

### 获取管理员详情功能实现

#### service/types.go

##### 新增GetUserDetailReply数据结构

#### models/sys_user.go

##### 新增GetUserDetail函数

#### service/user.go

##### 新增GetUserDetail函数

#### router/app.go

##### 添加管理员详情信息路由地址

### 更新管理员

#### service/types.go

##### 新增UpdateUserRequest结构体

#### service/user.go

##### 新增UpdateUser函数

#### router/app.go

##### 新增更新管理员信息路由地址

### 删除管理员

#### service/user.go

##### 新增DeleteUser函数

#### router/app.go文件下

##### 新增删除管理员路由地址

## 角色模块功能实现

### 角色模块搭建

#### models/sys_role.go

##### 设计角色数据结构

### 获取角色列表数据

#### models/sys_role.go

##### 新增GetRoleList函数

#### service/types.go

##### 编写GetRoleListRequest数据结构

##### 编写GetRoleListReply数据结构

#### service/role.go

##### 新建role.go文件

##### 新增GetRoleList函数

#### router/app.go

##### 新增获取角色列表的路由

### 添加角色功能

#### service/types.go

##### 新增AddRoleRequest结构体

#### service/role.go

##### 新增AddRole函数

#### router/app.go

##### 添加路由地址

### 获取角色详情功能

#### service/types.go

##### 新增GetRoleDetailReply结构

#### models/sys_role.go

##### 新增GetRoleDetail函数

#### service/role.go

##### 新增GetRoleDetail函数

#### router/app.go

##### 新增路由地址

### 更新角色信息功能

#### service/types.go

##### 新增UpdateRoleRequest结构体，代码如下所示：

#### service/role.go

##### 新增UpdateRole函数

#### router/app.go

##### 新增路由地址

### 删除角色信息功能

#### service/role.go

##### 新增DeleteRole函数

#### router/app.go

##### 新增路由地址

### 修改角色管理员功能

#### service/role.go

##### 新增PatchRoleAdmin函数

#### router/app.go

##### 添加路由地址

## 菜单模块功能实现

### 编写菜单数据结构

#### models/sys_menu.go

### 设置自动建表

#### models/init.go文件

### 获取菜单列表数据

#### service/types.go

##### 新增MenuReply结构体

##### 新增AllMenu结构体

#### models/sys_menu.go

##### 新增GetMenuList函数

#### service/menu.go

##### 新增GetMenuList函数

##### 新增Menus的函数

##### 新增allMenuToMenuReply函数

##### 新增getChildrenMenu函数

#### router/app.go

##### 添加路由地址

### 新增菜单功能

#### service/types.go

##### 新增AddMenuRequest结构体

#### service/menu.go

##### 新增AddMenu函数

#### router/app.go

##### 添加路由地址

### 更新菜单信息功能

#### service/types.go

##### 新增UpdateMenuRequest结构体

#### service/menu.go

##### 新增UpdateMenu函数

#### router/app.go

##### 新增路由地址

### 删除菜单信息功能

#### service/menu.go

##### 新增DeleteMenu函数

#### router/app.go

##### 新增路由地址

## 角色和菜单授权功能实现

### 添加用户时绑定角色功能

#### router/app.go

##### 新增路由

#### service/role.go

##### 新增AllRole函数

#### models/sys_user.go

##### 修改SysUser结构体

#### service/user.go

##### 修改AddUser函数

#### service/types.go

##### 修改AddUserRequest结构体

##### 新增AllRoleListReply结构体

### 编辑用户时绑定角色

#### service/user.go

##### 修改GetUserDetail函数

##### 修改UpdateUser函数

### 添加角色时绑定菜单

#### models/role_menu.go

##### 新建role_menu.go文件

##### 新增RoleMenu结构体

##### 设置角色和菜单数据表的名称

#### models/init.go

##### 设置自动生成表结构

#### service/types.go

##### 修改AddRoleRequest结构体

#### service/role.go

##### 修改AddRole函数

### 编辑角色时分配菜单

#### service/role.go

##### 修改GetRoleDetail函数

##### 修改UpdateRole函数

#### models/role_menu.go

##### 新增GetRoleMenuId函数

### 根据登录用户的角色获取相应菜单功能

#### define/define.go

##### 修改UserClaim结构体

#### models/role_menu.go

##### 新增GetRoleMenus函数

#### service/menu.go

##### 修改Menus函数，代码如下所示：

#### middleware/cors.go

##### 修改Cors函数

### 登录认证功能

#### helper/helper.go

##### 新增AnalyzeToken函数

##### 修改GenerateToken函数

#### middleware/auth.go

##### 新增LoginAuthCheck函数

#### router/app.go

##### 添加auth.go中间件

#### service/login.go

##### 修改LoginPassword函数



### 
