package controllers

import (
	"errors"
	"fmt"
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

func (ic *ItemController) InserItem(id uint) (models.Item, error) {
	var newData models.Item
	newData.UpdatedAt = time.Now()
	fmt.Print("\nMasukkan Data Barang ")
	fmt.Print("Masukkan Nama ")
	fmt.Scanln(&newData.ItemName)
	fmt.Print("Masukkan Jumlah ")
	fmt.Scanln(&newData.ItemStock)
	fmt.Print("Masukkan Harga ")
	fmt.Scanln(&newData.Price)
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
	fmt.Print("\nPendaftaran Customer Baru")
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
	newDataTrx.Employee = id_customer

	var newDataTrxItem models.ItemTransaction

	fmt.Print("Masukkan Barang Yang Dibeli ")
	fmt.Scanln(&newDataTrxItem.TrxId)
	fmt.Print("Masukkan Jumlah Barang Yang dibeli ")
	fmt.Scanln(&newDataTrxItem.Amount)

	newDataTrx.Employee = newDataTrx.ID

	result, err := tc.model.InsertTransaction("public", newDataTrx, newDataTrxItem)

	if err != nil && !result {
		return models.Transaction{}, models.ItemTransaction{}, errors.New("terjadi masalah ketika membuat transaksi")
	}
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
