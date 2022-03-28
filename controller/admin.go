package controller

import (
	"html/template"
	"log"
	"net/http"
	"py.org/MyGoWeb/config"
)

type AdminIndexController struct{}

func NewAdminIndexController() AdminIndexController {
	return AdminIndexController{}
}

func (c *AdminIndexController) Index(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		log.Fatalln("must GET.")
	}
	t, err := template.ParseFiles(config.TemplatesPath + "admin.html")
	if err != nil {
		log.Fatalln(err)
	}
	err = t.Execute(w, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
