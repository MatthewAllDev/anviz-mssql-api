package db

import (
	"anviz-mssql-api/models"

	"gorm.io/gorm"
)

func GetDevices(conn *gorm.DB, pagination Pagination) ([]models.Device, error) {
	var devices []models.Device
	err := paginate(conn.Order("clientid"), pagination).Find(&devices).Error
	return devices, err
}

func CreateDevice(conn *gorm.DB, device *models.Device) error {
	return conn.Create(device).Error
}

func UpdateDevice(conn *gorm.DB, id uint, updates map[string]any) error {
	return updateExisting(conn, &models.Device{}, "clientid = ?", []any{id}, updates)
}

func DeleteDevice(conn *gorm.DB, id uint) error {
	result := conn.Delete(&models.Device{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
