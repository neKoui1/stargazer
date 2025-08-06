package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"stargazer/internal/config"
	"stargazer/internal/database/postgresql"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf(
			"main: failed to load config: %v",
			err,
		)
	}

	if !cfg.App.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	if err := postgresql.ConnectPostgresql(&cfg.Database); err != nil {
		log.Fatalf(
			"main: failed to connect to postgresql: %v",
			err,
		)
	}

	if err := postgresql.AutoMigrate(); err != nil {
		log.Fatalf(
			"main: failed to auto migrate: %v",
			err,
		)
	}

	r := gin.Default()
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.Server.Port),
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("main: failed to start server: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	sig := <-quit
	log.Printf("\nreceive signal: %v\n", sig)
	log.Println("Shut down server ... ")
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("HTTP Server shut down failure: ", err)
	}
	log.Println("Server exit successfully")
	postgresql.Close()
	log.Println("PostgreSQL connection closed")
}
