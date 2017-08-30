package controllers

import (
	"whale/models"
	//"encoding/json"
	"github.com/astaxie/beego"
)

// Operations about Image
type ImageController struct {
	beego.Controller
}

// @Title get all container
// @Description get all container
// @Param	vm_ip		query 	string	true		"宿主机地址"
// @Success 200 {"status": 200}
// @Failure 403 body is empty
// @router /getall [get]
func (c *ImageController) GetAllImage() {
	vm_ip := c.GetString("vm_ip")
	c.Data["json"] = models.GetImages(vm_ip)
	c.ServeJSON()
}