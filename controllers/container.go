package controllers

import (
	"whale/models"
	//"encoding/json"
	"github.com/astaxie/beego"
)

// Operations about Container
type ContainerController struct {
	beego.Controller
}

// @Title create container
// @Description create container
// @Param	cont_name		query 	string	true		"容器名称"
// @Param	vm_ip		query 	string	true		"宿主机地址"
// @Param	cont_image		query 	string	true		"容器的image"
// @Param	id_para		query 	string	true		"env参数"
// @Success 200 {"status": 200}
// @Failure 403 body is empty
// @router /create [get]
func (c *ContainerController) CreateContainer() {
	cont_name := c.GetString("cont_name")
	vm_ip := c.GetString("vm_ip")
	cont_image := c.GetString("cont_image")
	env := c.GetString("id_para")
	c.Data["json"] = models.Create(vm_ip, cont_name, cont_image, env)
	c.ServeJSON()
}
