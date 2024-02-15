package datasources

type DummyDataSource struct {
}

func (d *DummyDataSource) ReadAllGroups() (*WordGroups, error) {
	return &Groups, nil
}

func NewDummyDataSource() DataSource {
	return &DummyDataSource{}
}
