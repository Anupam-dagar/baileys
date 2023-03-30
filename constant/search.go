package constant

const (
	AndOperatorDelimiter = ";"
	Root                 = "root"
	MaxPageSize          = 1000
	DefaultPageSize      = 100
	DefaultSortOrder     = "ASC"
)

type SortOrder string

const (
	ASC  SortOrder = "ASC"
	DESC SortOrder = "DESC"
)
