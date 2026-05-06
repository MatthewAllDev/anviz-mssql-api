package db

import (
	"anviz-mssql-api/models"
	"time"

	"gorm.io/gorm"
)

func GetRecords(conn *gorm.DB, recordID *uint, userIDs []string, startTime, endTime *time.Time, pagination Pagination) ([]models.Record, error) {
	var records []models.Record
	query := conn.Preload("User").Preload("Device").Order("checktime DESC").Order("logid DESC")
	if recordID != nil {
		query = query.Where("logid = ?", *recordID)
	}
	if len(userIDs) > 0 {
		query = query.Where("userid IN ?", userIDs)
	}
	if startTime != nil {
		query = query.Where("checktime >= ?", *startTime)
	}
	if endTime != nil {
		query = query.Where("checktime <= ?", *endTime)
	}
	if err := paginate(query, pagination).Find(&records).Error; err != nil {
		return records, err
	}
	return records, attachRecordCheckTypes(conn, records)
}

func GetRecordsSince(conn *gorm.DB, afterID uint, pagination Pagination) ([]models.Record, error) {
	var records []models.Record
	query := conn.
		Preload("User").
		Preload("Device").
		Where("logid > ?", afterID).
		Order("logid ASC")
	if err := paginate(query, pagination).Find(&records).Error; err != nil {
		return records, err
	}
	return records, attachRecordCheckTypes(conn, records)
}

func attachRecordCheckTypes(conn *gorm.DB, records []models.Record) error {
	if len(records) == 0 {
		return nil
	}

	ids := make([]uint, 0, len(records))
	seen := make(map[uint]struct{}, len(records))
	for _, record := range records {
		if _, ok := seen[record.CheckTypeID]; ok {
			continue
		}
		seen[record.CheckTypeID] = struct{}{}
		ids = append(ids, record.CheckTypeID)
	}

	var checkTypes []models.CheckType
	if err := conn.Where("statusid IN ?", ids).Find(&checkTypes).Error; err != nil {
		return err
	}

	byID := make(map[uint]*models.CheckType, len(checkTypes))
	for i := range checkTypes {
		byID[checkTypes[i].Id] = &checkTypes[i]
	}
	for i := range records {
		records[i].CheckType = byID[records[i].CheckTypeID]
	}
	return nil
}

func CreateRecord(conn *gorm.DB, record *models.Record) error {
	return conn.Create(record).Error
}

func UpdateRecord(conn *gorm.DB, id uint, updates map[string]any) error {
	return updateExisting(conn, &models.Record{}, "logid = ?", []any{id}, updates)
}

func DeleteRecord(conn *gorm.DB, id uint) error {
	result := conn.Delete(&models.Record{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
