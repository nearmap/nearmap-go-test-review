package datasource

import (
	"errors"
	"sync"
)

type DataSource interface {
	Value(key string) (any, error)
}

type LocalDataSource struct {
	mutex *sync.Mutex
	data  map[string]any
	db    *Database
	cache *DistributedCache
}

func NewLocalDataSource(db *Database, cache *DistributedCache) *LocalDataSource {
	return &LocalDataSource{
		mutex: &sync.Mutex{},
		db:    db,
		cache: cache,
		data:  map[string]any{},
	}
}

func (lds *LocalDataSource) Value(key string) (any, error) {
	lds.mutex.Lock()
	result := lds.data[key]
	lds.mutex.Unlock()
	if result == nil {
		cacheResult, err := lds.cache.Value(key)
		if err == nil {
			if cacheResult == nil {
				dbResult, err := lds.db.Value(key)
				if err == nil {
					if dbResult == nil {
						return nil, errors.New("unknown key")
					} else {
						lds.cache.Store(key, dbResult)
						lds.mutex.Lock()
						lds.data[key] = dbResult
						lds.mutex.Unlock()
						result = dbResult
					}
				} else {
					return nil, errors.New("error retrieving value from db")
				}
			} else {
				lds.mutex.Lock()
				lds.data[key] = cacheResult
				lds.mutex.Unlock()
				result = cacheResult
			}
		} else {
			return nil, errors.New("error retrieving value from cache")
		}
	}
	return result, nil
}
