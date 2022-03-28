package controller

import (
	"net/http"
	"py.org/MyGoWeb/config"
	"py.org/MyGoWeb/core"
	"py.org/MyGoWeb/utils"
)

type ArticleController struct{}

func NewArticleController() ArticleController {
	return ArticleController{}
}

func (c *ArticleController) Index(w http.ResponseWriter, r *http.Request) {
	err := core.Template.Files(config.TemplatesPath+"article.html").Map(w, nil)
	utils.Common.Err(err)
}
