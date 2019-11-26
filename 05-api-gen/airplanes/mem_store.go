package airplanes

import "context"

type MemoryStore struct{

}

func (m *MemoryStore) Airplanes(ctx context.Context) ([]string, error){
	return []string{"Boeing 737", "Boeing 747", "Cessna 182", "Jabiru UL-450"}, nil
}
