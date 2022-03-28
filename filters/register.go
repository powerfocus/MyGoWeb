package filters

import (
	"py.org/MyGoWeb/core"
)

func init() {
	core.HHandler.Register("/", &IndexFilter{})
	core.HHandler.Register("/admin/**", &AdminFilter{})
	core.HHandler.Register("/admin/art", &ArticleFilter{})
}
