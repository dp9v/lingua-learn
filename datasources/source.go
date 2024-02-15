package datasources

type DataSource interface {
	ReadAllGroups() (*WordGroups, error)
}
