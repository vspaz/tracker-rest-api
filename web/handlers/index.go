package handlers

import (
	"net/http"
	"strconv"
)

func (r *Router) Index(response http.ResponseWriter, request *http.Request) {
	okBody := []byte("Ok")
	response.WriteHeader(http.StatusOK)
	response.Header().Set("Content-Length", strconv.Itoa(len(okBody)))
	_, err := response.Write(okBody)
	if err != nil {
		r.Logger.Errorf("error occurred: %s", err.Error())
	}
}
