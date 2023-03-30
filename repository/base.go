package repository

import (
	"context"
	"fmt"
	"github.com/Anupam-dagar/baileys/constant"
	"github.com/Anupam-dagar/baileys/interfaces"
	"github.com/Anupam-dagar/baileys/util"
	"github.com/Anupam-dagar/baileys/util/database"
	"github.com/Anupam-dagar/baileys/util/database/query_builder"
	"github.com/Anupam-dagar/baileys/util/search"
	"gorm.io/gorm"
	"strings"
)

type BaseRepository[T interfaces.Entity] interface {
	GetById(ctx context.Context, id string) (T, error)
	Create(ctx context.Context, entity *T) error
	Update(ctx context.Context, id string, entity *T) error
	Delete(ctx context.Context, id string) error
	Search(ctx context.Context, filters map[string]map[string]map[string]string, includes string, page int, pageSize int, sortParams string) ([]T, int, error)
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

	err = txn.Scopes(query_builder.ColumnValEqual("id", id)).First(&res).Error

	return res, err
}

func (br *baseRepository[T]) Create(ctx context.Context, entity *T) (err error) {
	txn := br.db.Debug().WithContext(ctx).Model(entity)

	return txn.Create(entity).Error
}

func (br *baseRepository[T]) Update(ctx context.Context, id string, entity *T) (err error) {
	txn := br.db.Debug().WithContext(ctx).Model(entity)

	return txn.Scopes(query_builder.ColumnValEqual("id", id)).Updates(entity).Error
}

func (br *baseRepository[T]) Delete(ctx context.Context, id string) (err error) {
	txn := br.db.Debug().WithContext(ctx).Model(br.repoModel)

	return util.SoftDeleteById(ctx, txn, id)
}

func (br *baseRepository[T]) Search(ctx context.Context, filterMap map[string]map[string]map[string]string, includes string, page int, pageSize int, sortParams string) (data []T, totalCount int, err error) {
	txn := br.db.Debug().WithContext(ctx).Model(br.repoModel)
	query := txn.Model(br.repoModel)

	for tableName, tableFilterMap := range filterMap {
		if tableName == constant.Root {
			tableName, err = database.GetTableName(br.db, br.repoModel)
			if err != nil {
				return nil, 0, err
			}
		} else {
			joinCondition := util.ReadTag(br.repoModel, tableName, "join")
			query = query.Joins(joinCondition)
			tableName = util.ReadTag(br.repoModel, tableName, "tableName")
		}

		for _, colVal := range tableFilterMap {
			for col, val := range colVal {
				key := fmt.Sprintf("\"%s\".%s", tableName, col)
				delete(colVal, col)
				colVal[key] = val
			}
		}

		wherePredicates, err := search.GetWherePredicates(tableFilterMap)
		if err != nil {
			return nil, 0, err
		}

		query = query.Scopes(wherePredicates...)
	}

	//get total count
	totalCount, err = util.GetTotalCount(query)
	if err != nil {
		return nil, 0, err
	}

	//add includes
	if includes != "" {
		for _, include := range strings.Split(includes, ",") {
			query = query.Preload(include)
		}
	}

	//add sort
	if sortParams != "" {
		sortPredicates := search.GetSortPredicates(sortParams)
		query = query.Scopes(sortPredicates...)
	}

	//add pagination
	paginationPredicates := search.GetPaginationPredicates(page, pageSize)
	query = query.Scopes(paginationPredicates...)

	//filter rows
	if err := query.Find(&data).Error; err != nil {
		return nil, 0, err
	}

	return data, totalCount, nil
}
