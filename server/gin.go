package server

import "github.com/gin-gonic/gin"

var ginEngine *gin.Engine

func InitGinApp() {
	ginEngine = gin.Default()
}

func GetGinEngine() *gin.Engine {
	return ginEngine
}
