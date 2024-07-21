package mqtt

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"iot/models"
	"log"
	"strings"
)

var topic = "/sys/#"
var MC mqtt.Client

func NewMqttServer(mqttBroker, clientId, password string) {
	opt := mqtt.NewClientOptions().AddBroker(mqttBroker).SetClientID("go-mqtt-server-client-id").
		SetUsername("get").SetPassword(password)
	//设置回调函数
	opt.SetDefaultPublishHandler(publishHandler)

	MC := mqtt.NewClient(opt)
	//连接
	if token := MC.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	//订阅主题
	if token := MC.Subscribe(topic, 0, nil); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	defer func() {
		//取消订阅
		if token := MC.Unsubscribe(topic); token.Wait() && token.Error() != nil {
			log.Println("[ERROR]: ", token.Error())
		}
		//关闭连接
		MC.Disconnect(100)
	}()
	select {}
}
func publishHandler(client mqtt.Client, message mqtt.Message) {
	fmt.Printf("MESSAGE : #{message.Payload()}\n")
	fmt.Printf("TOPIC : #{message.Topic()}\n")

	topicArrays := strings.Split(strings.TrimPrefix(message.Topic(), "/"), "/")
	if len(topicArrays) >= 4 {
		if topicArrays[3] == "ping" {
			err := models.UpdateDeviceOnlineTime(topicArrays[1], topicArrays[2])
			if err != nil {
				log.Println("[DB ERROR]: %v\n", err)
			}
		}
	}
}
