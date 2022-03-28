package core

type Router struct {
	RMap map[string]HttpHandle
}

func NewRouter() Router {
	return Router{
		RMap: make(map[string]HttpHandle, 5),
	}
}

func (c *Router) Has(key string) bool {
	_, ok := c.RMap[key]
	return ok
}

func (c *Router) Register(key string, cl HttpHandle) {
	if !c.Has(key) {
		c.RMap[key] = cl
	}
}

func (c *Router) Remove(key string) {
	if c.Has(key) {
		delete(c.RMap, key)
	}
}

func (c *Router) Do() {
	for p, cl := range c.RMap {
		Serv.Handle(p, cl)
	}
}
