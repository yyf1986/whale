package models

import (
	//"fmt"
	"sync"
	"net"
	"github.com/astaxie/beego/toolbox"
	"github.com/astaxie/beego/logs"
)

type AgentInfo struct {
	IP string
	Port int
	TotalCpu int
	TotalMem int
	AverCpu int
	AverMem int
	DockerStatus string
}

var AgentPool []*AgentInfo

var l sync.Mutex

func SetInfo(aa AgentInfo) {
	l.Lock()
	//time.Sleep(10 * time.Second)
	isNew := true
	for _, ai := range AgentPool {                                
		if ai.IP == aa.IP {
			isNew = false
		}
	}
	if isNew {
		AgentPool = append(AgentPool, &aa)
	} else {
		for _, ai := range AgentPool {
			if ai.IP == aa.IP {
				ai.Port = aa.Port
				ai.TotalCpu = aa.TotalCpu
				ai.TotalMem = aa.TotalMem
				ai.AverCpu = aa.AverCpu
				ai.AverMem = aa.AverMem
				ai.DockerStatus = aa.DockerStatus
			}
		}
	}
	l.Unlock()
}

func GetInfo() []*AgentInfo{
	return AgentPool
}

func GetInfo4Web() []*AgentInfo{
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
//检查端口
func checkPort(ip string, port int) bool{
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
	logs.Info("update agent info")
	l.Lock()
	var agentinfos []*AgentInfo
	for _, ai := range AgentPool {
		if checkPort(ai.IP,ai.Port) {
			agentinfos = append(agentinfos,ai)
		}
	}
	AgentPool = agentinfos
	l.Unlock()
}

func addTask4Update() {
	f := func() error { updateInfo(); return nil }
	tk := toolbox.NewTask("updateInfo", "0 */1 * * * *", f)
	toolbox.AddTask("updateInfo", tk)
}

func init() {
	addTask4Update()
}