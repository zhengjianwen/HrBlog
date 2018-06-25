package routers

import (
	"github.com/zhengjianwen/HrBlog/HrBlog/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
