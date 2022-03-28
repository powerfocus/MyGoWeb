package config

// PublicPath 公共文件目录
// TemplatesPath html模板文件目录
// Prefix 静态资源处理前缀
// Dir 静态资源目录
const (
	PublicPath    = "./public/"
	TemplatesPath = PublicPath + "./templates/"

	Prefix = "/static/"
	Dir    = PublicPath + "static"
)

const (
	SystemName = "MyGoWeb"
	Version    = 1.0
)
