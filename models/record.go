package models

import "time"

type Record struct {
	Id                 uint       `gorm:"column:Logid;primaryKey;type:integer" json:"id"`
	UserID             string     `gorm:"column:Userid;size:20;type:varchar(20)" json:"userId,omitempty"`
	User               *User      `gorm:"foreignKey:UserID;references:Id" json:"user,omitempty"`
	CheckTime          time.Time  `gorm:"column:CheckTime;type:datetime" json:"checkTime"`
	CheckTypeID        uint       `gorm:"column:CheckType;type:integer" json:"checkTypeId"`
	CheckType          *CheckType `gorm:"foreignKey:CheckTypeID;references:Id" json:"checkType"`
	DeviceID           uint       `gorm:"column:Sensorid;type:integer" json:"deviceId,omitempty"`
	Device             *Device    `gorm:"foreignKey:DeviceID;references:Id" json:"device,omitempty"`
	WorkTypeID         uint       `gorm:"column:WorkType;type:integer" json:"workTypeId,omitempty"`
	IdentificationCode uint       `gorm:"column:AttFlag;type:integer" json:"identificationCode,omitempty"`
	IsChecked          bool       `gorm:"column:Checked;type:bit" json:"isChecked"`
	IsExported         bool       `gorm:"column:Exported;type:bit" json:"isExported"`
	OpenDoorFlag       bool       `gorm:"column:OpenDoorFlag;type:bit" json:"openDoorFlag"`
	Temperature        float64    `gorm:"column:temperature;type:float" json:"temperature,omitempty"`
	WhyNoOpen          uint       `gorm:"column:whynoopen;type:integer" json:"whyNoOpen,omitempty"`
	Mask               uint       `gorm:"column:mask;type:integer" json:"mask,omitempty"`
}

func (Record) TableName() string {
	return "Checkinout"
}
