package models

import "time"

type User struct {
	Id           string      `gorm:"column:Userid;primaryKey;type:varchar(20)" json:"id"`
	Code         string      `gorm:"column:UserCode;size:20;type:varchar(20)" json:"code,omitempty"`
	Name         string      `gorm:"column:Name;size:50;type:varchar(50)" json:"name,omitempty"`
	Sex          string      `gorm:"column:Sex;size:10;type:varchar(10)" json:"sex,omitempty"`
	Password     string      `gorm:"column:Pwd;size:50;type:varchar(50)" json:"-"`
	DepartmentID uint        `gorm:"column:Deptid;type:integer" json:"departmentId,omitempty"`
	Department   *Department `gorm:"foreignKey:DepartmentID;references:Id" json:"department,omitempty"`
	Nation       string      `gorm:"column:Nation;size:50;type:varchar(50)" json:"nation,omitempty"`
	Birthday     time.Time   `gorm:"column:Birthday;type:smalldatetime" json:"-"`
	EmployDate   time.Time   `gorm:"column:EmployDate;type:smalldatetime" json:"employDate,omitempty"`
	Phone        string      `gorm:"column:Telephone;size:50;type:varchar(50)" json:"-"`
	Duty         string      `gorm:"column:Duty;size:50;type:varchar(50)" json:"duty,omitempty"`
	NativePlace  string      `gorm:"column:NativePlace;size:50;type:varchar(50)" json:"-"`
	IDCard       string      `gorm:"column:IDCard;size:50;type:varchar(50)" json:"-"`
	Address      string      `gorm:"column:Address;size:150;type:varchar(150)" json:"-"`
	Mobile       string      `gorm:"column:Mobile;size:50;type:varchar(50)" json:"-"`
	Educated     string      `gorm:"column:Educated;size:50;type:varchar(50)" json:"educated,omitempty"`
	Polity       string      `gorm:"column:Polity;size:50;type:varchar(50)" json:"polity,omitempty"`
	Specialty    string      `gorm:"column:Specialty;size:50;type:varchar(50)" json:"specialty,omitempty"`
	IsAtt        bool        `gorm:"column:IsAtt;type:bit" json:"isAtt"`
	IsOverTime   bool        `gorm:"column:Isovertime;type:bit" json:"isOverTime"`
	IsRest       bool        `gorm:"column:Isrest;type:bit" json:"isRest"`
	Remark       string      `gorm:"column:Remark;size:250;type:varchar(250)" json:"remark,omitempty"`
	MgFlag       int         `gorm:"column:MgFlag;type:smallint" json:"mgFlag,omitempty"`
	CardNum      string      `gorm:"column:CardNum;size:10;type:varchar(10)" json:"-"`
	Picture      []byte      `gorm:"column:Picture;type:image" json:"-"`
	UserFlag     int         `gorm:"column:UserFlag;type:integer" json:"userFlag,omitempty"`
	GroupID      int         `gorm:"column:Groupid;type:integer" json:"groupId,omitempty"`
	ClassFlag    int         `gorm:"column:ClassFlag;type:integer" json:"classFlag,omitempty"`
	OtherInfo    []byte      `gorm:"column:OtherInfo;type:image" json:"-"`
	AdminGroupID int         `gorm:"column:admingroupid;type:integer" json:"-"`
}

func (User) TableName() string {
	return "Userinfo"
}
