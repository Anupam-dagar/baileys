package types

import "github.com/gin-gonic/gin"

type RouteFunc = func() (*gin.RouterGroup, func(rg *gin.RouterGroup))
