package filters

import (
	"log"
	"net/http"
)

// IndexFilter implements interface Filters
type IndexFilter struct{}

func (c *IndexFilter) DoFilter(w http.ResponseWriter, r *http.Request) {
	log.Println("IndexFilter.")
}
