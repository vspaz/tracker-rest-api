package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"net/http"
)

func (r *Router) handleSingleRequest(response http.ResponseWriter, request *http.Request) {
	writeKey := chi.URLParam(request, "writeKey")
	var eventBody Event
	if err := json.NewDecoder(request.Body).Decode(&eventBody); err != nil {
		r.Logger.Errorf("failed to parse payload %s", err)
		render.Status(request, http.StatusBadRequest)
		render.JSON(response, request, map[string]string{"status": "400", "message": "Bad Request"})
		return
	}
	println(writeKey, eventBody)
	render.Status(request, http.StatusOK)
	render.JSON(response, request, map[string]string{"status": "200 OK", "message": "OK"})
}

func (r *Router) Track(response http.ResponseWriter, request *http.Request) {
	render.Status(request, http.StatusOK)
	render.JSON(response, request, map[string]string{"status": "200 OK", "message": "OK"})
}

func (r *Router) Identify(response http.ResponseWriter, request *http.Request) {
	render.Status(request, http.StatusOK)
	render.JSON(response, request, map[string]string{"status": "200 OK", "message": "OK"})
}

func (r *Router) Group(response http.ResponseWriter, request *http.Request) {
	render.Status(request, http.StatusOK)
	render.JSON(response, request, map[string]string{"status": "200 OK", "message": "OK"})
}

func (r *Router) Alias(response http.ResponseWriter, request *http.Request) {
	render.Status(request, http.StatusOK)
	render.JSON(response, request, map[string]string{"status": "200 OK", "message": "OK"})
}

func (r *Router) Page(response http.ResponseWriter, request *http.Request) {
	render.Status(request, http.StatusOK)
	render.JSON(response, request, map[string]string{"status": "200 OK", "message": "OK"})
}

func (r *Router) Screen(response http.ResponseWriter, request *http.Request) {
	render.Status(request, http.StatusOK)
	render.JSON(response, request, map[string]string{"status": "200 OK", "message": "OK"})
}

func (r *Router) Batch(response http.ResponseWriter, request *http.Request) {
	render.Status(request, http.StatusOK)
	render.JSON(response, request, map[string]string{"status": "200 OK", "message": "OK"})
}

func (r *Router) Import(response http.ResponseWriter, request *http.Request) {
	render.Status(request, http.StatusOK)
	render.JSON(response, request, map[string]string{"status": "200 OK", "message": "OK"})
}
