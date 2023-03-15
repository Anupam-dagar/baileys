package database

import (
	"baileys/constant"
	"fmt"
	"gorm.io/gorm"
)

func ColumnStrValStartsWith(columnName string, filter string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if filter == "" {
			return db
		}

		queryStr := fmt.Sprintf("%s LIKE ?", columnName)
		paramStr := filter + "%"
		return db.Where(queryStr, paramStr)
	}
}

func ColumnStrValStartsHasSubStr(columnName string, filter string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if filter == "" {
			return db
		}
		queryStr := fmt.Sprintf("%s ILIKE ?", columnName)
		paramStr := "%" + filter + "%"
		return db.Where(queryStr, paramStr)
	}
}

func ColumnValEqual(columnName string, val interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if val == "" {
			return db
		}

		return db.Where(fmt.Sprintf("%s = ?", columnName), val)
	}
}

func ColumnOrderBy(columnName string, isAscending bool) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		orderToken := "DESC"
		if isAscending {
			orderToken = "ASC"
		}
		return db.Order(fmt.Sprintf("%s %s", columnName, orderToken))
	}
}

func ColumnValIsNull(columnName string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		query := fmt.Sprintf("%s is null", columnName)
		return db.Where(query)
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

func ApplyWhereCondition(db *gorm.DB, query string, val interface{}, operation constant.QueryOperation) *gorm.DB {
	switch operation {
	case constant.QueryAnd:
		return db.Where(query, val)
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
