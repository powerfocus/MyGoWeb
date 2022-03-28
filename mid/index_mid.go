package mid

import (
	"log"
	"net/http"
)

type IndexMid struct{}

func NewIndexMid() IndexMid {
	return IndexMid{}

}

func (c *IndexMid) DoMid(w http.ResponseWriter, r *http.Request) {
	// 中间件添加的操作
	log.Println("IndexMid Do")
}
