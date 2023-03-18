package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"net/http"
)

func (r *Router) handleSingleRequest(response http.ResponseWriter, request *http.Request) {
	writeKey := chi.URLParam(request, "writeKey")
	var eventBody Message
	if err := json.NewDecoder(request.Body).Decode(&eventBody); err != nil {
		r.Logger.Errorf("failed to parse payload %s", err)
		render.Status(request, http.StatusBadRequest)
		render.JSON(response, request, map[string]string{"status": "400", "message": "Bad Request"})
		return
	}
	r.saveMessage(writeKey, eventBody)
	render.Status(request, http.StatusOK)
	render.JSON(response, request, map[string]string{"status": "200 OK", "message": "OK"})
}

func saveBatch(eventBatch *Batch) map[string]string {
	if eventBatch == nil {
		return map[string]string{"status": "400", "message": "Bad Request"}
	}
	eventBatch = fixWriteKey(eventBatch)

	responseStub := 1
	switch responseStub {
	case 3:
		return map[string]string{"status": "500", "message": "Internal Server Error"}
	case 2:
		return map[string]string{"status": "400", "message": "Bad Request"}
	case 1:
		return map[string]string{"status": "200", "message": "Partial success"}
	case 0:
		return map[string]string{"status": "200 OK", "message": "OK"}
	default:
		return map[string]string{"status": "500", "message": "Internal Server Error"}
	}
}

func fixWriteKey(messageBatch *Batch) *Batch {
	if len(messageBatch.WriteKey) > 0 {
		return messageBatch
	}
	for _, event := range messageBatch.Batch {
		if len(event.WriteKey) > 0 {
			messageBatch.WriteKey = event.WriteKey
			return messageBatch
		}
	}
	return messageBatch
}

func (r *Router) saveMessage(writeKey string, eventBody Message) {
	eventBach := &Batch{
		Batch:   []Message{eventBody},
		Message: Message{WriteKey: writeKey},
	}
	saveBatch(eventBach)
}

func (r *Router) Track(response http.ResponseWriter, request *http.Request) {
	r.handleSingleRequest(response, request)
}

func (r *Router) Identify(response http.ResponseWriter, request *http.Request) {
	r.handleSingleRequest(response, request)
}

func (r *Router) Group(response http.ResponseWriter, request *http.Request) {
	r.handleSingleRequest(response, request)
}

func (r *Router) Alias(response http.ResponseWriter, request *http.Request) {
	r.handleSingleRequest(response, request)
}

func (r *Router) Page(response http.ResponseWriter, request *http.Request) {
	r.handleSingleRequest(response, request)
}

func (r *Router) Screen(response http.ResponseWriter, request *http.Request) {
	r.handleSingleRequest(response, request)
}

func (r *Router) handleBatch(response http.ResponseWriter, request *http.Request) {
	render.Status(request, http.StatusOK)
	render.JSON(response, request, map[string]string{"status": "200 OK", "message": "OK"})
}

func (r *Router) Batch(response http.ResponseWriter, request *http.Request) {
	r.handleBatch(response, request)
}

func (r *Router) Import(response http.ResponseWriter, request *http.Request) {
	r.handleBatch(response, request)
}
