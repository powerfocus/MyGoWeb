package core

import (
	"log"
	"net/http"
	"py.org/MyGoWeb/utils"
)

var (
	M Middle
)

func init() {
	M = NewMiddle()
}

type Middle struct {
	Middles map[string]Mid
}

func NewMiddle() Middle {
	return Middle{
		Middles: make(map[string]Mid),
	}
}

func (c *Middle) Has(key string) bool {
	_, ok := c.Middles[key]
	return ok
}

func (c *Middle) Register(key string, mid Mid) {
	if !c.Has(key) {
		c.Middles[key] = mid
	}
}

func (c *Middle) Remove(key string) {
	if c.Has(key) {
		delete(c.Middles, key)
	}
}

func (c *Middle) MiddleCount() int {
	var count int
	for _, _ = range c.Middles {
		count++
	}
	return count
}

func (c *Middle) Do(path string, w http.ResponseWriter, r *http.Request) {
	if c.MiddleCount() > 0 {
		log.Println("start process middle")
		for name, mid := range c.Middles {
			if utils.Common.AntCheck(path, name) {
				mid.DoMid(w, r)
				break
			}
		}
	}
}
