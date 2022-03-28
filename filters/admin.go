package filters

import (
	"log"
	"net/http"
)

type AdminFilter struct{}

func (c *AdminFilter) DoFilter(w http.ResponseWriter, r *http.Request) {
	log.Println("AdminFilter.")
}
