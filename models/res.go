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
}

var AgentInfos []*AgentInfo

var l sync.Mutex

func SetInfo(aa AgentInfo) {
	l.Lock()
	//time.Sleep(10 * time.Second)
	isNew := true
	for _, ai := range AgentInfos {                                
		if ai.IP == aa.IP {
			isNew = false
		}
	}
	if isNew {
		AgentInfos = append(AgentInfos, &aa)
	} else {
		for _, ai := range AgentInfos {
			if ai.IP == aa.IP {
				ai.Port = aa.Port
				ai.TotalCpu = aa.TotalCpu
				ai.TotalMem = aa.TotalMem
				ai.AverCpu = aa.AverCpu
				ai.AverMem = aa.AverMem
			}
		}
	}
	l.Unlock()
}

func GetInfo() []*AgentInfo{
	return AgentInfos
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

func checkInfo() {
	logs.Info("run check task")
	l.Lock()
	var agentinfos []*AgentInfo
	for _, ai := range AgentInfos {
		if checkPort(ai.IP,ai.Port) {
			agentinfos= append(agentinfos,ai)
		}
	}
	AgentInfos = agentinfos
	l.Unlock()
}

func addTask4Check() {
	f := func() error { checkInfo(); return nil }
	tk := toolbox.NewTask("check", "*/10 * * * * *", f)
	toolbox.AddTask("check", tk)
}

func init() {
	addTask4Check()
}
