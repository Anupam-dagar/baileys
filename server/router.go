package server

import "github.com/Anupam-dagar/baileys/constant/types"

var routes []types.RouteFunc

func AddRoute(route types.RouteFunc) {
	routes = append(routes, route)
}

func GetRoutes() []types.RouteFunc {
	return routes
}
