package search

import (
	"errors"
	"fmt"
	"github.com/Anupam-dagar/baileys/constant"
	"github.com/Anupam-dagar/baileys/dto"
	"github.com/Anupam-dagar/baileys/util"
	"github.com/Anupam-dagar/baileys/util/database"
	"github.com/Anupam-dagar/baileys/util/database/query_builder"
	"gorm.io/gorm"
	"strings"
)

func ParseFilters(filters string) (filterMap dto.SearchFilters, err error) {
	filterMap = make(dto.SearchFilters)

	filterArray := strings.Split(filters, constant.AndOperatorDelimiter)
	for _, filter := range filterArray {
		if filter == "" {
			continue
		}

		var columnAndOperator, value string
		if strings.Contains(filter, constant.KeyValueDelimiter) {
			splitFilter := strings.SplitN(filter, constant.KeyValueDelimiter, 2)
			columnAndOperator = splitFilter[0]
			value = splitFilter[1]
		} else {
			columnAndOperator = filter
			value = ""
		}

		if !strings.Contains(columnAndOperator, constant.SearchOperatorDelimiter) {
			return nil, errors.New("invalid Search Operator")
		}

		column, operator := util.SplitStringFromBack(columnAndOperator, constant.SearchOperatorDelimiter)
		table := constant.Root
		if strings.Contains(column, constant.ColumnDelimiter) {
			colArray := strings.SplitN(column, constant.ColumnDelimiter, 2)
			table = colArray[0]
			column = colArray[1]
		} else if strings.Contains(column, constant.SearchOperatorDelimiter) {
			return nil, errors.New("child table struct specified without column. Invalid Operation")
		}

		// initialize the table key in filterMap if it doesn't exist
		if _, ok := filterMap[table]; !ok {
			filterMap[table] = make(map[string]map[string]string)
		}

		// initialize the operator key in filterMap[table] if it doesn't exist
		if _, ok := filterMap[table][operator]; !ok {
			filterMap[table][operator] = make(map[string]string)
		}

		// assign the value to filterMap[table][operator][column]
		filterMap[table][operator][column] = value
	}
	return filterMap, nil
}

func getWherePredicates(searchParams map[string]map[string]string) ([]func(db *gorm.DB) *gorm.DB, error) {
	var predicates []func(db *gorm.DB) *gorm.DB
	for operator, params := range searchParams {
		for column, value := range params {
			switch constant.SearchOperator(operator) {
			case constant.Equal:
				predicates = append(predicates, query_builder.ColumnValEqual(column, value))
			case constant.NotEqual:
				predicates = append(predicates, query_builder.ColumnValNotEqual(column, value))
			case constant.In:
				values := strings.Split(value, constant.CommaDelimiter)
				predicates = append(predicates, query_builder.ColumnValIn(column, values, constant.QueryAnd))
			case constant.NotIn:
				values := strings.Split(value, constant.CommaDelimiter)
				predicates = append(predicates, query_builder.ColumnValNotIn(column, values, constant.QueryNot))
			case constant.GreaterThan:
				predicates = append(predicates, query_builder.ColumnValGreaterThan(column, value, false))
			case constant.LessThan:
				predicates = append(predicates, query_builder.ColumnValLessThan(column, value, false))
			case constant.GreaterThanEqualTo:
				predicates = append(predicates, query_builder.ColumnValGreaterThan(column, value, true))
			case constant.LessThanEqualTo:
				predicates = append(predicates, query_builder.ColumnValLessThan(column, value, true))
			case constant.Like:
				predicates = append(predicates, query_builder.ColumnStrValStartsWith(column, value))
			case constant.IsNull:
				predicates = append(predicates, query_builder.ColumnValNull(column))
			case constant.IsNotNull:
				predicates = append(predicates, query_builder.ColumnValNotNull(column))
			default:
				return nil, errors.New("operator not defined")
			}
		}
	}
	return predicates, nil
}

func GetSortPredicates(sortParams string) []func(db *gorm.DB) *gorm.DB {
	var predicates []func(db *gorm.DB) *gorm.DB

	if sortParams == "" {
		sortParams = "id"
	}

	sortParamArray := strings.Split(sortParams, ",")
	for _, sortParam := range sortParamArray {
		order := constant.DefaultSortOrder
		column := sortParam
		if strings.Contains(sortParam, ":") {
			columnOrder := strings.Split(sortParam, ":")
			column = columnOrder[0]
			order = columnOrder[1]
		}
		predicates = append(predicates, query_builder.ColumnOrderBy(column, order))
	}
	return predicates
}

func GetPaginationPredicates(page int, pageSize int) []func(db *gorm.DB) *gorm.DB {
	var predicates []func(db *gorm.DB) *gorm.DB
	predicates = append(predicates, query_builder.Paginate(page, pageSize))
	return predicates
}

func GetWherePredicates(query *gorm.DB, filterMap dto.SearchFilters, repoModel interface{}) (predicates []func(db *gorm.DB) *gorm.DB, err error) {
	for tableName, tableFilterMap := range filterMap {
		if tableName == constant.Root {
			tableName, err = database.GetTableName(query, repoModel)
			if err != nil {
				return nil, err
			}
		} else {
			joinCondition, err := util.ReadTag(repoModel, tableName, "join")
			if err != nil {
				return nil, err
			}
			query = query.Joins(joinCondition)
			tableName, err = util.ReadTag(repoModel, tableName, "tableName")
			if err != nil {
				return nil, err
			}
		}

		for _, colVal := range tableFilterMap {
			for col, val := range colVal {
				key := fmt.Sprintf("\"%s\".%s", tableName, col)
				delete(colVal, col)
				colVal[key] = val
			}
		}

		tableWherePredicates, err := getWherePredicates(tableFilterMap)
		if err != nil {
			return nil, err
		}

		predicates = append(predicates, tableWherePredicates...)
	}
	return predicates, nil
}

func AddIncludes(includes string, query *gorm.DB) *gorm.DB {
	if includes == "" {
		return query
	}
	for _, include := range strings.Split(includes, ",") {
		query = query.Preload(include)
	}
	return query
}

func AddScopes(txn *gorm.DB, predicates ...[]func(db *gorm.DB) *gorm.DB) *gorm.DB {
	for _, predicate := range predicates {
		txn = txn.Scopes(predicate...)
	}
	return txn
}
