package storage

import (
	"time"

	"github.com/jinzhu/gorm"
)

// InitDB initialize DB.
func InitDB() *gorm.DB {
	db, err := gorm.Open("mysql", "root@/expenses?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db.SingularTable(true)
	db.AutoMigrate(&Expense{})
	return db
}

// Expense ...
//go:generate gokit entity -name Expense
type Expense struct {
	ID        int64     `gorm:"primary_key"`
	Account   string    `gorm:"varchar(32);not null;default ''"`
	Category  string    `gorm:"varchar(100);not null;default ''"`
	Currency  string    `gorm:"varchar(3);not null;default ''"`
	Amount    float64   `gorm:"not null;default 0"`
	Date      time.Time `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
}
