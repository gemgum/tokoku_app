package main

import (
	"fmt"
	"tokoku_app/configs"
	"tokoku_app/internal/models"
)

func main() {
	setup := configs.ImportSetting()
	connection, err := configs.ConnectDB(setup)
	if err != nil {
		fmt.Println("Stop program, masalah pada database", err.Error())
		return
	}

	connection.AutoMigrate(
		&models.Employee{},
		&models.Item{},
		&models.Customer{},
		&models.Transaction{},
		&models.ItemTransaction{})

}
