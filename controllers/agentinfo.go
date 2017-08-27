package controllers

import (
	"test/models"
	"encoding/json"
	//"fmt"
	"github.com/astaxie/beego"
)

// Operations about AgentInfo
type AgentInfoController struct {
	beego.Controller
}

// @Title post agentinfo
// @Description post agentinfo
// @Param	body		body 	models.AgentInfo	true		"body for agent info"
// @Success 200 {"status": 200}
// @Failure 403 body is empty
// @router / [post]
func (u *AgentInfoController) Post() {
	var agentinfo models.AgentInfo
	json.Unmarshal(u.Ctx.Input.RequestBody, &agentinfo)
	//fmt.Println(agentinfo)
	models.SetInfo(agentinfo)
	u.Data["json"] = map[string]int{"status": 200}
	u.ServeJSON()
}

// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.AgentInfo
// @router / [get]
func (u *AgentInfoController) GetAll() {
	users := models.GetInfo()
	u.Data["json"] = users
	u.ServeJSON()
}