# Defining routes

In Baileys, defining routes closely aligns with the standard route definition process in Gin. However, Baileys introduces an additional layer of functionality, requiring the inclusion of a special function to register routes seamlessly within the framework.

## Setting up route handler
A route handler serves as the central hub for your routes. It provides router groups for individual routes and takes care of the essential task of registering these routes in Baileys. Below is an example of a route handler file:
```go
// route.go
package route

import "github.com/Anupam-dagar/baileys/server"

func SetupRoutes() {
	// Retrieve the root router group
	rootRouterGroup := server.GetGinEngine().GetRootRouterGroup()

	// Add routes using specific functions (e.g., BRoute) defined in each route file
	server.AddRoute(BRoute(rootRouterGroup))
}
```

## Defining routes
To register a route, you must define a function that accepts a router group and returns a function of type `RouterFunc`. This returned function is used to register the route in Baileys via the previously mentioned route handler. Below is an example illustrating the route definition for the `poll` entity, as presented in the tutorial:
```go
package route

import (
	"github.com/Anupam-dagar/baileys/constant/types"
	"github.com/gin-gonic/gin"
	"polls/controller"
)

// PollRoutes defines the routes for the 'poll' entity, including your custom API endpoints supported by the Gin framework, along with the corresponding controller methods to be executed.
func PollRoutes(routerGroup *gin.RouterGroup) {
	router := routerGroup.Group("/polls")
	{
		controller.NewPollController(router)
	}
}

// BPollRoutes returns a function that, in turn, returns the router group for this route and the function defining the API endpoints. This function is utilized to register the route within Baileys.
func BPollRoutes(routerGroup *gin.RouterGroup) types.RouteFunc {
	return func() (*gin.RouterGroup, func(rg *gin.RouterGroup)) {
		return routerGroup, PollRoutes
	}
}
```