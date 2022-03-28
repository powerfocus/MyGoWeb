package main

import (
	"py.org/MyGoWeb/core"
	_ "py.org/MyGoWeb/routers"
)

func main() {
	core.NewBootstrap().Start("localhost:8080")
}
