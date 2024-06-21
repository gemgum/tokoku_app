package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type ItemTransaction struct {
	gorm.Model
	Amount uint
	Item   uint
	TrxId  uint
}

type Transaction struct {
	gorm.Model
	TrxDate          string `gorm:"type:date"`
	Employee         uint
	Customer         uint
	ItemTransactions []ItemTransaction `gorm:"foreignKey:TrxId"`
}
type Item struct {
	gorm.Model
	ItemName         string `gorm:"type:varchar(100)"`
	ItemStock        uint
	Price            uint
	Employee         uint
	ItemTransactions []ItemTransaction `gorm:"foreignKey:Item"`
}

type Customer struct {
	gorm.Model
	Name         string `gorm:"type:varchar(100)"`
	Phone        string
	Address      string
	Employee     uint
	Transactions []Transaction `gorm:"foreignKey:Customer"`
}

type ItemModel struct {
	db *gorm.DB
}

func NewItemModel(connection *gorm.DB) *ItemModel {
	return &ItemModel{
		db: connection,
	}
}

func (im *ItemModel) SelectSumarrytItem(s uint) (uint, error) {
	todos := make([]Item, 0)
	// query := fmt.Sprintf("SELECT * FROM \"%s\".\"items\" WHERE \"items\".\"deleted_at\" IS NULL;", "public")
	query := fmt.Sprintf(`SELECT item_stock FROM %s.items WHERE id = ? AND deleted_at IS NULL;`, "public")

	// err := im.db.Debug().Exec(query, &s.Dbschema).Error
	// var items []Item

	err := im.db.Debug().Raw(query, s).Scan(&todos).Error

	// todos = append(todos, items...)
	var rv uint
	for _, result := range todos {
		// fmt.Printf("Tanggal Transaksi: %v, Nama Barang: %s, Jumlah: %d, Pembeli: %d, Pegawai: %d\n",
		// 	result.UpdatedAt, result.ItemName, result.ItemStock, result.Price, result.Employee)

		fmt.Printf("Jumlah: %d\n",
			result.ItemStock)
		rv = result.ItemStock
	}
	// fmt.Println(items)
	if err != nil {
		// return Todo{}, err
		return 0, err

	}
	return rv, nil
}

func (im *ItemModel) InsertItem(schema string, item Item) (bool, error) {
	query := fmt.Sprintf(`INSERT INTO "%s"."items" 
	("created_at","updated_at", "item_name", "item_stock", "price", "employee") 
	VALUES (?, ?, ?, ?, ?, ?);`, schema)
	res := im.db.Debug().Exec(query, &item.UpdatedAt, &item.UpdatedAt, &item.ItemName, &item.ItemStock, &item.Price, &item.Employee)
	err := res.Error
	if err != nil {
		return false, err
	}

	rowsAffected := res.RowsAffected
	if rowsAffected <= 0 {
		err = fmt.Errorf("no rows affected")
		return false, err

	}

	return true, nil
}

func (cm *ItemModel) InsertCustomer(schema string, cust Customer) (bool, error) {
	query := fmt.Sprintf(`INSERT INTO "%s"."customers" 
	("created_at","updated_at", "name", "phone", "address", "employee") 
	VALUES (?, ?, ?, ?, ?, ?);`, schema)
	res := cm.db.Debug().Exec(query, &cust.UpdatedAt, &cust.UpdatedAt, &cust.Name, &cust.Phone, &cust.Address, &cust.Employee)
	err := res.Error
	if err != nil {
		return false, err
	}

	rowsAffected := res.RowsAffected
	if rowsAffected <= 0 {
		err = fmt.Errorf("no rows affected")
		return false, err

	}

	return true, nil
}

func (cm *ItemModel) InsertTransaction(schema string, trx Transaction) (int, bool, error) {
	// var trxId uint
	query := fmt.Sprintf(`INSERT INTO "%s"."transactions" 
	("created_at","updated_at", "trx_date", "customer", "employee") 
	VALUES (?, ?, ?, ?, ?) RETURNING id;`, schema)
	var TrxId int
	res := cm.db.Debug().Raw(query, &trx.UpdatedAt, &trx.UpdatedAt, &trx.UpdatedAt, &trx.Customer, &trx.Employee).Scan(&TrxId)
	err := res.Error
	if err != nil {
		return 0, false, err
	}
	rowsAffected := res.RowsAffected
	if rowsAffected <= 0 {
		err = fmt.Errorf("no rows affected")
		return 0, false, err

	}
	return TrxId, true, nil
}

