package util

import (
	"baileys/constant"
	"baileys/util/database"
	"context"
	"gorm.io/gorm"
)

func SoftDeleteById[T string](ctx context.Context, txn *gorm.DB, id T) (err error) {
	deletedAt := GetNowTimeMillis()
	entityMap := map[string]interface{}{
		constant.ColDeletedAt: deletedAt,
	}

	return txn.Scopes(
		database.ColumnValEqual(constant.ColId, id),
	).Updates(entityMap).Error
}
