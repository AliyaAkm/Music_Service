package main

import (
	"PracticeCrud/internal/cache"
	"PracticeCrud/internal/middleware"
	"PracticeCrud/internal/migrations"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"

	"PracticeCrud/internal/config"
	httpHandler "PracticeCrud/internal/delivery/http"
	"PracticeCrud/internal/repo/postgres"
	usecase2 "PracticeCrud/internal/usecase"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println(".env not found, using environment variables")
	}

	cfg, err := config.ReadEnv()
	if err != nil {
		log.Fatal("Error reading config:", err)
	}
	migrations.Run(cfg.Db)
	db := NewDBConnection(cfg.Db)

	repo := postgres.NewSongRepo(db)
	songUS := usecase2.NewSongUsecase(repo)

	userRepo := postgres.NewUserRepo(db)
	userUsecase := usecase2.NewUserUsecase(userRepo)

	topSongRepo := postgres.NewTopSongRepo(db)
	rdb := cache.NewRedisCache(cfg.RedisAddr)
	topSongUS := usecase2.NewTopSongUsecase(topSongRepo, rdb)

	router := mux.NewRouter()

	router.Handle("/metrics", promhttp.Handler())

	apiRouter := router.PathPrefix("/").Subrouter()
	httpHandler.RegisterSongRoutes(apiRouter, songUS, topSongUS)

	authMiddleware := middleware.BasicAuth(userUsecase)
	apiRouter.Use(authMiddleware)
	apiRouter.Use(middleware.LoggerMiddleware)

	log.Println("Server running on port:", cfg.HTTPPort)
	if err := http.ListenAndServe(":"+cfg.HTTPPort, router); err != nil {
		log.Fatal("Failed to run server:", err)
	}
}
