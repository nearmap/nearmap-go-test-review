package datasource

import (
	"sync"
	"time"
)

type Database struct {
	mutex sync.Mutex
	data  map[string]any
}

func NewDatabase(initialState map[string]any) *Database {
	return &Database{
		data: initialState,
	}
}

func (db *Database) Value(key string) (any, error) {
	// simulate 500ms roundtrip to the database
	time.Sleep(500 * time.Millisecond)

	db.mutex.Lock()
	result := db.data[key]
	db.mutex.Unlock()

	return result, nil
}

func (db *Database) Store(key string, value any) error {
	// simulate 500ms roundtrip to the database
	time.Sleep(500 * time.Millisecond)

	db.mutex.Lock()
	db.data[key] = value
	db.mutex.Unlock()

	return nil
}
