package datasource

// DataSource defines the interface for loading data from a data source.
type DataSource interface {
	// Get returns the value of the given key.
	// key support dot notation.
	// Example:
	//  - Get("port")
	//  - Get("redis.port")
	//  - Get("address.city.houses.0.id")
	Get(key string) any
}
