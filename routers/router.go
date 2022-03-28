package routers

import (
	"py.org/MyGoWeb/controller"
	"py.org/MyGoWeb/core"
	_ "py.org/MyGoWeb/filters"
	_ "py.org/MyGoWeb/mid"
)

func init() {
	/*core.Serv.Handle("/", indexController.Home)
	core.Serv.Handle("/user/form", indexController.Form)
	core.Serv.Handle("/user/act", indexController.Act)

	core.Serv.Handle("/admin", adminIndexController.Index)

	core.Serv.Handle("/admin/art", articleController.Index)*/

	indexController := controller.NewIndexController()
	adminIndexController := controller.NewAdminIndexController()
	articleController := controller.NewArticleController()
	R := core.NewRouter()
	defer R.Do()
	R.Register("/", indexController.Home)
	R.Register("/user/form", indexController.Form)
	R.Register("/user/act", indexController.Act)
	R.Register("/admin", adminIndexController.Index)
	R.Register("/admin/art", articleController.Index)
}
