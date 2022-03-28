package core

import (
	"log"
	"net/http"
	"py.org/MyGoWeb/utils"
	_ "py.org/MyGoWeb/utils"
)

type HttpHandle func(w http.ResponseWriter, r *http.Request)

// Filter 过滤器
type Filter struct {
	Mid     Middle
	Filters map[string]Filters
}

func NewFilter() Filter {
	log.Println("filters initialize")
	return Filter{
		Mid:     NewMiddle(),
		Filters: make(map[string]Filters),
	}
}

func (c *Filter) Has(key string) bool {
	_, ok := c.Filters[key]
	return ok
}

func (c *Filter) Register(key string, handle Filters) {
	if !c.Has(key) {
		c.Filters[key] = handle
	}
}

func (c *Filter) Remove(key string) {
	if c.Has(key) {
		delete(c.Filters, key)
	}
}

func (c *Filter) Handle(handle HttpHandle) HttpHandle {
	return func(w http.ResponseWriter, r *http.Request) {
		for path, h := range c.Filters {
			if utils.Common.AntCheck(r.RequestURI, path) {
				h.DoFilter(w, r)
				break
			}
		}
		c.Mid.Do(r.RequestURI, w, r)
		handle(w, r)
	}
}
