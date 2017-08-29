package models

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/logs"
	"strconv"
)

var (
	create = "/v1/container/create"
	del    = "/v1/container/del"
	stop   = "/v1/container/stop"
	start  = "/v1/container/start"
	getall = "/v1/container/getall"
)

type Resp struct {
	Status  int    `json:"status"`
	Errinfo string `json:"errinfo"`
}

func genResp(abc string) Resp {
	var resp Resp
	if err := json.Unmarshal([]byte(abc), &resp); err != nil {
		fmt.Println(err.Error())
	}
	//fmt.Println(resp)
	return resp
}

func Create(ip, container_name, image_name, env string) Resp {
	port := checkIp(ip)
	if port != 0 {
		url := "http://" + ip + ":" + strconv.Itoa(port) + create
		fmt.Println(url)
		req := httplib.Get(url)
		req.Param("container_name", container_name)
		req.Param("image_name", image_name)
		req.Param("env", env)
		str, _ := req.String()
		logs.Info(str)
		return genResp(str)
	} else {
		return Resp{404,"该机器没有注册"}
	}
}

func Del(ip, container_id string) Resp {
	port := checkIp(ip)
	if port != 0 {
		url := "http://" + ip + ":" + strconv.Itoa(port) + del
		fmt.Println(url)
		req := httplib.Get(url)
		req.Param("container_id", container_id)
		str, _ := req.String()
		logs.Info(str)
		return genResp(str)
	} else {
		return Resp{404,"该机器没有注册"}
	}
}

func Start(ip, container_id string) Resp {
	port := checkIp(ip)
	if port != 0 {
		url := "http://" + ip + ":" + strconv.Itoa(port) + start
		fmt.Println(url)
		req := httplib.Get(url)
		req.Param("container_id", container_id)
		str, _ := req.String()
		logs.Info(str)
		return genResp(str)
	} else {
		return Resp{404,"该机器没有注册"}
	}
}

func Stop(ip, container_id string) Resp {
	port := checkIp(ip)
	if port != 0 {
		url := "http://" + ip + ":" + strconv.Itoa(port) + stop
		fmt.Println(url)
		req := httplib.Get(url)
		req.Param("container_id", container_id)
		str, _ := req.String()
		logs.Info(str)
		return genResp(str)
	} else {
		return Resp{404,"该机器没有注册"}
	}

}

type Container struct {
	ContainerName   string
	ContainerID     string
	Image           string
	ContainerStatus string
}

func genResp2(abc string) []Container {
	var resp []Container
	if err := json.Unmarshal([]byte(abc), &resp); err != nil {
		fmt.Println(err.Error())
	}
	//fmt.Println(resp)
	return resp
}
//返回所有容器信息
func GetAll(ip string) []Container {
	port := checkIp(ip)
	if port != 0 {
		url := "http://" + ip + ":" + strconv.Itoa(port) + getall
		fmt.Println(url)
		req := httplib.Get(url)
		str, _ := req.String()
		logs.Info(str)
		return genResp2(str)
	} else {
		return []Container{}
	}
}
