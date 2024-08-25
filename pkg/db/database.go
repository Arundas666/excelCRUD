package db

import (
	"context"
	"fmt"
	"vrwizards/pkg/models"

	"log"
	"os"

	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB    *gorm.DB
	Redis *redis.Client
	Ctx   = context.Background()
)


func SetupDatabase() {
    user := os.Getenv("DB_USER")
    password := os.Getenv("DB_PASSWORD")
    host := os.Getenv("DB_HOST")
    port := os.Getenv("DB_PORT")
    dbname := os.Getenv("DB_NAME")
    sslMode := os.Getenv("DB_SSL_MODE")

    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?tls=%s", user, password, host, port, dbname, sslMode)
    fmt.Println("DSN:", dsn) // Debugging line

    var err error
    DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Error connecting to database: %v", err)
    }

    sqlDB, err := DB.DB()
    if err != nil {
        log.Fatalf("Error getting database connection: %v", err)
    }

    if err = sqlDB.Ping(); err != nil {
        log.Fatalf("Error pinging database: %v", err)
    }
	DB.AutoMigrate(models.Record{})

    log.Println("Database connected successfully")
}
func SetupRedis() {
	opt, _ := redis.ParseURL(os.Getenv("REDIS_HOST"))
	Redis = redis.NewClient(opt)

	_, err := Redis.Ping(Ctx).Result()
	if err != nil {
		log.Fatalf("Error connecting to Redis: %v", err)
	}

	log.Println("Redis connected successfully")
}
