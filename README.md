# iot_platform

>基于go-zero(微服务) 实现的物联网平台

#技术栈

+ 后端:go-zero
+ 前端:vue
+ 硬件:arduino、esp8266


#适用场景
共享单车、共享充电宝、外卖柜

## Topic
1. 心跳
```text
/sys/产品唯一标识/设备唯一标识/ping
```

##功能模块
+ [ ] 用户模块
    +[x] 登录 
+ [ ] 后台管理模块
  + [] 设备管理
    + [] 设备列表
+ [ ] 开放平台模块
+ [ ] 设备服务模块

5.搭建EMQX环境
```shell
docker run -d --name emqx -p 1883:1883 -p 8083:8083 -p 8084:8084 -p 8883:8883 -p 19083:18083 emqx/emqx:5.0.12
```
      
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
# 3.生成 user rpc服务代码
goctl rpc protoc user.proto --go_out=./types --go-grpc_out=./types --zrpc_out=. --style go_zero
# 4.生成 device rpc服务代码
goctl rpc protoc device.proto --go_out=./types --go-grpc_out=./types --zrpc_out=. --style go_zero
```

+ 启动服务

```shell
go run 服务名称.go -f 配置文件地址
# 1.启动user服务
go run user.go -f etc/user-api.yaml
# 2.启动admin服务
go run admin.go -f etc/admin-api.yaml
# 3.启动user rpc服务
go run user.go -f etc/user.yaml
```