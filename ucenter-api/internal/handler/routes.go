package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest"
)

type Routers struct {
	server     *rest.Server
	middleware []rest.Middleware
}

/**
 * 注册路由
 */
func NewRouters(server *rest.Server) *Routers {
	return &Routers{
		server: server,
	}
}

func (r *Routers) GET(path string, handlerFunc http.HandlerFunc) {
	r.server.AddRoutes(
		rest.WithMiddlewares(
			r.middleware,
			rest.Route{
				Method:  http.MethodGet,
				Path:    path,
				Handler: handlerFunc,
			},
		),
	)
}

func (r *Routers) POST(path string, handlerFunc http.HandlerFunc) {
	r.server.AddRoutes(
		rest.WithMiddlewares(
			r.middleware,
			rest.Route{
				Method:  http.MethodPost,
				Path:    path,
				Handler: handlerFunc,
			},
		),
	)
}

func (r *Routers) Group() *Routers {
	return &Routers{server: r.server}
}
