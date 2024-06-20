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

func (ic *ItemController) DeleteTransaction() error {
	var newData models.Transaction
	newData.UpdatedAt = time.Now()
	fmt.Print("\nHapus Data Transaksi ")
	fmt.Print("Masukkan ID Transaksi : ")
	fmt.Scanln(&newData.ID)
	result, err := ic.model.DeleteTransaction("public", newData)
	if err != nil {
		return fmt.Errorf("terjadi masalah ketika menghapus transaksi: %v", err)
	}
	if !result {
		return errors.New("gagal menghapus transaksi, id tidak ditemukan")
	}
	return nil
}

func (ic *ItemController) DeleteItemTransaction() error {
	var newData models.ItemTransaction
	// newData.ID = id
	newData.UpdatedAt = time.Now()
	fmt.Print("\nHapus Data Transaksi Barang ")
	fmt.Print("Masukkan ID Transaksi Barang : ")
	fmt.Scanln(&newData.ID)
	result, err := ic.model.DeleteItemTransaction("public", newData)
	if err != nil {
		return fmt.Errorf("terjadi masalah ketika menghapus transaksi: %v", err)
	}
	if !result {
		return errors.New("gagal menghapus transaksi, id tidak ditemukan")
	}
	return nil
}

func (cc *ItemController) DeleteCustomer() error {
	var newData models.Customer
	// newData.ID = id
	newData.UpdatedAt = time.Now()
	fmt.Print("\nHapus Data Customer ")
	fmt.Print("Masukkan ID Customer : ")
	fmt.Scanln(&newData.ID)
	result, err := cc.model.DeleteCustomerData("public", newData)
	if err != nil {
		return fmt.Errorf("terjadi masalah ketika menghapus transaksi: %v", err)
	}
	if !result {
		return errors.New("gagal menghapus transaksi, id tidak ditemukan")
	}
	return nil
}

func (ec *ItemController) DeleteEmployee() error {
	var newData models.Employee
	// newData.ID = id
	newData.UpdatedAt = time.Now()
	fmt.Print("\nHapus Data Employee ")
	fmt.Print("Masukkan ID Employee : ")
	fmt.Scanln(&newData.ID)
	result, err := ec.model.DeleteEmployee("public", newData)
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

func (tc *ItemController) ShowTransaction(idEmployee uint) ([]models.TransactionResult, error) {
	var trx models.Transaction
	fmt.Print("Menampilkan Data Transaksi ")
	// fmt.Scanln(&showTodoData.Activity)
	trx.Employee = idEmployee

	fmt.Print("\nMasukkan ID Customer : ")
	val, err := strconv.Atoi(getScanner())
	if err != nil {
		fmt.Println(err)
	}
	trx.Customer = uint(val)

	// trx.Customer = idCustomer
	trxRrv, err := tc.model.ShowTransaction(trx)
	if err != nil {
		return trxRrv, err
	}
	// fmt.Println("type", reflect.TypeOf(showTodoData))
	fmt.Println()
	fmt.Println()
	for _, result := range trxRrv {
		fmt.Printf("Tanggal Transaksi: %v, Nama Barang: %s, Jumlah: %d, Pembeli: %s, Pegawai: %s\n",
			result.TanggalTransaksi, result.NamaBarang, result.Jumlah, result.Pembeli, result.Pegawai)
	}
	fmt.Println()
	fmt.Println()
	return trxRrv, nil
}

func (ic *ItemController) EditItemStock() (models.Item, error) {

	var newData models.Item
	newData.UpdatedAt = time.Now()
	fmt.Print("\nMasukkan ID Barang : ")

	val, err := strconv.Atoi(getScanner())
	if err != nil {
		fmt.Println(err)
	}
	newData.ID = uint(val)

	fmt.Print("Masukkan Jumlah Barang : ")

	val, err = strconv.Atoi(getScanner())
	if err != nil {
		fmt.Println(err)
	}
	newData.ItemStock = uint(val)

	result, err := ic.model.EditItemStock("public", newData)
	if err != nil && !result {
		return models.Item{}, errors.New("terjadi masalah ketika memasukan data")
	}
	return newData, nil
}

func (ic *ItemController) EditItem() (models.Item, error) {

	var newData models.Item
	newData.UpdatedAt = time.Now()
	fmt.Print("\nMasukkan ID Barang : ")

	val, err := strconv.Atoi(getScanner())
	if err != nil {
		fmt.Println(err)
	}
	newData.ID = uint(val)

	fmt.Print("Edit Nama Barang : ")

	newData.ItemName = getScanner()

	fmt.Print("Edit Harga Barang : ")

	val, err = strconv.Atoi(getScanner())
	if err != nil {
		fmt.Println(err)
	}
	newData.Price = uint(val)

	fmt.Print("Edit Jumlah Barang : ")

	val, err = strconv.Atoi(getScanner())
	if err != nil {
		fmt.Println(err)
	}
	newData.ItemStock = uint(val)

	result, err := ic.model.EditItem("public", newData)
	if err != nil && !result {
		return models.Item{}, errors.New("terjadi masalah ketika memasukan data")
	}
	return newData, nil
}
