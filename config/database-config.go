package configs

import (
	"fmt"
	"os"

	"github.com/Kontribute/kontribute-web-backend/entity"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//setupDatabaseConnectiuon is creating a new connection in opur database

func SetupDatabaseConnection() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		panic("Failed to load env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed To create Connection To Database")
	}

	//lets fill in the model here
	db.AutoMigrate(&entity.Goal{}, &entity.User{})
	return db

}

/**
closeDatabaseConnection method is closing a connection between your app and your database
*/
func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to close connection to the database")
	}
	dbSQL.Close()
}
