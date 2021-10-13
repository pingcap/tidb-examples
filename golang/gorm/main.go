package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Order struct {
	Oid   int     `gorm:"primary_key;autoIncrement:true"`
	Uid   int     `gorm:"column:uid"`
	Price float64 `gorm:"column:price"`
}

type GenderModel string

const (
	Female GenderModel = "Female"
	Male   GenderModel = "Male"
)

type User struct {
	Uid    int         `gorm:"primary_key;autoIncrement:true"`
	Name   string      `gorm:"column:name"`
	Gender GenderModel `sql:"type:gender_model"`
}

func PrintResult(tx *gorm.DB, result []Order) {
	if tx.Error == nil && tx.RowsAffected > 0 {
		for _, order := range result {
			fmt.Printf("%+v\n", order)
		}
	}
}

type JoinResult struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func PrintJoinResult(tx *gorm.DB, result []JoinResult) {
	if tx.Error == nil && tx.RowsAffected > 0 {
		for _, order := range result {
			fmt.Printf("%+v\n", order)
		}
	}
}

func main() {
	dsn := "root:@tcp(localhost:4000)/gorm?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Create table
	db.AutoMigrate(&Order{})
	db.AutoMigrate(&User{})

	// Insert
	db.Create(&Order{Uid: 1, Price: 100})
	db.Create(&Order{Uid: 2, Price: 200})
	db.Create(&Order{Uid: 2, Price: 300})
	db.Create(&Order{Uid: 3, Price: 400})
	db.Create(&Order{Uid: 4, Price: 500})

	db.Create(&User{Name: "Alice", Gender: Female})
	db.Create(&User{Name: "John", Gender: Male})
	db.Create(&User{Name: "Ben", Gender: Male})
	db.Create(&User{Name: "Aileen", Gender: Female})

	// Delete
	db.Delete(&Order{}, 1)
	db.Where("uid = ?", 2).Delete(&Order{})

	// Update
	db.Model(&Order{}).Where("uid = ?", 2).Update("price", gorm.Expr("price * ? + ?", 2, 100))

	var orders []Order
	// Get all records
	result := db.Find(&orders)
	PrintResult(result, orders)

	// Get records with conditions
	result = db.Where("uid IN ?", []int{2, 3}).Find(&orders)
	PrintResult(result, orders)

	result = db.Where("price >= ?", 300).Find(&orders)
	PrintResult(result, orders)

	result = db.Raw("SELECT * FROM orders WHERE price = ?", 500).Scan(&orders)
	PrintResult(result, orders)

	var join_result []JoinResult
	// join orders and users
	result = db.Table("users").Select("orders.price as price, users.name as name").Joins("INNER JOIN orders ON orders.uid = users.uid").Where("users.uid = ?", 4).Find(&join_result)
	PrintJoinResult(result, join_result)
}
