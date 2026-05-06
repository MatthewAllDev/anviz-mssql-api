package db

import "gorm.io/gorm"

func updateExisting(conn *gorm.DB, model any, condition string, args []any, updates map[string]any) error {
	var count int64
	if err := conn.Model(model).Where(condition, args...).Count(&count).Error; err != nil {
		return err
	}
	if count == 0 {
		return gorm.ErrRecordNotFound
	}
	return conn.Model(model).Where(condition, args...).Updates(updates).Error
}
