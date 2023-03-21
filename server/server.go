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
	"syscall"
	"time"
)

func RunServer(
	ginEngine *gin.Engine,
) {
	database.InitDatabase()

	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	port := configuration.GetStringConfig("server.port")
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: ginEngine,
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
