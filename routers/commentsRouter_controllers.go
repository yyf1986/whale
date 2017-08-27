package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["whale/controllers:AgentInfoController"] = append(beego.GlobalControllerRouter["whale/controllers:AgentInfoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["whale/controllers:AgentInfoController"] = append(beego.GlobalControllerRouter["whale/controllers:AgentInfoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["whale/controllers:AgentInfoController"] = append(beego.GlobalControllerRouter["whale/controllers:AgentInfoController"],
		beego.ControllerComments{
			Method: "Get4Web",
			Router: `/get4web`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
