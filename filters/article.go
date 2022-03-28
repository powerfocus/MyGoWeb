package filters

import (
	"log"
	"net/http"
)

type ArticleFilter struct{}

func (c *ArticleFilter) DoFilter(w http.ResponseWriter, r *http.Request) {
	log.Println("ArticleFilter")
}
