package handlers

var userUpdateFields = map[string]string{
	"code":         "UserCode",
	"name":         "Name",
	"sex":          "Sex",
	"departmentId": "Deptid",
	"nation":       "Nation",
	"employDate":   "EmployDate",
	"duty":         "Duty",
	"educated":     "Educated",
	"polity":       "Polity",
	"specialty":    "Specialty",
	"isAtt":        "IsAtt",
	"isOverTime":   "Isovertime",
	"isRest":       "Isrest",
	"remark":       "Remark",
	"mgFlag":       "MgFlag",
	"userFlag":     "UserFlag",
	"groupId":      "Groupid",
	"classFlag":    "ClassFlag",
}

var departmentUpdateFields = map[string]string{
	"name":            "DeptName",
	"supDepartmentId": "SupDeptid",
}

var deviceUpdateFields = map[string]string{
	"name":         "ClientName",
	"linkMode":     "Linkmode",
	"ip":           "IPaddress",
	"port":         "CommPort",
	"clientNumber": "ClientNumber",
	"baudRate":     "Baudrate",
	"recStatus":    "RecStatus",
	"floorId":      "Floorid",
	"machineType":  "MachineType",
	"deviceType":   "DeviceType",
	"deviceFlag":   "deviceflag",
	"timezone":     "timezone",
}

var recordUpdateFields = map[string]string{
	"userId":             "Userid",
	"checkTime":          "CheckTime",
	"checkTypeId":        "CheckType",
	"deviceId":           "Sensorid",
	"workTypeId":         "WorkType",
	"identificationCode": "AttFlag",
	"isChecked":          "Checked",
	"isExported":         "Exported",
	"openDoorFlag":       "OpenDoorFlag",
	"temperature":        "temperature",
	"whyNoOpen":          "whynoopen",
	"mask":               "mask",
}

var checkTypeUpdateFields = map[string]string{
	"char": "StatusChar",
	"name": "StatusText",
}
