package test

import (
	"fmt"
	"iot/helper"
	"testing"
)

var adminServiceAddr = "http://127.0.0.1:14010"

func TestDeviceList(t *testing.T) {
	resp, err := helper.HttpGet(adminServiceAddr + "/device/list?page=1&size=20&name=")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(resp))

}
