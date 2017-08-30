// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"whale/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/reg",
			beego.NSInclude(
				&controllers.AgentInfoController{},
			),
		),
		beego.NSNamespace("/container",
			beego.NSInclude(
				&controllers.ContainerController{},
			),
		),
		beego.NSNamespace("/image",
			beego.NSInclude(
				&controllers.ImageController{},
			),
		),
		beego.NSNamespace("/res",
			beego.NSInclude(
				&controllers.ResController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
