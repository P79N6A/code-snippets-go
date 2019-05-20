package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Product struct {
	gorm.Model
	Code string
	Price uint
}


func main() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Code: "L1212", Price: 1000})

	// Read
	var product Product
	db.First(&product, 1)
	fmt.Println(product)
	db.First(&product, "code = ?", "L1212")
	fmt.Println(product)

	// Update
	db.Model(&product).Update("Price", 2000)
	fmt.Println(product)

	// Delete
	db.Delete(&product)
}
