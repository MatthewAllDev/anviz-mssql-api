package db

import (
	"anviz-mssql-api/models"

	"gorm.io/gorm"
)

func GetDepartments(conn *gorm.DB, pagination Pagination) ([]models.Department, error) {
	var departments []models.Department
	err := paginate(conn.Order("deptid"), pagination).Find(&departments).Error
	return departments, err
}

func CreateDepartment(conn *gorm.DB, dept *models.Department) error {
	return conn.Create(dept).Error
}

func UpdateDepartment(conn *gorm.DB, id uint, updates map[string]any) error {
	return updateExisting(conn, &models.Department{}, "deptid = ?", []any{id}, updates)
}

func DeleteDepartment(conn *gorm.DB, id uint) error {
	result := conn.Delete(&models.Department{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
