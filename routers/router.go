package routers

import (
	"py.org/MyGoWeb/controller"
	"py.org/MyGoWeb/core"
	_ "py.org/MyGoWeb/filters"
	_ "py.org/MyGoWeb/mid"
)

func init() {
	R := core.NewRouter()
	defer R.Do()

	indexController := controller.NewIndexController()
	adminIndexController := controller.NewAdminIndexController()
	articleController := controller.NewArticleController()
	R.Register("/", indexController.Home)
	R.Register("/user/form", indexController.Form)
	R.Register("/user/act", indexController.Act)
	R.Register("/admin", adminIndexController.Index)
	R.Register("/admin/art", articleController.Index)
}
