package datasources

type DataSource interface {
	ReadAllGroups() *WordGroups
	SaveGroups(groups []string)
}
