package core

import "net/http"

// HttpHandler 请求处理必须实现此接口
type HttpHandler interface {
	Handle(w http.ResponseWriter, r *http.Request)
}

// Filters 过滤器必须实现此接口
type Filters interface {
	DoFilter(w http.ResponseWriter, r *http.Request)
}

// Mid 中间操作必须实现此接口
type Mid interface {
	DoMid(w http.ResponseWriter, r *http.Request)
}
