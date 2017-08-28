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
// @Param	vm_ip		query 	string	true		"宿主机地址"
// @Param	cont_name		query 	string	true		"容器名称"
// @Param	cont_image		query 	string	true		"容器的image"
// @Param	env		query 	string	true		"env参数，多个以逗号分隔"
// @Success 200 {"status": 200}
// @Failure 403 body is empty
// @router /create [get]
func (c *ContainerController) CreateContainer() {
	cont_name := c.GetString("cont_name")
	vm_ip := c.GetString("vm_ip")
	cont_image := c.GetString("cont_image")
	env := c.GetString("env")
	c.Data["json"] = models.Create(vm_ip, cont_name, cont_image, env)
	c.ServeJSON()
}

// @Title delete container
// @Description delete container
// @Param	vm_ip		query 	string	true		"宿主机地址"
// @Param	cont_name		query 	string	true		"容器名称"
// @Success 200 {"status": 200}
// @Failure 403 body is empty
// @router /del [get]
func (c *ContainerController) DelContainer() {
	cont_name := c.GetString("cont_name")
	vm_ip := c.GetString("vm_ip")
	c.Data["json"] = models.Del(vm_ip, cont_name)
	c.ServeJSON()
}

// @Title start container
// @Description start container
// @Param	vm_ip		query 	string	true		"宿主机地址"
// @Param	cont_name		query 	string	true		"容器名称"
// @Success 200 {"status": 200}
// @Failure 403 body is empty
// @router /start [get]
func (c *ContainerController) StartContainer() {
	cont_name := c.GetString("cont_name")
	vm_ip := c.GetString("vm_ip")
	c.Data["json"] = models.Start(vm_ip, cont_name)
	c.ServeJSON()
}

// @Title stop container
// @Description stop container
// @Param	vm_ip		query 	string	true		"宿主机地址"
// @Param	cont_name		query 	string	true		"容器名称"
// @Success 200 {"status": 200}
// @Failure 403 body is empty
// @router /stop [get]
func (c *ContainerController) StopContainer() {
	cont_name := c.GetString("cont_name")
	vm_ip := c.GetString("vm_ip")
	c.Data["json"] = models.Stop(vm_ip, cont_name)
	c.ServeJSON()
}
