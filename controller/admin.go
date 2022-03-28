package controller

import (
	"log"
	"net/http"
	"py.org/MyGoWeb/config"
	"py.org/MyGoWeb/core"
)

type AdminIndexController struct{}

func NewAdminIndexController() AdminIndexController {
	return AdminIndexController{}
}

func (c *AdminIndexController) Index(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		log.Fatalln("must GET.")
	}

	/*t, err := template.ParseFiles(config.TemplatesPath + "admin.html")
	if err != nil {
		log.Fatalln(err)
	}
	err = t.Execute(w, nil)
	if err != nil {
		log.Fatalln(err)
	}*/
	err := core.Template.Files(config.TemplatesPath+"admin.html").Map(w, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
