package models

type Employee struct {
	Id           int64  `json:"id" gorm:"primary_key"`
	EmployeeName string `json:"employeename"`

	EmployeeDesignation string `json:"employeedesignation"`
}
