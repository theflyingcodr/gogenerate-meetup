// This code was created by a generator 2021-11-07T21:11:51Z
// CAUTION - If the generator is re-ran it will override the contents of this file

package airplanes

import "context"

// memoryStore is an in memory store that returns static data
// but could implment a full sql data store for example.
type memoryStore struct {
	airplanes []Airplane
}

// NewMemoryStore will setup and a return a new in-memory data store.
// This is useful for testing only.
func NewMemoryStore() *memoryStore {
	return &memoryStore{
		airplanes: []Airplane{
			{
				ID:   1,
				Name: "test",
			},
		},
	}
}

// Airplanes will return all Airplanes currently stored.
func (m *memoryStore) Airplanes(ctx context.Context) ([]Airplane, error) {
	return m.airplanes, nil
}

// Airplanes will add a new airplane to the data store.
func (m *memoryStore) AirplaneCreate(ctx context.Context, req Airplane) (*Airplane, error) {
	req.ID = len(m.airplanes) + 1
	m.airplanes = append(m.airplanes, req)
	return &req, nil
}
