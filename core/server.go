package core

import (
	"log"
	"net/http"
	"py.org/MyGoWeb/config"
)

type Server struct {
	Serv *http.ServeMux
}

func NewServer() Server {
	log.Println("server initialize")
	s := Server{
		Serv: http.NewServeMux(),
	}
	s.Serv.Handle(config.Prefix, http.StripPrefix(config.Prefix, http.FileServer(http.Dir(config.Dir))))
	return s
}

func (c *Server) Handle(path string, handle HttpHandle) {
	c.Serv.HandleFunc(path, HHandler.Handle(handle))
}
