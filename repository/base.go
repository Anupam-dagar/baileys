package repository

import (
	"context"
	"github.com/Anupam-dagar/baileys/interfaces"
	"github.com/Anupam-dagar/baileys/util"
	"github.com/Anupam-dagar/baileys/util/database"
	"gorm.io/gorm"
)

type BaseRepository[T interfaces.Entity] interface {
	GetById(ctx context.Context, id string) (T, error)
	Create(ctx context.Context, entity *T) error
	Update(ctx context.Context, id string, entity *T) error
	Delete(ctx context.Context, id string) error
	Search(ctx context.Context) ([]T, error)
}

type baseRepository[T interfaces.Entity] struct {
	db        *gorm.DB
	repoModel interface{}
}

func NewBaseRepository[T interfaces.Entity]() BaseRepository[T] {
	br := new(baseRepository[T])
	br.db = database.GetDatabase()

	var tModel T
	br.repoModel = tModel.GetModel()

	return br
}

func (br *baseRepository[T]) GetById(ctx context.Context, id string) (res T, err error) {
	txn := br.db.Debug().WithContext(ctx).Model(&res)

	err = txn.Scopes(database.ColumnValEqual("id", id)).First(&res).Error

	return res, err
}

func (br *baseRepository[T]) Create(ctx context.Context, entity *T) (err error) {
	txn := br.db.Debug().WithContext(ctx).Model(entity)

	return txn.Create(entity).Error
}

func (br *baseRepository[T]) Update(ctx context.Context, id string, entity *T) (err error) {
	txn := br.db.Debug().WithContext(ctx).Model(entity)

	return txn.Scopes(database.ColumnValEqual("id", id)).Updates(entity).Error
}

func (br *baseRepository[T]) Delete(ctx context.Context, id string) (err error) {
	txn := br.db.Debug().WithContext(ctx).Model(br.repoModel)

	return util.SoftDeleteById(ctx, txn, id)
}

func (br *baseRepository[T]) Search(ctx context.Context) (data []T, err error) {
	txn := br.db.Debug().WithContext(ctx).Model(br.repoModel)

	err = txn.Find(&data).Error

	return data, err
}
