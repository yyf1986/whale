package models

import (
	"encoding/json"
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

type Port struct {
	Port int `json:"port"`
}

func genPort(abc string) Port {
	var resp Port
	if err := json.Unmarshal([]byte(abc), &resp); err != nil {
		fmt.Println(err.Error())
	}
	return resp
}

//根据ip地址查询是否在注册列表中，有返回对应的端口号，没有返回0
func checkIp(ip, do string) (int, string) {
	var port int
	var url string
	for _, ai := range AgentPool {
		if ai.IP == ip {
			port = ai.Port
			url = ai.Url[do]
		}
	}
	return port, url
}

func CreatePort(ip string) Port {
	port, urlpath := checkIp(ip, "CreatePort")
	if port != 0 {
		url := "http://" + ip + ":" + strconv.Itoa(port) + urlpath
		fmt.Println(url)
		req := httplib.Get(url)
		str, _ := req.String()
		logs.Info(str)
		return genPort(str)
	} else {
		return Port{0}
	}
}

func DelPort(ip, p string) string {
	port, urlpath := checkIp(ip, "DelPort")
	if port != 0 {
		url := "http://" + ip + ":" + strconv.Itoa(port) + urlpath
		fmt.Println(url)
		req := httplib.Get(url)
		req.Param("port", p)
		str, _ := req.String()
		logs.Info(str)
		return str
	} else {
		return `{"status":404,"errinfo":该机器没有注册"}`
	}
}

func GetAllPorts(ip string) string {
	port, urlpath := checkIp(ip, "GetAllPorts")
	if port != 0 {
		url := "http://" + ip + ":" + strconv.Itoa(port) + urlpath
		fmt.Println(url)
		req := httplib.Get(url)
		str, _ := req.String()
		logs.Info(str)
		return str
	} else {
		return `{"status":404,"errinfo":该机器没有注册"}`
	}
}

func DelAllPorts(ip string) string {
	port, urlpath := checkIp(ip, "DelAllPorts")
	if port != 0 {
		url := "http://" + ip + ":" + strconv.Itoa(port) + urlpath
		fmt.Println(url)
		req := httplib.Get(url)
		str, _ := req.String()
		logs.Info(str)
		return str
	} else {
		return `{"status":404,"errinfo":该机器没有注册"}`
	}
}
