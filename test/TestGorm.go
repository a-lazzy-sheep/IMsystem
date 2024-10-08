package main

import (
	"ginchat/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


func main() {
	db, err := gorm.Open(mysql.Open("root:123456@tcp(localhost:3306)/ginchat?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&models.Community{})

	// Create
	db.Create(&models.Community{Name: "打倒李小东", OwnerId: 4})

	// Read
	// user := models.UserBasic{}

	// db.First(&user, 2)                 // find product with integer primary key
	// db.First(&product, "code = ?", "D42") // find product with code D42

	// Update - update product's price to 200
	// db.Model(&user).Update("Identity", 200)
	// Update - update multiple fields
	// db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	// db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - delete product
	// db.Delete(&user, 2)
}
