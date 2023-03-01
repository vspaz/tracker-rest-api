package handlers

import (
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

type Router struct {
	Logger *logrus.Logger
	mux    *chi.Mux
}
