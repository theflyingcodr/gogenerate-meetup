package main

import "context"

//go:generate 04-simple-mock-gen -type=AirplaneReader

// AirplaneReader will return all airplanes from a data source.
type AirplaneReader interface {
	Airplanes(ctx context.Context) ([]string, error)
}
