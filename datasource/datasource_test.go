package datasource

import (
	"reflect"
	"testing"
)

func TestLocalDataSource(t *testing.T) {
	db := NewDatabase(map[string]any{
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
	cache := NewDistributedCache()
	dataSource := NewLocalDataSource(db, cache)

	tcs := []struct {
		desc      string
		lds       *LocalDataSource
		key       string
		expected  any
		expectErr bool
	}{
		{
			desc:      "get value from cache",
			lds:       dataSource,
			key:       "key0",
			expected:  "value0",
			expectErr: false,
		},
		{
			desc:      "key doesn't exist",
			lds:       dataSource,
			key:       "key10",
			expected:  nil,
			expectErr: true,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.desc, func(t *testing.T) {
			val, err := tc.lds.Value(tc.key)
			if (err != nil) != tc.expectErr {
				t.Fatalf("expected error to be %v, was %v", tc.expectErr, err)
				return
			}
			if !reflect.DeepEqual(val, tc.expected) {
				t.Errorf("expected value to be %v, was %v", tc.expected, val)
			}
		})
	}
}
