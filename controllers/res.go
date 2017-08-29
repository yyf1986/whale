package controllers

import (
	"whale/models"
	//"encoding/json"
	"github.com/astaxie/beego"
	//"strconv"
)

// Operations about resource
type ResController struct {
	beego.Controller
}

// @Title get port
// @Description get port
// @Param	vm_ip		query 	string	true		"宿主机地址"
// @Success 200 {"port": 200}
// @Failure 403 body is empty
// @router /createport [get]
func (c *ResController) CreatePort() {
	vm_ip := c.GetString("vm_ip")
	c.Data["json"] = models.CreatePort(vm_ip)
	c.ServeJSON()
}

// @Title del port
// @Description del port
// @Param	vm_ip		query 	string	true		"宿主机地址"
// @Param	port		query 	string	true		"宿主机端口"
// @Success 200 {"port": 200}
// @Failure 403 body is empty
// @router /delport [get]
func (c *ResController) DelPort() {
	vm_ip := c.GetString("vm_ip")
	port := c.GetString("port")
	c.Data["json"] = models.DelPort(vm_ip, port)
	c.ServeJSON()
}

// @Title get all ports
// @Description get all ports
// @Param	vm_ip		query 	string	true		"宿主机地址"
// @Success 200 {"port": 200}
// @Failure 403 body is empty
// @router /getallports [get]
func (c *ResController) GetAllPorts() {
	vm_ip := c.GetString("vm_ip")
	c.Data["json"] = models.GetAllPorts(vm_ip)
	c.ServeJSON()
}

// @Title del all ports
// @Description del all ports
// @Param	vm_ip		query 	string	true		"宿主机地址"
// @Success 200 {"port": 200}
// @Failure 403 body is empty
// @router /delallports [get]
func (c *ResController) DelAllPorts() {
	vm_ip := c.GetString("vm_ip")
	c.Data["json"] = models.DelAllPorts(vm_ip)
	c.ServeJSON()
}
