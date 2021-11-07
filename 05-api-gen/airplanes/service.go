// This code was created by a generator 2021-11-07T21:11:51Z
// CAUTION - If the generator is re-ran it will override the contents of this file

package airplanes

import (
	"context"
	"github.com/pkg/errors"
)

type airplaneSvc struct {
	rdr AirplaneReader
	wtr AirplaneWriter
}

// NewAirplane is the constructor used to setup the service.
func NewAirplaneSvc(rdrWtr AirplaneReaderWriter) *airplaneSvc {
	return &airplaneSvc{rdr: rdrWtr, wtr: rdrWtr}
}

// Airplanes will return all Airplanes currently stored.
func (s *airplaneSvc) Airplanes(ctx context.Context) ([]Airplane, error) {
	airplanes, err := s.rdr.Airplanes(ctx)
	if err != nil {
		return nil, errors.WithMessage(err, "oh no, it died")
	}
	// TODO add additional logic
	return airplanes, nil
}

// AirplaneCreate will return will validate and add a new airplane.
func (s *airplaneSvc) AirplaneCreate(ctx context.Context, req Airplane) (*Airplane, error) {
	if req.Name == "" {
		return nil, errors.New("Name must be supplied")
	}
	resp, err := s.wtr.AirplaneCreate(ctx, req)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to add")
	}
	// TODO add additional logic

	return resp, nil
}
