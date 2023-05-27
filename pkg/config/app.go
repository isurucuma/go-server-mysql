package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

// Connect this method is exported from the config module bacically this is used to get connected
// with a database.
// this function started with a capital letter telling that this module should export this function
func Connect() {
	d, err := gorm.Open("mysql", "root:root123@/go-books?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

// GetDB simply returns the db object as a pointer
func GetDB() *gorm.DB {
	return db
}
