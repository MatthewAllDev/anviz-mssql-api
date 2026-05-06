package db

import (
	"anviz-mssql-api/models"

	"gorm.io/gorm"
)

func GetUsers(conn *gorm.DB, userID *string, departmentID *uint, pagination Pagination) ([]models.User, error) {
	var users []models.User
	query := conn.Preload("Department").Order("userid")
	if userID != nil {
		query = query.Where("userid = ?", *userID)
	}
	if departmentID != nil {
		query = query.Where("deptid = ?", *departmentID)
	}
	err := paginate(query, pagination).Find(&users).Error
	return users, err
}

func CreateUser(conn *gorm.DB, user *models.User) error {
	return conn.Create(user).Error
}

func UpdateUser(conn *gorm.DB, id string, updates map[string]any) error {
	return updateExisting(conn, &models.User{}, "userid = ?", []any{id}, updates)
}

func DeleteUser(conn *gorm.DB, id string) error {
	result := conn.Delete(&models.User{}, "userid = ?", id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
