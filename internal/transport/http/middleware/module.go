package middleware

import (
	"SUP/internal/service"
	"SUP/pkg/config"
	"net/http"
)

type Middleware struct {
	config  *config.Config
	service service.ServiceInterface
}

type MiddlewareInterface interface {
	TimeDuration(next http.Handler) http.Handler
	JWT(next http.Handler) http.Handler
}

func NewMiddleware(config *config.Config, svc service.ServiceInterface) MiddlewareInterface {
	return &Middleware{
		config:  config,
		service: svc,
	}
}
