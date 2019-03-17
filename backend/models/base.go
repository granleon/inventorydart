package models

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var db *gorm.DB

func init() {
	env := godotenv.Load()
	if env != nil {
		fmt.Println(env)
	}

	username := os.Genenv("db_user")
	password := os.Genenv("db_pass")
	dbName := os.Genenv("db_name")
	dbHost := os.Genenv("db_host")

	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)
	fmt.Println(dbUri)

	conn, err := gorm.Open("postgres", dbUri)
	if err != nil {
		fmt.Println(err)
	}
	db = conn
	db.Debug().AutoMigrate(&Account{}, &Contact{})
}

// GetDB returns a handle to the DB object
func GetDB() *gorm.DB {
	return db
}
