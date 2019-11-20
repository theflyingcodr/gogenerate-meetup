package main

import "context"

//go:generate 04-simple-mock-gen -type=AirplaneStorer

// AirplaneStorer will return all airplanes from a data store.
type AirplaneStorer interface{
	Airplane(ctx context.Context, id int64) (string, error)
	Airplanes(ctx context.Context) ([]string, error)
}
