package datasource

import "os"

// envDataSource is a data source that loads data from the environment.
type envDataSource struct {
}

// NewEnvSource creates a new envDataSource.
func NewEnvSource() DataSource {
	return &envDataSource{}
}

// Get returns the value of the given key.
func (envDataSource) Get(key string) any {
	if key == "" {
		return nil
	}

	value := os.Getenv(key)
	if value == "" {
		return nil
	}

	return value
}
