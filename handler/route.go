package handler

import (
	"architecture-test/registry"
	"github.com/gin-gonic/gin"
)

type routes = []*route

type route struct {
	path    string
	method  string
	handler func(c *gin.Context, registry registry.Registry)
}

func RegistRoute(router *gin.Engine) {
	regi := registry.NewRegistryImpl()
	ctr := NewController()

	routes := routes{
		{path: "/users", method: "GET", handler: ctr.GetUsers},
	}

	for _, r := range routes {
		handler := func(c *gin.Context) {
			r.handler(c, regi)
		}

		switch r.method {
		case "GET":
			router.GET(r.path, handler)
		case "POST":
			router.POST(r.path, handler)
		case "PUT":
			router.PUT(r.path, handler)
		case "DELETE":
			router.DELETE(r.path, handler)
		}
	}
}
