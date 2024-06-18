package models

import (
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

func (um *EmployeeModel) Login(email string, password string) (Employee, error) {
	var result Employee
	err := um.db.Where("email = ? AND password = ?", email, password).First(&result).Error
	if err != nil {
		return Employee{}, err
	}
	return result, nil
}

func (um *EmployeeModel) Register(newEmployee Employee) (bool, error) {
	err := um.db.Create(&newEmployee).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
