package models

type CheckType struct {
	Id   uint   `gorm:"column:Statusid;primaryKey;type:integer" json:"id"`
	Char string `gorm:"column:StatusChar;size:2;type:varchar(2)" json:"char,omitempty"`
	Name string `gorm:"column:StatusText;size:50;type:varchar(50)" json:"name,omitempty"`
}

func (CheckType) TableName() string {
	return "Status"
}
