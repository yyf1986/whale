package models

import (
	"strconv"
	"fmt"
	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/logs"
	"encoding/json"
)

var (
	getimages = "/v1/image/getall"
)

func genImage(abc string) []string {
	var resp []string
	if err := json.Unmarshal([]byte(abc), &resp); err != nil {
		fmt.Println(err.Error())
	}
	//fmt.Println(resp)
	return resp
}

func GetImages(ip string) []string {
	port := checkIp(ip)
	if port != 0 {
		url := "http://" + ip + ":" + strconv.Itoa(port) + getimages
		fmt.Println(url)
		req := httplib.Get(url)
		str, _ := req.String()
		logs.Info(str)
		return genImage(str)
	} else {
		return []string{}
	}
}
