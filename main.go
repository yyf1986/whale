package main

import (
	_ "whale/routers"
	
	"flag"
	
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/toolbox"
)

func main() {
	port := flag.Int("port", 12346, "http port")
	flag.Parse()
	
	//默认配置
	beego.BConfig.AppName = "whale"
	beego.BConfig.RunMode = "dev"
	beego.BConfig.ServerName = "whale"
	beego.BConfig.Listen.HTTPPort = *port
	beego.BConfig.CopyRequestBody = true
	beego.BConfig.WebConfig.AutoRender = false
	
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	toolbox.StartTask()
	beego.Run()
}
