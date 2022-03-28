#MyGoWeb

>目录结构和功能

core 系统文件，不能随便更改

config 系统配置

controller 控制器目录 存放用户自定义控制器文件

filters 过滤器 register.go过滤器注册程序

mid 中间操作 register.go中间操作注册程序

public 直接访问文件目录 

public/templates html模板目录

routers 路由程序 用户自定义路由

utils 工具程序

>操作步骤

1. routers/router.go 创建路由

2. controller 创建控制器，到路由程序中把路由处理程序绑定到控制器

3. public/templates 创建html模板

4. 运行程序