package models

import (
	"github.com/astaxie/beego/httplib"
)

var (
	create = "/v1/create"
	del = "/v1/del"
	stop = "/v1/stop"
	start = "/v1/start"
	get = "/v1/get"

)

func Create(ip,port string) {
	url := "http://" + ip + ":/" + port + create
	req := httplib.Get(url)
}