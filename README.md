# iot_platform

>基于go-zero(微服务) 实现的物联网平台

#技术栈

+ 后端:go-zero
+ 前端:vue
+ 硬件:arduino、esp8266


#适用场景
共享单车、共享充电宝、外卖柜

##功能模块
+ [ ] 用户模块
    +[x] 登录 
+ [ ] 后台管理模块
  + [] 设备管理
    + [] 设备列表
+ [ ] 开放平台模块
+ [ ] 设备服务模块
      
## 命令
+ 创建API服务
```shell
goctl api new 服务名称
# 1.创建user服务
goctl api new user
# 2.创建 admin服务
goctl api new admin
```
+ 生成服务代码
```shell
goctl api go -api 服务名称.api -dir . -style go_zero
# 1.生成user服务代码
goctl api go -api user.api -dir . -style go_zero
# 2.生成admin服务代码
goctl api go -api admin.api -dir . -style go_zero
```

+ 启动服务

```shell
go run 服务名称.go -f 配置文件地址
# 1.启动user服务
go run user.go -f etc/user-api.yaml
# 2.启动admin服务
go run admin.go -f etc/admin-api.yaml
```