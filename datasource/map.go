package datasource

import "github.com/go-zoox/core-utils/object"

// mapDataSource is a tag data source that loads data from a map.
type mapDataSource struct {
	data map[string]any
}

// NewMapDataSource creates a new MapDataSource.
func NewMapDataSource(data map[string]any) DataSource {
	return &mapDataSource{data: data}
}

// Get returns the value of the given key.
// key support dot notation.
// Example:
//   - Get("port", "port")
//   - Get("redis.port", "port")
//   - Get("address.city.houses.0.id", "id")
func (m *mapDataSource) Get(path, key string) any {
	return object.Get(m.data, path)
}
