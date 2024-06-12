package test

import (
	"encoding/json"
	"fmt"
	"iot/define"
	"iot/helper"
	"testing"
)

var userServiceAddr = "http://127.0.0.1:8888"

func TestUserLogin(t *testing.T) {
	m := define.M{
		"userName": "get",
		"password": "1234",
	}
	data, _ := json.Marshal(m)
	rep, err := helper.HttpPost(userServiceAddr+"/user/login", data)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(rep))
}

//func TestMd5(t *testing.T) {
//	var s = "1234"
//	md5 := helper.Md5(s)
//	fmt.Println(md5)
//}
