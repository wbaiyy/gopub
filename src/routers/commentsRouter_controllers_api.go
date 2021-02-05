package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["controllers/api:ProjectController"] = append(beego.GlobalControllerRouter["controllers/api:ProjectController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			Params:           nil})

	beego.GlobalControllerRouter["controllers/api:ProjectController"] = append(beego.GlobalControllerRouter["controllers/api:ProjectController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params:           nil})

	beego.GlobalControllerRouter["controllers/api:ProjectController"] = append(beego.GlobalControllerRouter["controllers/api:ProjectController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			Params:           nil})

	beego.GlobalControllerRouter["controllers/api:ProjectController"] = append(beego.GlobalControllerRouter["controllers/api:ProjectController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params:           nil})

	beego.GlobalControllerRouter["controllers/api:ProjectController"] = append(beego.GlobalControllerRouter["controllers/api:ProjectController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params:           nil})

	beego.GlobalControllerRouter["controllers/api:TaskController"] = append(beego.GlobalControllerRouter["controllers/api:TaskController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			Params:           nil})

	beego.GlobalControllerRouter["controllers/api:TaskController"] = append(beego.GlobalControllerRouter["controllers/api:TaskController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			Params:           nil})

	beego.GlobalControllerRouter["controllers/api:TaskController"] = append(beego.GlobalControllerRouter["controllers/api:TaskController"],
		beego.ControllerComments{
			Method:           "GetAllAndProName",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			Params:           nil})

	beego.GlobalControllerRouter["controllers/api:TaskController"] = append(beego.GlobalControllerRouter["controllers/api:TaskController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			Params:           nil})

	beego.GlobalControllerRouter["controllers/api:TaskController"] = append(beego.GlobalControllerRouter["controllers/api:TaskController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			Params:           nil})

	beego.GlobalControllerRouter["controllers/api:TaskController"] = append(beego.GlobalControllerRouter["controllers/api:TaskController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			Params:           nil})

	beego.GlobalControllerRouter["controllers/api:TokenController"] = append(beego.GlobalControllerRouter["controllers/api:TokenController"],
		beego.ControllerComments{
			Method:           "IssueToken",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			Params:           nil})

}
