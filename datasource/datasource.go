package datasource

// DataSource defines the interface for loading data from a data source.
type DataSource interface {
	// Get returns the value of the given key.
	// key support dot notation.
	// Example:
	//  - Get("port", "port")
	//  - Get("redis.port", "port")
	//  - Get("address.city.houses.0.id", "id")
	Get(path string, key string) any
}
