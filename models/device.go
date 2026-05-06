package models

type Device struct {
	Id           uint   `gorm:"column:Clientid;primaryKey;type:integer" json:"id"`
	Name         string `gorm:"column:ClientName;size:50;type:varchar(50)" json:"name,omitempty"`
	LinkMode     uint   `gorm:"column:Linkmode;type:smallint" json:"linkMode,omitempty"`
	IP           string `gorm:"column:IPaddress;size:255;type:varchar(255)" json:"ip,omitempty"`
	Port         uint   `gorm:"column:CommPort;type:integer" json:"port,omitempty"`
	ClientNumber uint   `gorm:"column:ClientNumber;type:integer" json:"clientNumber,omitempty"`
	BaudRate     uint   `gorm:"column:Baudrate;type:integer" json:"baudRate,omitempty"`
	RecStatus    uint   `gorm:"column:RecStatus;type:integer" json:"recStatus,omitempty"`
	FloorID      uint   `gorm:"column:Floorid;type:integer" json:"floorId,omitempty"`
	MachineType  uint   `gorm:"column:MachineType;type:integer" json:"machineType,omitempty"`
	DeviceType   uint   `gorm:"column:DeviceType;type:integer" json:"deviceType,omitempty"`
	CommPassword string `gorm:"column:CommPWD;size:50;type:varchar(50)" json:"-"`
	DeviceFlag   uint   `gorm:"column:deviceflag;type:integer" json:"deviceFlag,omitempty"`
	Timezone     string `gorm:"column:timezone;size:255;type:varchar(255)" json:"timezone,omitempty"`
}

func (Device) TableName() string {
	return "FingerClient"
}
