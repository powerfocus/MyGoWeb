package core

import (
	"net/http"
	"py.org/MyGoWeb/utils"
)

type Bootstrap struct{}

func NewBootstrap() *Bootstrap {
	return &Bootstrap{}
}

func (c *Bootstrap) Start(Addr string) {
	server := http.Server{Addr: Addr, Handler: Serv.Serv}
	err := server.ListenAndServe()
	utils.Common.Err(err)
}