func (iit *ItemModel) InsertItemTransaction(schema string, trx_id uint, trx_item ItemTransaction) (bool, error) {

	query := fmt.Sprintf(`INSERT INTO "%s"."item_transactions" 
	("created_at","updated_at", "amount", "item", "trx_id") 
	VALUES (?, ?, ?, ?, ?);`, schema)
	trx_item.TrxId = trx_id
	res := iit.db.Debug().Exec(query, &trx_item.UpdatedAt, &trx_item.UpdatedAt, &trx_item.Amount, &trx_item.Item, &trx_item.TrxId)

	err := res.Error
	if err != nil {
		return false, err
	}
	rowsAffected := res.RowsAffected
	if rowsAffected <= 0 {
		err = fmt.Errorf("no rows affected")
		return false, err

	}
	return true, nil

}

func (cm *ItemModel) DecreItem(schema string, item Item) (bool, error) {

	query := fmt.Sprintf(`UPDATE "%s"."items"
	SET "updated_at" = ?, "item_stock"= ? WHERE "items"."id" = ?;`, schema)
	res := cm.db.Debug().Exec(query, &item.UpdatedAt, &item.ItemStock, &item.ID)

	err := res.Error
	if err != nil {
		return false, err
	}
	rowsAffected := res.RowsAffected
	if rowsAffected <= 0 {
		err = fmt.Errorf("no rows affected")
		return false, err

	}
	return true, nil

}

func (im *ItemModel) ItemUpdateInfo(schema string, item Item) (bool, error) {
	query := fmt.Sprintf(`UPDATE "%s"."items"
	 SET "updated_at", "item_name", "item_stock", "price", "employee" WHERE "items"."id")
	  VALUES (?, ?, ?, ?, ?;`, schema)
	res := im.db.Debug().Exec(query, &item.UpdatedAt, &item.ItemName, &item.ItemStock, &item.Price, &item.Employee, &item.ID)
	err := res.Error
	if err != nil {
		return false, err
	}
	rowsAffected := res.RowsAffected
	if rowsAffected > 0 {
		err = fmt.Errorf("no rows affected")
		return false, err

	}
	return true, nil
}

func (im *ItemModel) ItemEdit(schema string, item Item) (bool, error) {
	query := fmt.Sprintf(`UPDATE "%s"."items"
	 SET "updated_at", "item_name", "item_stock", "price", "employee" WHERE "items"."id")
	  VALUES (?, ?, ?, ?, ?;`, schema)
	res := im.db.Debug().Exec(query, &item.UpdatedAt, &item.ItemName, &item.ItemStock, &item.Price, &item.Employee, &item.ID)
	err := res.Error
	if err != nil {
		return false, err
	}
	rowsAffected := res.RowsAffected
	if rowsAffected > 0 {
		err = fmt.Errorf("no rows affected")
		return false, err

	}
	return true, nil
}

func (im *ItemModel) DeleteTransaction(schema string, trx Transaction) (bool, error) {
	query := fmt.Sprintf(`UPDATE "%s"."transactions" SET "deleted_at"= ? 
	WHERE id = ? AND "deleted_at" IS NULL;`, schema)
	res := im.db.Debug().Exec(query, &trx.UpdatedAt, &trx.ID)
	err := res.Error
	if err != nil {
		return false, err
	}
	rowsAffected := res.RowsAffected
	if rowsAffected <= 0 {
		err = fmt.Errorf("no rows affected")
		return false, err

	}
	return true, nil
}

func (im *ItemModel) DeleteItemTransaction(schema string, itemTrx ItemTransaction) (bool, error) {
	query := fmt.Sprintf(`UPDATE "%s"."item_transactions" SET "deleted_at"= ? 
	WHERE id = ? AND "deleted_at" IS NULL;`, schema)
	res := im.db.Debug().Exec(query, &itemTrx.UpdatedAt, &itemTrx.ID)
	err := res.Error
	if err != nil {
		return false, err
	}
	rowsAffected := res.RowsAffected
	if rowsAffected <= 0 {
		err = fmt.Errorf("no rows affected")
		return false, err

	}
	return true, nil
}

func (cm *ItemModel) DeleteCustomerData(schema string, cust Customer) (bool, error) {
	query := fmt.Sprintf(`UPDATE "%s"."customers" SET "deleted_at"= ? 
	WHERE id = ? AND "deleted_at" IS NULL;`, schema)
	res := cm.db.Debug().Exec(query, &cust.UpdatedAt, &cust.ID)
	err := res.Error
	if err != nil {
		return false, err
	}
	rowsAffected := res.RowsAffected
	if rowsAffected <= 0 {
		err = fmt.Errorf("no rows affected")
		return false, err

	}
	return true, nil
}

