package db

import "gorm.io/gorm"

type Pagination struct {
	Limit  int
	Offset int
}

func paginate(query *gorm.DB, pagination Pagination) *gorm.DB {
	return query.Limit(pagination.Limit).Offset(pagination.Offset)
}
