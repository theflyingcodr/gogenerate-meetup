package main

import (
	"net/http"

	"github.com/theflyingcodr/gogenerate-meetup/05-api-gen/airplanes"
)

func main() {
	airplanes.NewHttpHandler(airplanes.NewAirplaneSvc(airplanes.NewMemoryStore())).Register()
	http.ListenAndServe(":8080", nil)
}
