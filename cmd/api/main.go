// cmd/api/main.go

package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/gaurav4288/go_tutorial/blogging_app/internal/config"
	"github.com/gaurav4288/go_tutorial/blogging_app/internal/handlers"
	"github.com/gaurav4288/go_tutorial/blogging_app/internal/pkg/database"
	"github.com/gaurav4288/go_tutorial/blogging_app/internal/router"
)

func main() {
	// 1. Load config (env vars, DB creds, JWT secret, port)
	cfg := config.Load()

	// 2. Connect to Postgres
	db, err := database.NewPostgresDB(database.Config{
		Host:     cfg.DBHost,
		Port:     cfg.DBPort,
		User:     cfg.DBUser,
		Password: cfg.DBPassword,
		DBName:   cfg.DBName,
	})
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}
	defer db.Close()

	// 3. Wire up repositories
	// userRepo := repository.NewUserRepository(db)
	// postRepo := repository.NewPostRepository(db)

	// 4. Wire up services (business logic layer)
	// authService := service.UserService(userRepo)
	// userService := service.NewUserService(userRepo)
	// postService := service.NewPostService(postRepo)

	// 5. Wire up handlers (HTTP layer)
	authHandler := &handlers.AuthHandler{}
	// userHandler := handlers.NewUserHandler(userService)
	// postHandler := handlers.NewPostHandler(postService)

	// 6. Set gin mode based on environment
	if os.Getenv("APP_ENV") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// 7. Set up router with all handlers + middleware
	r := router.NewRouter(router.Handlers{
		Auth: authHandler,
	}, os.Getenv("JWT_SECRET"))

	// 8. Configure HTTP server
	srv := &http.Server{
		Addr:         ":" + cfg.DBPort,
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// 9. Start server in a goroutine so it doesn't block graceful shutdown handling
	go func() {
		log.Printf("server starting on port %s", cfg.DBPort)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("failed to start server: %v", err)
		}
	}()

	// 10. Graceful shutdown on interrupt/terminate signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("server forced to shutdown: %v", err)
	}

	log.Println("server exited gracefully")
}
