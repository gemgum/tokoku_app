package main

import (
	"TOKOKU_APP/configs"
	"TOKOKU_APP/internal/controllers"
	"TOKOKU_APP/internal/models"
	"bufio"
	"fmt"
	"os"
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

	scanner := bufio.NewScanner(os.Stdin)
	var currentUser *models.Employee

	for {
		fmt.Println("selamat datang di Tokoku ^^, silahkan pilih menu dibawah ini")
		fmt.Println("1. login")
		fmt.Println("2. register")
		fmt.Println("3. Hapus Item")
		fmt.Println("4. logout")
		fmt.Print("masukkan pilihan : ")
		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			fmt.Println("===Silahkan Login===")
			// currentUser = controllers.Login(connection)
			if currentUser != nil {
				fmt.Println("Login Berhasil, Selamat Datang ^^")
			} else {
				fmt.Println("Login Gagal, Silahkan periksa kembali")
			}

		case "2":
			fmt.Println("===Silahkan Register===")
			// controllers.Register(connection)
			if currentUser != nil {
				fmt.Println("Register gagal")
			} else {
				fmt.Println("Register Berhasil, Selamat Datang ^^")
			}

		case "3":
			if currentUser == nil {
				fmt.Println("Silahkan login terlebih dahulu")
				continue
			}
			fmt.Println("===Hapus Item===")
			fmt.Print("Masukkan ID item yang ingin di hapus : ")
			scanner.Scan()
			// id := scanner.Text()
			// controllers.DeleteItem(connection, id)
			if err != nil {
				fmt.Println("Hapus item gagal", err.Error())
			} else {
				fmt.Println("Hapus item Berhasil")
			}

		case "4":
			fmt.Println("Terima Kasih ^^")
			os.Exit(0)

		default:
			fmt.Println("Pilihan tidak tersedia")
		}

	}

}
