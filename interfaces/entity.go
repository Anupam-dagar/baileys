package interfaces

type Entity interface {
	SetCol(col string, val interface{}) error
	GetModel() interface{}
}
