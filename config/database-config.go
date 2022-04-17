package config

import (
	"fmt"
	"os"
	"github.com/joho/godotenv"
	"github.com/Wordyka/golang_gin_gorm_JWT.git/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// melakukan configurasi dan membuat koneksi pada database
func SetupDatabaseConnection() *gorm.DB {
	errEnv := godotenv.Load()	// melakukan load .env / environment file 
	if errEnv != nil {	// jika tidak berhasil di load
		panic("Failed to load env file")
	}
	// melakukan passing suatu string variabel pada environment
	dbUser := os.Getenv("DB_USER")	
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")


	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})	// membuka koneksi pada database berdasarkan environment yang dibuat
	if err != nil {
		panic("Failed to create a connection to database")
	}
	db.AutoMigrate(&entity.Book{}, &entity.User{})	// melakukan migrasi agar table dibuat pada database
	return db
}

// menutup koneksi database jika koneksi fail
func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()	// memanggil database 
	if err != nil {	// jika database yang dipanggil tidak memiliki error
		panic("Failed to close a connection to database")
	}
	dbSQL.Close()	// menutup koneksi database
}
