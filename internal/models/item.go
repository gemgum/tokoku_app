package models

import (
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
	ItemStock        bool   `gorm:"type:varchar(100)"`
	Price            uint
	Employee         uint
	ItemTransactions []ItemTransaction `gorm:"foreignKey:Item"`
}

type Customer struct {
	gorm.Model
	Name         string `gorm:"type:varchar(100)"`
	Phone        uint
	Address      uint
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
