package constant

const (
	Root             = "root"
	MinPageSize      = -1
	MaxPageSize      = 1000
	DefaultPageSize  = 100
	DefaultSortOrder = "ASC"
)

const (
	AndOperatorDelimiter    = ";"
	KeyValueDelimiter       = ":"
	SearchOperatorDelimiter = "."
	CommaDelimiter          = ","
	ColumnDelimiter         = "-"
)

type SearchOperator string

const (
	Equal              SearchOperator = "eq"
	NotEqual           SearchOperator = "neq"
	IsNull             SearchOperator = "isNull"
	IsNotNull          SearchOperator = "isNotNull"
	GreaterThan        SearchOperator = "gt"
	GreaterThanEqualTo SearchOperator = "gte"
	LessThan           SearchOperator = "lt"
	LessThanEqualTo    SearchOperator = "lte"
	Like               SearchOperator = "like"
	In                 SearchOperator = "in"
	NotIn              SearchOperator = "nin"
)
