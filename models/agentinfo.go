package models

import (
	//"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/toolbox"
	"net"
	"sync"
)

type AgentInfo struct {
	IP           string
	Port         int
	TotalCpu     int
	TotalMem     int
	AverCpu      int
	AverMem      int
	DockerStatus string
	Url          map[string]string
}

var AgentPool []*AgentInfo

var l sync.Mutex

//如果是新增的话，直接增加
//如果是已有的，则更新信息
func SetInfo(aa AgentInfo) {
	l.Lock()
	//time.Sleep(10 * time.Second)
	isNew := true
	logs.Info(aa.IP + " start to reg")
	for _, ai := range AgentPool {
		if ai.IP == aa.IP {
			isNew = false
		}
	}
	if isNew {
		AgentPool = append(AgentPool, &aa)
		logs.Info(aa.IP + " reg sucess")
	} else {
		for _, ai := range AgentPool {
			if ai.IP == aa.IP {
				ai.Port = aa.Port
				ai.TotalCpu = aa.TotalCpu
				ai.TotalMem = aa.TotalMem
				ai.AverCpu = aa.AverCpu
				ai.AverMem = aa.AverMem
				ai.DockerStatus = aa.DockerStatus
				logs.Info(aa.IP + " update sucess")
			}
		}
	}
	l.Unlock()
}

func GetInfo() []*AgentInfo {
	return AgentPool
}

func GetInfo4Web() []*AgentInfo {
	var webagentinfos []*AgentInfo
	l.Lock()
	for _, ap := range AgentPool {
		if ap.DockerStatus == "Y" {
			webagentinfos = append(webagentinfos, ap)
		}
	}
	l.Unlock()
	return webagentinfos
}

//通过注册的IP和端口查看服务是否可用
func checkPort(ip string, port int) bool {
	tcpAddr := net.TCPAddr{
		IP:   net.ParseIP(ip),
		Port: port,
	}
	conn, err := net.DialTCP("tcp", nil, &tcpAddr)
	if err == nil {
		conn.Close()
		return true

	} else {
		return false
	}
}

//端口不能访问，会把此agent从资源池中剔除，60s更新一次
func updateInfo() {
	logs.Info("Task start to update agent info")
	l.Lock()
	isok := true
	var agentinfos []*AgentInfo
	for _, ai := range AgentPool {
		if !checkPort(ai.IP, ai.Port) {
			isok = false
			break
		}
	}
	if !isok {
		for _, ai := range AgentPool {
			if checkPort(ai.IP, ai.Port) {
				agentinfos = append(agentinfos, ai)
			}
		}
		AgentPool = agentinfos
		logs.Info("Task end to update agent info")
		l.Unlock()
	} else {
		logs.Info("Task agent info do not need to update")
		l.Unlock()
	}
}

func addTask4Update() {
	f := func() error { updateInfo(); return nil }
	tk := toolbox.NewTask("updateInfo", "0 */1 * * * *", f)
	toolbox.AddTask("updateInfo", tk)
}

func init() {
	addTask4Update()
}
