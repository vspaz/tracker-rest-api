package handlers

import (
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"github.com/vspaz/tracker-rest-api/config"
)

type Router struct {
	Logger *logrus.Logger
	Conf   *config.Conf
	mux    *chi.Mux
}

func NewRouter(conf *config.Conf, logger *logrus.Logger) *Router {
	return &Router{
		Logger: logger,
		Conf:   conf,
		mux:    chi.NewRouter(),
	}
}

func (r *Router) RegisterHandlers() {
	r.mux.Get("/ping", r.GetHealthStatus)
	r.mux.Post("/ping", r.GetHealthStatus)
}
