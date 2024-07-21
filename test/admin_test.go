package test

import (
	"encoding/json"
	"fmt"
	"iot/helper"
	"testing"
)

var adminServiceAddr = "http://127.0.0.1:14010"

func TestDeviceList(t *testing.T) {
	m := map[string]string{
		"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwiaWRlbnRpdHkiOiIxIiwibmFtZSI6ImdldCIsInBhc3N3b3JkIjoiIiwiZXhwIjoxNzIwNzk3NzYzfQ.tCDUEuNdVxM6YMSld2N7sJTVb67L95NhkzXSsPxeFeE",
	}
	header, _ := json.Marshal(m)
	resp, err := helper.HttpGet(adminServiceAddr+"/device/list?page=1&size=20&name=", header...)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(resp))

}
