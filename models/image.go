package models

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/logs"
	"strconv"
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
	port, urlpath := checkIp(ip, "GetImages")
	if port != 0 {
		url := "http://" + ip + ":" + strconv.Itoa(port) + urlpath
		fmt.Println(url)
		req := httplib.Get(url)
		str, _ := req.String()
		logs.Info(str)
		return genImage(str)
	} else {
		return []string{}
	}
}
