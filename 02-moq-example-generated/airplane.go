package airplanes

import "context"

// AirplaneReader will return all airplanes from a data source.
type AirplaneReader interface {
	Airplanes(ctx context.Context) ([]string, error)
}
