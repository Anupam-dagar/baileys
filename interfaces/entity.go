package interfaces

type Entity interface {
	SetCol(col string, val interface{}) error
	ValidateFields(fields []string) bool
	GetModel() interface{}
	GetStructData() (map[string]interface{}, error)
}
