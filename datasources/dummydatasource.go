package datasources

type DummyDataSource struct {
}

func (d *DummyDataSource) ReadAllGroups() *WordGroups {
	return &Groups
}

func (d *DummyDataSource) SaveGroups(groups []string) {
	//TODO implement me
	panic("implement me")
}

func NewDummyDataSource() DataSource {
	return &DummyDataSource{}
}
