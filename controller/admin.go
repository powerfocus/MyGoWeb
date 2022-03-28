package controller

import (
	"log"
	"net/http"
	"py.org/MyGoWeb/config"
	"py.org/MyGoWeb/core"
	"py.org/MyGoWeb/utils"
)

type AdminIndexController struct{}

func NewAdminIndexController() AdminIndexController {
	return AdminIndexController{}
}

func (c *AdminIndexController) Index(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		log.Fatalln("must GET.")
	}

	err := core.Template.Files(config.TemplatesPath+"admin.html").Map(w, nil)
	utils.Common.Err(err)
}
