package datasource

import "github.com/go-zoox/core-utils/object"

// MapDataSource is a tag data source that loads data from a map.
type MapDataSource struct {
	data map[string]any
}

// NewMapDataSource creates a new MapDataSource.
func NewMapDataSource(data map[string]any) *MapDataSource {
	return &MapDataSource{data: data}
}

// Get returns the value of the given key.
// key support dot notation.
// Example:
//  - Get("port")
//  - Get("redis.port")
//  - Get("address.city.houses.0.id")
func (m *MapDataSource) Get(key string) any {
	return object.Get(m.data, key)
}
