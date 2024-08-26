package utils

import (
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitConfig() {
	viper.SetConfigName("app")
	viper.AddConfigPath("config")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	log.Printf("config app: %v", viper.Get("app"))
}

func InitMySQL() {
	dsn := viper.GetString("mysql.dns")
	if dsn == "" {
		log.Printf("mysql dns is not configured")
	}
	DB, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	log.Printf("config mysql: %v", viper.Get("mysql"))
}
