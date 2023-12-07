package datasource

import (
	"sync"
	"time"
)

type DistributedCache struct {
	mutex sync.Mutex
	data  map[string]any
}

func NewDistributedCache() *DistributedCache {
	return &DistributedCache{
		data: map[string]any{},
	}
}

func (dc *DistributedCache) Value(key string) (any, error) {
	// simulate 100ms roundtrip to the distributed cache
	time.Sleep(100 * time.Millisecond)

	dc.mutex.Lock()
	result := dc.data[key]
	dc.mutex.Unlock()

	return result, nil
}

func (dc *DistributedCache) Store(key string, value any) error {
	// simulate 100ms roundtrip to the distributed cache
	time.Sleep(100 * time.Millisecond)

	dc.mutex.Lock()
	dc.data[key] = value
	dc.mutex.Unlock()

	return nil
}
