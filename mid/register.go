package mid

import "py.org/MyGoWeb/core"

func init() {
	indexMid := NewIndexMid()
	adminMid := NewAdminMid()

	core.HHandler.Mid.Register("/", &indexMid)
	core.HHandler.Mid.Register("/admin/**", &adminMid)
}