func (cm *ItemModel) DeleteEmployee(schema string, emp Employee) (bool, error) {
	query := fmt.Sprintf(`UPDATE "%s"."employees" SET "deleted_at"= ? 
	WHERE id = ? AND "deleted_at" IS NULL;`, schema)
	res := cm.db.Debug().Exec(query, &emp.UpdatedAt, &emp.ID)
	err := res.Error
	if err != nil {
		return false, err
	}
	rowsAffected := res.RowsAffected
	if rowsAffected <= 0 {
		err = fmt.Errorf("no rows affected")
		return false, err

	}
	return true, nil
}

func (im *ItemModel) RemoveItem(schema string, item Item) (bool, error) {

	query := fmt.Sprintf(`UPDATE "%s"."items" SET "deleted_at"= ? 
	WHERE id = ? AND "deleted_at" IS NULL;`, schema)
	res := im.db.Debug().Exec(query, &item.UpdatedAt, &item.ID)
	err := res.Error
	if err != nil {
		return false, err
	}
	rowsAffected := res.RowsAffected
	if rowsAffected <= 0 {
		err = fmt.Errorf("no rows affected")
		return false, err
	}
	return true, nil
}

type TransactionResult struct {
	TanggalTransaksi time.Time `json:"tanggal_transaksi"`
	NamaBarang       string    `json:"nama_barang"`
	Jumlah           int       `json:"jumlah"`
	Pembeli          string    `json:"pembeli"`
	Pegawai          string    `json:"pegawai"`
}

func (tm *ItemModel) ShowTransactionTodayByCustomer(trx Transaction) ([]TransactionResult, error) {
	trxRv := make([]TransactionResult, 0)
	query := `select 
     it.created_at as tanggal_transaksi,
     it.item as nama_barang,
     it.amount as jumlah,
     t.customer as pembeli,
     t.employee as pegawai
 	 from transactions t
 	 join employees e on e.id = t.employee 
 	 join item_transactions it on it.trx_id = t.id
 	 where t.customer = ? and t.employee = ? and DATE(it.created_at) = CURRENT_DATE;`
	err := tm.db.Debug().Raw(query, &trx.Customer, &trx.Employee).Scan(&trxRv).Error
	fmt.Println(query)
	if err != nil {
		// return Todo{}, err
		return nil, err

	}
	return trxRv, nil
}

func (tm *ItemModel) ShowTransactionAll(trx Transaction) ([]TransactionResult, error) {
	trxRv := make([]TransactionResult, 0)
	query := `select 
     it.created_at as tanggal_transaksi,
     it.item as nama_barang,
     it.amount as jumlah,
     t.customer as pembeli,
     t.employee as pegawai
 	 from transactions t
 	 join employees e on e.id = t.employee 
 	 join item_transactions it on it.trx_id = t.id
 	 where t.customer = ? ;`
	err := tm.db.Debug().Raw(query, &trx.Customer, &trx.Employee).Scan(&trxRv).Error
	fmt.Println(query)
	if err != nil {
		// return Todo{}, err
		return nil, err

	}
	return trxRv, nil
}

func (im *ItemModel) EditItemStock(schema string, item Item) (bool, error) {
	query := fmt.Sprintf(`UPDATE "%s"."items" SET "item_stock"= ?, "updated_at"= ?
	WHERE id = ? AND "deleted_at" IS NULL;`, schema)
	res := im.db.Debug().Exec(query, &item.ItemStock, &item.UpdatedAt, &item.ID)
	err := res.Error
	if err != nil {
		return false, err
	}
	rowsAffected := res.RowsAffected
	if rowsAffected <= 0 {
		err = fmt.Errorf("no rows affected")
		return false, err
	}
	return true, nil
}

func (im *ItemModel) EditItem(schema string, item Item) (bool, error) {
	query := fmt.Sprintf(`UPDATE "%s"."items" SET "updated_at"= ?, "item_name"= ?, "price"= ?, "item_stock"= ?
	WHERE id = ? AND "deleted_at" IS NULL;`, schema)
	res := im.db.Debug().Exec(query, &item.UpdatedAt, &item.ItemName, &item.Price, &item.ItemStock, &item.ID)
	err := res.Error
	if err != nil {
		return false, err
	}
	rowsAffected := res.RowsAffected
	if rowsAffected <= 0 {
		err = fmt.Errorf("no rows affected")
		return false, err
	}
	return true, nil
}
