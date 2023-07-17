package server

import (
	"context"
	"fmt"
	"github.com/Anupam-dagar/baileys/configuration"
	"github.com/Anupam-dagar/baileys/util/database"
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
	"net/http"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type ginEngine struct {
	engine *gin.Engine
}

type GinEngineInterface interface {
	InitGinApp(setupRoutes func()) GinEngineInterface
	RunServer()
	GetRootRouterGroup() *gin.RouterGroup
}

var ge *ginEngine
var geOnce sync.Once

func NewGinEngine() GinEngineInterface {
	geOnce.Do(func() {
		ge = new(ginEngine)
		ge.engine = gin.Default()
	})

	return ge
}

func GetGinEngine() GinEngineInterface {
	return ge
}

func (ge *ginEngine) InitGinApp(setupRoutes func()) GinEngineInterface {
	configuration.Init()
	database.InitDatabase()

	setupRoutes()

	return ge
}

func (ge *ginEngine) GetRootRouterGroup() *gin.RouterGroup {
	return ge.engine.Group(configuration.GetStringConfig("server.base_api_path"))
}

func (ge *ginEngine) RunServer() {
	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	ginRoutes := GetRoutes()
	for _, route := range ginRoutes {
		routerGroup, routeFunc := route()
		routeFunc(routerGroup)
	}

	port := configuration.GetStringConfig("server.port")
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: ge.engine,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		log.Infof("Starting server on port: %s", port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen error: %s\n", err)
		}
	}()

	// Listen for the interrupt signal.
	<-ctx.Done()

	// Restore default behavior on the interrupt signal and notify user of shutdown.
	stop()
	log.Info("shutting down gracefully, press Ctrl+C again to force")
	database.DisconnectDatabase()

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Error("Server forced to shutdown: ", err)
	}

	log.Info("Server exiting")
}
