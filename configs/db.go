package configs

import (
	"database/sql"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	_ "github.com/go-sql-driver/mysql"
)

var orm *gorm.DB
var db *sql.DB

func InitDB() {

	var err error

	// Initialize database connection
	dsn := "root:@tcp(127.0.0.1:3306)/techiebutler?parseTime=true"

	orm, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	fmt.Println("Connected to the database")
}

func GetDB() *sql.DB {
	return db
}

func GetORM() *gorm.DB {
	return orm
}
