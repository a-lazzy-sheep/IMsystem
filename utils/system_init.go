package utils

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB
var RDS *redis.Client

func InitConfig() {
	viper.SetConfigName("app")
	viper.AddConfigPath("config")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	log.Printf("config app: %v", viper.Get("app"))
}

func InitMySQL() {

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)
	dsn := viper.GetString("mysql.dns")
	if dsn == "" {
		log.Printf("mysql dns is not configured")
	}
	DB, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: newLogger})
	log.Printf("config mysql: %v", viper.Get("mysql"))
}

func InitRedis() {
	addr := viper.GetString("redis.addr")
	if addr == "" {
		log.Printf("redis addr is not configured")
	}
	passwd := viper.GetString("redis.passwd")
	if passwd == "" {
		log.Printf("redis passwd is not configured")
	}
	db := viper.GetInt("redis.db")
	if db == 0 {
		log.Printf("redis db is not configured")
	}
	RDS = redis.NewClient(&redis.Options{
		Addr:	addr,
		Password: passwd,
		DB:	db,
	})
	// 增加错误处理机制
	_, err := RDS.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
}

const (
	PublicKey = "websocket"
)

// Publish 消息到redis
func Publish(ctx context.Context, channel string, message string) error {
	err := RDS.Publish(ctx, channel, message).Err()
	if err!= nil {
		return err
	}
	return nil
}

// Subscribe 订阅redis消息
func Subscribe(ctx context.Context, channel string) (string, error) {
	pubsub := RDS.Subscribe(ctx, channel)
	msg, err := pubsub.ReceiveMessage(ctx)
	if err!= nil {
		log.Println("Error subscribing:", err)
		return "", err
	}
	return msg.Payload, err
}

