package utils

import (
	"fmt"
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
	log.Printf("config mysql: %v", viper.Get("mysql"))
}

func InitMySQL() (*gorm.DB, error) {
	dsn := viper.GetString("mysql.dns")
	if dsn == "" {
		return nil, fmt.Errorf("mysql dns is not configured")
	}
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}
	return DB, nil
}
