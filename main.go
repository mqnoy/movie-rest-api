package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/mqnoy/movie-rest-api/config"
	"github.com/mqnoy/movie-rest-api/handler"

	_movieHttpDelivery "github.com/mqnoy/movie-rest-api/movie/delivery/http"
	_movieRepoMysql "github.com/mqnoy/movie-rest-api/movie/repository/mysql"
	_movieUseCase "github.com/mqnoy/movie-rest-api/movie/usecase"
)

func main() {
	// TODO: supported with logger
	// TODO: supported with error handler

	// loading env
	cfg, err := config.Get(".")
	if err != nil {
		log.Fatalf("error while loading .env err: %v\n", err)
	}

	// database setup
	db := config.InitDatabase(cfg)

	// router setup
	g := gin.New()
	api := g.Group("/api")

	// middleware setup
	api.Use(gin.Recovery())

	// health check route
	api.GET("/health", handler.HealthCheck)

	// fallback route
	g.NoRoute(handler.FallbackHandler)

	// Initialize repository
	movieRepoMysql := _movieRepoMysql.New(db)

	// Initialize useCase
	movieUseCase := _movieUseCase.New(movieRepoMysql)

	// Initialize handler
	_movieHttpDelivery.New(api, movieUseCase)

	g.Run(cfg.ServerAddress())
}
