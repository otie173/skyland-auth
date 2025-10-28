package main

import (
	"log"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/otie173/skyland-auth/api/handler"
	"github.com/otie173/skyland-auth/api/router"
	"github.com/otie173/skyland-auth/internal/config"
	"github.com/otie173/skyland-auth/internal/domain/services"
	"github.com/otie173/skyland-auth/internal/infrastructure/postgresql"
	"github.com/otie173/skyland-auth/internal/infrastructure/redis"
	"github.com/otie173/skyland-auth/internal/pkg/server"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Error! Problems with loading config: %v\n", err)
	}
	log.Println(cfg)

	db, err := postgresql.NewConnection(cfg)
	if err != nil {
		log.Fatalf("Error! Problems with connecting to database: %v\n", err)
	}
	userRepo := postgresql.NewUserRepository(db)

	cache := redis.NewClient(cfg)
	tokenRepo := redis.NewTokenRepository(cache)

	authService := services.NewAuthService(userRepo, tokenRepo)

	handler := handler.New(authService)
	router := router.New(handler)

	router.Use(middleware.Logger)
	router.SetupRoutes()

	server := server.New(cfg.Address, handler, router)
	server.Run()
}
