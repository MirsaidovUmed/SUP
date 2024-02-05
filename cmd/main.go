package main

import (
	"SUP/internal/repository"
	"SUP/internal/service"
	"SUP/internal/transport/http/handlers"
	"SUP/internal/transport/http/middleware"
	"SUP/internal/transport/http/router"
	"SUP/pkg/config"
	"SUP/pkg/database"
	"SUP/pkg/http"
	"SUP/pkg/logger"
)

func main() {
	conf := config.NewConfig()

	logger := logger.NewLogger(conf)

	db := database.NewDatabase(conf, logger)
	repo := repository.NewRepository(db, logger)
	svc := service.NewService(repo, conf, logger)

	handlers := handlers.NewHandler(svc, logger)

	middle := middleware.NewMiddleware(conf, svc)

	router := router.InitRouter(handlers, middle)

	server := http.NewServer(conf.ServerAddress, conf.ServerPort, router)

	server.Run()
}
