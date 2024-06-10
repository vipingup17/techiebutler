package employee

import (
	"errors"
	"fmt"
	"log"

	"techiebutler/configs"

	"gorm.io/gorm"
)

func createNewEmployee(name, position string, salary float64) error {

	employee := EmployeeModel{
		Name:     name,
		Position: position,
		Salary:   salary,
	}

	if err := configs.GetORM().Create(&employee).Error; err != nil {
		return err
	}

	return nil
}

func getEmployeeByID(id int) (*Employee, error) {

	var employeeModel EmployeeModel

	if err := configs.GetORM().First(&employeeModel, id).Error; err != nil {
		return nil, err
	}

	employee := Employee{
		Name:     employeeModel.Name,
		Position: employeeModel.Position,
		Salary:   employeeModel.Salary,
	}

	return &employee, nil
}

func updateEmployee(id int, name, position string, salary float64) error {

	return configs.GetORM().Transaction(func(tx *gorm.DB) error {
		var employeeModel EmployeeModel
		if err := tx.First(&employeeModel, id).Error; err != nil {
			return err
		}

		if name != "" {
			employeeModel.Name = name
		}
		if position != "" {
			employeeModel.Position = position
		}
		if salary != 0 {
			employeeModel.Salary = salary
		}

		if err := tx.Save(&employeeModel).Error; err != nil {
			return err
		}

		return nil
	})
}

func deleteEmployee(id int) error {

	result := configs.GetORM().Delete(&EmployeeModel{}, id)
	if result.Error != nil {
		return errors.New("error deleting employee")
	}

	return nil
}

func getAllEmployeesFromDB(offset, limit int) ([]Employee, error) {
	var employees []Employee

	query := fmt.Sprintf(`
		SELECT name, position, salary
		FROM employee
		ORDER BY id ASC LIMIT %d OFFSET %d`, limit, offset)

	rows, err := configs.GetDB().Query(query)
	if err != nil {
		log.Printf("Error executing query: %v", err)
		return nil, errors.New("error fetching employees from db")
	}
	defer rows.Close()

	for rows.Next() {
		var employee Employee
		err := rows.Scan(&employee.Name, &employee.Position, &employee.Salary)
		if err != nil {
			log.Printf("Error scanning row: %v", err)
			return nil, errors.New("error fetching employees from db")
		}
		employees = append(employees, employee)
	}
	if err := rows.Err(); err != nil {
		fmt.Println("Error iterating over rows: %v", err)
	}

	return employees, nil
}
