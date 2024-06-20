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

	trxId, result, err := tc.model.InsertTransaction("public", newDataTrx)

	if err != nil && !result {
		return models.Transaction{}, models.ItemTransaction{}, errors.New("terjadi masalah ketika membuat transaksi")
	}
	var newDataTrxItem models.ItemTransaction
	newDataTrxItem.UpdatedAt = time.Now()
	var lanjutBeli string = "Y"
	for lanjutBeli == "Y" {
		fmt.Print("\nMasukkan Barang Yang Dibeli ")
		fmt.Scanln(&newDataTrxItem.Item)

		fmt.Print("Masukkan Jumlah Barang Yang dibeli ")
		fmt.Scanln(&newDataTrxItem.Amount)

		result, err := tc.model.InsertItemTransaction("public", uint(trxId), newDataTrxItem)
		if err != nil && !result {
			return models.Transaction{}, models.ItemTransaction{}, errors.New("terjadi masalah ketika membuat transaksi item")
		}
		fmt.Print("Lanjutkan membeli ? (Y/n) ")
		fmt.Scanln(&lanjutBeli)
	}
	// newDataTrx.Employee = newDataTrx.ID

	return newDataTrx, newDataTrxItem, nil
}

func (ic *ItemController) ItemUpdateInfo(id uint) (models.Item, error) {
	var newUpdateItem models.Item
	newUpdateItem.UpdatedAt = time.Now()
	fmt.Print("\nEdit Data Item ")
	fmt.Print("Masukkan Nama ")
	fmt.Scanln(&newUpdateItem.ItemName)
	fmt.Print("Masukkan Jumlah ")
	fmt.Scanln(&newUpdateItem.ItemStock)
	fmt.Print("Masukkan Harga ")
	fmt.Scanln(&newUpdateItem.Price)
	newUpdateItem.Employee = id
	result, err := ic.model.ItemUpdateInfo("public", newUpdateItem)
	if err != nil && !result {
		return models.Item{}, errors.New("terjadi masalah ketika memasukan data")
	}
	return newUpdateItem, nil
}

func (ic *ItemController) ItemEdit(id uint) (models.Item, error) {
	var newItemEdit models.Item
	newItemEdit.UpdatedAt = time.Now()
	fmt.Print("\nEdit Data Item ")
	fmt.Print("Masukkan Nama ")
	fmt.Scanln(&newItemEdit.ItemName)
	fmt.Print("Masukkan Jumlah ")
	fmt.Scanln(&newItemEdit.ItemStock)
	fmt.Print("Masukkan Harga ")
	fmt.Scanln(&newItemEdit.Price)
	newItemEdit.Employee = id
	result, err := ic.model.ItemEdit("public", newItemEdit)
	if err != nil && !result {
		return models.Item{}, errors.New("terjadi masalah ketika memasukan data")
	}
	return newItemEdit, nil
}

func (ic *ItemController) DeleteTransaction(id uint) error {
	var newData models.Item
	newData.ID = id
	result, err := ic.model.DeleteTransaction("public", newData)
	if err != nil {
		return fmt.Errorf("terjadi masalah ketika menghapus transaksi: %v", err)
	}
	if !result {
		return errors.New("gagal menghapus transaksi, id tidak ditemukan")
	}
	return nil
}

func (ic *ItemController) DeleteItemTransaction(id uint) error {
	var newData models.Item
	newData.ID = id
	result, err := ic.model.DeleteItemTransaction("public", newData)
	if err != nil {
		return fmt.Errorf("terjadi masalah ketika menghapus transaksi: %v", err)
	}
	if !result {
		return errors.New("gagal menghapus transaksi, id tidak ditemukan")
	}
	return nil
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
