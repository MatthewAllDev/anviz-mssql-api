package models

type Department struct {
	Id              uint        `gorm:"column:Deptid;primaryKey;type:integer" json:"id"`
	Name            string      `gorm:"column:DeptName;size:50;type:varchar(50)" json:"name,omitempty"`
	SupDepartmentID uint        `gorm:"column:SupDeptid;type:integer" json:"supDepartmentId,omitempty"`
	SupDepartment   *Department `gorm:"foreignKey:SupDepartmentID;references:Id" json:"supDepartment,omitempty"`
}

func (Department) TableName() string {
	return "Dept"
}
