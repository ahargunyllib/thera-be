package repository

import (
	"fmt"
	"slices"
	"strings"

	"github.com/ahargunyllib/thera-be/domain/dto"
)

func normalizeQuery(query *dto.GetPatientsQuery) {
	if query.Limit < 1 {
		query.Limit = 10
	}
	if query.Page < 1 {
		query.Page = 1
	}
	allowedColumns := []string{"id", "full_name", "created_at"}
	if !slices.Contains(allowedColumns, query.SortBy) {
		query.SortBy = "created_at"
	}
	if query.SortOrder != "asc" && query.SortOrder != "desc" {
		query.SortOrder = "asc"
	}
}

func appendPagination(qb *strings.Builder, args *[]any, sortBy, sortOrder string, limit, page int) {
	qb.WriteString(" ORDER BY ")
	qb.WriteString(sortBy)
	qb.WriteString(" ")
	qb.WriteString(sortOrder)

	qb.WriteString(" LIMIT ")
	qb.WriteString(fmt.Sprintf("$%d", len(*args)+1))
	qb.WriteString(" OFFSET ")
	qb.WriteString(fmt.Sprintf("$%d", len(*args)+2))

	offset := (page - 1) * limit
	*args = append(*args, limit, offset)
}
