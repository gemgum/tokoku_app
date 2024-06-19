package controllers

import (
	"errors"
	"fmt"
	"time"
	"tokoku_app/internal/models"
)

type EmployeeController struct {
	model *models.EmployeeModel
}

func NewEmployeeController(m *models.EmployeeModel) *EmployeeController {
	return &EmployeeController{
		model: m,
	}
}

func (uc *EmployeeController) Login() (models.Employee, error) {
	var email, password string
	fmt.Print("Masukkan email ")
	fmt.Scanln(&email)
	fmt.Print("Masukkan password ")
	fmt.Scanln(&password)
	result, err := uc.model.Login(email, password)
	if err != nil {
		return models.Employee{}, errors.New("terjadi masalah ketika login")
	}
	return result, nil
}

func (uc *EmployeeController) Register() (models.Employee, error) {
	var newData models.Employee
	newData.UpdatedAt = time.Now()
	fmt.Print("Masukkan Nama ")
	fmt.Scanln(&newData.Name)
	fmt.Print("Masukkan Email ")
	fmt.Scanln(&newData.Email)
	fmt.Print("Masukkan Password ")
	fmt.Scanln(&newData.Password)
	result, err := uc.model.Register("public", newData)
	if err != nil && !result {
		return models.Employee{}, errors.New("terjadi masalah ketika registrasi")
	}
	return newData, nil
}
