package database

import (
	"fmt"
	"log"

	"github.com/agusbasari29/simple-xjx-tasking.git/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDatabaseConnection() *gorm.DB {
	dbConfig := config.DbConfig()
	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s ",
		dbConfig.Host,
		dbConfig.User,
		dbConfig.Pass,
		dbConfig.Port,
		dbConfig.DbName,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to open database connection!")
	}
	return db
}

func CloseDatabaseConnection(db *gorm.DB) {
	conn, err := db.DB()
	if err != nil {
		panic("Failed to close database connection!")
	}
	conn.Close()
}
