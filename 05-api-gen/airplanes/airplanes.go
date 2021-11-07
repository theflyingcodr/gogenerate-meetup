// This code was created by a generator 2021-11-07T21:11:51Z
// CAUTION - If the generator is re-ran it will override the contents of this file

package airplanes

import "context"

//go:generate 04-simple-mock-gen -type=airplaneReader
//go:generate 04-simple-mock-gen -type=airplaneWriter

// Airplane defines a single object, validators etc can be added to this.
type Airplane struct {
	ID   int
	Name string
	// TODO - add more properties
}

// AirplaneService will return all airplane from a data source.
type AirplaneService interface {
	Airplanes(ctx context.Context) ([]Airplane, error)
	AirplaneCreate(ctx context.Context, req Airplane) (*Airplane, error)
}

// AirplaneReader will return all airplane from a data source.
type AirplaneReader interface {
	Airplanes(ctx context.Context) ([]Airplane, error)
}

// AirplaneWriter will add and update airplanes to the data source.
type AirplaneWriter interface {
	// AirplaneCreate will add a single Airplane to a data source.
	AirplaneCreate(ctx context.Context, req Airplane) (*Airplane, error)
}

type AirplaneReaderWriter interface {
	AirplaneReader
	AirplaneWriter
}
