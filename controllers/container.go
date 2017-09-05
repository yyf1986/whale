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
// @Param	container_name		query 	string	true		"容器名称"
// @Param	image_name		query 	string	true		"容器的image"
// @Param	env		query 	string	true		"env参数，多个以逗号分隔"
// @Success 200 {"status": 200}
// @Failure 403 body is empty
// @router /create [get]
func (c *ContainerController) CreateContainer() {
	cont_name := c.GetString("container_name")
	vm_ip := c.GetString("vm_ip")
	cont_image := c.GetString("image_name")
	env := c.GetString("env")
	c.Data["json"] = models.Create(vm_ip, cont_name, cont_image, env)
	c.ServeJSON()
}

// @Title delete container
// @Description delete container
// @Param	vm_ip		query 	string	true		"宿主机地址"
// @Param	container_id		query 	string	true		"容器id"
// @Success 200 {"status": 200}
// @Failure 403 body is empty
// @router /del [get]
func (c *ContainerController) DelContainer() {
	container_id := c.GetString("container_id")
	vm_ip := c.GetString("vm_ip")
	c.Data["json"] = models.Del(vm_ip, container_id)
	c.ServeJSON()
}

// @Title start container
// @Description start container
// @Param	vm_ip		query 	string	true		"宿主机地址"
// @Param	container_id		query 	string	true		"容器id"
// @Success 200 {"status": 200}
// @Failure 403 body is empty
// @router /start [get]
func (c *ContainerController) StartContainer() {
	container_id := c.GetString("container_id")
	vm_ip := c.GetString("vm_ip")
	c.Data["json"] = models.Start(vm_ip, container_id)
	c.ServeJSON()
}

// @Title stop container
// @Description stop container
// @Param	vm_ip		query 	string	true		"宿主机地址"
// @Param	container_id		query 	string	true		"容器id"
// @Success 200 {"status": 200}
// @Failure 403 body is empty
// @router /stop [get]
func (c *ContainerController) StopContainer() {
	container_id := c.GetString("container_id")
	vm_ip := c.GetString("vm_ip")
	c.Data["json"] = models.Stop(vm_ip, container_id)
	c.ServeJSON()
}

// @Title get all container
// @Description get all container
// @Param	vm_ip		query 	string	true		"宿主机地址"
// @Success 200 {"status": 200}
// @Failure 403 body is empty
// @router /getall [get]
func (c *ContainerController) GetAllContainer() {
	vm_ip := c.GetString("vm_ip")
	c.Data["json"] = models.GetAll(vm_ip)
	c.ServeJSON()
}
