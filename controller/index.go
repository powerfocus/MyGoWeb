package controller

import (
	"log"
	"net/http"
	"py.org/MyGoWeb/config"
	"py.org/MyGoWeb/core"
	"strings"
)

type IndexController struct{}

func NewIndexController() IndexController {
	return IndexController{}
}

func (c *IndexController) Home(w http.ResponseWriter, r *http.Request) {
	err := core.Template.Files(config.TemplatesPath+"index.html").Map(w, nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func (c *IndexController) Form(w http.ResponseWriter, r *http.Request) {
	err := core.Template.Files(config.TemplatesPath+"user.html").Map(w, nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func (c *IndexController) Act(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		log.Fatalln("must POST.")
	}
	err := r.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	userName := r.PostFormValue("username")
	pwd := r.PostFormValue("pwd")

	if strings.EqualFold("admin", userName) && strings.EqualFold("123", pwd) {
		http.Redirect(w, r, "/admin", http.StatusFound)
	} else {
		data := make(map[string]interface{}, 5)
		data["msg"] = "系统发生错误！登录失败，请确认用户名和密码后重试。"

		err := core.Template.Files(config.TemplatesPath+"err.html").Map(w, data)
		if err != nil {
			log.Fatalln(err)
		}
	}
}
