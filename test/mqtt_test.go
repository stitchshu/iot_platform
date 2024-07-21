package test

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"testing"
	"time"
)

func TestMqtt(t *testing.T) {
	opt := mqtt.NewClientOptions().AddBroker("tcp://127.0.0.1:8080").SetClientID("go-test")

	//设置回调函数
	opt.SetDefaultPublishHandler(func(client mqtt.Client, message mqtt.Message) {
		fmt.Printf("TOPIC: %s\n", message.Topic())
		fmt.Printf("MESSAGE: %s\n", message.Payload())
	})
	c := mqtt.NewClient(opt)
	//连接
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		t.Fatal(token.Error())
	}

	//订阅主题
	if token := c.Subscribe("/topic/#", 0, nil); token.Wait() && token.Error() != nil {
		t.Fatal(token.Error())
	}

	//发布消息
	if token := c.Publish("/topic/1/1", 0, false, "hello"); token.Wait() && token.Error() != nil {
		t.Fatal(token.Error())
	}
	//取消订阅
	if token := c.Unsubscribe("/topic/#"); token.Wait() && token.Error() != nil {
		t.Fatal(token.Error())
	}

	time.Sleep(time.Second * 100)
	//关闭连接
	c.Disconnect(100)

}
