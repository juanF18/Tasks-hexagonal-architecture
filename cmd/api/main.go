package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"test-hex-architecture/internal/shared/config"
	"test-hex-architecture/internal/shared/db"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {

	if err := config.LoadEnv(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	ctx := context.Background()
	mongoResource, err := db.NewMongo(ctx)

	if err != nil {
		log.Fatalf("Mongo init %v", err)
	}
	defer mongoResource.Disconnect(ctx)

	// Init router
	r := gin.Default()

	// TODO: registar handlers aqui cuanod agreguemos el repositorio y casos de uso

	//Arrancar servidor
	addr := ":" + config.HTTPPort()
	SrvErr := make(chan error, 1)
	go func() { SrvErr <- r.Run(addr) }()

	// Señales de apagado
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err := <-SrvErr:
		log.Fatalf("Server error: %v", err)
	case sig := <-quit:
		log.Printf("Shutting down server... Reason: %v", sig)
		shutDownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		_ = mongoResource.Disconnect(shutDownCtx)
	}
}
