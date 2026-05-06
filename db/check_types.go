package db

import (
	"anviz-mssql-api/models"

	"gorm.io/gorm"
)

func GetCheckTypes(conn *gorm.DB, pagination Pagination) ([]models.CheckType, error) {
	var checkTypes []models.CheckType
	err := paginate(conn.Order("statusid"), pagination).Find(&checkTypes).Error
	return checkTypes, err
}

func UpdateCheckType(conn *gorm.DB, id uint, updates map[string]any) error {
	return updateExisting(conn, &models.CheckType{}, "statusid = ?", []any{id}, updates)
}
