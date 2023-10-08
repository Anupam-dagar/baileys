# Defining an entity

An entity is a Go struct that represents a database table.

## Entity in baileys
Baileys operates under the hood with the GORM library and requires a GORM-compatible struct. To ensure your entity struct aligns with Baileys, adhere to the following criteria:
1. **Embed `BaseModel`**: Your struct must incorporate the `BaseModel` from the `entity` package in Baileys.
2. **Define a Pointer Type**: You should create a type for the struct, representing a pointer to the struct.
3. **Implementing `Entity` Interface**: Your entity struct must implement the methods specified in the `Entity` interface.

## Implementing `Entity` interface
The entity struct should feature two functions: `GetModel` and `SetCol` with the following implementations:
1. `GetModel` - return pointer to the struct
```go
func (me *MyEntity) GetModel() interface{} {
    return &MyEntity{}
}
```
2. `SetCol` - set the provided value to the provided column in the struct
```go
func (me *MyEntity) SetCol(field string, val interface{}) error {
	return reflections.SetField(p, field, val)
}
```

## Creating an entity
To craft an entity equipped with Baileys capabilities, your entity struct must meet the requisites of the `Entity` interface and embed `BaseModel`. Below is an example of crafting an entity for the `poll` entity from the tutorial:
```go
package entity

import (
	"github.com/Anupam-dagar/baileys/entity"
	"github.com/oleiade/reflections"
)

type Poll struct {
	entity.BaseModel
	Title string `gorm:"column:title" json:"title"`
}

type PollPtr = *Poll

func (p *Poll) GetModel() interface{} {
	return &Poll{}
}

func (p *Poll) SetCol(field string, val interface{}) error {
	return reflections.SetField(p, field, val)
}
```