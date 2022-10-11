package database

import (
	"autoReload/models"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// config db
const (
	username = "root"
	password = ""
	hostname = "127.0.0.1:3306"
	dbname   = "autoreload"
)

func DBInit() *gorm.DB {

	// conn db
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true&loc=Local", username, password, hostname, dbname))
	if err != nil {
		panic(err)
	}

	// auto migrate models to database table
	db.AutoMigrate(models.AutoReload{})

	return db
}
