package employee

type EmployeeModel struct {
	ID       int     `gorm:"primaryKey"`
	Name     string  `gorm:"column:name"`
	Position string  `gorm:"column:position"`
	Salary   float64 `gorm:"column:salary"`
}

func (EmployeeModel) TableName() string {
	return "employee"
}
