package mid

import (
	"log"
	"net/http"
)

type AdminMid struct{}

func NewAdminMid() AdminMid {
	return AdminMid{}
}

func (c *AdminMid) DoMid(w http.ResponseWriter, r *http.Request) {
	log.Println("admin middle")
}
