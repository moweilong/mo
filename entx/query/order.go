package query

import (
	"strings"

	"entgo.io/ent/dialect/sql"
)

// QueryCommandToOrderConditions 查询命令转换为排序条件
func QueryCommandToOrderConditions(orderBys []string) (func(s *sql.Selector), error) {
	if len(orderBys) == 0 {
		return nil, nil
	}

	return func(s *sql.Selector) {
		for _, v := range orderBys {
			if strings.HasPrefix(v, "-") {
				// 降序
				key := v[1:]
				if len(key) == 0 {
					continue
				}

				BuildOrderSelect(s, key, true)
			} else {
				// 升序
				if len(v) == 0 {
					continue
				}

				BuildOrderSelect(s, v, false)
			}
		}
	}, nil
}

func BuildOrderSelect(s *sql.Selector, field string, desc bool) {
	if desc {
		s.OrderBy(sql.Desc(s.C(field)))
	} else {
		s.OrderBy(sql.Asc(s.C(field)))
	}
}

func BuildOrderSelector(orderBys []string, defaultOrderField string) (func(s *sql.Selector), error) {
	if len(orderBys) == 0 {
		return func(s *sql.Selector) {
			BuildOrderSelect(s, defaultOrderField, true)
		}, nil
	} else {
		return QueryCommandToOrderConditions(orderBys)
	}
}
