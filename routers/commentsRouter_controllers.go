package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["smark_api/controllers:IndexController"] = append(beego.GlobalControllerRouter["smark_api/controllers:IndexController"],
		beego.ControllerComments{
			Method: "GetAllVideo",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["smark_api/controllers:ObjectController"] = append(beego.GlobalControllerRouter["smark_api/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["smark_api/controllers:ObjectController"] = append(beego.GlobalControllerRouter["smark_api/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["smark_api/controllers:ObjectController"] = append(beego.GlobalControllerRouter["smark_api/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["smark_api/controllers:ObjectController"] = append(beego.GlobalControllerRouter["smark_api/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["smark_api/controllers:ObjectController"] = append(beego.GlobalControllerRouter["smark_api/controllers:ObjectController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:objectId`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["smark_api/controllers:TabController"] = append(beego.GlobalControllerRouter["smark_api/controllers:TabController"],
		beego.ControllerComments{
			Method: "GetAllTab",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["smark_api/controllers:UserController"] = append(beego.GlobalControllerRouter["smark_api/controllers:UserController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["smark_api/controllers:UserController"] = append(beego.GlobalControllerRouter["smark_api/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["smark_api/controllers:UserController"] = append(beego.GlobalControllerRouter["smark_api/controllers:UserController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["smark_api/controllers:UserController"] = append(beego.GlobalControllerRouter["smark_api/controllers:UserController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["smark_api/controllers:UserController"] = append(beego.GlobalControllerRouter["smark_api/controllers:UserController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:uid`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["smark_api/controllers:UserController"] = append(beego.GlobalControllerRouter["smark_api/controllers:UserController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["smark_api/controllers:UserController"] = append(beego.GlobalControllerRouter["smark_api/controllers:UserController"],
		beego.ControllerComments{
			Method: "Logout",
			Router: `/logout`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["smark_api/controllers:VideoController"] = append(beego.GlobalControllerRouter["smark_api/controllers:VideoController"],
		beego.ControllerComments{
			Method: "GetAllVideo",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["smark_api/controllers:VideoController"] = append(beego.GlobalControllerRouter["smark_api/controllers:VideoController"],
		beego.ControllerComments{
			Method: "GetVideo",
			Router: `/:videoId`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
