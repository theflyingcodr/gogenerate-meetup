package airplanes

import "context"

// AirplaneStorer will return all airplanes from a data store.
type AirplaneStorer interface{
	Airplanes(ctx context.Context) ([]string, error)
}
