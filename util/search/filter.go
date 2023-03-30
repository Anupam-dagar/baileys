package search

import (
	"errors"
	"github.com/Anupam-dagar/baileys/constant"
	"github.com/Anupam-dagar/baileys/util"
	"github.com/Anupam-dagar/baileys/util/database/query_builder"
	"gorm.io/gorm"
	"strings"
)

func ParseFilters(filters string) (filterMap map[string]map[string]map[string]string, err error) {
	filterMap = make(map[string]map[string]map[string]string)
	filterArray := strings.Split(filters, constant.AndOperatorDelimiter)
	for _, filter := range filterArray {
		if filter == "" {
			continue
		}

		var columnAndOperator, value string
		if strings.Contains(filter, ":") {
			splitFilter := strings.SplitN(filter, ":", 2)
			columnAndOperator = splitFilter[0]
			value = splitFilter[1]
		} else {
			columnAndOperator = filter
			value = ""
		}

		if !strings.Contains(columnAndOperator, ".") {
			return nil, errors.New("invalid Search Operator")
		}

		column, operator := util.SplitStringFromBack(columnAndOperator, ".")
		table := "root"
		if strings.Contains(column, "-") {
			colArray := strings.SplitN(column, "-", 2)
			table = colArray[0]
			column = colArray[1]
		} else if strings.Contains(column, ".") {
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

func GetWherePredicates(searchParams map[string]map[string]string) ([]func(db *gorm.DB) *gorm.DB, error) {
	var predicates []func(db *gorm.DB) *gorm.DB
	for operator, params := range searchParams {
		for column, value := range params {
			switch operator {
			case "eq":
				predicates = append(predicates, query_builder.ColumnValEqual(column, value))
			case "neq":
				predicates = append(predicates, query_builder.ColumnValNotEqual(column, value))
			case "in":
				values := strings.Split(value, ",")
				predicates = append(predicates, query_builder.ColumnValIn(column, values, constant.QueryAnd))
			case "nin":
				values := strings.Split(value, ",")
				predicates = append(predicates, query_builder.ColumnValNotIn(column, values, constant.QueryNot))
			case "gt":
				predicates = append(predicates, query_builder.ColumnValGreaterThan(column, value, false))
			case "lt":
				predicates = append(predicates, query_builder.ColumnValLessThan(column, value, false))
			case "ge":
				predicates = append(predicates, query_builder.ColumnValGreaterThan(column, value, true))
			case "le":
				predicates = append(predicates, query_builder.ColumnValLessThan(column, value, true))
			case "like":
				predicates = append(predicates, query_builder.ColumnStrValStartsWith(column, value))
			case "null":
				predicates = append(predicates, query_builder.ColumnValNull(column))
			case "nn":
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

	sortParamArray := strings.Split(sortParams, ",")
	for _, sortParam := range sortParamArray {
		var order constant.SortOrder
		var column string
		if strings.Contains(sortParam, ":") {
			columnOrder := strings.Split(sortParam, ":")
			column = columnOrder[0]
			order = constant.SortOrder(columnOrder[1])
		} else {
			column = sortParam
			order = constant.DefaultSortOrder
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
