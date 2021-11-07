package airplanes

import (
	"context"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestAirplaneSvc_Airplanes(t *testing.T) {
	tests := map[string]struct {
		storeFunc func(ctx context.Context) ([]string, error)
		len       int
		expected  error
	}{
		"valid output should not return an error": {
			storeFunc: func(ctx context.Context) (strings []string, e error) {
				return []string{"I'm a plane, honest"}, nil
			},
			len: 1,
		}, "none found should return error": {
			storeFunc: func(ctx context.Context) (strings []string, e error) {
				return []string{}, nil
			},
			len:      0,
			expected: errors.New("none found mate"),
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			store := &AirplaneReaderMock{
				AirplanesFunc: test.storeFunc,
			}
			svc := newAirplaneSvc(store)
			aa, err := svc.Airplanes(context.Background())
			assert.Equal(t, test.len, len(aa))
			assert.Equal(t, 1, len(store.AirplanesCalls()))
			if test.expected == nil {
				assert.NoError(t, err)
				return
			}
			assert.Equal(t, test.expected.Error(), err.Error())
		})
	}
}
