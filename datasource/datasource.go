package datasource

// DataSource defines the interface for loading data from a data source.
type DataSource interface {
	Get(key string) any
}
