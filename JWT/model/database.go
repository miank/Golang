package model

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

var secretKey = envVariable("SECRET")

func SetDBClient() {
	var (
		host     = envVariable("DB_HOST")
		port     = envVariable("DB_PORT")
		user     = envVariable("DB_USER")
		password = envVariable("DB_PASSWORD")
		dbName   = envVariable("DB_NAME")
	)

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	DB.AutoMigrate(&User{})
	if err != nil {
		panic(err)
	}

	fmt.Println("Database connected")

}

func envVariable(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		panic(fmt.Sprintf("Environment variable %s not set", key))
	}
	return value
}

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
