package models

import (
	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/logs"
	"fmt"
	"strconv"
)

var (
	create  = "/v1/container/create"
	del     = "/v1/container/del"
	stop    = "/v1/container/stop"
	start   = "/v1/container/start"
	get     = "/v1/container/getall"
	getport = "v1/getport"
)

func Create(ip ,container_name,image_name, env string) string {
	var port int
	for _, ai := range AgentPool{
		if ai.IP == ip {
			port = ai.Port
		}
	}
	url := "http://" + ip + ":" + strconv.Itoa(port) + create
	fmt.Println(url)
	req := httplib.Get(url)
	req.Param("container_name", container_name)
	req.Param("image_name", image_name)
	req.Param("env", env)
	str, _ := req.String()
	logs.Info(str)
	return str
}
