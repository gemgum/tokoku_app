package controllers

import (
	"errors"
	"time"
	"tokoku_app/internal/models"
)

type AdminController struct {
	model *models.AdminModel
}

func NewAdminController(m *models.AdminModel) *AdminController {
	return &AdminController{
		model: m,
	}
}

func (ac *AdminController) IntializeAdminAccount() (models.Employee, error) {
	var newData models.Employee
	newData.UpdatedAt = time.Now()
	newData.Name = "admin"
	newData.Email = "admin"
	newData.Password = "admin"
	newData.ID = 0
	result, err := ac.model.IntializeAdminAccount("public", newData)
	if err != nil && !result {
		return models.Employee{}, errors.New("terjadi masalah ketika menginisialisasi admin")
	}
	return newData, nil
}
