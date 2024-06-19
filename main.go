package main

import (
	"TOKOKU_APP/configs"
	"TOKOKU_APP/internal/controllers"
	"TOKOKU_APP/internal/models"
	"fmt"
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

	im := models.NewItemModel(connection)
	ic := controllers.NewItemController(im)

	// em := models.NewEmployeeModel(connection)
	// cm := controllers.NewEmployeeController(em)

	// data, err := cm.Login()
	// uc := controllers.(um)
	// im.SelectItem(setup)

	// cm.Register()
	ic.InserItem(1)

}
