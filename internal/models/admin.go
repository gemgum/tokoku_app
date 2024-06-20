package models

import (
	"fmt"

	"gorm.io/gorm"
)

type AdminModel struct {
	db *gorm.DB
}

func NewAdminModel(connection *gorm.DB) *AdminModel {
	return &AdminModel{
		db: connection,
	}
}

func (am *AdminModel) IntializeAdminAccount(schema string, newEmployee Employee) (bool, error) {
	query := fmt.Sprintf(`INSERT INTO "%s"."employees" 
	("created_at","updated_at", "name", "password", "email", "id") 
	VALUES (?, ?, ?, ?, ?, ?);`, schema)
	res := am.db.Debug().Exec(query, &newEmployee.UpdatedAt, &newEmployee.UpdatedAt,
		&newEmployee.Name,
		&newEmployee.Password,
		&newEmployee.Email,
		&newEmployee.ID)
	err := res.Error
	if err != nil {
		// return Todo{}, err
		return false, err
	}
	rowsAffected := res.RowsAffected
	if rowsAffected <= 0 {
		// return Todo{}, err
		err = fmt.Errorf("no rows affected")
		return false, err
	}
	return true, nil
}
