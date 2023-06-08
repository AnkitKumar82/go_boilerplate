package controllers

import (
	"net/http"
)

type Route struct {
	Path    string
	Handler http.Handler
}

func InitialzeControllers(mux *http.ServeMux) {
	mux.Handle("/v1/health", LoggingMiddleware(CustomHeaderMiddleware(http.HandlerFunc(HealthGet))))
}
