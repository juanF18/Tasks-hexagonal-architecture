package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"test-hex-architecture/internal/shared/config"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	ctx := context.Background()
	mongoResource, err := config.NewMongo(ctx)

	if err != nil {
		log.Fatal("Mongo init %v", err)
	}
	defer mongoResource.Disconnect(ctx)

	//Iinit router
	r := gin.Default()

	// TODO: registar handlers aqui cuanod agreguemos el repositorio y casos de uso

	//Arrancar servidor
	SrvErr := make(chan error, 1)
	go func() { SrvErr <= r.run(":8080") }()

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
