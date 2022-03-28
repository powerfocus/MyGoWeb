package core

import (
	"log"
	"net/http"
)

type Bootstrap struct{}

func NewBootstrap() *Bootstrap {
	return &Bootstrap{}
}

func (c *Bootstrap) Start(Addr string) {
	server := http.Server{Addr: Addr, Handler: Serv.Serv}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
}
