package query

import (
	"entgo.io/ent/dialect/sql"

	_ "github.com/go-kratos/kratos/v2/encoding/json"
)

// BuildQuerySelector 构建分页过滤查询器
func BuildQuerySelector(
	andFilterJsonString, orFilterJsonString string,
	page, pageSize int32, noPaging bool,
	orderBys []string, defaultOrderField string,
	selectFields []string,
) (whereSelectors []func(s *sql.Selector), querySelectors []func(s *sql.Selector), err error) {
	whereSelectors, err = BuildFilterSelector(andFilterJsonString, orFilterJsonString)
	if err != nil {
		return nil, nil, err
	}

	var orderSelector func(s *sql.Selector)
	orderSelector, err = BuildOrderSelector(orderBys, defaultOrderField)
	if err != nil {
		return nil, nil, err
	}

	pageSelector := BuildPaginationSelector(page, pageSize, noPaging)

	var fieldSelector func(s *sql.Selector)
	fieldSelector, err = BuildFieldSelector(selectFields)

	if len(whereSelectors) > 0 {
		querySelectors = append(querySelectors, whereSelectors...)
	}

	if orderSelector != nil {
		querySelectors = append(querySelectors, orderSelector)
	}
	if pageSelector != nil {
		querySelectors = append(querySelectors, pageSelector)
	}
	if fieldSelector != nil {
		querySelectors = append(querySelectors, fieldSelector)
	}

	return
}
