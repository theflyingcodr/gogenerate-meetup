package airplanes

import (
	"encoding/json"
	"fmt"
	api "github.com/theflyingcodr/gogenerate-meetup/05-api-gen"
	"net/http"
)

type httpHandler struct{
	svc AirplaneStorer
}

func NewHttpHandler(svc AirplaneStorer) *httpHandler {
	return &httpHandler{svc:svc}
}

func (h *httpHandler) Register(){
	http.HandleFunc(api.RouteAirplanesGET, h.Airplanes)
}

func (h *httpHandler) Airplanes(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet{
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
	w.Header().Add("content-type","application/json")

	aa, err := h.svc.Airplanes(r.Context())
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(aa); err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
	}
	return
}
