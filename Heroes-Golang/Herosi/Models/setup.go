//models/setup.go

package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

//SetupModels is an exported function
func SetupModels() *gorm.DB {
	db, err := gorm.Open("sqlite3", "herosi_db.db")

	if err != nil {
		panic("Failed to connect to database!")
	}

	db.AutoMigrate(&Hero{})

	return db
}
