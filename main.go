package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/nearmap/nearmap-go-test/datasource"
)

func main() {
	db := datasource.NewDatabase(map[string]any{
		"key0": "value0",
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
		"key4": "value4",
		"key5": "value5",
		"key6": "value6",
		"key7": "value7",
		"key8": "value8",
		"key9": "value9",
	})
	cache := datasource.NewDistributedCache()

	dataSource := datasource.NewLocalDataSource(db, cache)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 50; j++ {
				key := fmt.Sprintf("key%d", rand.Intn(10))
				start := time.Now()
				value, err := dataSource.Value(key)
				if err != nil {
					fmt.Printf("Received error")
				} else {
					fmt.Println("Request", key, "response", value.(string), "time:", time.Since(start).Milliseconds(), "ms")
				}
			}
		}()
	}

	wg.Wait()
}
