package datasource

type Getter interface {
	Get(string) any
}

type getter2DataSource struct {
	get Getter
}

func (g *getter2DataSource) Get(path, key string) any {
	return g.get.Get(path)
}

func GetterToDataSource(s Getter) DataSource {
	return &getter2DataSource{s}
}
