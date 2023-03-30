package query_builder

import (
	"fmt"
	"github.com/Anupam-dagar/baileys/constant"
	"gorm.io/gorm"
)

func ColumnStrValStartsWith(columnName string, value string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if value == "" {
			return db
		}

		query := fmt.Sprintf("%s ILIKE ?", columnName)
		value = "%" + value
		return ApplyWhereCondition(db, query, value, constant.QueryAnd)
	}
}

func ColumnStrValStartsHasSubStr(columnName string, value string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if value == "" {
			return db
		}
		query := fmt.Sprintf("%s ILIKE ?", columnName)
		value := "%" + value + "%"
		return ApplyWhereCondition(db, query, value, constant.QueryAnd)
	}
}

func ColumnValEqual(columnName string, val interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if val == "" {
			return db
		}
		query := fmt.Sprintf("%s = ?", columnName)
		return ApplyWhereCondition(db, query, val, constant.QueryAnd)
	}
}

func ColumnValNotEqual(columnName string, val interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if val == "" {
			return db
		}
		query := fmt.Sprintf("%s = ?", columnName)
		return ApplyWhereCondition(db, query, val, constant.QueryNot)
	}
}

func ColumnValGreaterThan(columnName string, value interface{}, applyEqual bool) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if value == "" {
			return db
		}

		if applyEqual {
			query := fmt.Sprintf("%s >= ?", columnName)
			return ApplyWhereCondition(db, query, value, constant.QueryAnd)
		}

		query := fmt.Sprintf("%s > ?", columnName)
		return ApplyWhereCondition(db, query, value, constant.QueryAnd)
	}
}

func ColumnValLessThan(columnName string, value interface{}, applyEqual bool) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if value == "" {
			return db
		}

		if applyEqual {
			query := fmt.Sprintf("%s <= ?", columnName)
			return ApplyWhereCondition(db, query, value, constant.QueryAnd)
		}

		query := fmt.Sprintf("%s < ?", columnName)
		return ApplyWhereCondition(db, query, value, constant.QueryAnd)
	}
}

func ColumnValNull(columnName string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		query := fmt.Sprintf("%s IS NULL", columnName)
		return db.Where(query)
	}
}

func ColumnValNotNull(columnName string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		query := fmt.Sprintf("%s IS NULL", columnName)
		return db.Not(query)
	}
}

func ColumnOrderBy(columnName string, orderToken constant.SortOrder) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Order(fmt.Sprintf("%s %s", columnName, orderToken))
	}
}

func InnerJoinTable(tableName string, condition string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		query := fmt.Sprintf("inner join %s on %s", tableName, condition)
		return db.Joins(query)
	}
}

func LeftJoinTable(tableName string, condition string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		query := fmt.Sprintf("left join %s on %s", tableName, condition)
		return db.Joins(query)
	}
}

func BuildJoinCondition(columnOne string, columnTwo string) string {
	return fmt.Sprintf("%s = %s", columnOne, columnTwo)
}

func ColumnValIn[T any](columnName string, vals []T, operation constant.QueryOperation) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(vals) == 0 {
			return db
		}

		query := fmt.Sprintf("%s IN ?", columnName)
		return ApplyWhereCondition(db, query, vals, operation)
	}
}

func ColumnValNotIn[T any](columnName string, vals []T, operation constant.QueryOperation) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(vals) == 0 {
			return db
		}

		query := fmt.Sprintf("%s IN ?", columnName)
		return ApplyWhereCondition(db, query, vals, operation)
	}
}

func ApplyWhereCondition(db *gorm.DB, query string, val interface{}, operation constant.QueryOperation) *gorm.DB {
	switch operation {
	case constant.QueryAnd:
		return db.Where(query, val)
	case constant.QueryNot:
		return db.Not(query, val)
	default:
		return db.Or(query, val)
	}
}

func BoolColumnValEqual(columnName string, val *bool, operation constant.QueryOperation) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if val == nil {
			return db
		}

		query := fmt.Sprintf("%s = ?", columnName)
		return ApplyWhereCondition(db, query, val, operation)
	}
}

func Paginate(page int, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if pageSize > constant.MaxPageSize {
			pageSize = constant.DefaultPageSize
		}
		offset := page * pageSize

		return db.Offset(offset).Limit(pageSize)
	}
}
