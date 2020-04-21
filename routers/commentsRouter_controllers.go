package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["draw/controllers:DrawController"] = append(beego.GlobalControllerRouter["draw/controllers:DrawController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["draw/controllers:DrawController"] = append(beego.GlobalControllerRouter["draw/controllers:DrawController"],
        beego.ControllerComments{
            Method: "Edit",
            Router: `/Edit`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["draw/controllers:DrawController"] = append(beego.GlobalControllerRouter["draw/controllers:DrawController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/GetAll`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["draw/controllers:DrawController"] = append(beego.GlobalControllerRouter["draw/controllers:DrawController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/Login`,
            AllowHTTPMethods: []string{"POST"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["draw/controllers:DrawController"] = append(beego.GlobalControllerRouter["draw/controllers:DrawController"],
        beego.ControllerComments{
            Method: "SaveConfig",
            Router: `/SaveConfig`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["draw/controllers:DrawController"] = append(beego.GlobalControllerRouter["draw/controllers:DrawController"],
        beego.ControllerComments{
            Method: "GetInfo",
            Router: `/getInfo`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
