package test

import (
	"ginchat/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


func TestDB() {
	db, err := gorm.Open(mysql.Open(""), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&models.UserBasic{})

	// Create
	// db.Create(&models.UserBasic{Name: "Bob", Password: "Bob123"})

	// Read
	user := models.UserBasic{}

	db.First(&user, 2)                 // find product with integer primary key
	// db.First(&product, "code = ?", "D42") // find product with code D42

	// Update - update product's price to 200
	db.Model(&user).Update("Identity", 200)
	// Update - update multiple fields
	// db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	// db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - delete product
	db.Delete(&user, 2)
}
