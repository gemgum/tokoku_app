package models

import (
	"fmt"

	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	Name         string
	Password     string
	Email        string
	Items        []Item        `gorm:"foreignKey:Employee"`
	Customers    []Customer    `gorm:"foreignKey:Employee"`
	Transactions []Transaction `gorm:"foreignKey:Employee"`
}

type EmployeeModel struct {
	db *gorm.DB
}

func NewEmployeeModel(connection *gorm.DB) *EmployeeModel {
	return &EmployeeModel{
		db: connection,
	}
}

func (em *EmployeeModel) Login(email string, password string) (Employee, error) {
	var result Employee
	err := em.db.Where("email = ? AND password = ?", email, password).First(&result).Error
	if err != nil {
		return Employee{}, err
	}
	return result, nil
}

func (em *EmployeeModel) Register(schema string, newEmployee Employee) (bool, error) {
	query := fmt.Sprintf(`INSERT INTO "%s"."employees" 
	("created_at","updated_at", "name", "password", "email") 
	VALUES (?, ?, ?, ?, ?);`, schema)

	// query = ` UPDATE "be23"."todos" SET "deleted_at"= ?
	// WHERE (owner = ? AND activity = ?) AND "todos"."deleted_at" IS NULL `

	res := em.db.Debug().Exec(query, &newEmployee.UpdatedAt, &newEmployee.UpdatedAt,
		&newEmployee.Name,
		&newEmployee.Password,
		&newEmployee.Email)
	// var items []Item

	// err := im.db.Debug().Raw(query).Scan(&items).Error

	// fmt.Println(items)

	err := res.Error
	if err != nil {
		// return Todo{}, err
		return false, err

	}

	rowsAffected := res.RowsAffected
	if rowsAffected > 0 {
		// return Todo{}, err
		err = fmt.Errorf("no rows affected")
		return false, err

	}
	return true, nil
}

func (em *EmployeeModel) DeleteCustData(schema string, emp Employee) (bool, error) {
	query := fmt.Sprintf(`DELETE FROM "%s"."employees" WHERE "employees"."id" = ?;`, schema)

	// query = ` UPDATE "be23"."todos" SET "deleted_at"= ?
	// WHERE (owner = ? AND activity = ?) AND "todos"."deleted_at" IS NULL `

	res := em.db.Debug().Exec(query, emp.ID)

	err := res.Error
	if err != nil {
		// return Todo{}, err
		return false, err

	}

	rowsAffected := res.RowsAffected
	if rowsAffected > 0 {
		// return Todo{}, err
		err = fmt.Errorf("no rows affected")
		return false, err

	}
	return true, nil
}

func (em *EmployeeModel) DeleteEmployeData(schema string, emp Employee) (bool, error) {
	query := fmt.Sprintf(`DELETE FROM "%s"."employees" WHERE "employees"."id" = ?;`, schema)

	// query = ` UPDATE "be23"."todos" SET "deleted_at"= ?
	// WHERE (owner = ? AND activity = ?) AND "todos"."deleted_at" IS NULL `

	res := em.db.Debug().Exec(query, emp.ID)

	err := res.Error
	if err != nil {
		// return Todo{}, err
		return false, err

	}

	rowsAffected := res.RowsAffected
	if rowsAffected > 0 {
		// return Todo{}, err
		err = fmt.Errorf("no rows affected")
		return false, err

	}
	return true, nil
}
