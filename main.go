package main

import (
	"bufio"
	"fmt"
	"os"
	"tokoku_app/configs"
	"tokoku_app/internal/controllers"
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

	im := models.NewItemModel(connection)
	ic := controllers.NewItemController(im)

	em := models.NewEmployeeModel(connection)
	ec := controllers.NewEmployeeController(em)

	am := models.NewAdminModel(connection)
	ac := controllers.NewAdminController(am)
	ac.IntializeAdminAccount()

	scanner := bufio.NewScanner(os.Stdin)
	// var currentUser *models.Employee
	for {
		fmt.Println("selamat datang di Tokoku ^^, silahkan pilih menu dibawah ini")
		fmt.Println("1. login")
		// fmt.Println("2. register")
		// fmt.Println("3. Hapus Item")
		fmt.Println("4. logout")
		fmt.Print("masukkan pilihan : ")
		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			fmt.Println("===Silahkan Login===")
			var isLogin = true
			data, err := ec.Login()
			if err != nil {
				fmt.Println("Terjadi error pada saat login, error: ", err.Error())
				return
			}
			for isLogin {
				fmt.Println("Selamat datang ", data.Name, ",")
				fmt.Println("Pilih menu")
				fmt.Println("1. Tambah Barang")
				fmt.Println("2. Update Barang")
				fmt.Println("3. Pembelian Barang")
				fmt.Println("4. Pendaftaran Customer")
				fmt.Println("5. Tampilkan Nota Pembelian")
				if data.ID == 0 {
					fmt.Println("6. Hapus Data")
				}
				fmt.Println("9. Keluar")
				fmt.Print("Masukkan input: ")
				scanner.Scan()
				choice2 := scanner.Text()
				switch choice2 {
				case "1":
					_, err := ic.InserItem(data.ID)
					if err != nil {
						fmt.Println(err)
					}
				case "3":
					fmt.Print("Masukan ID Customer ")
					var customerId uint
					fmt.Scanln(&customerId)
					_, _, err := ic.InsertTransaction(data.ID, customerId)
					if err != nil {
						fmt.Println(err)
					}

				case "4":
					_, err := ic.InsertCustomer(data.ID)
					if err != nil {
						fmt.Println(err)
					}
				case "5":
					_, err := ic.ShowTransaction(data.ID)
					if err != nil {
						fmt.Println(err)
					}
					// fmt.Println(trxData)
				case "6":

					for choice2 == "6" {

						fmt.Println("Selamat datang di menu admin")
						fmt.Println("Pilih menu")
						fmt.Println("1. Pendaftaran Pegawai")
						fmt.Println("2. Hapus Barang")
						fmt.Println("3. Hapus Transaksi")
						fmt.Println("4. Hapus Transaksi Barang")
						fmt.Println("5. Hapus Data Customer")
						fmt.Println("6. Hapus Data Pegawai")
						fmt.Println("9. Keluar")
						fmt.Print("Masukan Pilihan ")
						scanner.Scan()
						choice3 := scanner.Text()
						switch choice3 {
						case "1":
							_, err := ec.Register()
							if err != nil {
								fmt.Println(err)
							}
						case "2":
							_, err := ic.RemoveItem()
							if err != nil {
								fmt.Println(err)
							}
						case "3":
							err := ic.DeleteTransaction()
							if err != nil {
								fmt.Println(err)
							}
						case "4":
							err := ic.DeleteItemTransaction()
							if err != nil {
								fmt.Println(err)
							}
						case "5":
							err := ic.DeleteCustomer()
							if err != nil {
								fmt.Println(err)
							}
						case "6":
							err := ic.DeleteEmployee()
							if err != nil {
								fmt.Println(err)
							}
						}

					}
				}

			}
		// case "2":
		// 	fmt.Println("===Silahkan Register===")
		// 	// controllers.Register(connection)
		// 	if currentUser != nil {
		// 		fmt.Println("Register gagal")
		// 	} else {
		// 		fmt.Println("Register Berhasil, Selamat Datang ^^")
		// 	}

		// case "3":
		// 	if currentUser == nil {
		// 		fmt.Println("Silahkan login terlebih dahulu")
		// 		continue
		// 	}
		// 	fmt.Println("===Hapus Item===")
		// 	fmt.Print("Masukkan ID item yang ingin di hapus : ")
		// 	scanner.Scan()
		// 	if err != nil {
		// 		fmt.Println("Hapus item gagal", err.Error())
		// 	} else {
		// 		fmt.Println("Hapus item Berhasil")
		// 	}

		case "4":
			fmt.Println("Terima Kasih ^^")
			os.Exit(0)

		default:
			fmt.Println("Pilihan tidak tersedia")
		}

	}

}
