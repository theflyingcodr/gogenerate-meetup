package airplanes

import "context"

//go:generate 04-simple-mock-gen -type=AirplaneStorer

// AirplaneStorer will return all airplanes from a data store.
type AirplaneStorer interface{
	Airplanes(ctx context.Context) ([]string, error)
}

