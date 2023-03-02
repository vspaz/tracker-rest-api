package handlers

import (
	"github.com/go-chi/render"
	"net/http"
)

func (r *Router) Track(response http.ResponseWriter, request *http.Request) {
	render.Status(request, http.StatusOK)
	render.JSON(response, request, map[string]string{"status": "200 OK", "message": "OK"})
}

func (r *Router) Identify(response http.ResponseWriter, request *http.Request) {
	render.Status(request, http.StatusOK)
	render.JSON(response, request, map[string]string{"status": "200 OK", "message": "OK"})
}
