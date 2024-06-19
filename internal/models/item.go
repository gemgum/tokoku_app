package models

import (
	"fmt"

	"tokoku_app/configs"

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

func (im *ItemModel) SelectItem(s configs.Setting) ([]Item, error) {
	todos := make([]Item, 0)
	query := fmt.Sprintf("SELECT * FROM \"%s\".\"items\" WHERE \"items\".\"deleted_at\" IS NULL;", "public")
	// err := im.db.Debug().Exec(query, &s.Dbschema).Error
	var items []Item

	err := im.db.Debug().Raw(query).Scan(&items).Error

	fmt.Println(items)
	if err != nil {
		// return Todo{}, err
		return nil, err

	}
	return todos, nil
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
	if rowsAffected > 0 {
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
	if rowsAffected > 0 {
		err = fmt.Errorf("no rows affected")
		return false, err

	}

	return true, nil
}

func (cm *ItemModel) InsertTransaction(schema string, trx Transaction, trx_item ItemTransaction) (bool, error) {
	// var trxId uint
	query := fmt.Sprintf(`INSERT INTO "%s"."transactions" 
	("created_at","updated_at", "trx_date", "customer", "employee") 
	VALUES (?, ?, ?, ?, ?) RETURNING id;`, schema)
	res := cm.db.Debug().Raw(query, &trx.UpdatedAt, &trx.UpdatedAt, &trx.UpdatedAt, &trx.Customer, &trx.Employee).Scan(&trx_item.TrxId)
	err := res.Error
	if err != nil {
		return false, err
	}
	rowsAffected := res.RowsAffected
	if rowsAffected <= 0 {
		err = fmt.Errorf("no rows affected")
		return false, err

	}

	query = fmt.Sprintf(`INSERT INTO "%s"."item_transactions" 
	("created_at","updated_at", "amount", "item", "trx_id") 
	VALUES (?, ?, ?, ?, ?);`, schema)
	res = cm.db.Debug().Exec(query, &trx_item.UpdatedAt, &trx_item.UpdatedAt, &trx_item.Amount, &trx_item.Item, &trx_item.TrxId)

	err = res.Error
	if err != nil {
		return false, err
	}
	rowsAffected = res.RowsAffected
	if rowsAffected <= 0 {
		err = fmt.Errorf("no rows affected")
		return false, err

	}
	return true, nil
}
