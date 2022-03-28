package core

import (
	"html/template"
	"io"
	"py.org/MyGoWeb/utils"
)

var (
	Template Tpl
)

func init() {
	Template = NewTpl()
}

type Tpl struct {
	t       *template.Template
	DataMap map[string]interface{}
}

func NewTpl() Tpl {
	return Tpl{
		DataMap: make(map[string]interface{}, 10),
	}
}

func (c *Tpl) Files(filenames ...string) *Tpl {
	t, err := template.ParseFiles(filenames...)
	c.t = t
	utils.Common.Err(err)
	return c
}

func (c *Tpl) Map(w io.Writer, data map[string]interface{}) error {
	return c.t.Execute(w, data)
}
