package controllers

import (
	"whale/models"
	"encoding/json"
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
	models.SetInfo(agentinfo)
	u.Data["json"] = map[string]int{"status": 200}
	u.ServeJSON()
}

// @Title GetAll
// @Description get all agent
// @Success 200 {object} models.AgentInfo
// @router / [get]
func (u *AgentInfoController) GetAll() {
	agentinfos := models.GetInfo()
	u.Data["json"] = agentinfos
	u.ServeJSON()
}

// @Title Get4Web
// @Description get all agent which dockerstatus is y
// @Success 200 {object} models.AgentInfo
// @router /get4web [get]
func (u *AgentInfoController) Get4Web() {
	agentinfos := models.GetInfo4Web()
	u.Data["json"] = agentinfos
	u.ServeJSON()
}