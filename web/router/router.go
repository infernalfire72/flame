package router

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

type webRouter struct {
	*router.Router
}

func NewRouter() webRouter {
	return webRouter{router.New()}
}

func (r webRouter) Get(path string, next func(WebCtx)) {
	r.GET(path, func(ctx *fasthttp.RequestCtx) {
		next(WebCtx{ctx})
	})
}

func (r webRouter) Post(path string, next func(WebCtx)) {
	r.POST(path, func(c *fasthttp.RequestCtx) {
		next(WebCtx{c})
	})
}

func (r *webRouter) Group(name string) PathGroup {
	return PathGroup{
		Name:   name,
		router: r,
	}
}

type PathGroup struct {
	Name   string
	router *webRouter
}

func (g *PathGroup) Get(path string, cb func(WebCtx)) {
	g.router.Get(g.Name+path, cb)
}

func (g *PathGroup) Post(path string, cb func(WebCtx)) {
	g.router.Post(g.Name+path, cb)
}

type WebCtx struct {
	*fasthttp.RequestCtx
}

func (c *WebCtx) Error(reason string) {
	c.WriteString("error: " + reason)
}
