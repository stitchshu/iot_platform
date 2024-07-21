package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"gorm.io/gorm"
	"iot/admin/internal/config"
	"iot/models"
	"iot/user/rpc/types/user"
	"iot/user/rpc/user_client"
)

type ServiceContext struct {
	Config   config.Config
	DB       *gorm.DB
	RpcUser  user_client.User
	AuthUser *user.UserAuthResponse
}

func NewServiceContext(c config.Config) *ServiceContext {
	models.NewDB()
	return &ServiceContext{
		Config:  c,
		DB:      models.DB,
		RpcUser: user_client.NewUser(zrpc.MustNewClient(c.RpcClientConf)),
	}
}
