package controllers

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"
	"tokoku_app/internal/models"
	// "todo/internal/models"
)

type ItemController struct {
	model *models.ItemModel
}

func NewItemController(i *models.ItemModel) *ItemController {
	return &ItemController{
		model: i,
	}
}

func getScanner() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	choice := scanner.Text()
	return choice
}
func (ic *ItemController) InserItem(id uint) (models.Item, error) {

	var newData models.Item
	newData.UpdatedAt = time.Now()
	fmt.Println("\nMasukkan Data Barang ")
	newData.ItemName = getScanner()

	fmt.Print("Masukkan Jumlah ")
	val, err := strconv.Atoi(getScanner())
	if err != nil {
		fmt.Println(err)
	}
	newData.ItemStock = uint(val)

	fmt.Print("Masukkan Harga ")
	val, err = strconv.Atoi(getScanner())
	if err != nil {
		fmt.Println(err)
	}
	newData.Price = uint(val)

	newData.Employee = id
	result, err := ic.model.InsertItem("public", newData)
	if err != nil && !result {
		return models.Item{}, errors.New("terjadi masalah ketika memasukan data")
	}
	return newData, nil
}

func (ic *ItemController) InsertCustomer(id uint) (models.Customer, error) {
	var newData models.Customer
	newData.UpdatedAt = time.Now()
	fmt.Print("\nPendaftaran Customer Baru\n")
	fmt.Print("Masukkan Nama ")
	fmt.Scanln(&newData.Name)
	fmt.Print("Masukkan Nomer Telfon ")
	fmt.Scanln(&newData.Phone)
	fmt.Print("Masukkan Alamat ")
	fmt.Scanln(&newData.Address)
	newData.Employee = id
	result, err := ic.model.InsertCustomer("public", newData)
	if err != nil && !result {
		return models.Customer{}, errors.New("terjadi masalah ketika pendaftaran")
	}
	return newData, nil
}

func (tc *ItemController) InsertTransaction(id_employee uint, id_customer uint) (models.Transaction, models.ItemTransaction, error) {
	var newDataTrx models.Transaction

	newDataTrx.UpdatedAt = time.Now()
	newDataTrx.Employee = id_employee
	newDataTrx.Customer = id_customer

	fmt.Println("Employee", newDataTrx.Employee)
	fmt.Println("Customer", newDataTrx.Customer)

	var newDataTrxItem models.ItemTransaction
	newDataTrxItem.UpdatedAt = time.Now()
	fmt.Print("Masukkan Barang Yang Dibeli ")
	fmt.Scanln(&newDataTrxItem.Item)

	fmt.Print("Masukkan Jumlah Barang Yang dibeli ")
	fmt.Scanln(&newDataTrxItem.Amount)

	// newDataTrx.Employee = newDataTrx.ID

	result, err := tc.model.InsertTransaction("public", newDataTrx, newDataTrxItem)

	if err != nil && !result {
		return models.Transaction{}, models.ItemTransaction{}, errors.New("terjadi masalah ketika membuat transaksi")
	}
	return newDataTrx, newDataTrxItem, nil
}

func (ic *ItemController) RemoveItem() (models.Item, error) {

	var newData models.Item
	newData.UpdatedAt = time.Now()
	fmt.Println("\nMasukkan ID Barang ")

	val, err := strconv.Atoi(getScanner())
	if err != nil {
		fmt.Println(err)
	}
	newData.ID = uint(val)
	result, err := ic.model.RemoveItem("public", newData)
	if err != nil && !result {
		return models.Item{}, errors.New("terjadi masalah ketika memasukan data")
	}
	return newData, nil
}
