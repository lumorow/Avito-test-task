package app

import (
	"Avito-test-task/config"
	"Avito-test-task/internal/logger"
	rt "Avito-test-task/internal/router"
	repo "Avito-test-task/pkg/postgres"
	"context"
	"fmt"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run(configPath string) {
	cfg, err := config.NewConfig(configPath)
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Logger
	logger.SetLogrus(cfg.Logger.Level)

	// Repositories
	log.Info("Initializing postgres...")
	pg, err := repo.NewPostgresDB(cfg)
	if err != nil {
		log.Fatal(fmt.Errorf("app - Run - pgdb.NewServices: %w", err))
	}
	defer pg.Close()

	// Init routers
	log.Info("Initializing handlers and router...")
	router := rt.NewRouter()
	router.InitRoutes()

	// HTTP server
	srv := &http.Server{
		Addr:    cfg.HTTP.Port,
		Handler: router.Rtr,
	}

	log.Info("Starting http server...")
	log.Debugf("Server port: %s", cfg.HTTP.Port)
	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	}
	log.Println("Server exiting")
}
