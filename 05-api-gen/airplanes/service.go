package airplanes

import (
	"context"
	"github.com/pkg/errors"
)

type airplaneSvc struct{
	store AirplaneStorer
}

func NewAirplaneSvc(store AirplaneStorer) *airplaneSvc {
	return &airplaneSvc{store: store}
}

func (a *airplaneSvc) Airplanes(ctx context.Context) ([]string, error){
	aa, err := a.store.Airplanes(ctx)
	if err != nil{
		return nil, errors.WithMessage(err, "oh no, it died")
	}
	if len(aa) == 0{
		return nil, errors.New("none found mate")
	}
	if len(aa) > 100{
		return nil, errors.New("I dunno why but i hate when there are more than 100")
	}
	return aa, nil
}
