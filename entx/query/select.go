package query

import (
	"entgo.io/ent/dialect/sql"
	"github.com/moweilong/mo/stringcase"
)

func BuildFieldSelect(s *sql.Selector, fields []string) {
	if len(fields) > 0 {
		for i, field := range fields {
			switch {
			case field == "id_" || field == "_id":
				field = "id"
			}
			fields[i] = stringcase.ToSnakeCase(field)
		}
		s.Select(fields...)
	}
}

func BuildFieldSelector(fields []string) (func(s *sql.Selector), error) {
	if len(fields) > 0 {
		return func(s *sql.Selector) {
			BuildFieldSelect(s, fields)
		}, nil
	} else {
		return nil, nil
	}
}
