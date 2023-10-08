package server

import (
	"context"
	"github.com/Anupam-dagar/baileys/constant/types"
	"github.com/gin-gonic/gin"
)

var routes []types.RouteFunc

func AddRoute(route types.RouteFunc) {
	routes = append(routes, route)
}

func GetRoutes() []types.RouteFunc {
	return routes
}

func SetToContext(ctx *gin.Context, key string, value any) {
	ctx.Request = ctx.Request.WithContext(context.WithValue(ctx.Request.Context(), key, value))
}
