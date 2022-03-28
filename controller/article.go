package controller

import (
	"html/template"
	"log"
	"net/http"
	"py.org/MyGoWeb/config"
)

type ArticleController struct{}

func NewArticleController() ArticleController {
	return ArticleController{}
}

func (c *ArticleController) Index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(config.TemplatesPath + "article.html")
	if err != nil {
		log.Fatalln(err)
	}
	_ = t.Execute(w, nil)
}
