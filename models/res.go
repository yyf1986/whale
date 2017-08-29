package models

import (
	"fmt"
	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/logs"
	"strconv"
)

var (
	getport     = "/v1/res/createport"
	delport     = "/v1/res/delport"
	getallports = "/v1/res/getallports"
	delallports = "/v1/res/delallports"
)

func CreatePort(ip string) string {
	var port int
	for _, ai := range AgentPool {
		if ai.IP == ip {
			port = ai.Port
		}
	}
	url := "http://" + ip + ":" + strconv.Itoa(port) + getport
	fmt.Println(url)
	req := httplib.Get(url)
	str, _ := req.String()
	logs.Info(str)
	return str
}

func DelPort(ip, p string) string {
	var port int
	for _, ai := range AgentPool {
		if ai.IP == ip {
			port = ai.Port
		}
	}
	url := "http://" + ip + ":" + strconv.Itoa(port) + delport
	fmt.Println(url)
	req := httplib.Get(url)
	req.Param("port", p)
	str, _ := req.String()
	logs.Info(str)
	return str
}

func GetAllPorts(ip string) string {
	var port int
	for _, ai := range AgentPool {
		if ai.IP == ip {
			port = ai.Port
		}
	}
	url := "http://" + ip + ":" + strconv.Itoa(port) + getallports
	fmt.Println(url)
	req := httplib.Get(url)
	str, _ := req.String()
	logs.Info(str)
	return str
}

func DelAllPorts(ip string) string {
	var port int
	for _, ai := range AgentPool {
		if ai.IP == ip {
			port = ai.Port
		}
	}
	url := "http://" + ip + ":" + strconv.Itoa(port) + delallports
	fmt.Println(url)
	req := httplib.Get(url)
	str, _ := req.String()
	logs.Info(str)
	return str
}
