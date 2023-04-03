package util

import (
	"context"
	"github.com/Anupam-dagar/baileys/constant"
	"github.com/Anupam-dagar/baileys/util/database/query_builder"
	"gorm.io/gorm"
)

func SoftDeleteById[T string](ctx context.Context, txn *gorm.DB, id T) (err error) {
	deletedAt := GetNowTimeMillis()
	entityMap := map[string]interface{}{
		constant.ColDeletedAt: deletedAt,
	}

	return txn.Scopes(
		query_builder.ColumnValEqual(constant.ColId, id),
	).Updates(entityMap).Error
}

func GetTotalCount(txn *gorm.DB) (totalCount int, err error) {
	var count int64
	err = txn.Count(&count).Error
	return int(count), err
}
