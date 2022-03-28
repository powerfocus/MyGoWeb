package main

import (
	"log"
	"net/http"
	"py.org/MyGoWeb/core"
	_ "py.org/MyGoWeb/routers"
)

func main() {
	server := http.Server{Addr: "localhost:8080", Handler: core.Serv.Serv}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
}
