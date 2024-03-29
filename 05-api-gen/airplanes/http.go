// This code was created by a generator 2021-11-07T21:11:51Z
// CAUTION - If the generator is re-ran it will override the contents of this file

package airplanes

import (
	"encoding/json"
	"fmt"
	api "github.com/theflyingcodr/gogenerate-meetup/05-api-gen"
	"net/http"
)

type httpHandler struct {
	svc AirplaneService
}

// NewHttpHandler will setup and return an http handler for Airplanes.
func NewHttpHandler(svc AirplaneService) *httpHandler {
	return &httpHandler{svc: svc}
}

// Register will setup rest endpoints.
func (h *httpHandler) Register() {
	http.HandleFunc(api.RouteAirplanes, h.Airplanes)
}

// Airplanes will handle http requests for Airplanes.
func (h *httpHandler) Airplanes(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	switch r.Method {
	case http.MethodGet:
		h.airplanes(w, r)
	case http.MethodPost:
		h.create(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

	return
}

func (h *httpHandler) airplanes(w http.ResponseWriter, r *http.Request) {
	resp, err := h.svc.Airplanes(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
	}
}

// create will handle http requests for adding Airplanes.
func (h *httpHandler) create(w http.ResponseWriter, r *http.Request) {
	var req Airplane
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}
	resp, err := h.svc.AirplaneCreate(r.Context(), req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
		return
	}
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
	}
	return
}
