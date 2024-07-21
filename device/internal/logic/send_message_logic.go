package logic

import (
	"context"
	"errors"
	"iot/device/internal/mqtt"

	"iot/device/internal/svc"
	"iot/device/types/device"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendMessageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendMessageLogic {
	return &SendMessageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendMessageLogic) SendMessage(in *device.SendMessageRequest) (*device.SendMessageResponse, error) {
	if in.ProductKey == " " || in.DeviceKey == "" || in.Data == "" {
		return nil, errors.New("参数异常")
	}
	topic := "/sys/" + in.ProductKey + "/" + in.DeviceKey + "/receive"
	//发布消息
	if token := mqtt.MC.Publish(topic, 0, false, "hello"); token.Wait() && token.Error() != nil {
		logx.Error("[PUBLISH ERROR]:", token.Error())
	}

	return &device.SendMessageResponse{}, nil
}
